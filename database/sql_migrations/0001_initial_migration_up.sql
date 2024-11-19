-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS games (
    id INTEGER PRIMARY KEY,
    genre_id INTEGER REFERENCES genre(id),
    title VARCHAR(255) NOT NULL,
    price DECIMAL(20,2) NOT NULL,
    stock INTEGER ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255)
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS transaction (
    id INTEGER PRIMARY KEY,
    user_id INTEGER REFERENCES user(id),
    game_id INTEGER REFERENCES games(id),
    quantity INTEGER,
    total_price DECIMAL(20,2)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255)
)

CREATE TABLE IF NOT EXISTS genre (
    id INTEGER PRIMARY KEY,
    genre VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255)
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255)
) 

-- +migrate StatementEnd