SET NAMES utf8mb4;
SET UNIQUE_CHECKS = 0;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for setting_config
-- ----------------------------
CREATE TABLE `setting_config` (
                                  `group_id` bigint(20) NOT NULL COMMENT '组id',
                                  `key` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '配置键名',
                                  `value` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '配置值',
                                  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '配置名称',
                                  `input_type` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '数据输入类型',
                                  `config_select_data` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '配置选项数据',
                                  `sort` smallint(5)  NOT NULL DEFAULT '0' COMMENT '排序',
                                  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                  PRIMARY KEY (`key`),
                                  KEY `setting_config_group_id_index` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='参数配置信息表';

-- ----------------------------
-- Table structure for setting_config_group
-- ----------------------------
CREATE TABLE `setting_config_group` (
                                        `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                        `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '配置组名称',
                                        `code` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '配置组标识',
                                        `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                                        `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                                        `created_at` timestamp NULL DEFAULT NULL,
                                        `updated_at` timestamp NULL DEFAULT NULL,
                                        `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='参数配置分组表';

-- ----------------------------
-- Table structure for setting_crontab
-- ----------------------------
CREATE TABLE `setting_crontab` (
                                   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                   `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '任务名称',
                                   `type` smallint(6) DEFAULT '4' COMMENT '任务类型 (1 command, 2 class, 3 url, 4 eval)',
                                   `target` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '调用任务字符串',
                                   `parameter` varchar(1000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '调用任务参数',
                                   `rule` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '任务执行表达式',
                                   `singleton` smallint(6) DEFAULT '1' COMMENT '是否单次执行 (1 是 2 不是)',
                                   `status` smallint(6) DEFAULT '1' COMMENT '状态 (1正常 2停用)',
                                   `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                                   `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                                   `created_at` timestamp NULL DEFAULT NULL,
                                   `updated_at` timestamp NULL DEFAULT NULL,
                                   `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='定时任务信息表';

-- ----------------------------
-- Table structure for setting_crontab_log
-- ----------------------------
CREATE TABLE `setting_crontab_log` (
                                       `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                       `crontab_id` bigint(20) NOT NULL COMMENT '任务ID',
                                       `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '任务名称',
                                       `target` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '任务调用目标字符串',
                                       `parameter` varchar(1000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '任务调用参数',
                                       `exception_info` varchar(2000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '异常信息',
                                       `status` smallint(6) DEFAULT '1' COMMENT '执行状态 (1成功 2失败)',
                                       `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                       PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=117 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='定时任务执行日志表';

-- ----------------------------
-- Table structure for setting_generate_columns
-- ----------------------------
CREATE TABLE `setting_generate_columns` (
                                            `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                            `table_id` bigint(20) NOT NULL COMMENT '所属表ID',
                                            `column_name` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字段名称',
                                            `column_comment` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字段注释',
                                            `column_type` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字段类型',
                                            `is_pk` smallint(6) DEFAULT '1' COMMENT '1 非主键 2 主键',
                                            `is_required` smallint(6) DEFAULT '1' COMMENT '1 非必填 2 必填',
                                            `is_insert` smallint(6) DEFAULT '1' COMMENT '1 非插入字段 2 插入字段',
                                            `is_edit` smallint(6) DEFAULT '1' COMMENT '1 非编辑字段 2 编辑字段',
                                            `is_list` smallint(6) DEFAULT '1' COMMENT '1 非列表显示字段 2 列表显示字段',
                                            `is_query` smallint(6) DEFAULT '1' COMMENT '1 非查询字段 2 查询字段',
                                            `is_sort` smallint(6) DEFAULT '1' COMMENT '1 不排序 2 排序字段',
                                            `query_type` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT 'eq' COMMENT '查询方式 eq 等于, neq 不等于, gt 大于, lt 小于, like 范围',
                                            `view_type` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT 'text' COMMENT '页面控件，text, textarea, password, select, checkbox, radio, date, upload, ma-upload（封装的上传控件）',
                                            `dict_type` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典类型',
                                            `allow_roles` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '允许查看该字段的角色',
                                            `options` varchar(1000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字段其他设置',
                                            `extra` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字段扩展信息',
                                            `sort` tinyint(3)  DEFAULT '0' COMMENT '排序',
                                            `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                                            `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                                            `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                            `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                                            `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='代码生成业务字段信息表';

-- ----------------------------
-- Table structure for setting_generate_tables
-- ----------------------------
CREATE TABLE `setting_generate_tables` (
                                           `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                           `table_name` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '表名称',
                                           `table_comment` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '表注释',
                                           `module_name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '所属模块',
                                           `namespace` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '命名空间',
                                           `menu_name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '生成菜单名',
                                           `belong_menu_id` bigint(20) DEFAULT NULL COMMENT '所属菜单',
                                           `package_name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'controller,api包名',
                                           `type` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '生成类型，single 单表CRUD，tree 树表CRUD，parent_sub父子表CRUD',
                                           `generate_type` smallint(6) DEFAULT '1' COMMENT '1 压缩包下载 2 生成到模块',
                                           `generate_menus` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '生成菜单列表',
                                           `build_menu` smallint(6) DEFAULT '1' COMMENT '是否构建菜单',
                                           `component_type` smallint(6) DEFAULT '1' COMMENT '组件显示方式',
                                           `options` varchar(1500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '其他业务选项',
                                           `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                                           `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                                           `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                           `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                                           `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                           `source` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'db连接群组',
                                           PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='代码生成业务信息表';

-- ----------------------------
-- Table structure for system_api
-- ----------------------------
CREATE TABLE `system_api` (
                              `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                              `group_id` bigint(20) NOT NULL COMMENT '接口组ID',
                              `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口名称',
                              `access_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口访问名称',
                              `auth_mode` smallint(6) NOT NULL DEFAULT '1' COMMENT '认证模式 (1简易 2复杂)',
                              `request_mode` char(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'A' COMMENT '请求模式 (A 所有 P POST G GET)',
                              `status` smallint(6) DEFAULT '1' COMMENT '状态 (1正常 2停用)',
                              `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                              `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                              `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                              `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                              `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                              `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                              PRIMARY KEY (`id`),
                              KEY `system_api_group_id_index` (`group_id`),
                              KEY `system_api_access_name_index` (`access_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='接口表';

-- ----------------------------
-- Table structure for system_api_group
-- ----------------------------
CREATE TABLE `system_api_group` (
                                    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                    `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口组名称',
                                    `status` smallint(6) DEFAULT '1' COMMENT '状态 (1正常 2停用)',
                                    `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                                    `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                                    `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                    `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                                    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                    `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='接口分组表';

-- ----------------------------
-- Table structure for system_api_log
-- ----------------------------
CREATE TABLE `system_api_log` (
                                  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                  `api_id` bigint(20) NOT NULL COMMENT 'api ID',
                                  `api_name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口名称',
                                  `access_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口访问名称',
                                  `request_data` text COLLATE utf8mb4_unicode_ci COMMENT '请求数据',
                                  `response_code` varchar(5) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '响应状态码',
                                  `response_data` LONGTEXT COLLATE utf8mb4_unicode_ci COMMENT '响应数据',
                                  `ip` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '访问IP地址',
                                  `ip_location` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'IP所属地',
                                  `access_time` timestamp NULL DEFAULT NULL COMMENT '访问时间',
                                  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                  PRIMARY KEY (`id`),
                                  KEY `system_api_log_api_id_index` (`api_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='接口日志表';

-- ----------------------------
-- Table structure for system_app
-- ----------------------------
CREATE TABLE `system_app` (
                              `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                              `group_id` bigint(20) NOT NULL COMMENT '应用组ID',
                              `app_name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用名称',
                              `app_id` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用ID',
                              `app_secret` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用密钥',
                              `status` smallint(6) DEFAULT '1' COMMENT '状态 (1正常 2停用)',
                              `description` text COLLATE utf8mb4_unicode_ci COMMENT '应用介绍',
                              `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                              `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                              `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                              `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                              `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                              `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                              PRIMARY KEY (`id`),
                              KEY `system_app_group_id_app_id_app_secret_index` (`group_id`,`app_id`,`app_secret`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='应用表';

-- ----------------------------
-- Table structure for system_app_api
-- ----------------------------
CREATE TABLE `system_app_api` (
                                  `app_id` bigint(20) NOT NULL COMMENT '应用ID',
                                  `api_id` bigint(20) NOT NULL COMMENT 'API—ID',
                                  PRIMARY KEY (`app_id`,`api_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='应用和api关联表';

-- ----------------------------
-- Table structure for system_app_group
-- ----------------------------
CREATE TABLE `system_app_group` (
                                    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                    `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用组名称',
                                    `status` smallint(6) DEFAULT '1' COMMENT '状态 (1正常 2停用)',
                                    `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                                    `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                                    `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                    `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                                    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                    `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='应用分组表';

-- ----------------------------
-- Table structure for system_dept
-- ----------------------------
CREATE TABLE `system_dept` (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                               `parent_id` bigint(20) NOT NULL COMMENT '父ID',
                               `level` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '组级集合',
                               `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '部门名称',
                               `leader` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '负责人',
                               `phone` varchar(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '联系电话',
                               `status` smallint(6) DEFAULT '1' COMMENT '状态 (1正常 2停用)',
                               `sort` smallint(5)  DEFAULT '0' COMMENT '排序',
                               `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                               `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                               `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                               `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                               `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                               `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                               PRIMARY KEY (`id`),
                               KEY `system_dept_parent_id_index` (`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='部门信息表';

-- ----------------------------
-- Table structure for system_dept_leader
-- ----------------------------
CREATE TABLE `system_dept_leader` (
                                      `dept_id` bigint(20) NOT NULL COMMENT '部门主键',
                                      `user_id` bigint(20) NOT NULL COMMENT '用户主键',
                                      `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
                                      `created_at` timestamp NOT NULL COMMENT '添加时间',
                                      PRIMARY KEY (`dept_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='部门领导表';

-- ----------------------------
-- Table structure for system_dict_data
-- ----------------------------
CREATE TABLE `system_dict_data` (
                                    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                    `type_id` bigint(20) NOT NULL COMMENT '字典类型ID',
                                    `label` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典标签',
                                    `value` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典值',
                                    `code` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典标示',
                                    `sort` smallint(5)  DEFAULT '0' COMMENT '排序',
                                    `status` smallint(6) DEFAULT '1' COMMENT '状态 (1正常 2停用)',
                                    `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                                    `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                                    `created_at` timestamp NULL DEFAULT NULL,
                                    `updated_at` timestamp NULL DEFAULT NULL,
                                    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                    `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                    PRIMARY KEY (`id`),
                                    KEY `system_dict_data_type_id_index` (`type_id`)
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='字典数据表';

-- ----------------------------
-- Table structure for system_dict_type
-- ----------------------------
CREATE TABLE `system_dict_type` (
                                    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                    `name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典名称',
                                    `code` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典标示',
                                    `status` smallint(6) DEFAULT '1' COMMENT '状态 (1正常 2停用)',
                                    `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                                    `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                                    `created_at` timestamp NULL DEFAULT NULL,
                                    `updated_at` timestamp NULL DEFAULT NULL,
                                    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                    `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='字典类型表';

-- ----------------------------
-- Table structure for system_login_log
-- ----------------------------
CREATE TABLE `system_login_log` (
                                    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                    `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
                                    `ip` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录IP地址',
                                    `ip_location` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'IP所属地',
                                    `os` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '操作系统',
                                    `browser` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '浏览器',
                                    `status` smallint(6) NOT NULL DEFAULT '1' COMMENT '登录状态 (1成功 2失败)',
                                    `message` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '提示消息',
                                    `login_time` timestamp NOT NULL COMMENT '登录时间',
                                    `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                    PRIMARY KEY (`id`),
                                    KEY `system_login_log_username_index` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='登录日志表';

-- ----------------------------
-- Table structure for system_menu
-- ----------------------------
CREATE TABLE `system_menu` (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                               `parent_id` bigint(20) NOT NULL COMMENT '父ID',
                               `level` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '组级集合',
                               `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单名称',
                               `code` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单标识代码',
                               `icon` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '菜单图标',
                               `route` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '路由地址',
                               `component` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '组件路径',
                               `redirect` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '跳转地址',
                               `is_hidden` smallint(6) NOT NULL DEFAULT '1' COMMENT '是否隐藏 (1是 2否)',
                               `type` char(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单类型, (M菜单 B按钮 L链接 I iframe)',
                               `status` smallint(6) DEFAULT '1' COMMENT '状态 (1正常 2停用)',
                               `sort` smallint(5)  DEFAULT '0' COMMENT '排序',
                               `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                               `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                               `created_at` timestamp NULL DEFAULT NULL,
                               `updated_at` timestamp NULL DEFAULT NULL,
                               `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                               `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4519 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单信息表';

-- ----------------------------
-- Table structure for system_modules
-- ----------------------------
CREATE TABLE `system_modules` (
                                  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                                  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                                  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                                  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '模块名称',
                                  `label` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '模块标记',
                                  `description` text COLLATE utf8mb4_unicode_ci COMMENT '描述',
                                  `installed` smallint(6) DEFAULT '0' COMMENT '是否安装1-否，2-是',
                                  `status` smallint(6) DEFAULT '0' COMMENT '状态 (1正常 2停用)',
                                  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='modules';

-- ----------------------------
-- Table structure for system_notice
-- ----------------------------
CREATE TABLE `system_notice` (
                                 `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                 `message_id` bigint(20) NOT NULL COMMENT '消息ID',
                                 `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',
                                 `type` smallint(6) NOT NULL COMMENT '公告类型（1通知 2公告）',
                                 `content` text COLLATE utf8mb4_unicode_ci COMMENT '公告内容',
                                 `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                                 `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                                 `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                 `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                                 `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                 `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                 `receive_users` text COLLATE utf8mb4_unicode_ci COMMENT '接收用户id,隔开',
                                 PRIMARY KEY (`id`),
                                 KEY `system_notice_message_id_index` (`message_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统公告表';

-- ----------------------------
-- Table structure for system_oper_log
-- ----------------------------
CREATE TABLE `system_oper_log` (
                                   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                   `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
                                   `method` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求方式',
                                   `router` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求路由',
                                   `service_name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '业务名称',
                                   `ip` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求IP地址',
                                   `ip_location` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'IP所属地',
                                   `request_data` text COLLATE utf8mb4_unicode_ci COMMENT '请求数据',
                                   `response_code` varchar(5) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '响应状态码',
                                   `response_data` LONGTEXT COLLATE utf8mb4_unicode_ci COMMENT '响应数据',
                                   `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                                   `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                                   `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                   `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                                   `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                   `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                   PRIMARY KEY (`id`),
                                   KEY `system_oper_log_username_index` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='操作日志表';

-- ----------------------------
-- Table structure for system_post
-- ----------------------------
CREATE TABLE `system_post` (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                               `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '岗位名称',
                               `code` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '岗位代码',
                               `sort` smallint(5)  DEFAULT '0' COMMENT '排序',
                               `status` smallint(6) DEFAULT '1' COMMENT '状态 (1正常 2停用)',
                               `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                               `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                               `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                               `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                               `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                               `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='岗位信息表';

-- ----------------------------
-- Table structure for system_queue_message
-- ----------------------------
CREATE TABLE `system_queue_message` (
                                        `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                        `content_type` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '内容类型',
                                        `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '消息标题',
                                        `send_by` bigint(20) DEFAULT NULL COMMENT '发送人',
                                        `content` longtext COLLATE utf8mb4_unicode_ci COMMENT '消息内容',
                                        `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                                        `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                                        `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                        `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                                        `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                        PRIMARY KEY (`id`),
                                        KEY `system_queue_message_content_type_index` (`content_type`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='队列消息表';

-- ----------------------------
-- Table structure for system_queue_message_receive
-- ----------------------------
CREATE TABLE `system_queue_message_receive` (
                                                `message_id` bigint(20) NOT NULL COMMENT '队列消息主键',
                                                `user_id` bigint(20) NOT NULL COMMENT '接收用户主键',
                                                `read_status` smallint(6) DEFAULT '1' COMMENT '已读状态 (1未读 2已读)',
                                                PRIMARY KEY (`message_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='队列消息接收人表';

-- ----------------------------
-- Table structure for system_role
-- ----------------------------
CREATE TABLE `system_role` (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                               `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
                               `code` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色代码',
                               `data_scope` smallint(6) DEFAULT '1' COMMENT '数据范围（1：全部数据权限 2：自定义数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：本人数据权限）',
                               `status` smallint(6) DEFAULT '1' COMMENT '状态 (1正常 2停用)',
                               `sort` smallint(5)  DEFAULT '0' COMMENT '排序',
                               `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                               `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                               `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                               `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                               `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                               `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色信息表';

-- ----------------------------
-- Table structure for system_role_dept
-- ----------------------------
CREATE TABLE `system_role_dept` (
                                    `role_id` bigint(20) NOT NULL COMMENT '角色主键',
                                    `dept_id` bigint(20) NOT NULL COMMENT '部门主键',
                                    PRIMARY KEY (`role_id`,`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色与部门关联表';

-- ----------------------------
-- Table structure for system_role_menu
-- ----------------------------
CREATE TABLE `system_role_menu` (
                                    `role_id` bigint(20) NOT NULL COMMENT '角色主键',
                                    `menu_id` bigint(20) NOT NULL COMMENT '菜单主键',
                                    PRIMARY KEY (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色与菜单关联表';

-- ----------------------------
-- Table structure for system_uploadfile
-- ----------------------------
CREATE TABLE `system_uploadfile` (
                                     `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                     `storage_mode` smallint(6) DEFAULT '1' COMMENT '存储模式 (1 本地 2 阿里云 3 七牛云 4 腾讯云)',
                                     `origin_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '原文件名',
                                     `object_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '新文件名',
                                     `hash` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件hash',
                                     `mime_type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '资源类型',
                                     `storage_path` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '存储目录',
                                     `suffix` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件后缀',
                                     `size_byte` bigint(20) DEFAULT NULL COMMENT '字节数',
                                     `size_info` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件大小',
                                     `url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'url地址',
                                     `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                                     `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                                     `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                                     `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                                     `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                     `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                                     PRIMARY KEY (`id`),
                                     UNIQUE KEY `system_uploadfile_hash_unique` (`hash`),
                                     KEY `system_uploadfile_storage_path_index` (`storage_path`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='上传文件信息表';

-- ----------------------------
-- Table structure for system_user
-- ----------------------------
CREATE TABLE `system_user` (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID，主键',
                               `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
                               `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
                               `user_type` varchar(3) COLLATE utf8mb4_unicode_ci DEFAULT '100' COMMENT '用户类型：(100系统用户)',
                               `nickname` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户昵称',
                               `phone` varchar(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机',
                               `email` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户邮箱',
                               `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户头像',
                               `signed` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '个人签名',
                               `dashboard` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '后台首页类型',
                               `status` smallint(6) DEFAULT '1' COMMENT '状态 (1正常 2停用)',
                               `login_ip` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '最后登陆IP',
                               `login_time` timestamp NULL DEFAULT NULL COMMENT '最后登陆时间',
                               `backend_setting` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '后台设置数据',
                               `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
                               `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
                               `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                               `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                               `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                               `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `system_user_username_unique` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户信息表';

-- ----------------------------
-- Table structure for system_user_dept
-- ----------------------------
CREATE TABLE `system_user_dept` (
                                    `user_id` bigint(20) NOT NULL COMMENT '用户主键',
                                    `dept_id` bigint(20) NOT NULL COMMENT '部门主键',
                                    PRIMARY KEY (`user_id`,`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户与部门关联表';

-- ----------------------------
-- Table structure for system_user_post
-- ----------------------------
CREATE TABLE `system_user_post` (
                                    `user_id` bigint(20) NOT NULL COMMENT '用户主键',
                                    `post_id` bigint(20) NOT NULL COMMENT '岗位主键',
                                    PRIMARY KEY (`user_id`,`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户与岗位关联表';

-- ----------------------------
-- Table structure for system_user_role
-- ----------------------------
CREATE TABLE `system_user_role` (
                                    `user_id` bigint(20) NOT NULL COMMENT '用户主键',
                                    `role_id` bigint(20) NOT NULL COMMENT '角色主键',
                                    PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户与角色关联表';

SET UNIQUE_CHECKS = 1;
SET FOREIGN_KEY_CHECKS = 1;
