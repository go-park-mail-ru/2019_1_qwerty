CREATE SCHEMA IF NOT EXISTS public;

-- DROP SCHEMA IF EXISTS public CASCADE;
-- CREATE SCHEMA public;

CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE IF NOT EXISTS users (
    nickname        CITEXT PRIMARY KEY COLLATE ucs_basic,
    "password"      TEXT,
    avatar          TEXT DEFAULT ''
);

DROP TABLE IF EXISTS scores CASCADE;
CREATE TABLE IF NOT EXISTS scores (
    player          CITEXT PRIMARY KEY REFERENCES users,
    score           INTEGER DEFAULT 0
);

DROP TRIGGER IF EXISTS check_score ON scores;
