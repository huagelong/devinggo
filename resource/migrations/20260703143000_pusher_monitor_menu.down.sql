DELETE FROM "system_menu"
WHERE "code" IN (
  'system:pusherMonitor:apps',
  'system:pusherMonitor:overview',
  'system:pusherMonitor:channels',
  'system:pusherMonitor:connections',
  'system:pusherMonitor:terminate',
  'system:pusherMonitor:debugEvent',
  'system:pusherMonitor'
);
