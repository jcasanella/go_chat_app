CREATE TABLE IF NOT EXISTS user (
        id VARCHAR(255) NOT NULL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        username VARCHAR(255) NULL,
        password VARCHARR(255) NULL
    );