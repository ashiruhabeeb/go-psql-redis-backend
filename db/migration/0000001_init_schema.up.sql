-- PSQL Extension: install uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Table: user
CREATE TABLE IF NOT EXISTS "users" (
    userid      uuid         NOT NULL,
    firstname   VARCHAR(225) NOT NULL,
    lastname    VARCHAR(225) NOT NULL,
    username    VARCHAR(225) NOT NULL UNIQUE,
    email       VARCHAR(225) NOT NULL UNIQUE,
    password    VARCHAR(225) NOT NULL,
    phone       VARCHAR(20) NOT NULL UNIQUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT (now()),
    updated_at  TIMESTAMPTZ,
    PRIMARY KEY(userid)
);

-- Table: user address
CREATE TABLE IF NOT EXISTS "address" (
    address_id   uuid         NOT NULL,
    owner_id     uuid         NOT NULL,
    house_number VARCHAR(225) NOT NULL,
    street_name  VARCHAR(225) NOT NULL,
    local_area   VARCHAR(225) NOT NULL,
    state        VARCHAR(225) NOT NULL,
    country      VARCHAR(225) NOT NULL,
    PRIMARY KEY(address_id)
);

-- Table data definition commands
ALTER TABLE "address" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("userid") ON UPDATE CASCADE ON DELETE CASCADE;

-- Users table updated_at field automatic update function creatiom
CREATE  FUNCTION update_updated_at_users()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Users table updated_at field automatic update trigger creation
CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE
    ON
        users
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_users();
