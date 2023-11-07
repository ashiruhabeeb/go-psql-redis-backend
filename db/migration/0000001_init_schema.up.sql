-- PSQL Extension: install uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Table: user
CREATE TABLE IF NOT EXISTS "users" (
    userid      uuid    DEFAULT uuid_generate_v4 (),
    firstname   VARCHAR(225) NOT NULL,
    lastname    VARCHAR(225) NOT NULL,
    username    VARCHAR(225) NOT NULL UNIQUE,
    email       VARCHAR(225) NOT NULL UNIQUE,
    password    VARCHAR(225) NOT NULL,
    phone       VARCHAR(20) NOT NULL UNIQUE,
    address_id  uuid NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT (now()),
    updated_at  TIMESTAMPTZ,
    PRIMARY KEY(userid)
);

-- Table: user address
CREATE TABLE IF NOT EXISTS "address" (
    address_id   uuid DEFAULT uuid_generate_v4 (),
    house_number VARCHAR(225) NOT NULL,
    street_name  VARCHAR(225) NOT NULL,
    local_area   VARCHAR(225) NOT NULL,
    state        VARCHAR(225) NOT NULL,
    country      VARCHAR(225) NOT NULL,
    PRIMARY KEY(address_id)
);

-- Table data definition commands
ALTER TABLE "users" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("address_id") ON UPDATE CASCADE ON DELETE CASCADE;
