CREATE TABLE "product" (        
    name text NOT NULL,
    price bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (name)
);