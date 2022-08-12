CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id                      UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    login VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE users_data
(
    id                      UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    user_id                 UUID NOT NULL REFERENCES users (id),
    text                    TEXT,
    number                  VARCHAR(255),
    type                    INTEGER,
    meta                    VARCHAR(255),
    created_at              TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON users_data
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();