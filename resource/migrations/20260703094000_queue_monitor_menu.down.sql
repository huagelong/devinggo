DELETE FROM "system_menu"
WHERE "code" IN (
  'system:queueMonitor:queues',
  'system:queueMonitor:overview',
  'system:queueMonitor:tasks',
  'system:queueMonitor:schedulerEntries',
  'system:queueMonitor'
);
