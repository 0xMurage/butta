
--
-- Name: migrations; Type: TABLE;  Owner: -
--

CREATE TABLE migrations (
    version character varying(128) NOT NULL
);


--
-- Name: river_client; Type: TABLE;  Owner: -
--

CREATE UNLOGGED TABLE river_client (
    id text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    metadata jsonb DEFAULT '{}'::jsonb NOT NULL,
    paused_at timestamp with time zone,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT name_length CHECK (((char_length(id) > 0) AND (char_length(id) < 128)))
);


--
-- Name: river_client_queue; Type: TABLE;  Owner: -
--

CREATE UNLOGGED TABLE river_client_queue (
    river_client_id text NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    max_workers bigint DEFAULT 0 NOT NULL,
    metadata jsonb DEFAULT '{}'::jsonb NOT NULL,
    num_jobs_completed bigint DEFAULT 0 NOT NULL,
    num_jobs_running bigint DEFAULT 0 NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT name_length CHECK (((char_length(name) > 0) AND (char_length(name) < 128))),
    CONSTRAINT num_jobs_completed_zero_or_positive CHECK ((num_jobs_completed >= 0)),
    CONSTRAINT num_jobs_running_zero_or_positive CHECK ((num_jobs_running >= 0))
);


--
-- Name: river_job; Type: TABLE;  Owner: -
--

CREATE TABLE river_job (
    id bigint NOT NULL,
    state river_job_state DEFAULT 'available'::soi.river_job_state NOT NULL,
    attempt smallint DEFAULT 0 NOT NULL,
    max_attempts smallint NOT NULL,
    attempted_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    finalized_at timestamp with time zone,
    scheduled_at timestamp with time zone DEFAULT now() NOT NULL,
    priority smallint DEFAULT 1 NOT NULL,
    args jsonb NOT NULL,
    attempted_by text[],
    errors jsonb[],
    kind text NOT NULL,
    metadata jsonb DEFAULT '{}'::jsonb NOT NULL,
    queue text DEFAULT 'default'::text NOT NULL,
    tags character varying(255)[] DEFAULT '{}'::character varying[] NOT NULL,
    unique_key bytea,
    unique_states bit(8),
    CONSTRAINT finalized_or_finalized_at_null CHECK ((((finalized_at IS NULL) AND (state <> ALL (ARRAY['cancelled'::river_job_state, 'completed'::soi.river_job_state, 'discarded'::soi.river_job_state]))) OR ((finalized_at IS NOT NULL) AND (state = ANY (ARRAY['cancelled'::soi.river_job_state, 'completed'::soi.river_job_state, 'discarded'::soi.river_job_state]))))),
    CONSTRAINT kind_length CHECK (((char_length(kind) > 0) AND (char_length(kind) < 128))),
    CONSTRAINT max_attempts_is_positive CHECK ((max_attempts > 0)),
    CONSTRAINT priority_in_range CHECK (((priority >= 1) AND (priority <= 4))),
    CONSTRAINT queue_length CHECK (((char_length(queue) > 0) AND (char_length(queue) < 128)))
);


--
-- Name: river_job_id_seq; Type: SEQUENCE;  Owner: -
--

CREATE SEQUENCE river_job_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: river_job_id_seq; Type: SEQUENCE OWNED BY;  Owner: -
--

ALTER SEQUENCE river_job_id_seq OWNED BY soi.river_job.id;


--
-- Name: river_leader; Type: TABLE;  Owner: -
--

CREATE UNLOGGED TABLE river_leader (
    elected_at timestamp with time zone NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    leader_id text NOT NULL,
    name text DEFAULT 'default'::text NOT NULL,
    CONSTRAINT leader_id_length CHECK (((char_length(leader_id) > 0) AND (char_length(leader_id) < 128))),
    CONSTRAINT name_length CHECK ((name = 'default'::text))
);


--
-- Name: river_migration; Type: TABLE;  Owner: -
--

CREATE TABLE river_migration (
    line text NOT NULL,
    version bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT line_length CHECK (((char_length(line) > 0) AND (char_length(line) < 128))),
    CONSTRAINT version_gte_1 CHECK ((version >= 1))
);


--
-- Name: river_queue; Type: TABLE;  Owner: -
--

CREATE TABLE river_queue (
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    metadata jsonb DEFAULT '{}'::jsonb NOT NULL,
    paused_at timestamp with time zone,
    updated_at timestamp with time zone NOT NULL
);


--
-- Name: user_auth_methods; Type: TABLE;  Owner: -
--

