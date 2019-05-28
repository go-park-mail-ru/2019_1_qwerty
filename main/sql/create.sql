CREATE SCHEMA IF NOT EXISTS public;

-- DROP SCHEMA IF EXISTS public CASCADE;
-- CREATE SCHEMA public;

CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE IF NOT EXISTS users (
    nickname        CITEXT PRIMARY KEY COLLATE ucs_basic,
    "password"      TEXT,
    avatar          TEXT DEFAULT ''
);


CREATE TABLE IF NOT EXISTS scores (
    player          CITEXT PRIMARY KEY REFERENCES users,
    score           INTEGER DEFAULT 0
);

CREATE OR REPLACE FUNCTION check_score()
    RETURNS TRIGGER AS
    $$
    BEGIN
        IF new.score > old.score
            THEN UPDATE scores SET score = new.score WHERE player = new.player;
            RETURN new;
        END IF;
    END;
    $$
    LANGUAGE 'plpgsql';

CREATE TRIGGER check_score
    AFTER UPDATE ON scores
    FOR EACH ROW EXECUTE PROCEDURE check_score()
