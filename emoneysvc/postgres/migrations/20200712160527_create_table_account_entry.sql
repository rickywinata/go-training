-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "account_entry" (
    id uuid,
    account_id uuid,
    amount integer,
    booked_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES "account" (id) ON DELETE CASCADE
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "account_entry";
