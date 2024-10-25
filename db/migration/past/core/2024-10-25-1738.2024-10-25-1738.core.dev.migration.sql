ALTER SCHEMA public OWNER TO goku;

REVOKE ALL ON SCHEMA public FROM pg_database_owner;

REVOKE USAGE ON SCHEMA public FROM PUBLIC;

GRANT ALL ON SCHEMA public TO PUBLIC;

GRANT ALL ON SCHEMA public TO goku;

COMMENT ON SCHEMA public IS 'General Public Schema';

CREATE TABLE public.tb_file (
    id uuid NOT NULL,
    file_name text NOT NULL,
    storage_client text NOT NULL,
    storage_client_identifier text NOT NULL,
    size_bytes int NOT NULL,
    mime_type text NOT NULL,
    file_hash text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone);

CREATE TABLE public.tb_task (
    id uuid NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    method text NOT NULL,
    method_request_data jsonb NOT NULL,
    enabled boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone);

CREATE TABLE public.tb_task_run (
    id uuid NOT NULL,
    started_at timestamp without time zone NOT NULL,
    completed_at timestamp without time zone NOT NULL,
    status text NOT NULL,
    method_request_data jsonb NOT NULL,
    method_response_data jsonb NOT NULL,
    triggered_by text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone,
    task_id uuid NOT NULL);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file__file_hash__empty_str CHECK (file_hash <> ''::TEXT);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file__file_name__empty_str CHECK (file_name <> ''::TEXT);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file__mime_type__empty_str CHECK (mime_type <> ''::TEXT);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file__storage_client_identifier__empty_str CHECK (storage_client_identifier <> ''::TEXT);

ALTER TABLE public.tb_file ADD CONSTRAINT chk_file_storage_client_enum CHECK (storage_client = ANY (ARRAY['LocalFile'::text, 'Database'::text, 'CloudAmazonS3'::text]));

ALTER TABLE public.tb_file ADD CONSTRAINT pk_file_id PRIMARY KEY (id);

ALTER TABLE public.tb_task ADD CONSTRAINT chk_task__description__empty_str CHECK (description <> ''::TEXT);

ALTER TABLE public.tb_task ADD CONSTRAINT chk_task__name__empty_str CHECK (name <> ''::TEXT);

ALTER TABLE public.tb_task ADD CONSTRAINT chk_task_method_enum CHECK (method = ANY (ARRAY['Auth_LoginUser'::text, 'Auth_RegisterUser'::text, 'Auth_AuthenticateToken'::text, 'Core_FileUpload'::text, 'Auth_Secret_Add'::text, 'Auth_Secret_Update'::text, 'Auth_Secret_Get'::text, 'Auth_Secret_List'::text, 'Auth_Secret_QueryByText'::text, 'Core_File_Add'::text, 'Core_File_Update'::text, 'Core_File_Get'::text, 'Core_File_List'::text, 'Core_File_QueryByText'::text, 'Core_Task_Add'::text, 'Core_Task_Update'::text, 'Core_Task_Get'::text, 'Core_Task_List'::text, 'Core_Task_QueryByText'::text, 'Core_TaskRun_Add'::text, 'Core_TaskRun_Update'::text, 'Core_TaskRun_Get'::text, 'Core_TaskRun_List'::text, 'Core_TaskRun_QueryByText'::text, 'User_User_Add'::text, 'User_User_Update'::text, 'User_User_Get'::text, 'User_User_List'::text, 'User_User_QueryByText'::text, 'User_Team_Add'::text, 'User_Team_Update'::text, 'User_Team_Get'::text, 'User_Team_List'::text, 'User_Team_QueryByText'::text, 'Core_File_HookCreatePre'::text, 'Core_Task_HookSavePre'::text, 'Core_TaskRun_HookCreatePre'::text, 'User_User_HookInit'::text, 'Core_Task_ActionRun'::text, 'Core_TaskRun_ActionRun'::text]));

ALTER TABLE public.tb_task ADD CONSTRAINT pk_task_id PRIMARY KEY (id);

ALTER TABLE public.tb_task_run ADD CONSTRAINT chk_task_run_status_enum CHECK (status = ANY (ARRAY['Pending'::text, 'Running'::text, 'Success'::text, 'Failed'::text]));

ALTER TABLE public.tb_task_run ADD CONSTRAINT chk_task_run_triggered_by_enum CHECK (triggered_by = ANY (ARRAY['Cron'::text, 'Manual'::text]));

ALTER TABLE public.tb_task_run ADD CONSTRAINT pk_task_run_id PRIMARY KEY (id);

ALTER TABLE public.tb_task_run ADD CONSTRAINT fk_task_run__task_id__task FOREIGN KEY (task_id) REFERENCES public.tb_task (id) ON DELETE CASCADE;

DROP EXTENSION plpgsql;

