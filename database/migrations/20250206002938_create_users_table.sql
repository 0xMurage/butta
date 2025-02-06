-- migrate:up
CREATE TABLE users
(
    id         uuid PRIMARY KEY         NOT NULL DEFAULT gen_random_uuid(),
    firstname  TEXT,
    lastname   TEXT,
    email      TEXT                     NOT NULL UNIQUE,
    status     TEXT                     NOT NULL CHECK ( status IN ('ACTIVE', 'SUSPENDED', 'INVITED')),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- migrate:down
DROP TABLE users;