-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "account" (        
    id uuid,
    balance integer,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    primary key (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "account";
