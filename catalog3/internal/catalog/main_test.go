package catalog

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rickywinata/go-training/catalog3/internal/database"
)

var testDatabaseInstance *database.TestInstance

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	testDatabaseInstance = database.MustTestInstance(database.MigrationsDir())
	defer testDatabaseInstance.Close()
	return m.Run()
}
