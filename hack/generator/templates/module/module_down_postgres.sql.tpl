-- Drop {{.moduleName}} module tables
-- Author: devinggo
-- Date: {{.date}}

DROP TABLE IF EXISTS {{.moduleName}}_example;

DELETE FROM system_modules WHERE name = '{{.moduleName}}';
