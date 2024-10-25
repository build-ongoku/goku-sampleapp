ALTER SCHEMA public OWNER TO goku;

REVOKE ALL ON SCHEMA public FROM pg_database_owner;

REVOKE USAGE ON SCHEMA public FROM PUBLIC;

GRANT ALL ON SCHEMA public TO PUBLIC;

GRANT ALL ON SCHEMA public TO goku;

COMMENT ON SCHEMA public IS 'General Public Schema';

CREATE TABLE public.tb_team (
    id uuid NOT NULL,
    name text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone);

CREATE TABLE public.tb_user (
    id uuid NOT NULL,
    email text NOT NULL,
    past_team_ids uuid[],
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone,
    teams_ids uuid[]);

CREATE TABLE public.tb_user_addresses (
    parent_id uuid NOT NULL,
    id uuid NOT NULL,
    line_1 text NOT NULL,
    line_2 text,
    city text NOT NULL,
    state text NOT NULL,
    postal_code text NOT NULL,
    country text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone);

CREATE TABLE public.tb_user_name (
    parent_id uuid NOT NULL,
    id uuid NOT NULL,
    first_name text NOT NULL,
    middle_initial text,
    last_name text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone);

ALTER TABLE public.tb_team ADD CONSTRAINT chk_team__name__empty_str CHECK (name <> ''::TEXT);

ALTER TABLE public.tb_team ADD CONSTRAINT pk_team_id PRIMARY KEY (id);

ALTER TABLE public.tb_user ADD CONSTRAINT pk_user_id PRIMARY KEY (id);

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT chk_addresses__city__empty_str CHECK (city <> ''::TEXT);

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT chk_addresses__line_1__empty_str CHECK (line_1 <> ''::TEXT);

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT chk_addresses__postal_code__empty_str CHECK (postal_code <> ''::TEXT);

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT chk_addresses__state__empty_str CHECK (state <> ''::TEXT);

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT chk_addresses_country_enum CHECK (country = ANY (ARRAY['USA'::text, 'Canada'::text, 'Mexico'::text]));

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT pk_user_addresses_id PRIMARY KEY (id);

ALTER TABLE public.tb_user_name ADD CONSTRAINT chk_name__first_name__empty_str CHECK (first_name <> ''::TEXT);

ALTER TABLE public.tb_user_name ADD CONSTRAINT chk_name__last_name__empty_str CHECK (last_name <> ''::TEXT);

ALTER TABLE public.tb_user_name ADD CONSTRAINT pk_user_name_id PRIMARY KEY (id);

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT fk_user_addresses__parent_id__user FOREIGN KEY (parent_id) REFERENCES public.tb_user (id) ON DELETE CASCADE;

ALTER TABLE public.tb_user_name ADD CONSTRAINT fk_user_name__parent_id__user FOREIGN KEY (parent_id) REFERENCES public.tb_user (id) ON DELETE CASCADE;

DROP EXTENSION plpgsql;

