DELETE FROM "system_menu"
WHERE "code" IN ('system:dbMonitor:groups', 'system:dbMonitor:monitor', 'system:dbMonitor');
