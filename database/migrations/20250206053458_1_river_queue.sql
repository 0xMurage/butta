-- migrate:up

-- River main migration 001 [up]
CREATE TABLE river_migration(
  id bigserial PRIMARY KEY,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  version bigint NOT NULL,
  CONSTRAINT version CHECK (version >= 1)
);

CREATE UNIQUE INDEX ON river_migration USING btree(version);


-- migrate:down

-- River main migration 001 [down]
DROP TABLE river_migration;
