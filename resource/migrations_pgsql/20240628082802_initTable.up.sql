-- 关闭约束检查
SET session_replication_role = replica;

-- ----------------------------
-- setting_config
-- ----------------------------
CREATE TABLE setting_config (
    group_id BIGINT NOT NULL,
    key VARCHAR(32) NOT NULL,
    value VARCHAR(255),
    name VARCHAR(255),
    input_type VARCHAR(32),
    config_select_data VARCHAR(500),
    sort SMALLINT NOT NULL DEFAULT 0,
    remark VARCHAR(255),
    PRIMARY KEY (key)
);
COMMENT ON TABLE setting_config IS '参数配置信息表';
COMMENT ON COLUMN setting_config.group_id IS '组id';
COMMENT ON COLUMN setting_config.key IS '配置键名';
COMMENT ON COLUMN setting_config.value IS '配置值';
-- ...其他列注释类似...

-- ----------------------------
-- setting_config_group
-- ----------------------------
CREATE TABLE setting_config_group (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    code VARCHAR(64) NOT NULL,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    remark VARCHAR(255)
);
COMMENT ON TABLE setting_config_group IS '参数配置分组表';

-- ----------------------------
-- setting_crontab
-- ----------------------------
CREATE TABLE setting_crontab (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type SMALLINT DEFAULT 4,
    target VARCHAR(500),
    parameter jsonb,
    rule VARCHAR(32) NOT NULL,
    singleton SMALLINT DEFAULT 1,
    status SMALLINT DEFAULT 1,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    remark VARCHAR(255)
);
COMMENT ON TABLE setting_crontab IS '定时任务信息表';

-- ----------------------------
-- setting_crontab_log
-- ----------------------------
CREATE TABLE setting_crontab_log (
    id BIGSERIAL PRIMARY KEY,
    crontab_id BIGINT NOT NULL,
    name VARCHAR(255),
    target VARCHAR(500),
    parameter VARCHAR(1000),
    exception_info VARCHAR(2000),
    status SMALLINT DEFAULT 1,
    created_at TIMESTAMP
);
COMMENT ON TABLE setting_crontab_log IS '定时任务执行日志表';

-- ----------------------------
-- setting_generate_columns
-- ----------------------------
CREATE TABLE setting_generate_columns (
    id BIGSERIAL PRIMARY KEY,
    table_id BIGINT NOT NULL,
    column_name VARCHAR(200),
    column_comment VARCHAR(255),
    column_type VARCHAR(50),
    is_pk SMALLINT DEFAULT 1,
    is_required SMALLINT DEFAULT 1,
    is_insert SMALLINT DEFAULT 1,
    is_edit SMALLINT DEFAULT 1,
    is_list SMALLINT DEFAULT 1,
    is_query SMALLINT DEFAULT 1,
    is_sort SMALLINT DEFAULT 1,
    query_type VARCHAR(100) DEFAULT 'eq',
    view_type VARCHAR(100) DEFAULT 'text',
    dict_type VARCHAR(200),
    allow_roles VARCHAR(255),
    options VARCHAR(1000),
    extra VARCHAR(255),
    sort SMALLINT DEFAULT 0,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    remark VARCHAR(255)
);
COMMENT ON TABLE setting_generate_columns IS '代码生成业务字段信息表';

-- ----------------------------
-- setting_generate_tables
-- ----------------------------
CREATE TABLE setting_generate_tables (
    id BIGSERIAL PRIMARY KEY,
    table_name VARCHAR(200),
    table_comment VARCHAR(500),
    module_name VARCHAR(100),
    namespace VARCHAR(255),
    menu_name VARCHAR(100),
    belong_menu_id BIGINT,
    package_name VARCHAR(100),
    type VARCHAR(100),
    generate_type SMALLINT DEFAULT 1,
    generate_menus VARCHAR(255),
    build_menu SMALLINT DEFAULT 1,
    component_type SMALLINT DEFAULT 1,
    options VARCHAR(1500),
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    remark VARCHAR(255),
    source VARCHAR(100)
);
COMMENT ON TABLE setting_generate_tables IS '代码生成业务信息表';

