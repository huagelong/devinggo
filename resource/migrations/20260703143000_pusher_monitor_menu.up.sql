INSERT INTO "system_menu"(
  "id","parent_id","level","name","code","icon","route","component","redirect",
  "is_hidden","type","status","sort","created_by","updated_by",
  "created_at","updated_at","deleted_at","remark"
) VALUES
  (3770, 2000, ',0,1000,2000,', 'Pusher监控', 'system:pusherMonitor', 'lucide:radio-tower', 'pusherMonitor', 'system/monitor/pusher/index', '', 2, 'M', 1, 101, 0, 1, NOW(), NOW(), NULL, ''),
  (3771, 3770, ',0,1000,2000,3770,', 'Pusher App列表', 'system:pusherMonitor:apps', '', '', '', '', 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
  (3772, 3770, ',0,1000,2000,3770,', 'Pusher概览', 'system:pusherMonitor:overview', '', '', '', '', 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
  (3773, 3770, ',0,1000,2000,3770,', 'Pusher频道', 'system:pusherMonitor:channels', '', '', '', '', 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
  (3774, 3770, ',0,1000,2000,3770,', 'Pusher连接', 'system:pusherMonitor:connections', '', '', '', '', 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
  (3775, 3770, ',0,1000,2000,3770,', 'Pusher批量断连', 'system:pusherMonitor:terminate', '', '', '', '', 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
  (3776, 3770, ',0,1000,2000,3770,', 'Pusher事件调试', 'system:pusherMonitor:debugEvent', '', '', '', '', 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, '');
