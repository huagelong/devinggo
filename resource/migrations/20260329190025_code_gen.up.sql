-- ----------------------------
-- code_gen_tables - 代码生成主表
-- ----------------------------
CREATE TABLE code_gen_tables (
    id BIGSERIAL PRIMARY KEY,
    table_name VARCHAR(100) NOT NULL,
    table_comment VARCHAR(500) NOT NULL DEFAULT '',
    remark VARCHAR(500),
    module_name VARCHAR(50),
    belong_menu_id BIGINT,
    type VARCHAR(20) NOT NULL DEFAULT 'single',
    menu_name VARCHAR(100),
    component_type SMALLINT NOT NULL DEFAULT 1,
    tpl_type VARCHAR(50) NOT NULL DEFAULT 'default',
    tree_id VARCHAR(50),
    tree_parent_id VARCHAR(50),
    tree_name VARCHAR(50),
    tag_id VARCHAR(100),
    tag_name VARCHAR(100),
    tag_view_name VARCHAR(50),
    generate_menus VARCHAR(500),
    options JSONB,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by BIGINT,
    updated_by BIGINT,
    status SMALLINT NOT NULL DEFAULT 1,
    sort SMALLINT NOT NULL DEFAULT 0,
    UNIQUE (table_name)
);
COMMENT ON TABLE code_gen_tables IS '代码生成主表';
COMMENT ON COLUMN code_gen_tables.table_name IS '表名称';
COMMENT ON COLUMN code_gen_tables.table_comment IS '表描述';
COMMENT ON COLUMN code_gen_tables.remark IS '备注信息';
COMMENT ON COLUMN code_gen_tables.module_name IS '所属模块';
COMMENT ON COLUMN code_gen_tables.belong_menu_id IS '所属菜单ID';
COMMENT ON COLUMN code_gen_tables.type IS '生成类型: single=单表, tree=树表';
COMMENT ON COLUMN code_gen_tables.menu_name IS '菜单名称';
COMMENT ON COLUMN code_gen_tables.component_type IS '组件类型: 1=模态框, 2=抽屉, 3=Tag页';
COMMENT ON COLUMN code_gen_tables.tpl_type IS '模板类型: default';
COMMENT ON COLUMN code_gen_tables.tree_id IS '树表主ID字段';
COMMENT ON COLUMN code_gen_tables.tree_parent_id IS '树表父ID字段';
COMMENT ON COLUMN code_gen_tables.tree_name IS '树表显示名称字段';
COMMENT ON COLUMN code_gen_tables.tag_id IS 'Tag页ID';
COMMENT ON COLUMN code_gen_tables.tag_name IS 'Tag页名称';
COMMENT ON COLUMN code_gen_tables.tag_view_name IS 'Tag页显示字段';
COMMENT ON COLUMN code_gen_tables.generate_menus IS '生成的菜单按钮';
COMMENT ON COLUMN code_gen_tables.options IS '扩展配置';
COMMENT ON COLUMN code_gen_tables.created_at IS '创建时间';
COMMENT ON COLUMN code_gen_tables.updated_at IS '更新时间';
COMMENT ON COLUMN code_gen_tables.deleted_at IS '删除时间';
COMMENT ON COLUMN code_gen_tables.created_by IS '创建者ID';
COMMENT ON COLUMN code_gen_tables.updated_by IS '更新者ID';
COMMENT ON COLUMN code_gen_tables.status IS '状态: 1=正常, 0=停用';
COMMENT ON COLUMN code_gen_tables.sort IS '排序';