-- ----------------------------
-- system_api
-- ----------------------------
CREATE TABLE system_api (
    id BIGSERIAL PRIMARY KEY,
    group_id BIGINT NOT NULL,
    name VARCHAR(32) NOT NULL,
    access_name VARCHAR(64) NOT NULL,
    auth_mode SMALLINT NOT NULL DEFAULT 1,
    request_mode CHAR(1) NOT NULL DEFAULT 'A',
    status SMALLINT DEFAULT 1,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255)
);
CREATE INDEX system_api_group_id_index ON system_api (group_id);
CREATE INDEX system_api_access_name_index ON system_api (access_name);
COMMENT ON TABLE system_api IS '接口表';

-- ----------------------------
-- system_api_group
-- ----------------------------
CREATE TABLE system_api_group (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    status SMALLINT DEFAULT 1,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255)
);
COMMENT ON TABLE system_api_group IS '接口分组表';

-- ----------------------------
-- system_api_log
-- ----------------------------
CREATE TABLE system_api_log (
    id BIGSERIAL PRIMARY KEY,
    api_id BIGINT NOT NULL,
    api_name VARCHAR(32) NOT NULL,
    access_name VARCHAR(64) NOT NULL,
    request_data TEXT,
    response_code VARCHAR(5),
    response_data TEXT,
    ip VARCHAR(45),
    ip_location VARCHAR(255),
    access_time TIMESTAMP,
    remark VARCHAR(255)
);
CREATE INDEX system_api_log_api_id_index ON system_api_log (api_id);
COMMENT ON TABLE system_api_log IS '接口日志表';

-- ----------------------------
-- system_app
-- ----------------------------
CREATE TABLE system_app (
    id BIGSERIAL PRIMARY KEY,
    group_id BIGINT NOT NULL,
    app_name VARCHAR(32) NOT NULL,
    app_id VARCHAR(16) NOT NULL,
    app_secret VARCHAR(128) NOT NULL,
    status SMALLINT DEFAULT 1,
    description TEXT,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255)
);
CREATE INDEX system_app_composite_idx ON system_app (group_id, app_id, app_secret);
COMMENT ON TABLE system_app IS '应用表';

-- ----------------------------
-- system_app_api
-- ----------------------------
CREATE TABLE system_app_api (
    app_id BIGINT NOT NULL,
    api_id BIGINT NOT NULL,
    PRIMARY KEY (app_id, api_id)
);
COMMENT ON TABLE system_app_api IS '应用和API关联表';

-- ----------------------------
-- system_app_group
-- ----------------------------
CREATE TABLE system_app_group (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    status SMALLINT DEFAULT 1,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255)
);
COMMENT ON TABLE system_app_group IS '应用分组表';

-- ----------------------------
-- system_dept
-- ----------------------------
CREATE TABLE system_dept (
    id BIGSERIAL PRIMARY KEY,
    parent_id BIGINT NOT NULL,
    level VARCHAR(500) NOT NULL,
    name VARCHAR(30) NOT NULL,
    leader VARCHAR(20),
    phone VARCHAR(11),
    status SMALLINT DEFAULT 1,
    sort SMALLINT DEFAULT 0,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255)
);
CREATE INDEX system_dept_parent_id_index ON system_dept (parent_id);
COMMENT ON TABLE system_dept IS '部门信息表';

-- ----------------------------
-- system_dept_leader
-- ----------------------------
CREATE TABLE system_dept_leader (
    dept_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    username VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY (dept_id, user_id)
);
COMMENT ON TABLE system_dept_leader IS '部门领导表';

