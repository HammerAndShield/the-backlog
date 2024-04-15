-- Down migration script for PostgreSQL
BEGIN;

-- Order is important due to foreign key constraints
DROP TABLE IF EXISTS user_games CASCADE;
DROP TABLE IF EXISTS game_releases CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS games CASCADE;
DROP TABLE IF EXISTS platforms CASCADE;
DROP TABLE IF EXISTS developers CASCADE;

COMMIT;