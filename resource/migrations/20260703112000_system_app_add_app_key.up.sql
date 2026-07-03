ALTER TABLE public.system_app
    ADD COLUMN IF NOT EXISTS app_key varchar(32);

UPDATE public.system_app
SET app_key = app_id
WHERE app_key IS NULL OR app_key = '';

ALTER TABLE public.system_app
    ALTER COLUMN app_key SET NOT NULL;

CREATE UNIQUE INDEX IF NOT EXISTS system_app_app_id_unique_idx
    ON public.system_app (app_id)
    WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX IF NOT EXISTS system_app_app_key_unique_idx
    ON public.system_app (app_key)
    WHERE deleted_at IS NULL;
