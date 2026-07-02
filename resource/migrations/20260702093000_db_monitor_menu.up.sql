INSERT INTO "system_menu"(
  "id","parent_id","level","name","code","icon","route","component","redirect",
  "is_hidden","type","status","sort","created_by","updated_by",
  "created_at","updated_at","deleted_at","remark"
) VALUES
  (3750, 2000, ',0,1000,2000,', '数据库监控', 'system:dbMonitor', 'lucide:database', 'dbMonitor', 'system/monitor/db/index', '', 2, 'M', 1, 99, 0, 1, NOW(), NOW(), NULL, ''),
  (3751, 3750, ',0,1000,2000,3750,', '数据库分组列表', 'system:dbMonitor:groups', '', '', '', '', 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
  (3752, 3750, ',0,1000,2000,3750,', '数据库监控数据', 'system:dbMonitor:monitor', '', '', '', '', 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, '');
