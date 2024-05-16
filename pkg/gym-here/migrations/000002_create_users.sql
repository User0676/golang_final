CREATE TABLE IF NOT EXISTS users(
                       id SERIAL PRIMARY KEY,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP,
                       username VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL
);