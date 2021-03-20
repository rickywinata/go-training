CREATE TABLE "product" (        
    name text,
    price integer,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (name)
);