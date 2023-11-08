CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO users (
    id,
    name,
    username,
    password,
    created_at,
    updated_at
)
VALUES
    (
        uuid_generate_v4 (),
        'test_user',
        'test',
        'test',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    );