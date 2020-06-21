-- --------------------------------------------------------
-- 主机:                           F:\project\python\django-workflow\one-workflow\backend\core.db
-- 服务器版本:                        3.30.1
-- 服务器操作系统:                      
-- HeidiSQL 版本:                  11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES  */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 main 的数据库结构
--CREATE DATABASE IF NOT EXISTS "main";
--USE "main";
;
DROP TABLE IF EXISTS `workflows_transition`;
DROP TABLE IF EXISTS `workflows_state_user_participant`;
DROP TABLE IF EXISTS `workflows_state_role_participant`;
DROP TABLE IF EXISTS `workflows_state_group_participant`;
DROP TABLE IF EXISTS `workflows_state_fields`;
DROP TABLE IF EXISTS `workflows_state`;
DROP TABLE IF EXISTS `workflows_customfield`;
DROP TABLE IF EXISTS `workflows_workflow`;
DROP TABLE IF EXISTS `workflows_workflowtype`;

CREATE TABLE IF NOT EXISTS `workflows_workflowtype` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT "主键",
	`create_time` DATETIME NOT NULL,
	`update_time` DATETIME NOT NULL,
	`memo` TEXT NOT NULL,
	`name` VARCHAR(50) NOT NULL COLLATE 'utf8mb4_general_ci',
	`code` VARCHAR(32) NOT NULL COLLATE 'utf8mb4_general_ci',
	`order_id` INT(11) NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	UNIQUE INDEX `uniq_workflows_workflowtype_code` (`code`)
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;

CREATE TABLE  IF NOT EXISTS `workflows_workflow` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT "主键",
	`create_time` DATETIME NOT NULL,
	`update_time` DATETIME NOT NULL,
	`memo` TEXT NOT NULL,
	`name` VARCHAR(50) NOT NULL COLLATE 'utf8mb4_general_ci',
	`ticket_sn_prefix` VARCHAR(20) NOT NULL COLLATE 'utf8mb4_general_ci',
	`status` TINYINT NOT NULL,
	`view_permission_check` TINYINT NOT NULL,
	`limit_expression` TEXT NOT NULL,
	`display_form_str` TEXT NOT NULL,
	`title_template` VARCHAR(50) NULL COLLATE 'utf8mb4_general_ci',
	`type_id` BIGINT NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	INDEX `workflows_workflow_type_id` (`type_id`),
	CONSTRAINT `FK_workflows_workflow_fields_type_id_workflows_workflowtype` FOREIGN KEY (`type_id`) REFERENCES `workflows_workflowtype` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;

CREATE TABLE IF NOT EXISTS  `workflows_customfield` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT "主键",
	`create_time` DATETIME NOT NULL,
	`update_time` DATETIME NOT NULL,
	`memo` TEXT NOT NULL,
	`field_attribute` TINYINT NOT NULL,
	`field_type` VARCHAR(1) NOT NULL  COLLATE 'utf8mb4_general_ci',
	`field_key` VARCHAR(50) NOT NULL COLLATE 'utf8mb4_general_ci',
	`field_name` VARCHAR(50) NOT NULL COLLATE 'utf8mb4_general_ci',
	`order_id` INT(11) NOT NULL,
	`default_value` VARCHAR(100) NULL COLLATE 'utf8mb4_general_ci',
	`field_template` TEXT NOT NULL COLLATE 'utf8mb4_general_ci',
	`boolean_field_display` VARCHAR(100) NOT NULL COLLATE 'utf8mb4_general_ci',
	`field_choice` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_general_ci',
	`label` VARCHAR(100) NOT NULL COLLATE 'utf8mb4_general_ci',
	`workflow_id` BIGINT NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	INDEX `workflows_customfield_workflow_id` (`workflow_id`),
	CONSTRAINT `FK_workflows_customfield_workflow_id_workflows_workflow` FOREIGN KEY (`workflow_id`) REFERENCES `workflows_workflow` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;

