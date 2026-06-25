-- Add deleted_at column to setting_crontab table for soft delete recycle bin
ALTER TABLE setting_crontab ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP;

COMMENT ON COLUMN setting_crontab.deleted_at IS '删除时间';
