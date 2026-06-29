-- <%.ChineseName%>菜单SQL
-- 生成时间: <%.Date%>
-- id 由数据库自动生成（bigserial），无需手动指定

-- 父级菜单（如果需要）
-- INSERT INTO "system_menu" ("parent_id", "level", "name", "code", "icon", "route", "component", "redirect",
--                            "is_hidden", "type", "status", "sort", "created_by", "updated_by", "created_at", "updated_at",
--                            "deleted_at", "remark")
-- VALUES (0, ',0,', '<%.ChineseName%>管理', '<%.ModuleName%>:<%.VarName%>', 'lucide:database', '<%.ModuleName%>/<%.VarName%>', '', '',
--         2, 'M', 1, 99, 0, 1, NOW(), NOW(), NULL, '');

-- 列表页面菜单（id 自动生成）
INSERT INTO "system_menu" ("parent_id", "level", "name", "code", "icon", "route", "component", "redirect",
                           "is_hidden", "type", "status", "sort", "created_by", "updated_by", "created_at", "updated_at",
                           "deleted_at", "remark")
VALUES (<%.ParentId%>, '<%.Level%>', '<%.ChineseName%>列表', '<%.ModuleName%>:<%.VarName%>', 'lucide:list', '<%.ModuleName%>/<%.VarName%>', '<%.ModuleName%>/<%.VarName%>/index', '',
        2, 'M', 1, 0, 0, 1, NOW(), NOW(), NULL, '');

-- 按钮权限（使用 lastval() 获取刚插入的父菜单 id）
INSERT INTO "system_menu" ("parent_id", "level", "name", "code", "icon", "route", "component", "redirect",
                           "is_hidden", "type", "status", "sort", "created_by", "updated_by", "created_at", "updated_at",
                           "deleted_at", "remark")
VALUES
(lastval(), ',0,' || lastval() || ',', '<%.ChineseName%>列表', '<%.ModuleName%>:<%.VarName%>:index', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(lastval(), ',0,' || lastval() || ',', '<%.ChineseName%>回收站', '<%.ModuleName%>:<%.VarName%>:recycle', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(lastval(), ',0,' || lastval() || ',', '<%.ChineseName%>保存', '<%.ModuleName%>:<%.VarName%>:save', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(lastval(), ',0,' || lastval() || ',', '<%.ChineseName%>更新', '<%.ModuleName%>:<%.VarName%>:update', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(lastval(), ',0,' || lastval() || ',', '<%.ChineseName%>删除', '<%.ModuleName%>:<%.VarName%>:delete', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(lastval(), ',0,' || lastval() || ',', '<%.ChineseName%>读取', '<%.ModuleName%>:<%.VarName%>:read', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(lastval(), ',0,' || lastval() || ',', '<%.ChineseName%>恢复', '<%.ModuleName%>:<%.VarName%>:recovery', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(lastval(), ',0,' || lastval() || ',', '<%.ChineseName%>真实删除', '<%.ModuleName%>:<%.VarName%>:realDelete', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(lastval(), ',0,' || lastval() || ',', '<%.ChineseName%>导入', '<%.ModuleName%>:<%.VarName%>:import', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(lastval(), ',0,' || lastval() || ',', '<%.ChineseName%>导出', '<%.ModuleName%>:<%.VarName%>:export', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(lastval(), ',0,' || lastval() || ',', '<%.ChineseName%>状态改变', '<%.ModuleName%>:<%.VarName%>:changeStatus', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, '');