-- ----------------------------
-- system_dict_data
-- ----------------------------
CREATE TABLE system_dict_data (
    id BIGSERIAL PRIMARY KEY,
    type_id BIGINT NOT NULL,
    label VARCHAR(50),
    value VARCHAR(100),
    code VARCHAR(100),
    sort SMALLINT DEFAULT 0,
    status SMALLINT DEFAULT 1,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255)
);
CREATE INDEX system_dict_data_type_id_index ON system_dict_data (type_id);
COMMENT ON TABLE system_dict_data IS '字典数据表';

-- ----------------------------
-- system_dict_type
-- ----------------------------
CREATE TABLE system_dict_type (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50),
    code VARCHAR(100),
    status SMALLINT DEFAULT 1,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255)
);
COMMENT ON TABLE system_dict_type IS '字典类型表';

-- ----------------------------
-- system_login_log
-- ----------------------------
CREATE TABLE system_login_log (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(20) NOT NULL,
    ip VARCHAR(45),
    ip_location VARCHAR(255),
    os VARCHAR(50),
    browser VARCHAR(50),
    status SMALLINT NOT NULL DEFAULT 1,
    message VARCHAR(50),
    login_time TIMESTAMP NOT NULL,
    remark VARCHAR(255)
);
CREATE INDEX system_login_log_username_index ON system_login_log (username);
COMMENT ON TABLE system_login_log IS '登录日志表';

-- ----------------------------
-- system_menu
-- ----------------------------
CREATE TABLE system_menu (
    id BIGSERIAL PRIMARY KEY,
    parent_id BIGINT NOT NULL,
    level VARCHAR(500) NOT NULL,
    name VARCHAR(50) NOT NULL,
    code VARCHAR(100) NOT NULL,
    icon VARCHAR(50),
    route VARCHAR(200),
    component VARCHAR(255),
    redirect VARCHAR(255),
    is_hidden SMALLINT NOT NULL DEFAULT 1,
    type CHAR(1) NOT NULL DEFAULT '',
    status SMALLINT DEFAULT 1,
    sort SMALLINT DEFAULT 0,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255)
);
COMMENT ON TABLE system_menu IS '菜单信息表';

-- ----------------------------
-- system_modules
-- ----------------------------
CREATE TABLE system_modules (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    created_by BIGINT,
    updated_by BIGINT,
    name VARCHAR(100),
    label VARCHAR(100),
    description TEXT,
    installed SMALLINT DEFAULT 0,
    status SMALLINT DEFAULT 0,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE system_modules IS 'modules';

-- ----------------------------
-- system_notice
-- ----------------------------
CREATE TABLE system_notice (
    id BIGSERIAL PRIMARY KEY,
    message_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    type SMALLINT NOT NULL,
    content TEXT,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255),
    receive_users TEXT
);
CREATE INDEX system_notice_message_id_index ON system_notice (message_id);
COMMENT ON TABLE system_notice IS '系统公告表';

-- ----------------------------
-- system_oper_log
-- ----------------------------
CREATE TABLE system_oper_log (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(20) NOT NULL,
    method VARCHAR(20) NOT NULL,
    router VARCHAR(500) NOT NULL,
    service_name VARCHAR(30) NOT NULL,
    ip VARCHAR(45),
    ip_location VARCHAR(255),
    request_data TEXT,
    response_code VARCHAR(5),
    response_data TEXT,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255)
);
CREATE INDEX system_oper_log_username_index ON system_oper_log (username);
COMMENT ON TABLE system_oper_log IS '操作日志表';

-- ----------------------------
-- system_post
-- ----------------------------
CREATE TABLE system_post (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    code VARCHAR(100) NOT NULL,
    sort SMALLINT DEFAULT 0,
    status SMALLINT DEFAULT 1,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255)
);
COMMENT ON TABLE system_post IS '岗位信息表';

