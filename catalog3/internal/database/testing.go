package database

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	// databaseName is the name of the template database to clone.
	databaseName = "test-db-template"

	// databaseUser and databasePassword are the username and password for
	// connecting to the database. These values are only used for testing.
	databaseUser     = "test-user"
	databasePassword = "testing123"

	// database container to use.
	postgresImageRepo = "postgres"
	postgresImageTag  = "13-alpine"
)

// TestInstance is a wrapper around the Docker-based database instance.
type TestInstance struct {
	pool      *dockertest.Pool
	container *dockertest.Resource
	db        *sql.DB
	connURL   *url.URL
	connLock  sync.Mutex
}

// MustTestInstance is NewTestInstance but it calls os.Exit(1) on error.
func MustTestInstance(migrationsDir string) *TestInstance {
	i, err := NewTestInstance(migrationsDir)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return i
}

// NewTestInstance creates a new Docker-based database instance.
func NewTestInstance(migrationsDir string) (*TestInstance, error) {
	// Uses a sensible default on windows (tcp/http) and linux/osx (socket).
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("could not connect to docker: %s", err)
	}

	// Pulls an image, creates a container based on it and runs it.
	container, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: postgresImageRepo,
		Tag:        postgresImageTag,
		Env: []string{
			"POSTGRES_DB=" + databaseName,
			"POSTGRES_USER=" + databaseUser,
			"POSTGRES_PASSWORD=" + databasePassword,
		},
	}, func(config *docker.HostConfig) {
		// Set AutoRemove to true so that stopped container goes away by itself.
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		return nil, fmt.Errorf("could not start resource: %s", err)
	}

	// Stop the container after its been running for too long. No since test suite
	// should take super long.
	if err := container.Expire(120); err != nil {
		return nil, fmt.Errorf("failed to expire database container: %w", err)
	}

	// Build the connection URL.
	connectionURL := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(databaseUser, databasePassword),
		Host:     container.GetHostPort("5432/tcp"),
		Path:     databaseName,
		RawQuery: "sslmode=disable",
	}

	// Open db connection.
	db, err := sql.Open("postgres", connectionURL.String())
	if err != nil {
		return nil, err
	}

	// Exponential backoff-retry, because the application in the container might not be ready to accept connections yet.
	if err := pool.Retry(func() error {
		return db.Ping()
	}); err != nil {
		return nil, fmt.Errorf("could not connect to docker: %s", err)
	}

	if err := MigrateUp(connectionURL.String(), migrationsDir); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &TestInstance{
		pool:      pool,
		container: container,
		db:        db,
		connURL:   connectionURL,
	}, nil
}

func MigrateUp(postgresURL string, migrationsDir string) error {
	m, err := migrate.New(
		"file://"+migrationsDir,
		postgresURL)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to migrate up: %w", err)
	}
	srcErr, dbErr := m.Close()
	if srcErr != nil {
		return fmt.Errorf("migrate source error: %w", srcErr)
	}
	if dbErr != nil {
		return fmt.Errorf("migrate database error: %w", dbErr)
	}
	return nil
}

// NewDatabase creates a new sqlx database connection.
func (i *TestInstance) NewDatabase(t *testing.T) *sqlx.DB {
	t.Helper()

	newDatabaseName := randomName()

	q := fmt.Sprintf(`CREATE DATABASE "%s" WITH TEMPLATE "%s";`, newDatabaseName, databaseName)

	// Unfortunately postgres does not allow parallel database creation from the
	// same template, so this is guarded with a lock.
	i.connLock.Lock()
	defer i.connLock.Unlock()

	if _, err := i.db.Exec(q); err != nil {
		t.Fatalf("failed to clone template database: %s", err)
	}

	connectionURL := i.connURL.ResolveReference(&url.URL{Path: newDatabaseName})
	connectionURL.RawQuery = "sslmode=disable"

	dbx, err := sqlx.Connect("postgres", connectionURL.String())
	if err != nil {
		t.Fatalf("failed to connect to the new db: %s", err)
	}

	return dbx
}

// Close purges all resources in the instance.
func (i *TestInstance) Close() {
	if err := i.pool.Purge(i.container); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func randomName() string {
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err) // out of randomness, should never happen
	}

	return fmt.Sprintf("%x", buf)
}
