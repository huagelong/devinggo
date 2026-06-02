-- Create {{.moduleName}} module tables
-- Author: devinggo
-- Date: {{.date}}

-- Example table for {{.moduleName}} module
CREATE TABLE IF NOT EXISTS {{.moduleName}}_example (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL DEFAULT '',
    description TEXT,
    status SMALLINT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Create index
CREATE INDEX idx_{{.moduleName}}_example_status ON {{.moduleName}}_example(status);
CREATE INDEX idx_{{.moduleName}}_example_deleted_at ON {{.moduleName}}_example(deleted_at);

-- Add comment
COMMENT ON TABLE {{.moduleName}}_example IS '{{.moduleNameCap}}模块示例表';
COMMENT ON COLUMN {{.moduleName}}_example.id IS 'ID';
COMMENT ON COLUMN {{.moduleName}}_example.name IS '名称';
COMMENT ON COLUMN {{.moduleName}}_example.description IS '描述';
COMMENT ON COLUMN {{.moduleName}}_example.status IS '状态:1=启用,0=禁用';
COMMENT ON COLUMN {{.moduleName}}_example.created_at IS '创建时间';
COMMENT ON COLUMN {{.moduleName}}_example.updated_at IS '更新时间';
COMMENT ON COLUMN {{.moduleName}}_example.deleted_at IS '删除时间';