CREATE TABLE user_auth_methods (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_profile_id uuid NOT NULL,
    auth_provider_id bigint NOT NULL,
    login_identifier text NOT NULL,
    secret_hash text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: user_auth_providers; Type: TABLE;  Owner: -
--

CREATE TABLE user_auth_providers (
    id integer NOT NULL,
    name text NOT NULL,
    enabled boolean DEFAULT true NOT NULL,
    properties jsonb DEFAULT '{}'::jsonb NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: user_auth_providers_id_seq; Type: SEQUENCE;  Owner: -
--

CREATE SEQUENCE user_auth_providers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_auth_providers_id_seq; Type: SEQUENCE OWNED BY;  Owner: -
--

ALTER SEQUENCE user_auth_providers_id_seq OWNED BY soi.user_auth_providers.id;


--
-- Name: users; Type: TABLE;  Owner: -
--

CREATE TABLE users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    firstname text,
    lastname text,
    email text NOT NULL,
    status text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT users_status_check CHECK ((status = ANY (ARRAY['ACTIVE'::text, 'SUSPENDED'::text, 'INVITED'::text])))
);


--
-- Name: river_job id; Type: DEFAULT;  Owner: -
--

ALTER TABLE ONLY river_job ALTER COLUMN id SET DEFAULT nextval('soi.river_job_id_seq'::regclass);


--
-- Name: user_auth_providers id; Type: DEFAULT;  Owner: -
--

ALTER TABLE ONLY user_auth_providers ALTER COLUMN id SET DEFAULT nextval('soi.user_auth_providers_id_seq'::regclass);


--
-- Name: migrations migrations_pkey; Type: CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (version);


--
-- Name: river_client river_client_pkey; Type: CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY river_client
    ADD CONSTRAINT river_client_pkey PRIMARY KEY (id);


--
-- Name: river_client_queue river_client_queue_pkey; Type: CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY river_client_queue
    ADD CONSTRAINT river_client_queue_pkey PRIMARY KEY (river_client_id, name);


--
-- Name: river_job river_job_pkey; Type: CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY river_job
    ADD CONSTRAINT river_job_pkey PRIMARY KEY (id);


--
-- Name: river_leader river_leader_pkey; Type: CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY river_leader
    ADD CONSTRAINT river_leader_pkey PRIMARY KEY (name);


--
-- Name: river_migration river_migration_pkey1; Type: CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY river_migration
    ADD CONSTRAINT river_migration_pkey1 PRIMARY KEY (line, version);


--
-- Name: river_queue river_queue_pkey; Type: CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY river_queue
    ADD CONSTRAINT river_queue_pkey PRIMARY KEY (name);


--
-- Name: user_auth_methods user_auth_methods_auth_provider_id_login_identifier_key; Type: CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY user_auth_methods
    ADD CONSTRAINT user_auth_methods_auth_provider_id_login_identifier_key UNIQUE (auth_provider_id, login_identifier);


--
-- Name: user_auth_methods user_auth_methods_pkey; Type: CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY user_auth_methods
    ADD CONSTRAINT user_auth_methods_pkey PRIMARY KEY (id);


--
-- Name: user_auth_providers user_auth_providers_name_key; Type: CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY user_auth_providers
    ADD CONSTRAINT user_auth_providers_name_key UNIQUE (name);


--
-- Name: user_auth_providers user_auth_providers_pkey; Type: CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY user_auth_providers
    ADD CONSTRAINT user_auth_providers_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: river_job_args_index; Type: INDEX;  Owner: -
--

CREATE INDEX river_job_args_index ON river_job USING gin (args);


--
-- Name: river_job_kind; Type: INDEX;  Owner: -
--

CREATE INDEX river_job_kind ON river_job USING btree (kind);


--
-- Name: river_job_metadata_index; Type: INDEX;  Owner: -
--

CREATE INDEX river_job_metadata_index ON river_job USING gin (metadata);


--
-- Name: river_job_prioritized_fetching_index; Type: INDEX;  Owner: -
--

CREATE INDEX river_job_prioritized_fetching_index ON river_job USING btree (state, queue, priority, scheduled_at, id);


--
-- Name: river_job_state_and_finalized_at_index; Type: INDEX;  Owner: -
--

CREATE INDEX river_job_state_and_finalized_at_index ON river_job USING btree (state, finalized_at) WHERE (finalized_at IS NOT NULL);


--
-- Name: river_job_unique_idx; Type: INDEX;  Owner: -
--

CREATE UNIQUE INDEX river_job_unique_idx ON river_job USING btree (unique_key) WHERE ((unique_key IS NOT NULL) AND (unique_states IS NOT NULL) AND soi.river_job_state_in_bitmask(unique_states, state));


--
-- Name: river_client_queue river_client_queue_river_client_id_fkey; Type: FK CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY river_client_queue
    ADD CONSTRAINT river_client_queue_river_client_id_fkey FOREIGN KEY (river_client_id) REFERENCES river_client(id) ON DELETE CASCADE;


--
-- Name: user_auth_methods user_auth_methods_auth_provider_id_fkey; Type: FK CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY user_auth_methods
    ADD CONSTRAINT user_auth_methods_auth_provider_id_fkey FOREIGN KEY (auth_provider_id) REFERENCES user_auth_providers(id) ON DELETE RESTRICT;


--
-- Name: user_auth_methods user_auth_methods_user_profile_id_fkey; Type: FK CONSTRAINT;  Owner: -
--

ALTER TABLE ONLY user_auth_methods
    ADD CONSTRAINT user_auth_methods_user_profile_id_fkey FOREIGN KEY (user_profile_id) REFERENCES users(id) ON DELETE RESTRICT;


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO migrations (version) VALUES
    ('20250206002935'),
    ('20250206002938'),
    ('20250206053458'),
    ('20250206053501'),
    ('20250206053503'),
    ('20250206053505'),
    ('20250206053507'),
    ('20250206053509'),
    ('20250219034909'),
    ('20250219034913');
