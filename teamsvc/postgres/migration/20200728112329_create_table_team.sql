-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "team" (        
    id uuid,
    owner_user_id uuid NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (id, owner_user_id),
    FOREIGN KEY (owner_user_id) REFERENCES "user" (id) ON DELETE SET NULL,
    UNIQUE (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "team";

