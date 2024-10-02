ALTER SCHEMA public OWNER TO goku;

REVOKE ALL ON SCHEMA public FROM pg_database_owner;

REVOKE USAGE ON SCHEMA public FROM PUBLIC;

GRANT ALL ON SCHEMA public TO PUBLIC;

GRANT ALL ON SCHEMA public TO goku;

COMMENT ON SCHEMA public IS 'General Public Schema';

CREATE TABLE public.tb_team (
    id UUID NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP);

CREATE TABLE public.tb_user (
    id UUID NOT NULL,
    email TEXT NOT NULL,
    team_id UUID NOT NULL,
    past_team_ids UUID[],
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP);

CREATE TABLE public.tb_user_addresses (
    parent_id UUID NOT NULL,
    id UUID NOT NULL,
    line_1 TEXT NOT NULL,
    line_2 TEXT,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    postal_code TEXT NOT NULL,
    country TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP);

CREATE TABLE public.tb_user_name (
    parent_id UUID NOT NULL,
    id UUID NOT NULL,
    first_name TEXT NOT NULL,
    middle_initial TEXT,
    last_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP);

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

ALTER TABLE public.tb_user ADD CONSTRAINT fk_user__team_id__team FOREIGN KEY (team_id) REFERENCES public.tb_team (id) ON DELETE CASCADE;

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT fk_user_addresses__parent_id__user FOREIGN KEY (parent_id) REFERENCES public.tb_user (id) ON DELETE CASCADE;

ALTER TABLE public.tb_user_name ADD CONSTRAINT fk_user_name__parent_id__user FOREIGN KEY (parent_id) REFERENCES public.tb_user (id) ON DELETE CASCADE;

DROP EXTENSION plpgsql;

