-- Enable CITEXT for case-insensitive text fields
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE developers (
    developer_id BIGSERIAL PRIMARY KEY,
    name CITEXT UNIQUE NOT NULL,
    location VARCHAR(255) NOT NULL,
    founded_year INT NOT NULL,
    description TEXT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    version INT DEFAULT 1 NOT NULL
);

CREATE TABLE games (
    game_id BIGSERIAL PRIMARY KEY,
    name CITEXT UNIQUE NOT NULL,
    genre VARCHAR(100) NOT NULL,
    developer_id BIGINT NOT NULL,
    description TEXT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    version INT DEFAULT 1 NOT NULL,
    FOREIGN KEY (developer_id) REFERENCES developers(developer_id)
);

CREATE TABLE platforms (
    name CITEXT PRIMARY KEY,
    developer_id BIGINT NOT NULL,
    company CITEXT NOT NULL,
    release_date TIMESTAMP NOT NULL,
    description TEXT NOT NULL,
    platform_image_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    version INT DEFAULT 1 NOT NULL,
    FOREIGN KEY (developer_id) REFERENCES developers(developer_id)
);

CREATE TABLE game_releases (
    game_id BIGINT NOT NULL,
    platform_name CITEXT NOT NULL,
    release_date TIMESTAMP NOT NULL,
    average_time INT NOT NULL,
    completionist_time INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    version INT DEFAULT 1 NOT NULL,
    PRIMARY KEY (game_id, platform_name),
    FOREIGN KEY (game_id) REFERENCES games(game_id),
    FOREIGN KEY (platform_name) REFERENCES platforms(name)
);

CREATE TABLE users (
    user_name CITEXT PRIMARY KEY,
    email CITEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    profile_image_url VARCHAR(255) NOT NULL,
    activated BOOLEAN NOT NULL,
    version INT DEFAULT 1 NOT NULL
);

CREATE TABLE user_games (
    user_name CITEXT NOT NULL,
    game_id BIGINT NOT NULL,
    platform_name CITEXT NOT NULL,
    completed BOOLEAN NOT NULL,
    in_progress BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    version INT DEFAULT 1 NOT NULL,
    PRIMARY KEY (user_name, game_id, platform_name),
    FOREIGN KEY (user_name) REFERENCES users(user_name),
    FOREIGN KEY (game_id, platform_name) REFERENCES game_releases(game_id, platform_name)
);
