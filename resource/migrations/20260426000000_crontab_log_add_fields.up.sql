-- Add start_time, end_time, output columns to setting_crontab_log table
ALTER TABLE setting_crontab_log ADD COLUMN IF NOT EXISTS start_time TIMESTAMP;
ALTER TABLE setting_crontab_log ADD COLUMN IF NOT EXISTS end_time TIMESTAMP;
ALTER TABLE setting_crontab_log ADD COLUMN IF NOT EXISTS output TEXT;

COMMENT ON COLUMN setting_crontab_log.start_time IS '任务开始执行时间';
COMMENT ON COLUMN setting_crontab_log.end_time IS '任务执行结束时间';
COMMENT ON COLUMN setting_crontab_log.output IS '任务执行输出';
