CREATE TABLE IF NOT EXISTS credit_cards(
    id serial PRIMARY KEY,
    name VARCHAR (255) NOT NULL,
    cardholder_name VARCHAR (255) NULL,
    type VARCHAR (255) NULL,
    cardholder_name VARCHAR (255) NULL,
    expire_date VARCHAR (255) NULL,
    valid_from VARCHAR (255) NULL,
    number VARCHAR (255) NULL,
    additional_data TEXT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
    );

CREATE TABLE IF NOT EXISTS logins(
    id serial PRIMARY KEY,
    name VARCHAR (255) NOT NULL,
    username VARCHAR (255) NULL,
    website VARCHAR (255) NULL,
    password VARCHAR (255) NULL,
    additional_data TEXT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
    );

CREATE TABLE IF NOT EXISTS notes(
    id serial PRIMARY KEY,
    name VARCHAR (255) NOT NULL,
    notes TEXT NULL,
    additional_data TEXT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
    );
---- create above / drop below ----

DROP TABLE credit_cards;
DROP TABLE logins;
DROP TABLE notes;