CREATE TABLE IF NOT EXISTS  `workflows_state` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT "主键",
	`create_time` DATETIME NOT NULL,
	`update_time` DATETIME NOT NULL,
	`memo` TEXT NOT NULL COLLATE 'utf8mb4_general_ci',
	`name` VARCHAR(50) NOT NULL COLLATE 'utf8mb4_general_ci',
	`is_hidden` TINYINT NOT NULL,
	`order_id` INT(11) NOT NULL,
	`state_type` VARCHAR(1) NOT NULL COLLATE 'utf8mb4_general_ci',
	`enable_retreat` TINYINT NOT NULL,
	`participant_type` VARCHAR(1) NOT NULL COLLATE 'utf8mb4_general_ci',
	`workflow_id` BIGINT NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	INDEX `workflows_state_workflow_id` (`workflow_id`),
	CONSTRAINT `FK_workflows_state_workflows_workflow` FOREIGN KEY (`workflow_id`) REFERENCES `workflows_workflow` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;


CREATE TABLE IF NOT EXISTS  `workflows_state_fields` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT "主键",
	`state_id` BIGINT NOT NULL,
	`customfield_id` BIGINT NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	INDEX `workflows_state_fields_customfield_id` (`customfield_id`),
	INDEX `workflows_state_fields_state_id` (`state_id`),
	UNIQUE INDEX `workflows_state_fields_state_id_customfield_id_uniq` (`state_id`, `customfield_id`),
	CONSTRAINT `FK_workflows_state_fields_customfield_id_workflows_customfield` FOREIGN KEY (`customfield_id`) REFERENCES `workflows_customfield` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_workflows_state_fields_state_id_workflows_state` FOREIGN KEY (`state_id`) REFERENCES `workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;



