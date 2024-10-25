ALTER SCHEMA public OWNER TO goku;

REVOKE ALL ON SCHEMA public FROM pg_database_owner;

REVOKE USAGE ON SCHEMA public FROM PUBLIC;

GRANT ALL ON SCHEMA public TO PUBLIC;

GRANT ALL ON SCHEMA public TO goku;

COMMENT ON SCHEMA public IS 'General Public Schema';

CREATE TABLE public.tb_secret (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    type text NOT NULL,
    secret text NOT NULL,
    expires_at timestamp without time zone,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone);

ALTER TABLE public.tb_secret ADD CONSTRAINT chk_secret__secret__empty_str CHECK (secret <> ''::TEXT);

ALTER TABLE public.tb_secret ADD CONSTRAINT chk_secret_type_enum CHECK (type = ANY (ARRAY['Password'::text, 'GithubToken'::text]));

ALTER TABLE public.tb_secret ADD CONSTRAINT pk_secret_id PRIMARY KEY (id);

DROP EXTENSION plpgsql;

