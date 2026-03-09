ALTER TABLE public.setting_generate_tables ADD tpl_type varchar DEFAULT 'default' NULL;
COMMENT ON COLUMN public.setting_generate_tables.tpl_type IS 'Vue模板类型: default(Arco Design) / ruoyi(RuoYi)';
