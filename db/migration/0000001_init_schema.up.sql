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
    craeted_on  TIMESTAMPTZ NOT NULL DEFAULT (now()),
    updated_on  TIMESTAMPTZ,
    PRIMARY KEY(userid)
);

-- Table: user address
CREATE TABLE IF NOT EXISTS "address" (
    address_id    uuid    DEFAULT uuid_generate_v4 (),
    housenumber VARCHAR(225) NOT NULL,
    streetname  VARCHAR(225) NOT NULL,
    localarea   VARCHAR(225) NOT NULL,
    state       VARCHAR(225) NOT NULL,
    country     VARCHAR(225) NOT NULL,
    PRIMARY KEY(address_id)
);

-- Table data definition commands
ALTER TABLE "users" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("addres_id") ON UPDATE CASCADE ON DELETE CASCADE;
