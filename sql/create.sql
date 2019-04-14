CREATE SCHEMA IF NOT EXISTS public;

-- DROP SCHEMA IF EXISTS public CASCADE;
-- CREATE SCHEMA public;

CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE IF NOT EXISTS users (
    nickname        CITEXT PRIMARY KEY COLLATE ucs_basic,
    email           CITEXT UNIQUE,
    "password"  text,
    avatar          TEXT DEFAULT 'default.jpg'
);

CREATE TABLE IF NOT EXISTS games (
    id              SERIAL8 PRIMARY KEY,
    created         TIMESTAMPTZ(3) NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS scores (
    game        BIGINT REFERENCES games,
    player      CITEXT REFERENCES users,
    score       BIGINT DEFAULT 0,
    PRIMARY KEY (game, player)
);