INSERT INTO "system_menu"(
  "id","parent_id","level","name","code","icon","route","component","redirect",
  "is_hidden","type","status","sort","created_by","updated_by",
  "created_at","updated_at","deleted_at","remark"
) VALUES
  (3760, 2000, ',0,1000,2000,', '队列监控', 'system:queueMonitor', 'lucide:workflow', 'queueMonitor', 'system/monitor/queue/index', '', 2, 'M', 1, 100, 0, 1, NOW(), NOW(), NULL, ''),
  (3761, 3760, ',0,1000,2000,3760,', '队列列表', 'system:queueMonitor:queues', '', '', '', '', 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
  (3762, 3760, ',0,1000,2000,3760,', '队列概览', 'system:queueMonitor:overview', '', '', '', '', 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
  (3763, 3760, ',0,1000,2000,3760,', '队列任务', 'system:queueMonitor:tasks', '', '', '', '', 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
  (3764, 3760, ',0,1000,2000,3760,', '队列调度', 'system:queueMonitor:schedulerEntries', '', '', '', '', 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, '');
