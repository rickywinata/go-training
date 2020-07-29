//+build integration

package service

import (
	"context"
	"os"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rickywinata/go-training/teamsvc/postgres/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestSuccessCreateTeam(t *testing.T) {
	db := sqlx.MustConnect("postgres", os.Getenv("POSTGRES_URI"))
	defer db.Close()

	usr := &model.User{
		ID:   uuid.New().String(),
		Name: "user1",
	}

	tests := map[string]struct {
		cmd  *CreateTeamCommand
		want *model.Team
	}{
		"success": {
			cmd:  &CreateTeamCommand{OwnerUserID: usr.ID, Name: "team 1"},
			want: &model.Team{OwnerUserID: usr.ID, Name: "team 1"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			c := qt.New(t)
			insertUser(t, db, usr)
			defer resetTables(t, db)

			svc := NewTeamService(db)
			got, err := svc.CreateTeam(context.TODO(), tc.cmd)
			if err != nil {
				t.Fatal(err)
			}

			gotTmMembers, err := got.TeamMembers().All(context.TODO(), db)
			if err != nil {
				t.Fatal(err)
			}

			c.Assert(got, qt.CmpEquals(cmpopts.IgnoreFields(
				model.Team{}, "ID", "L", "R", "CreatedAt", "UpdatedAt"),
			), tc.want)

			c.Assert(gotTmMembers, qt.CmpEquals(cmpopts.IgnoreFields(
				model.TeamMember{}, "L", "R", "CreatedAt", "UpdatedAt"),
			), model.TeamMemberSlice{
				{TeamID: got.ID, MemberUserID: got.OwnerUserID},
			})
		})
	}
}

func insertUser(t *testing.T, db *sqlx.DB, usr *model.User) {
	if err := usr.Insert(context.TODO(), db, boil.Infer()); err != nil {
		t.Fatal(err)
	}
}

func resetTables(t *testing.T, db *sqlx.DB) {
	db.MustExec(`TRUNCATE "user", "team", "team_member"`)
}