-- ----------------------------
-- system_queue_message
-- ----------------------------
CREATE TABLE system_queue_message (
    id BIGSERIAL PRIMARY KEY,
    content_type VARCHAR(64),
    title VARCHAR(255),
    send_by BIGINT,
    content TEXT,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    remark VARCHAR(255)
);
CREATE INDEX system_queue_message_content_type_idx ON system_queue_message (content_type);
COMMENT ON TABLE system_queue_message IS '队列消息表';

-- ----------------------------
-- system_queue_message_receive
-- ----------------------------
CREATE TABLE system_queue_message_receive (
    message_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    read_status SMALLINT DEFAULT 1,
    PRIMARY KEY (message_id, user_id))
;
COMMENT ON TABLE system_queue_message_receive IS '队列消息接收人表';

-- ----------------------------
-- system_role
-- ----------------------------
CREATE TABLE system_role (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    code VARCHAR(100) NOT NULL,
    data_scope SMALLINT DEFAULT 1,
    status SMALLINT DEFAULT 1,
    sort SMALLINT DEFAULT 0,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255)
);
COMMENT ON TABLE system_role IS '角色信息表';

-- ----------------------------
-- system_role_dept
-- ----------------------------
CREATE TABLE system_role_dept (
    role_id BIGINT NOT NULL,
    dept_id BIGINT NOT NULL,
    PRIMARY KEY (role_id, dept_id))
;
COMMENT ON TABLE system_role_dept IS '角色与部门关联表';

-- ----------------------------
-- system_role_menu
-- ----------------------------
CREATE TABLE system_role_menu (
    role_id BIGINT NOT NULL,
    menu_id BIGINT NOT NULL,
    PRIMARY KEY (role_id, menu_id))
;
COMMENT ON TABLE system_role_menu IS '角色与菜单关联表';

-- ----------------------------
-- system_uploadfile
-- ----------------------------
CREATE TABLE system_uploadfile (
    id BIGSERIAL PRIMARY KEY,
    storage_mode SMALLINT DEFAULT 1,
    origin_name VARCHAR(255),
    object_name VARCHAR(50),
    hash VARCHAR(64),
    mime_type VARCHAR(255),
    storage_path VARCHAR(100),
    suffix VARCHAR(10),
    size_byte BIGINT,
    size_info VARCHAR(50),
    url VARCHAR(255),
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255)
);
CREATE UNIQUE INDEX system_uploadfile_hash_unique ON system_uploadfile (hash);
CREATE INDEX system_uploadfile_storage_path_idx ON system_uploadfile (storage_path);
COMMENT ON TABLE system_uploadfile IS '上传文件信息表';

-- ----------------------------
-- system_user
-- ----------------------------
CREATE TABLE "system_user" (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(20) NOT NULL UNIQUE,
    "password" VARCHAR(100) NOT NULL,
    user_type VARCHAR(3) DEFAULT '100',
    nickname VARCHAR(30),
    phone VARCHAR(11),
    email VARCHAR(50),
    avatar VARCHAR(255),
    signed VARCHAR(255),
    dashboard VARCHAR(100),
    status SMALLINT DEFAULT 1,
    login_ip VARCHAR(45),
    login_time TIMESTAMP,
    backend_setting jsonb,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    remark VARCHAR(255)
);
COMMENT ON TABLE "system_user" IS '用户信息表';

-- ----------------------------
-- system_user_dept
-- ----------------------------
CREATE TABLE system_user_dept (
    user_id BIGINT NOT NULL,
    dept_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, dept_id)
);
COMMENT ON TABLE system_user_dept IS '用户与部门关联表';

-- ----------------------------
-- system_user_post
-- ----------------------------
CREATE TABLE system_user_post (
    user_id BIGINT NOT NULL,
    post_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, post_id)
);
COMMENT ON TABLE system_user_post IS '用户与岗位关联表';

-- ----------------------------
-- system_user_role
-- ----------------------------
CREATE TABLE system_user_role (
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, role_id)
);
COMMENT ON TABLE system_user_role IS '用户与角色关联表';

-- 重新启用约束检查
SET session_replication_role = DEFAULT;