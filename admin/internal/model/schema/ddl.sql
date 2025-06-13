-- deepkg_db.document definition

CREATE TABLE `document` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `doc_name` varchar(255) NOT NULL COMMENT '文件名',
  `doc_desc` text NOT NULL COMMENT '文件描述',
  `doc_path` varchar(255) NOT NULL COMMENT '文件路径',
  `dir_id` int(11) NOT NULL COMMENT '文件目录ID',
  `creator_id` int(11) NOT NULL COMMENT '创建者ID',
  PRIMARY KEY (`id`),
  KEY `idx_document_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- deepkg_db.document_dir definition

CREATE TABLE `document_dir` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `dir_name` varchar(255) NOT NULL COMMENT '目录名',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父目录ID',
  `sort_index` int(11) NOT NULL DEFAULT '0' COMMENT '排序索引',
  `remark` text NOT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_document_dir_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- deepkg_db.knowledge_graph_workspace definition

CREATE TABLE `knowledge_graph_workspace` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `work_space_name` varchar(255) NOT NULL COMMENT '知识库名称',
  `creator_id` int(11) NOT NULL COMMENT '创建者ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_knowledge_graph_workspace_work_space_name` (`work_space_name`),
  KEY `idx_knowledge_graph_workspace_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- deepkg_db.organization definition

CREATE TABLE `organization` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `org_name` varchar(255) NOT NULL COMMENT '组织名称',
  PRIMARY KEY (`id`),
  KEY `idx_organization_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- deepkg_db.schema_ontology definition

CREATE TABLE `schema_ontology` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `work_space_id` int(11) NOT NULL COMMENT '工作空间ID',
  `ontology_name` varchar(255) NOT NULL COMMENT '实体名称',
  `ontology_desc` text NOT NULL COMMENT '实体描述',
  `creator_id` int(11) NOT NULL COMMENT '创建者ID',
  PRIMARY KEY (`id`),
  KEY `idx_schema_ontology_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- deepkg_db.schema_ontology_prop definition

CREATE TABLE `schema_ontology_prop` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `work_space_id` int(11) NOT NULL COMMENT '工作空间ID',
  `ontology_id` int(11) NOT NULL COMMENT '实体ID',
  `prop_name` varchar(255) NOT NULL COMMENT '属性名称',
  `prop_desc` text NOT NULL COMMENT '属性描述',
  `creator_id` int(11) NOT NULL COMMENT '创建者ID',
  PRIMARY KEY (`id`),
  KEY `idx_schema_ontology_prop_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- deepkg_db.schema_triple definition

CREATE TABLE `schema_triple` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `work_space_id` int(11) NOT NULL COMMENT '工作空间ID',
  `source_ontology_id` int(11) NOT NULL COMMENT '源实体ID',
  `target_ontology_id` int(11) NOT NULL COMMENT '目标实体ID',
  `relationship` varchar(255) NOT NULL COMMENT '实体关系',
  `creator_id` int(11) NOT NULL COMMENT '创建者ID',
  PRIMARY KEY (`id`),
  KEY `idx_schema_triple_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- deepkg_db.`user` definition

CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_code` varchar(255) NOT NULL COMMENT '用户编码',
  `org_id` int(11) NOT NULL DEFAULT '0' COMMENT '组织ID',
  `account` varchar(255) NOT NULL COMMENT '账号用于登录',
  `username` varchar(255) NOT NULL COMMENT '用户名称',
  `password_hash` varchar(255) NOT NULL COMMENT '密码',
  `phone` varchar(255) NOT NULL DEFAULT '' COMMENT '电话',
  `mail` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `enable` tinyint(1) NOT NULL COMMENT '启用状态：1-启用，2-禁用',
  `role` tinyint(1) NOT NULL DEFAULT '0' COMMENT '角色：2-普通用户，1-管理员',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_user_user_code` (`user_code`),
  UNIQUE KEY `uni_user_account` (`account`),
  KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;