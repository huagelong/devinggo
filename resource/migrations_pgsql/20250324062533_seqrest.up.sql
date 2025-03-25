-- setting_config_group
SELECT setval(
               'setting_config_group_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "setting_config_group"),
               false
       );

-- setting_crontab
SELECT setval(
               'setting_crontab_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "setting_crontab"),
               false
       );

-- setting_crontab_log
SELECT setval(
               'setting_crontab_log_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "setting_crontab_log"),
               false
       );

-- setting_generate_columns
SELECT setval(
               'setting_generate_columns_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "setting_generate_columns"),
               false
       );

-- setting_generate_tables
SELECT setval(
               'setting_generate_tables_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "setting_generate_tables"),
               false
       );

-- system_api
SELECT setval(
               'system_api_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_api"),
               false
       );

-- system_api_group
SELECT setval(
               'system_api_group_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_api_group"),
               false
       );

-- system_api_log
SELECT setval(
               'system_api_log_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_api_log"),
               false
       );

-- system_app
SELECT setval(
               'system_app_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_app"),
               false
       );


-- system_app_group
SELECT setval(
               'system_app_group_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_app_group"),
               false
       );

-- system_dept
SELECT setval(
               'system_dept_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_dept"),
               false
       );


-- system_dict_data
SELECT setval(
               'system_dict_data_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_dict_data"),
               false
       );

-- system_dict_type
SELECT setval(
               'system_dict_type_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_dict_type"),
               false
       );

-- system_login_log
SELECT setval(
               'system_login_log_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_login_log"),
               false
       );

-- system_menu
SELECT setval(
               'system_menu_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_menu"),
               false
       );

-- system_modules
SELECT setval(
               'system_modules_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_modules"),
               false
       );

-- system_notice
SELECT setval(
               'system_notice_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_notice"),
               false
       );

-- system_oper_log
SELECT setval(
               'system_oper_log_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_oper_log"),
               false
       );

-- system_post
SELECT setval(
               'system_post_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_post"),
               false
       );

-- system_queue_message
SELECT setval(
               'system_queue_message_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_queue_message"),
               false
       );


-- system_role
SELECT setval(
               'system_role_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_role"),
               false
       );


-- system_uploadfile
SELECT setval(
               'system_uploadfile_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_uploadfile"),
               false
       );

-- system_user
SELECT setval(
               'system_user_id_seq',
               (SELECT COALESCE(MAX(id), 0) + 1 FROM "system_user"),
               false
       );