CREATE TABLE IF NOT EXISTS  `workflows_state_group_participant` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT "主键",
	`state_id` BIGINT NOT NULL,
	`group_id` INT(11) NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	INDEX `workflows_state_group_participant_group_id` (`group_id`),
	INDEX `workflows_state_group_participant_state_id` (`state_id`),
	UNIQUE INDEX `workflows_state_group_participant_state_id_group_id_uniq` (`state_id`, `group_id`),
	CONSTRAINT `FK_workflows_state_group_participant_group_id` FOREIGN KEY (`group_id`) REFERENCES `sys_dept` (`dept_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_workflows_state_group_participant_state_id` FOREIGN KEY (`state_id`) REFERENCES `workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;

CREATE TABLE IF NOT EXISTS  `workflows_state_role_participant` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT "主键",
	`state_id` BIGINT NOT NULL,
	`role_id` INT(11) NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	INDEX `workflows_state_role_participant_role_id` (`role_id`),
	INDEX `workflows_state_role_participant_state_id` (`state_id`),
	UNIQUE INDEX `workflows_state_role_participant_state_id_role_id_uniq` (`state_id`, `role_id`),
	CONSTRAINT `FK_workflows_state_role_participant_role_id` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`role_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_workflows_state_role_participant_state_id` FOREIGN KEY (`state_id`) REFERENCES `workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;

CREATE TABLE IF NOT EXISTS  `workflows_state_user_participant` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT "主键",
	`state_id` BIGINT NOT NULL,
	`user_id` INT(11) NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	INDEX `workflows_state_user_participant_user_id` (`user_id`),
	INDEX `workflows_state_user_participant_state_id` (`state_id`),
	UNIQUE INDEX `workflows_state_user_participant_state_id_user_id_uniq` (`state_id`, `user_id`),
	CONSTRAINT `FK_workflows_state_user_participant_user_id` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`user_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_workflows_state_user_participant_state_id` FOREIGN KEY (`state_id`) REFERENCES `workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;

CREATE TABLE IF NOT EXISTS  `workflows_transition` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT "主键",
	`create_time` DATETIME NOT NULL,
	`update_time` DATETIME NOT NULL,
	`memo` TEXT NOT NULL,
	`name` VARCHAR(1) NOT NULL COLLATE 'utf8mb4_general_ci',
	`transition_type` VARCHAR(1) NOT NULL,
	`timer` INT(11) NOT NULL,
	`condition_expression` TEXT NOT NULL,
	`attribute_type` VARCHAR(1) NOT NULL COLLATE 'utf8mb4_general_ci',
	`alert_enable` TINYINT NOT NULL,
	`alert_text` VARCHAR(100) NOT NULL COLLATE 'utf8mb4_general_ci',
	`dest_state_id` BIGINT NULL,
	`source_state_id` BIGINT NULL,
	`workflow_id` BIGINT NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	INDEX `workflows_transition_workflow_id` (`workflow_id`),
	INDEX `workflows_transition_source_state_id` (`source_state_id`),
	INDEX `workflows_transition_dest_state_id` (`dest_state_id`),
	CONSTRAINT `FK_workflows_transition_workflow_id` FOREIGN KEY (`workflow_id`) REFERENCES `workflows_workflow` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_workflows_transition_source_state_id` FOREIGN KEY (`source_state_id`) REFERENCES `workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_workflows_transition_dest_state_id` FOREIGN KEY (`dest_state_id`) REFERENCES `workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;



-- 导出  表 main.workflows_workflowtype 结构
-- CREATE TABLE IF NOT EXISTS "workflows_workflowtype" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "create_time" datetime NOT NULL, "update_time" datetime NOT NULL, "memo" text NOT NULL, "name" varchar(50) NOT NULL, "code" varchar(32) NOT NULL UNIQUE, "order_id" integer NOT NULL);

-- 正在导出表  main.workflows_workflowtype 的数据：-1 rows
/*!40000 ALTER TABLE "workflows_workflowtype" DISABLE KEYS */;
INSERT INTO `workflows_workflowtype` (`id`, `create_time`, `update_time`, `memo`, `name`, `code`, `order_id`) VALUES
	(1, '2020-06-13 06:08:00.378593', '2020-06-14 10:54:40.475119', 'mark', '行政', 'ad', 1);
INSERT INTO `workflows_workflowtype` (`id`, `create_time`, `update_time`, `memo`, `name`, `code`, `order_id`) VALUES
	(2, '2020-06-13 06:08:05.546763', '2020-06-14 11:06:14.757610', 'mark', '技术', 'it', 2);
INSERT INTO `workflows_workflowtype` (`id`, `create_time`, `update_time`, `memo`, `name`, `code`, `order_id`) VALUES
	(3, '2020-06-13 14:53:07.394805', '2020-06-13 14:53:07.394805', '公共事物中心', '盖章签呈', 'gzqc', 3);
/*!40000 ALTER TABLE "workflows_workflowtype" ENABLE KEYS */;

-- 导出  表 main.workflows_workflow 结构
-- CREATE TABLE IF NOT EXISTS "workflows_workflow" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "create_time" datetime NOT NULL, "update_time" datetime NOT NULL, "memo" text NOT NULL, "name" varchar(50) NOT NULL, "ticket_sn_prefix" varchar(20) NOT NULL, "status" bool NOT NULL, "view_permission_check" bool NOT NULL, "limit_expression" text NOT NULL, "display_form_str" text NOT NULL, "title_template" varchar(50) NULL, "type_id" integer NOT NULL REFERENCES "workflows_workflowtype" ("id") DEFERRABLE INITIALLY DEFERRED);

-- 正在导出表  main.workflows_workflow 的数据：-1 rows
/*!40000 ALTER TABLE "workflows_workflow" DISABLE KEYS */;
INSERT INTO `workflows_workflow` (`id`, `create_time`, `update_time`, `memo`, `name`, `ticket_sn_prefix`, `status`, `view_permission_check`, `limit_expression`, `display_form_str`, `title_template`, `type_id`) VALUES
	(1, '2020-06-13 06:08:00.510239', '2020-06-13 06:08:00.510239', '', '请假单', 'leave', 1, 1, '{}', '[]', '你有一个待办工单:{title}', 1);
INSERT INTO `workflows_workflow` (`id`, `create_time`, `update_time`, `memo`, `name`, `ticket_sn_prefix`, `status`, `view_permission_check`, `limit_expression`, `display_form_str`, `title_template`, `type_id`) VALUES
	(2, '2020-06-13 06:08:05.777146', '2020-06-13 06:08:05.777146', '', '发布单', 'deploy', 1, 1, '{}', '[]', '你有一个待办工单:{title}', 2);
INSERT INTO `workflows_workflow` (`id`, `create_time`, `update_time`, `memo`, `name`, `ticket_sn_prefix`, `status`, `view_permission_check`, `limit_expression`, `display_form_str`, `title_template`, `type_id`) VALUES
	(3, '2020-06-13 15:09:09.310868', '2020-06-13 15:09:09.459473', '盖章签呈', '盖章签呈', 'gzqc', 1, 1, '', '', '', 3);
/*!40000 ALTER TABLE "workflows_workflow" ENABLE KEYS */;

-- 导出  表 main.workflows_state 结构
-- CREATE TABLE IF NOT EXISTS "workflows_state" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "create_time" datetime NOT NULL, "update_time" datetime NOT NULL, "memo" text NOT NULL, "name" varchar(50) NOT NULL, "is_hidden" bool NOT NULL, "order_id" integer NOT NULL, "state_type" varchar(1) NOT NULL, "enable_retreat" bool NOT NULL, "participant_type" varchar(1) NOT NULL, "workflow_id" integer NOT NULL REFERENCES "workflows_workflow" ("id") DEFERRABLE INITIALLY DEFERRED);

-- 正在导出表  main.workflows_state 的数据：-1 rows
/*!40000 ALTER TABLE "workflows_state" DISABLE KEYS */;
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(1, '2020-06-13 06:08:02.379239', '2020-06-13 06:08:02.379239', '', '开始', 1, 1, '1', 0, '0', 1);
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(2, '2020-06-13 06:08:02.511883', '2020-06-13 06:08:02.511883', '', '关闭', 1, 99, '2', 0, '0', 1);
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(3, '2020-06-13 06:08:02.633557', '2020-06-13 06:08:02.633557', '', '申请人-编辑中', 0, 2, '0', 0, '0', 1);
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(4, '2020-06-13 06:08:02.978636', '2020-06-13 06:08:02.978636', '', '领导-审批中', 0, 3, '0', 0, '3', 1);
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(5, '2020-06-13 06:08:03.488270', '2020-06-13 06:08:03.488270', '', '人事-审批中', 0, 4, '0', 0, '2', 1);
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(6, '2020-06-13 06:08:03.969982', '2020-06-13 06:08:03.969982', '', '结束', 0, 98, '2', 0, '0', 1);
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(7, '2020-06-13 06:08:07.422742', '2020-06-13 06:08:07.422742', '', '开始', 1, 1, '1', 0, '0', 2);
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(8, '2020-06-13 06:08:07.568352', '2020-06-13 06:08:07.569349', '', '关闭', 1, 99, '2', 0, '0', 2);
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(9, '2020-06-13 06:08:07.939360', '2020-06-13 06:08:07.939360', '', '申请人-编辑中', 0, 2, '0', 0, '0', 2);
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(10, '2020-06-13 06:08:08.244543', '2020-06-13 06:08:08.244543', '', '领导-审批中', 0, 3, '0', 0, '3', 2);
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(11, '2020-06-13 06:08:08.676391', '2020-06-13 06:08:08.676391', '', '运维-执行中', 0, 4, '0', 0, '2', 2);
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(12, '2020-06-13 06:08:09.365586', '2020-06-13 06:08:09.365586', '', '结束', 0, 98, '2', 0, '0', 2);
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(13, '2020-06-13 15:09:09.659937', '2020-06-13 15:09:09.659937', '', '开始', 1, 1, '1', 0, '0', 3);
INSERT INTO `workflows_state` (`id`, `create_time`, `update_time`, `memo`, `name`, `is_hidden`, `order_id`, `state_type`, `enable_retreat`, `participant_type`, `workflow_id`) VALUES
	(14, '2020-06-13 15:09:09.888852', '2020-06-13 15:09:09.888852', '', '关闭', 1, 99, '2', 0, '0', 3);
/*!40000 ALTER TABLE "workflows_state" ENABLE KEYS */;



/*!40000 ALTER TABLE "workflows_customfield" ENABLE KEYS */;
-- 导出  表 main.workflows_customfield 结构
-- CREATE TABLE IF NOT EXISTS "workflows_customfield" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "create_time" datetime NOT NULL, "update_time" datetime NOT NULL, "memo" text NOT NULL, "field_attribute" bool NOT NULL, "field_type" varchar(1) NOT NULL, "field_key" varchar(50) NOT NULL, "field_name" varchar(50) NOT NULL, "order_id" integer NOT NULL, "default_value" varchar(100) NULL, "field_template" text NOT NULL, "boolean_field_display" varchar(100) NOT NULL, "field_choice" varchar(255) NOT NULL, "label" varchar(100) NOT NULL, "workflow_id" integer NOT NULL REFERENCES "workflows_workflow" ("id") DEFERRABLE INITIALLY DEFERRED);

-- 正在导出表  main.workflows_customfield 的数据：-1 rows
/*!40000 ALTER TABLE "workflows_customfield" DISABLE KEYS */;
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(1, '2020-06-13 06:08:00.658842', '2020-06-13 06:08:00.660836', '', 1, '1', 'create_user', '申请人', 1, NULL, '', '{}', '{}', '{}', 1);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(2, '2020-06-13 06:08:00.801460', '2020-06-13 06:08:00.801460', '', 1, '6', 'create_time', '申请时间', 2, NULL, '', '{}', '{}', '{}', 1);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(3, '2020-06-13 06:08:01.186430', '2020-06-13 06:08:01.186430', '', 1, '1', 'group', '部门', 3, NULL, '', '{}', '{}', '{}', 1);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(4, '2020-06-13 06:08:01.326098', '2020-06-13 06:08:01.326098', '', 1, '2', 'id', '工号', 4, NULL, '', '{}', '{}', '{}', 1);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(5, '2020-06-13 06:08:01.568408', '2020-06-13 06:08:01.568408', '', 0, '7', 'start_end_time', '请假时间', 10, NULL, '', '{}', '{}', '{}', 1);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(6, '2020-06-13 06:08:01.736956', '2020-06-13 06:08:01.736956', '', 0, '9', 'type', '请假类型', 30, NULL, '', '{}', '{`1`:`病假`, `2`:`产假`}', '{}', 1);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(7, '2020-06-13 06:08:01.889549', '2020-06-13 06:08:01.889549', '', 0, '8', 'memo', '事由说明', 50, NULL, '', '{}', '{}', '{}', 1);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(8, '2020-06-13 06:08:02.023192', '2020-06-13 06:08:02.023192', '', 0, '9', 'leader_radio', '领导审批', 60, NULL, '', '{}', '{`1`:`同意`, `2`:`不同意`}', '{}', 1);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(9, '2020-06-13 06:08:02.236620', '2020-06-13 06:08:02.236620', '', 0, '9', 'hr_radio', '人事审批', 80, NULL, '', '{}', '{`1`:`同意`, `2`:`不同意`}', '{}', 1);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(10, '2020-06-13 06:08:05.922758', '2020-06-13 06:08:05.922758', '', 1, '1', 'create_user', '申请人', 1, NULL, '', '{}', '{}', '{}', 2);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(11, '2020-06-13 06:08:06.122222', '2020-06-13 06:08:06.122222', '', 1, '6', 'create_time', '申请时间', 2, NULL, '', '{}', '{}', '{}', 2);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(12, '2020-06-13 06:08:06.277805', '2020-06-13 06:08:06.277805', '', 1, '1', 'group', '部门', 3, NULL, '', '{}', '{}', '{}', 2);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(13, '2020-06-13 06:08:06.411448', '2020-06-13 06:08:06.411448', '', 1, '2', 'id', '工号', 4, NULL, '', '{}', '{}', '{}', 2);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(14, '2020-06-13 06:08:06.557058', '2020-06-13 06:08:06.557058', '', 0, '6', 'start_time', '发布时间', 10, NULL, '', '{}', '{}', '{}', 2);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(15, '2020-06-13 06:08:06.701671', '2020-06-13 06:08:06.701671', '', 0, '9', 'type', '发布项目', 30, NULL, '', '{}', '{`1`:`前端`, `2`:`后端`}', '{}', 2);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(16, '2020-06-13 06:08:06.988902', '2020-06-13 06:08:06.988902', '', 0, '8', 'memo', '发布内容', 50, NULL, '', '{}', '{}', '{}', 2);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(17, '2020-06-13 06:08:07.134514', '2020-06-13 06:08:07.134514', '', 0, '9', 'leader_radio', '领导审批', 60, NULL, '', '{}', '{`1`:`同意`, `2`:`不同意`}', '{}', 2);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(18, '2020-06-13 06:08:07.267161', '2020-06-13 06:08:07.267161', '', 0, '9', 'ops_radio', '运维执行', 80, NULL, '', '{}', '{`1`:`已执行`, `2`:`未执行`}', '{}', 2);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(19, '2020-06-13 15:09:10.093305', '2020-06-13 15:09:10.093305', '', 1, '1', 'create_user', '申请人', 1, NULL, '', '{}', '{}', '{}', 3);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(20, '2020-06-13 15:09:10.225030', '2020-06-13 15:09:10.226028', '', 1, '6', 'create_time', '申请时间', 2, NULL, '', '{}', '{}', '{}', 3);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(21, '2020-06-13 15:09:10.369270', '2020-06-13 15:09:10.369270', '', 1, '1', 'group', '部门', 3, NULL, '', '{}', '{}', '{}', 3);
INSERT INTO `workflows_customfield` (`id`, `create_time`, `update_time`, `memo`, `field_attribute`, `field_type`, `field_key`, `field_name`, `order_id`, `default_value`, `field_template`, `boolean_field_display`, `field_choice`, `label`, `workflow_id`) VALUES
	(22, '2020-06-13 15:09:10.514508', '2020-06-13 15:09:10.514508', '', 1, '1', 'id', '工号', 4, NULL, '', '{}', '{}', '{}', 3);
	
-- 导出  表 main.workflows_state_fields 结构
-- CREATE TABLE IF NOT EXISTS "workflows_state_fields" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "state_id" integer NOT NULL REFERENCES "workflows_state" ("id") DEFERRABLE INITIALLY DEFERRED, "customfield_id" integer NOT NULL REFERENCES "workflows_customfield" ("id") DEFERRABLE INITIALLY DEFERRED);

-- 正在导出表  main.workflows_state_fields 的数据：-1 rows
/*!40000 ALTER TABLE "workflows_state_fields" DISABLE KEYS */;
INSERT INTO `workflows_state_fields` (`id`, `state_id`, `customfield_id`) VALUES
	(1, 3, 5);
INSERT INTO `workflows_state_fields` (`id`, `state_id`, `customfield_id`) VALUES
	(2, 3, 6);
INSERT INTO `workflows_state_fields` (`id`, `state_id`, `customfield_id`) VALUES
	(3, 3, 7);
INSERT INTO `workflows_state_fields` (`id`, `state_id`, `customfield_id`) VALUES
	(4, 4, 8);
INSERT INTO `workflows_state_fields` (`id`, `state_id`, `customfield_id`) VALUES
	(5, 5, 9);
INSERT INTO `workflows_state_fields` (`id`, `state_id`, `customfield_id`) VALUES
	(6, 9, 16);
INSERT INTO `workflows_state_fields` (`id`, `state_id`, `customfield_id`) VALUES
	(7, 9, 14);
INSERT INTO `workflows_state_fields` (`id`, `state_id`, `customfield_id`) VALUES
	(8, 9, 15);
INSERT INTO `workflows_state_fields` (`id`, `state_id`, `customfield_id`) VALUES
	(9, 10, 17);
INSERT INTO `workflows_state_fields` (`id`, `state_id`, `customfield_id`) VALUES
	(10, 11, 18);
/*!40000 ALTER TABLE "workflows_state_fields" ENABLE KEYS */;

-- 导出  表 main.workflows_state_group_participant 结构
-- CREATE TABLE IF NOT EXISTS "workflows_state_group_participant" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "state_id" integer NOT NULL REFERENCES "workflows_state" ("id") DEFERRABLE INITIALLY DEFERRED, "group_id" integer NOT NULL REFERENCES "systems_group" ("group_ptr_id") DEFERRABLE INITIALLY DEFERRED);

-- 正在导出表  main.workflows_state_group_participant 的数据：-1 rows
/*!40000 ALTER TABLE "workflows_state_group_participant" DISABLE KEYS */;
INSERT INTO `workflows_state_group_participant` (`id`, `state_id`, `group_id`) VALUES
	(1, 5, 7);
INSERT INTO `workflows_state_group_participant` (`id`, `state_id`, `group_id`) VALUES
	(2, 11, 7);
/*!40000 ALTER TABLE "workflows_state_group_participant" ENABLE KEYS */;

-- 导出  表 main.workflows_state_role_participant 结构
-- CREATE TABLE IF NOT EXISTS "workflows_state_role_participant" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "state_id" integer NOT NULL REFERENCES "workflows_state" ("id") DEFERRABLE INITIALLY DEFERRED, "role_id" integer NOT NULL REFERENCES "systems_role" ("id") DEFERRABLE INITIALLY DEFERRED);

-- 正在导出表  main.workflows_state_role_participant 的数据：-1 rows
/*!40000 ALTER TABLE "workflows_state_role_participant" DISABLE KEYS */;
INSERT INTO `workflows_state_role_participant` (`id`, `state_id`, `role_id`) VALUES
	(1, 4, 1);
INSERT INTO `workflows_state_role_participant` (`id`, `state_id`, `role_id`) VALUES
	(2, 4, 2);
INSERT INTO `workflows_state_role_participant` (`id`, `state_id`, `role_id`) VALUES
	(3, 4, 3);
INSERT INTO `workflows_state_role_participant` (`id`, `state_id`, `role_id`) VALUES
	(4, 10, 1);
INSERT INTO `workflows_state_role_participant` (`id`, `state_id`, `role_id`) VALUES
	(5, 10, 2);
INSERT INTO `workflows_state_role_participant` (`id`, `state_id`, `role_id`) VALUES
	(6, 10, 3);
/*!40000 ALTER TABLE "workflows_state_role_participant" ENABLE KEYS */;

-- 导出  表 main.workflows_state_user_participant 结构
-- CREATE TABLE IF NOT EXISTS "workflows_state_user_participant" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "state_id" integer NOT NULL REFERENCES "workflows_state" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" integer NOT NULL REFERENCES "systems_user" ("id") DEFERRABLE INITIALLY DEFERRED);

-- 正在导出表  main.workflows_state_user_participant 的数据：-1 rows
/*!40000 ALTER TABLE "workflows_state_user_participant" DISABLE KEYS */;
/*!40000 ALTER TABLE "workflows_state_user_participant" ENABLE KEYS */;

-- 导出  表 main.workflows_transition 结构
-- CREATE TABLE IF NOT EXISTS "workflows_transition" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "create_time" datetime NOT NULL, "update_time" datetime NOT NULL, "memo" text NOT NULL, "name" varchar(1) NOT NULL, "transition_type" varchar(1) NOT NULL, "timer" integer NOT NULL, "condition_expression" text NOT NULL, "attribute_type" varchar(1) NOT NULL, "alert_enable" bool NOT NULL, "alert_text" varchar(100) NOT NULL, "dest_state_id" integer NULL REFERENCES "workflows_state" ("id") DEFERRABLE INITIALLY DEFERRED, "source_state_id" integer NULL REFERENCES "workflows_state" ("id") DEFERRABLE INITIALLY DEFERRED, "workflow_id" integer NOT NULL REFERENCES "workflows_workflow" ("id") DEFERRABLE INITIALLY DEFERRED);

-- 正在导出表  main.workflows_transition 的数据：-1 rows
/*!40000 ALTER TABLE "workflows_transition" DISABLE KEYS */;
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(1, '2020-06-13 06:08:04.099634', '2020-06-13 06:08:04.099634', '', '0', '0', 0, '[]', '0', 0, '', 3, 1, 1);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(2, '2020-06-13 06:08:04.222309', '2020-06-13 06:08:04.222309', '', '1', '0', 0, '[]', '1', 0, '', 4, 1, 1);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(3, '2020-06-13 06:08:04.354950', '2020-06-13 06:08:04.354950', '', '0', '0', 0, '[]', '0', 0, '', 3, 3, 1);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(4, '2020-06-13 06:08:04.554418', '2020-06-13 06:08:04.554418', '', '1', '0', 0, '[]', '1', 0, '', 4, 3, 1);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(5, '2020-06-13 06:08:04.677089', '2020-06-13 06:08:04.677089', '', '3', '0', 0, '[]', '3', 0, '', 6, 3, 1);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(6, '2020-06-13 06:08:04.810733', '2020-06-13 06:08:04.810733', '', '2', '0', 0, '[]', '2', 0, '', 3, 4, 1);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(7, '2020-06-13 06:08:05.133867', '2020-06-13 06:08:05.133867', '', '1', '0', 0, '[]', '1', 0, '', 5, 4, 1);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(8, '2020-06-13 06:08:05.256538', '2020-06-13 06:08:05.256538', '', '2', '0', 0, '[]', '2', 0, '', 3, 5, 1);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(9, '2020-06-13 06:08:05.402150', '2020-06-13 06:08:05.402150', '', '4', '0', 0, '[]', '5', 0, '', 2, 5, 1);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(10, '2020-06-13 06:08:09.513148', '2020-06-13 06:08:09.513148', '', '0', '0', 0, '[]', '0', 0, '', 9, 7, 2);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(11, '2020-06-13 06:08:09.668732', '2020-06-13 06:08:09.668732', '', '1', '0', 0, '[]', '1', 0, '', 10, 7, 2);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(12, '2020-06-13 06:08:09.802375', '2020-06-13 06:08:09.802375', '', '0', '0', 0, '[]', '0', 0, '', 9, 9, 2);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(13, '2020-06-13 06:08:09.987878', '2020-06-13 06:08:09.987878', '', '1', '0', 0, '[]', '1', 0, '', 10, 9, 2);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(14, '2020-06-13 06:08:10.134487', '2020-06-13 06:08:10.134487', '', '3', '0', 0, '[]', '3', 0, '', 12, 9, 2);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(15, '2020-06-13 06:08:10.433686', '2020-06-13 06:08:10.433686', '', '2', '0', 0, '[]', '2', 0, '', 9, 10, 2);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(16, '2020-06-13 06:08:10.567328', '2020-06-13 06:08:10.567328', '', '1', '0', 0, '[]', '1', 0, '', 11, 10, 2);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(17, '2020-06-13 06:08:10.722912', '2020-06-13 06:08:10.722912', '', '2', '0', 0, '[]', '2', 0, '', 9, 11, 2);
INSERT INTO `workflows_transition` (`id`, `create_time`, `update_time`, `memo`, `name`, `transition_type`, `timer`, `condition_expression`, `attribute_type`, `alert_enable`, `alert_text`, `dest_state_id`, `source_state_id`, `workflow_id`) VALUES
	(18, '2020-06-13 06:08:10.878497', '2020-06-13 06:08:10.878497', '', '4', '0', 0, '[]', '5', 0, '', 8, 11, 2);
/*!40000 ALTER TABLE "workflows_transition" ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
