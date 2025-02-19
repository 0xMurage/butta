-- migrate:up
CREATE TABLE user_auth_methods
(
    id               uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_profile_id  uuid             NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
    auth_provider_id BIGINT           NOT NULL REFERENCES user_auth_providers (id) ON DELETE RESTRICT,
    login_identifier  TEXT             NOT NULL, -- identifier such as phone number, username or email
    secret_hash      TEXT,                      --hash which is secret e.g.  password hash
    created_at       TIMESTAMPTZ      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (auth_provider_id, login_identifier)
);


-- migrate:down

DROP TABLE user_auth_methods;