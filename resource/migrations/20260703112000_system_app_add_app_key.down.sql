DROP INDEX IF EXISTS public.system_app_app_key_unique_idx;
DROP INDEX IF EXISTS public.system_app_app_id_unique_idx;

ALTER TABLE public.system_app
    DROP COLUMN IF EXISTS app_key;
