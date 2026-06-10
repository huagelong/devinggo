-- <%.ChineseName%>菜单SQL
-- 生成时间: <%.Date%>
-- 请根据实际情况修改 parent_id 和 id 值

-- 父级菜单（如果需要）
-- INSERT INTO "system_menu" ("id", "parent_id", "level", "name", "code", "icon", "route", "component", "redirect",
--                            "is_hidden", "type", "status", "sort", "created_by", "updated_by", "created_at", "updated_at",
--                            "deleted_at", "remark")
-- VALUES (<%.ParentMenuId%>, 0, ',0,', '<%.ChineseName%>管理', '<%.ModuleName%>:<%.VarName%>', 'lucide:database', '<%.ModuleName%>/<%.VarName%>', '', '',
--         2, 'M', 1, 99, 0, 1, NOW(), NOW(), NULL, '');

-- 列表页面菜单
INSERT INTO "system_menu" ("id", "parent_id", "level", "name", "code", "icon", "route", "component", "redirect",
                           "is_hidden", "type", "status", "sort", "created_by", "updated_by", "created_at", "updated_at",
                           "deleted_at", "remark")
VALUES (<%.MenuId%>, <%.ParentId%>, '<%.Level%>', '<%.ChineseName%>列表', '<%.ModuleName%>:<%.VarName%>', 'lucide:list', '<%.ModuleName%>/<%.VarName%>', '<%.ModuleName%>/<%.VarName%>/index', '',
        2, 'M', 1, 0, 0, 1, NOW(), NOW(), NULL, '');

-- 按钮权限
INSERT INTO "system_menu" ("id", "parent_id", "level", "name", "code", "icon", "route", "component", "redirect",
                           "is_hidden", "type", "status", "sort", "created_by", "updated_by", "created_at", "updated_at",
                           "deleted_at", "remark")
VALUES
(<%.MenuId%> + 1, <%.MenuId%>, '<%.Level%><%.MenuId%>,', '<%.ChineseName%>列表', '<%.ModuleName%>:<%.VarName%>:index', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(<%.MenuId%> + 2, <%.MenuId%>, '<%.Level%><%.MenuId%>,', '<%.ChineseName%>回收站', '<%.ModuleName%>:<%.VarName%>:recycle', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(<%.MenuId%> + 3, <%.MenuId%>, '<%.Level%><%.MenuId%>,', '<%.ChineseName%>保存', '<%.ModuleName%>:<%.VarName%>:save', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(<%.MenuId%> + 4, <%.MenuId%>, '<%.Level%><%.MenuId%>,', '<%.ChineseName%>更新', '<%.ModuleName%>:<%.VarName%>:update', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(<%.MenuId%> + 5, <%.MenuId%>, '<%.Level%><%.MenuId%>,', '<%.ChineseName%>删除', '<%.ModuleName%>:<%.VarName%>:delete', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(<%.MenuId%> + 6, <%.MenuId%>, '<%.Level%><%.MenuId%>,', '<%.ChineseName%>读取', '<%.ModuleName%>:<%.VarName%>:read', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(<%.MenuId%> + 7, <%.MenuId%>, '<%.Level%><%.MenuId%>,', '<%.ChineseName%>恢复', '<%.ModuleName%>:<%.VarName%>:recovery', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(<%.MenuId%> + 8, <%.MenuId%>, '<%.Level%><%.MenuId%>,', '<%.ChineseName%>真实删除', '<%.ModuleName%>:<%.VarName%>:realDelete', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(<%.MenuId%> + 9, <%.MenuId%>, '<%.Level%><%.MenuId%>,', '<%.ChineseName%>导入', '<%.ModuleName%>:<%.VarName%>:import', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(<%.MenuId%> + 10, <%.MenuId%>, '<%.Level%><%.MenuId%>,', '<%.ChineseName%>导出', '<%.ModuleName%>:<%.VarName%>:export', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, ''),
(<%.MenuId%> + 11, <%.MenuId%>, '<%.Level%><%.MenuId%>,', '<%.ChineseName%>状态改变', '<%.ModuleName%>:<%.VarName%>:changeStatus', '', '', '', '',
 1, 'B', 1, 0, 0, 1, NOW(), NOW(), NULL, '');
