CREATE SCHEMA IF NOT EXISTS user_schema;

CREATE TABLE user_schema.users
(
    id             UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    username       VARCHAR(32)  NOT NULL UNIQUE,
    password_hash  VARCHAR(255) NOT NUll,
    salt           VARCHAR(255) NOT NUll,
    email          VARCHAR(255) NOT NULL UNIQUE,
    email_verified BOOLEAN          DEFAULT FALSE,
    first_name     VARCHAR(32),
    last_name      VARCHAR(32),
    date_of_birth  DATE,
    created_at     TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);