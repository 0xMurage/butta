-- migrate:up
CREATE TABLE user_auth_providers
(
    id         SERIAL PRIMARY KEY,
    name       TEXT        NOT NULL UNIQUE,
    enabled    BOOL        NOT NULL DEFAULT true,
    properties JSONB       NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO user_auth_providers(name)
VALUES ('basic-password-auth'),
       ('mobile-phone'),
       ('magic-link');

-- migrate:down
DROP TABLE user_auth_providers;