-- ----------------------------
-- code_gen_fields - 字段配置表
-- ----------------------------
CREATE TABLE code_gen_fields (
    id BIGSERIAL PRIMARY KEY,
    table_id BIGINT NOT NULL,
    column_name VARCHAR(100) NOT NULL,
    column_comment VARCHAR(500) NOT NULL DEFAULT '',
    column_type VARCHAR(100) NOT NULL,
    data_type VARCHAR(50),
    is_nullable VARCHAR(10),
    sort SMALLINT NOT NULL DEFAULT 0,
    is_required SMALLINT NOT NULL DEFAULT 0,
    is_insert SMALLINT NOT NULL DEFAULT 0,
    is_edit SMALLINT NOT NULL DEFAULT 0,
    is_list SMALLINT NOT NULL DEFAULT 0,
    is_query SMALLINT NOT NULL DEFAULT 0,
    is_sort SMALLINT NOT NULL DEFAULT 0,
    query_type VARCHAR(20),
    view_type VARCHAR(50) NOT NULL DEFAULT 'text',
    dict_type VARCHAR(100),
    allow_roles VARCHAR(500),
    options JSONB,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    UNIQUE (table_id, column_name)
);
COMMENT ON TABLE code_gen_fields IS '字段配置表';
COMMENT ON COLUMN code_gen_fields.table_id IS '所属表ID';
COMMENT ON COLUMN code_gen_fields.column_name IS '字段名称';
COMMENT ON COLUMN code_gen_fields.column_comment IS '字段描述';
COMMENT ON COLUMN code_gen_fields.column_type IS '物理类型';
COMMENT ON COLUMN code_gen_fields.data_type IS '数据类型';
COMMENT ON COLUMN code_gen_fields.is_nullable IS '是否可空';
COMMENT ON COLUMN code_gen_fields.sort IS '排序';
COMMENT ON COLUMN code_gen_fields.is_required IS '必填: 1=是, 2=否';
COMMENT ON COLUMN code_gen_fields.is_insert IS '新增: 1=是, 2=否';
COMMENT ON COLUMN code_gen_fields.is_edit IS '编辑: 1=是, 2=否';
COMMENT ON COLUMN code_gen_fields.is_list IS '列表显示: 1=是, 2=否';
COMMENT ON COLUMN code_gen_fields.is_query IS '查询: 1=是, 2=否';
COMMENT ON COLUMN code_gen_fields.is_sort IS '排序: 1=是, 2=否';
COMMENT ON COLUMN code_gen_fields.query_type IS '查询方式';
COMMENT ON COLUMN code_gen_fields.view_type IS '页面控件';
COMMENT ON COLUMN code_gen_fields.dict_type IS '数据字典';
COMMENT ON COLUMN code_gen_fields.allow_roles IS '允许角色';
COMMENT ON COLUMN code_gen_fields.options IS '组件扩展配置';
COMMENT ON COLUMN code_gen_fields.created_at IS '创建时间';
COMMENT ON COLUMN code_gen_fields.updated_at IS '更新时间';

CREATE INDEX code_gen_fields_table_id_index ON code_gen_fields (table_id);

-- ----------------------------
-- code_gen_buttons - 按钮权限表
-- ----------------------------
CREATE TABLE code_gen_buttons (
    id BIGSERIAL PRIMARY KEY,
    table_id BIGINT NOT NULL,
    button_code VARCHAR(50) NOT NULL,
    button_name VARCHAR(50) NOT NULL,
    button_comment VARCHAR(200),
    is_show SMALLINT NOT NULL DEFAULT 1,
    sort SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    UNIQUE (table_id, button_code)
);
COMMENT ON TABLE code_gen_buttons IS '按钮权限表';
COMMENT ON COLUMN code_gen_buttons.table_id IS '所属表ID';
COMMENT ON COLUMN code_gen_buttons.button_code IS '按钮编码';
COMMENT ON COLUMN code_gen_buttons.button_name IS '按钮名称';
COMMENT ON COLUMN code_gen_buttons.button_comment IS '按钮描述';
COMMENT ON COLUMN code_gen_buttons.is_show IS '是否显示: 1=显示, 2=隐藏';
COMMENT ON COLUMN code_gen_buttons.sort IS '排序';
COMMENT ON COLUMN code_gen_buttons.created_at IS '创建时间';
COMMENT ON COLUMN code_gen_buttons.updated_at IS '更新时间';

CREATE INDEX code_gen_buttons_table_id_index ON code_gen_buttons (table_id);
