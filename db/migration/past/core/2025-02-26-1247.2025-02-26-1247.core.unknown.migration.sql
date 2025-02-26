CREATE TABLE public.tb_file (
    id uuid NOT NULL,
    file_hash text,
    file_name text,
    mime_type text,
    size_bytes integer,
    storage_client text,
    storage_client_identifier text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone);

CREATE TABLE public.tb_secret_decryptable (
    id uuid NOT NULL,
    encrypted_value text,
    salt text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    secret_key_id uuid);

CREATE TABLE public.tb_secret_key (
    id uuid NOT NULL,
    private_key_format text,
    public_key text,
    public_key_format text,
    type text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file__file_hash__empty_str CHECK (file_hash <> ''::TEXT);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file__file_name__empty_str CHECK (file_name <> ''::TEXT);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file__mime_type__empty_str CHECK (mime_type <> ''::TEXT);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file__storage_client_identifier__empty_str CHECK (storage_client_identifier <> ''::TEXT);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file_storage_client_enum CHECK (storage_client = ANY (ARRAY['CloudAmazonS3'::text, 'Database'::text, 'LocalFile'::text]));

ALTER TABLE public.tb_file ADD CONSTRAINT pk_file_id PRIMARY KEY (id);

ALTER TABLE public.tb_secret_decryptable ADD CONSTRAINT chk_secret_decryptable__encrypted_value__empty_str CHECK (encrypted_value <> ''::TEXT);

ALTER TABLE public.tb_secret_decryptable ADD CONSTRAINT chk_secret_decryptable__salt__empty_str CHECK (salt <> ''::TEXT);

ALTER TABLE public.tb_secret_decryptable ADD CONSTRAINT pk_secret_decryptable_id PRIMARY KEY (id);

ALTER TABLE public.tb_secret_key ADD CONSTRAINT chk_secret_key__public_key__empty_str CHECK (public_key <> ''::TEXT);

ALTER TABLE public.tb_secret_key ADD CONSTRAINT chk_secret_key_private_key_format_enum CHECK (private_key_format = ANY (ARRAY['OpenSSH'::text, 'PEM'::text]));

ALTER TABLE public.tb_secret_key ADD CONSTRAINT chk_secret_key_public_key_format_enum CHECK (public_key_format = ANY (ARRAY['OpenSSH'::text, 'PEM'::text]));

ALTER TABLE public.tb_secret_key ADD CONSTRAINT chk_secret_key_type_enum CHECK (type = ANY (ARRAY['Ed25519'::text, 'RSA'::text]));

ALTER TABLE public.tb_secret_key ADD CONSTRAINT pk_secret_key_id PRIMARY KEY (id);

ALTER TABLE public.tb_secret_decryptable ADD CONSTRAINT fk_secret_decryptable__secret_key_id__secret_key FOREIGN KEY (secret_key_id) REFERENCES public.tb_secret_key (id) ON DELETE CASCADE;

DROP EXTENSION plpgsql;

