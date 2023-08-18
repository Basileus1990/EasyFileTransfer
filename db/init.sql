CREATE DATABASE easy_file_transfer;


CREATE TABLE users (
    email       VARCHAR(100) NOT NULL PRIMARY KEY,
    passhash    VARCHAR(256) NOT NULL
);

CREATE TABLE directories (
    id          SERIAL          PRIMARY KEY,
    name        VARCHAR(128)    NOT NULL,
    structure   JSON            NOT NULL,
    owner       VARCHAR(100)    NOT NULL REFERENCES users(email),
    created_at  TIMESTAMP       NOT NULL DEFAULT NOW()
);