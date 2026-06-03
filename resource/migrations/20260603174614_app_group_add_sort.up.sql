ALTER TABLE system_app_group ADD COLUMN sort SMALLINT DEFAULT 0;
COMMENT ON COLUMN system_app_group.sort IS '排序';
