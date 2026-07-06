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
COMMENT ON TABLE {{.moduleName}}_example IS '{{.moduleNameCap}} module example table';
COMMENT ON COLUMN {{.moduleName}}_example.id IS 'ID';
COMMENT ON COLUMN {{.moduleName}}_example.name IS 'name';
COMMENT ON COLUMN {{.moduleName}}_example.description IS 'description';
COMMENT ON COLUMN {{.moduleName}}_example.status IS 'status 1=enabled,0=disabled';
COMMENT ON COLUMN {{.moduleName}}_example.created_at IS 'created_at';
COMMENT ON COLUMN {{.moduleName}}_example.updated_at IS 'updated_at';
COMMENT ON COLUMN {{.moduleName}}_example.deleted_at IS 'deleted_at';

-- Enable module
INSERT INTO system_modules(created_at, updated_at, created_by, updated_by, name, label, description, installed, status, deleted_at)
SELECT CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1, 1, '{{.moduleName}}', '{{.moduleNameCap}} module', '{{.moduleNameCap}} module created by generator', 1, 1, NULL
WHERE NOT EXISTS (
    SELECT 1 FROM system_modules WHERE name = '{{.moduleName}}' AND deleted_at IS NULL
);
