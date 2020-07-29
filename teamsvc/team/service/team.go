package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rickywinata/go-training/teamsvc/postgres/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type (
	// CreateTeamCommand represents the parameters for creating a new team.
	CreateTeamCommand struct {
		OwnerUserID string `json:"owner_user_id"`
		Name        string `json:"name"`
	}
)

// TeamService is an interface for team-related operations.
type TeamService interface {
	CreateTeam(ctx context.Context, cmd *CreateTeamCommand) (*model.Team, error)
}

type teamService struct {
	db *sqlx.DB
}

// NewTeamService creates a new team service.
func NewTeamService(db *sqlx.DB) TeamService {
	return &teamService{db}
}

func (s *teamService) CreateTeam(ctx context.Context, cmd *CreateTeamCommand) (*model.Team, error) {
	tm := &model.Team{
		ID:          uuid.New().String(),
		OwnerUserID: cmd.OwnerUserID,
		Name:        cmd.Name,
	}
	if err := tm.Insert(ctx, s.db, boil.Infer()); err != nil {
		return nil, err
	}

	tmMember := &model.TeamMember{
		TeamID:       tm.ID,
		MemberUserID: tm.OwnerUserID,
	}
	if err := tmMember.Insert(ctx, s.db, boil.Infer()); err != nil {
		return nil, err
	}

	return tm, nil
}
