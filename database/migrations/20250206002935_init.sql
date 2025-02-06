-- migrate:up
CREATE
EXTENSION IF NOT EXISTS pgcrypto;

SET
intervalstyle = 'iso_8601'; -- Set time interval to be based on the ISO 8601 standard

SET
TIMEZONE = 'UTC'; -- Set session's timezone to UTC 0

-- migrate:down
