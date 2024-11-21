-- +migrate Up
-- +migrate StatementBegin


CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255) DEFAULT 'admin',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255) DEFAULT '',
    deleted_at TIMESTAMP
);

INSERT INTO genres (name) VALUES 
    ('Action'),
    ('Strategy'),
    ('Shooter'),
    ('Adventure'),
    ('Horror')
ON CONFLICT (name) DO NOTHING;

CREATE TABLE IF NOT EXISTS games (
    id SERIAL PRIMARY KEY,
    genre_id INTEGER REFERENCES genres(id), 
    title VARCHAR(255) NOT NULL,
    price DECIMAL(20,2) NOT NULL,
    stock INTEGER ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS carts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS cart_items (
    id SERIAL PRIMARY KEY,
    cart_id INTEGER REFERENCES carts(id),
    game_id INTEGER REFERENCES games(id),
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    total_price DECIMAL(20,2) NOT NULL CHECK (total_price >= 0)
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    cart_id INTEGER REFERENCES carts(id),
    total_transaction DECIMAL(20,2) DEFAULT 0,
    status VARCHAR(255) CHECK (status IN ('pending', 'completed', 'cancelled')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);




-- +migrate StatementEnd