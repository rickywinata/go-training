-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "team_member" (        
    team_id uuid,
    member_user_id uuid,
    name text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (team_id, member_user_id),
    FOREIGN KEY (team_id) REFERENCES "team" (id) ON DELETE SET NULL,
    FOREIGN KEY (member_user_id) REFERENCES "user" (id) ON DELETE SET NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "team_member";

