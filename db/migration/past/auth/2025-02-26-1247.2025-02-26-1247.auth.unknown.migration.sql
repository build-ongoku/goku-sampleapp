CREATE TABLE public.tb_secret (
    id uuid NOT NULL,
    expires_at timestamp without time zone,
    secret text,
    type text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    user_id uuid);

ALTER TABLE public.tb_secret ADD CONSTRAINT chk_secret__secret__empty_str CHECK (secret <> ''::TEXT);

ALTER TABLE public.tb_secret ADD CONSTRAINT chk_secret_type_enum CHECK (type = ANY (ARRAY['GithubToken'::text, 'Password'::text]));

ALTER TABLE public.tb_secret ADD CONSTRAINT pk_secret_id PRIMARY KEY (id);

DROP EXTENSION plpgsql;

