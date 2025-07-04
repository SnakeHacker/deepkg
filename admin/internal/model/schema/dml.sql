USE deepkg_db;
SET character_set_client=utf8mb4;
SET character_set_connection=utf8mb4;
SET character_set_server=utf8mb4;
SET character_set_results=utf8mb4;

INSERT INTO `deepkg_db`.`organization` (`id`, `created_at`, `updated_at`, `deleted_at`, `org_name`)
VALUES (1, '2025-03-17 21:24:19.918', '2025-03-17 21:24:19.918', NULL, '中国移动上海产业研究院');

INSERT INTO `deepkg_db`.`user` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_code`, `org_id`, `account`, `username`, `password_hash`, `phone`, `mail`, `enable`, `role`, `avatar`)
VALUES (1, '2025-03-17 21:27:05.869', '2025-03-17 21:27:05.869', 0, '671dffb9-2646-49a1-b917-bcf53d97ceea', 1, 'admin', '超级管理员', 'pbkdf2:sha256:150000$IgdfguYC$bd08f26f7de930fd65360c4fd49b6b365714692ce8d34acbccac9a83dee1a245', '13919999999', 'admin@chinamobile.com', 1, 1, '');

INSERT INTO deepkg_db.document (created_at,updated_at,deleted_at,doc_name,doc_desc,doc_path,dir_id,creator_id) VALUES
	 ('2025-06-13 13:59:55.498','2025-06-13 13:59:55.498',NULL,'data.txt','test','http://127.0.0.1:9003/file/ed4c9cac700f14d4a191f69fee9c8136.txt',1,1);
INSERT INTO deepkg_db.document_dir (created_at,updated_at,deleted_at,dir_name,parent_id,sort_index,remark) VALUES
	 ('2025-06-13 13:42:17.524','2025-06-13 13:42:17.524',NULL,'测试',0,0,'测试创建文件夹');
INSERT INTO deepkg_db.knowledge_graph_workspace (id, Created_at, updated_at, deleted_at, work_space_name, creator_id) VALUES
        (1,'2025-06-13 14:24:50.491', '2025-06-13 14:24:50.491', 0, '测试工作空间', 1);
INSERT INTO deepkg_db.extract_task (id,created_at,updated_at,deleted_at,task_name,remark,work_space_id,task_status,published,creator_id) VALUES
	 (1,'2025-06-13 14:54:28.564','2025-06-13 14:54:28.564',NULL,'测试01','测试',1,1,0,1),
	 (2,'2025-06-13 15:51:16.243','2025-06-13 15:51:16.243',NULL,'测试02','',1,1,0,1);
INSERT INTO deepkg_db.extract_task_document (created_at,updated_at,deleted_at,task_id,doc_id) VALUES
	 ('2025-06-13 14:54:28.572','2025-06-13 14:54:28.572',NULL,1,1),
	 ('2025-06-13 15:51:16.250','2025-06-13 15:51:16.250',NULL,2,1);
INSERT INTO deepkg_db.extract_task_triple (created_at,updated_at,deleted_at,task_id,triple_id) VALUES
	 ('2025-06-13 14:54:28.584','2025-06-13 14:54:28.584',NULL,1,1),
	 ('2025-06-13 14:54:28.591','2025-06-13 14:54:28.591',NULL,1,2),
	 ('2025-06-13 15:51:16.265','2025-06-13 15:51:16.265',NULL,2,1),
	 ('2025-06-13 15:51:16.279','2025-06-13 15:51:16.279',NULL,2,2),
	 ('2025-06-13 15:51:16.283','2025-06-13 15:51:16.283',NULL,2,3);
INSERT INTO deepkg_db.schema_ontology (created_at,updated_at,deleted_at,work_space_id,ontology_name,ontology_desc,creator_id) VALUES
	 ('2025-06-13 14:24:50.491','2025-06-13 14:24:50.491',NULL,1,'行业板块','',1),
	 ('2025-06-13 14:26:02.550','2025-06-13 14:26:02.550',NULL,1,'国家','',1),
	 ('2025-06-13 14:26:16.381','2025-06-13 14:26:16.381',NULL,1,'公司','',1),
	 ('2025-06-13 14:26:59.516','2025-06-13 14:26:59.516',NULL,1,'国家机关','',1),
	 ('2025-06-13 14:36:19.137','2025-06-13 14:36:19.137',NULL,1,'CPI','消费者价格指数',1),
	 ('2025-06-13 14:36:25.277','2025-06-13 14:36:25.277',NULL,1,'PPI','生产者价格指数',1);
INSERT INTO deepkg_db.schema_ontology_prop (created_at,updated_at,deleted_at,work_space_id,ontology_id,prop_name,prop_desc,creator_id) VALUES
	 ('2025-06-13 15:43:06.551','2025-06-13 15:43:06.551',NULL,1,3,'公司邮箱','公司邮箱',1);
INSERT INTO deepkg_db.schema_triple (created_at,updated_at,deleted_at,work_space_id,source_ontology_id,target_ontology_id,relationship,creator_id) VALUES
	 ('2025-06-13 14:38:40.113','2025-06-13 14:38:40.113',NULL,1,5,2,'发布',1),
	 ('2025-06-13 14:38:43.444','2025-06-13 14:38:43.444',NULL,1,6,2,'发布',1),
	 ('2025-06-13 15:50:10.871','2025-06-13 15:50:10.871',NULL,1,3,2,'属于',1);
