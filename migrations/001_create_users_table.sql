CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    login VARCHAR (255) NOT NULL UNIQUE,
    password VARCHAR (255) NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
    );
---- create above / drop below ----

DROP TABLE users;
