ALTER SCHEMA public OWNER TO goku;

REVOKE ALL ON SCHEMA public FROM pg_database_owner;

REVOKE USAGE ON SCHEMA public FROM PUBLIC;

GRANT ALL ON SCHEMA public TO PUBLIC;

GRANT ALL ON SCHEMA public TO goku;

COMMENT ON SCHEMA public IS 'General Public Schema';

CREATE TABLE public.tb_file (
    id UUID NOT NULL,
    file_name TEXT NOT NULL,
    storage_client TEXT NOT NULL,
    storage_client_identifier TEXT NOT NULL,
    size_bytes INT NOT NULL,
    mime_type TEXT NOT NULL,
    file_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file__file_hash__empty_str CHECK (file_hash <> ''::TEXT);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file__file_name__empty_str CHECK (file_name <> ''::TEXT);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file__mime_type__empty_str CHECK (mime_type <> ''::TEXT);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file__storage_client_identifier__empty_str CHECK (storage_client_identifier <> ''::TEXT);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file_storage_client_enum CHECK (storage_client = ANY (ARRAY['LocalFile'::text, 'Database'::text, 'CloudAmazonS3'::text]));

ALTER TABLE public.tb_file ADD CONSTRAINT pk_file_id PRIMARY KEY (id);

DROP EXTENSION plpgsql;

