DROP TABLE IF EXISTS users CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;
-- CREATE EXTENSION IF NOT EXISTS postgis;
-- CREATE EXTENSION IF NOT EXISTS postgis_topology;


-- CREATE TABLE IF NOT EXISTS (and remove first line)?
CREATE TABLE users (
    user_id      UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    first_name   VARCHAR(32)                 NOT NULL CHECK ( first_name <> '' ),
    last_name    VARCHAR(32)                 NOT NULL CHECK ( last_name <> '' ),
    email        VARCHAR(64) UNIQUE          NOT NULL CHECK ( email <> '' ),
    password     VARCHAR(250)                NOT NULL CHECK ( octet_length(password) <> 0 ),
    role         VARCHAR(10)                 NOT NULL DEFAULT 'user',
    about        VARCHAR(1024)                        DEFAULT '',
    avatar       VARCHAR(512),
    is_email_verified BOOLEAN                NOT NULL DEFAULT 'false',
    last_login   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP,

    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP
);
-- ,
--     login_date   TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
-- https://github.com/golang-migrate/migrate/blob/v4.6.2/database/postgres/TUTORIAL.md

CREATE TABLE files (
    file_id UUID PRIMARY KEY,
    file_owner_id UUID NOT NULL REFERENCES users (user_id),
    file VARCHAR(256) NOT NULL CHECK (file <> ''),
    filename VARCHAR(256) NOT NULL CHECK (filename <> ''),
    description VARCHAR(1024) DEFAULT '',
    extension VARCHAR(32) NOT NULL,
    size FLOAT NOT NULL,
    tags VARCHAR(256) DEFAULT '',

    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP
)