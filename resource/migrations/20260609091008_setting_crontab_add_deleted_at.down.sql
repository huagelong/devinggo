-- Remove deleted_at column from setting_crontab table
ALTER TABLE setting_crontab DROP COLUMN IF EXISTS deleted_at;
