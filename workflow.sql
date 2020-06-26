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

CREATE TABLE `workflows_workflowtype` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`memo` MEDIUMTEXT NOT NULL COMMENT '备注' COLLATE 'utf8mb4_general_ci',
	`name` VARCHAR(50) NOT NULL COMMENT '名称' COLLATE 'utf8mb4_general_ci',
	`code` VARCHAR(32) NOT NULL COMMENT '代码' COLLATE 'utf8mb4_general_ci',
	`order_id` INT(11) NOT NULL DEFAULT '1' COMMENT '状态顺序',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `uniq_workflows_workflowtype_code` (`code`) USING BTREE
)
COMMENT='工作流类型'
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;


CREATE TABLE `workflows_workflow` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`memo` MEDIUMTEXT NOT NULL COMMENT '备注' COLLATE 'utf8mb4_general_ci',
	`name` VARCHAR(50) NOT NULL COMMENT '名称' COLLATE 'utf8mb4_general_ci',
	`ticket_sn_prefix` VARCHAR(20) NOT NULL DEFAULT 'xxoo' COMMENT '工单流水号前缀' COLLATE 'utf8mb4_general_ci',
	`status` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '状态',
	`view_permission_check` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '查看权限校验，开启后，只允许工单的关联人(创建人、曾经的处理人)有权限查看工单',
	`limit_expression` MEDIUMTEXT NOT NULL COMMENT '限制表达式，限制周期({"period":24} 24小时), 限制次数({"count":1}在限制周期内只允许提交1次), 限制级别({"level":1} 针对(1单个用户 2全局)限制周期限制次数,默认特定用户);允许特定人员提交({"allow_persons":"zhangsan,lisi"}只允许张三提交工单,{"allow_depts":"1,2"}只允许部门id为1和2的用户提交工单，{"allow_roles":"1,2"}只允许角色id为1和2的用户提交工单)' COLLATE 'utf8mb4_general_ci',
	`display_form_str` MEDIUMTEXT NOT NULL COMMENT '展现表单字段，默认"[]"，用于用户只有对应工单查看权限时显示哪些字段,field_key的list的json,如["days","sn"],内置特殊字段participant_info.participant_name:当前处理人信息(部门名称、角色名称)，state.state_name:当前状态的状态名,workflow.workflow_name:工作流名称' COLLATE 'utf8mb4_general_ci',
	`title_template` VARCHAR(50) NULL DEFAULT '你有一个待办工单:{title}' COMMENT '标题模板，工单字段的值可以作为参数写到模板中，格式如：你有一个待办工单:{title}' COLLATE 'utf8mb4_general_ci',
	`type_id` BIGINT(20) NOT NULL COMMENT '工作流类型',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `workflows_workflow_type_id` (`type_id`) USING BTREE,
	CONSTRAINT `FK_workflows_workflow_fields_type_id_workflows_workflowtype` FOREIGN KEY (`type_id`) REFERENCES `goadmindb`.`workflows_workflowtype` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COMMENT='工作流'
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;


CREATE TABLE `workflows_customfield` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`memo` MEDIUMTEXT NOT NULL COMMENT '备注' COLLATE 'utf8mb4_general_ci',
	`field_attribute` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '字段是否内置',
	`field_type` TINYINT(4) NOT NULL DEFAULT '0' COMMENT '字段类型，1: \'字符串\',2: \'整形\',3: \'浮点型\',4: \'布尔\',5: \'日期\',6: \'日期时间\',7: \'范围日期\',8: \'文本域\',9: \'单选框\',10: \'下拉列表\',11: \'用户名\',12: \'多选框\',13: \'多选下拉\',14: \'多选用户名\',',
	`field_key` VARCHAR(50) NOT NULL COMMENT '字段标识，字段类型请尽量特殊，避免与系统中关键字冲突' COLLATE 'utf8mb4_general_ci',
	`field_name` VARCHAR(50) NOT NULL COMMENT '字段名称' COLLATE 'utf8mb4_general_ci',
	`order_id` INT(11) NOT NULL DEFAULT '0' COMMENT '排序',
	`default_value` VARCHAR(100) NULL DEFAULT NULL COMMENT '默认值，前端展示时，可以将此内容作为表单中的该字段的默认值' COLLATE 'utf8mb4_general_ci',
	`field_template` MEDIUMTEXT NOT NULL COMMENT '文本域模板，文本域类型字段前端显示时可以将此内容作为字段的placeholder' COLLATE 'utf8mb4_general_ci',
	`boolean_field_display` VARCHAR(100) NOT NULL DEFAULT '{}' COMMENT '布尔类型显示名，当为布尔类型时候，可以支持自定义显示形式。{"1":"是","0":"否"}或{"1":"需要","0":"不需要"}，注意数字也需要引号' COLLATE 'utf8mb4_general_ci',
	`field_choice` VARCHAR(255) NOT NULL DEFAULT '{}' COMMENT 'radio、checkbox、select的选项。radio,checkbox,select,multiselect类型可供选择的选项，格式为json如:{"1":"中国", "2":"美国"},注意数字也需要引号' COLLATE 'utf8mb4_general_ci',
	`label` VARCHAR(100) NOT NULL DEFAULT '{}' COMMENT '标签，自定义标签，json格式，调用方可根据标签自行处理特殊场景逻辑，loonflow只保存文本内容' COLLATE 'utf8mb4_general_ci',
	`workflow_id` BIGINT(20) NOT NULL COMMENT '工作流',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `workflows_customfield_workflow_id` (`workflow_id`) USING BTREE,
	CONSTRAINT `FK_workflows_customfield_workflow_id_workflows_workflow` FOREIGN KEY (`workflow_id`) REFERENCES `goadmindb`.`workflows_workflow` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COMMENT='自定义字段, 设定某个工作流有哪些自定义字段'
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;


CREATE TABLE `workflows_state` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`memo` MEDIUMTEXT NOT NULL COMMENT '备注' COLLATE 'utf8mb4_general_ci',
	`name` VARCHAR(50) NOT NULL COMMENT '名称' COLLATE 'utf8mb4_general_ci',
	`is_hidden` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '是否隐藏，设置为True时,获取工单步骤api中不显示此状态(当前处于此状态时除外)',
	`order_id` INT(11) NOT NULL DEFAULT '1' COMMENT '状态顺序',
	`state_type` TINYINT(4) NOT NULL DEFAULT '0' COMMENT '状态类型.0: \'普通状态\',1: \'初始状态\',2: \'结束状态\',',
	`enable_retreat` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '允许撤回，开启后允许工单创建人在此状态直接撤回工单到初始状态',
	`participant_type` TINYINT(4) NOT NULL DEFAULT '0' COMMENT '0: \'无处理人\',1: \'个人\',2: \'部门\',3: \'角色\',',
	`workflow_id` BIGINT(20) NOT NULL COMMENT '工作流',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `workflows_state_workflow_id` (`workflow_id`) USING BTREE,
	CONSTRAINT `FK_workflows_state_workflows_workflow` FOREIGN KEY (`workflow_id`) REFERENCES `goadmindb`.`workflows_workflow` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COMMENT='状态记录, 变量支持通过脚本获取'
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;


CREATE TABLE `workflows_state_fields` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`state_id` BIGINT(20) NOT NULL COMMENT '当前状态',
	`customfield_id` BIGINT(20) NOT NULL COMMENT '编辑字段',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `workflows_state_fields_state_id_customfield_id_uniq` (`state_id`, `customfield_id`) USING BTREE,
	INDEX `workflows_state_fields_customfield_id` (`customfield_id`) USING BTREE,
	INDEX `workflows_state_fields_state_id` (`state_id`) USING BTREE,
	CONSTRAINT `FK_workflows_state_fields_customfield_id_workflows_customfield` FOREIGN KEY (`customfield_id`) REFERENCES `goadmindb`.`workflows_customfield` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_workflows_state_fields_state_id_workflows_state` FOREIGN KEY (`state_id`) REFERENCES `goadmindb`.`workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COMMENT='可编辑字段'
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;


CREATE TABLE `workflows_state_group_participant` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`state_id` BIGINT(20) NOT NULL COMMENT '当前状态',
	`group_id` INT(11) NOT NULL COMMENT '参与组',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `workflows_state_group_participant_state_id_group_id_uniq` (`state_id`, `group_id`) USING BTREE,
	INDEX `workflows_state_group_participant_group_id` (`group_id`) USING BTREE,
	INDEX `workflows_state_group_participant_state_id` (`state_id`) USING BTREE,
	CONSTRAINT `FK_workflows_state_group_participant_group_id` FOREIGN KEY (`group_id`) REFERENCES `goadmindb`.`sys_dept` (`dept_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_workflows_state_group_participant_state_id` FOREIGN KEY (`state_id`) REFERENCES `goadmindb`.`workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COMMENT='参与组'
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;


CREATE TABLE `workflows_state_role_participant` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`state_id` BIGINT(20) NOT NULL COMMENT '当前状态',
	`role_id` INT(11) NOT NULL COMMENT '角色',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `workflows_state_role_participant_state_id_role_id_uniq` (`state_id`, `role_id`) USING BTREE,
	INDEX `workflows_state_role_participant_role_id` (`role_id`) USING BTREE,
	INDEX `workflows_state_role_participant_state_id` (`state_id`) USING BTREE,
	CONSTRAINT `FK_workflows_state_role_participant_role_id` FOREIGN KEY (`role_id`) REFERENCES `goadmindb`.`sys_role` (`role_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_workflows_state_role_participant_state_id` FOREIGN KEY (`state_id`) REFERENCES `goadmindb`.`workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COMMENT='参与角色'
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;

CREATE TABLE `workflows_state_user_participant` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`state_id` BIGINT(20) NOT NULL COMMENT '当前状态',
	`user_id` INT(11) NOT NULL COMMENT '当前用户',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `workflows_state_user_participant_state_id_user_id_uniq` (`state_id`, `user_id`) USING BTREE,
	INDEX `workflows_state_user_participant_user_id` (`user_id`) USING BTREE,
	INDEX `workflows_state_user_participant_state_id` (`state_id`) USING BTREE,
	CONSTRAINT `FK_workflows_state_user_participant_state_id` FOREIGN KEY (`state_id`) REFERENCES `goadmindb`.`workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_workflows_state_user_participant_user_id` FOREIGN KEY (`user_id`) REFERENCES `goadmindb`.`sys_user` (`user_id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COMMENT='参与用户'
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;


CREATE TABLE `workflows_transition` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`memo` MEDIUMTEXT NOT NULL COMMENT '备注' COLLATE 'utf8mb4_general_ci',
	`name` TINYINT(4) NOT NULL DEFAULT '1' COMMENT '名称类型， 0: \'保存\',1: \'转交下一步\',2: \'驳回\',3: \'撤销\',4: \'关闭\',',
	`transition_type` TINYINT(4) NOT NULL DEFAULT '0' COMMENT '流转类型.0: \'常规流转\',1: \'定时器流转\',',
	`timer` INT(11) NOT NULL DEFAULT '0' COMMENT '定时器(单位秒).流转类型设置为定时器流转时生效,单位秒。处于源状态X秒后如果状态都没有过变化则自动流转到目标状态',
	`condition_expression` MEDIUMTEXT NOT NULL COMMENT '条件表达式。流转条件表达式，根据表达式中的条件来确定流转的下个状态，格式为[{"expression":"{days} > 3 and {days}<10", "target_state_id":11}] 其中{}用于填充工单的字段key,运算时会换算成实际的值，当符合条件下个状态将变为target_state_id中的值,表达式只支持简单的运算或datetime/time运算.loonflow会以首次匹配成功的条件为准，所以多个条件不要有冲突' COLLATE 'utf8mb4_general_ci',
	`attribute_type` TINYINT(4) NOT NULL DEFAULT '0' COMMENT '属性类型，0: \'草稿中\',1: \'进行中\',2: \'被退回\',3: \'被撤销\',4: \'已完成\',5: \'已关闭\',',
	`alert_enable` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '点击弹窗提示',
	`alert_text` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '弹窗内容' COLLATE 'utf8mb4_general_ci',
	`dest_state_id` BIGINT(20) NULL DEFAULT NULL COMMENT '目的状态',
	`source_state_id` BIGINT(20) NULL DEFAULT NULL COMMENT '源状态',
	`workflow_id` BIGINT(20) NOT NULL COMMENT '工作流',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `workflows_transition_workflow_id` (`workflow_id`) USING BTREE,
	INDEX `workflows_transition_source_state_id` (`source_state_id`) USING BTREE,
	INDEX `workflows_transition_dest_state_id` (`dest_state_id`) USING BTREE,
	CONSTRAINT `FK_workflows_transition_dest_state_id` FOREIGN KEY (`dest_state_id`) REFERENCES `goadmindb`.`workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_workflows_transition_source_state_id` FOREIGN KEY (`source_state_id`) REFERENCES `goadmindb`.`workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_workflows_transition_workflow_id` FOREIGN KEY (`workflow_id`) REFERENCES `goadmindb`.`workflows_workflow` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COMMENT='工作流流转，定时器，条件(允许跳过)， 条件流转与定时器不可同时存在'
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


-- 导出  表 main.tickets_ticket 结构
DROP TABLE tickets_ticket;
CREATE TABLE `tickets_ticket` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`memo` TEXT(65535) NOT NULL COMMENT '备注' COLLATE 'utf8mb4_general_ci',
	`name` VARCHAR(112) NOT NULL COMMENT '标题' COLLATE 'utf8mb4_general_ci',
	`sn` VARCHAR(25) NOT NULL COMMENT '流水号,工单的流水号' COLLATE 'utf8mb4_general_ci',
	`participant` VARCHAR(50) NOT NULL COMMENT '当前处理人' COLLATE 'utf8mb4_general_ci',
	`customfield` MEDIUMTEXT NOT NULL COMMENT '所有表单数据' COLLATE 'utf8mb4_general_ci',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	`workflow_id` BIGINT(20) NULL DEFAULT NULL COMMENT '工作流',
	`transition_id` BIGINT(20) NULL DEFAULT NULL COMMENT '进行状态',
	`state_id` BIGINT(20) NULL DEFAULT NULL COMMENT '当前状态',
	`create_user_id` INT(11) NULL DEFAULT NULL COMMENT '创建者',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `tickets_ticket_workflow_id` (`workflow_id`) USING BTREE,
	INDEX `tickets_ticket_transition_id` (`transition_id`) USING BTREE,
	INDEX `tickets_ticket_state_id` (`state_id`) USING BTREE,
	INDEX `tickets_ticket_create_user_id` (`create_user_id`) USING BTREE,
	CONSTRAINT `FK_tickets_ticket_create_user_id` FOREIGN KEY (`create_user_id`) REFERENCES `goadmindb`.`sys_user` (`user_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_tickets_ticket_state_id` FOREIGN KEY (`state_id`) REFERENCES `goadmindb`.`workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_tickets_ticket_transition_id` FOREIGN KEY (`transition_id`) REFERENCES `goadmindb`.`workflows_transition` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_tickets_ticket_workflow_id` FOREIGN KEY (`workflow_id`) REFERENCES `goadmindb`.`workflows_workflow` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COMMENT='工单记录'
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;


-- 正在导出表  main.tickets_ticket 的数据：-1 rows
/*!40000 ALTER TABLE `tickets_ticket` DISABLE KEYS */;
INSERT INTO `tickets_ticket` (`id`, `create_time`, `update_time`, `memo`, `name`, `sn`, `participant`, `customfield`, `create_user_id`, `state_id`, `transition_id`, `workflow_id`) VALUES
	(1, '2020-06-13 10:33:36.882952', '2020-06-13 10:33:36.882952', '', '请假单-2020-06-13-18-32-48-202', 'leave_20200613183336522', '', '[{`customfield`:1,`field_key`:`create_user`,`field_value`:``},{`customfield`:2,`field_key`:`create_time`,`field_value`:``},{`customfield`:3,`field_key`:`group`,`field_value`:``},{`customfield`:4,`field_key`:`id`,`field_value`:``},{`customfield`:5,`field_key`:`start_end_time`,`field_value`:[`2020-06-16`,`2020-07-16`]},{`customfield`:6,`field_key`:`type`,`field_value`:`1`},{`customfield`:7,`field_key`:`memo`,`field_value`:`随便写写`},{`customfield`:8,`field_key`:`leader_radio`,`field_value`:``},{`customfield`:9,`field_key`:`hr_radio`,`field_value`:``}]', 1, 3, 1, 1);
INSERT INTO `tickets_ticket` (`id`, `create_time`, `update_time`, `memo`, `name`, `sn`, `participant`, `customfield`, `create_user_id`, `state_id`, `transition_id`, `workflow_id`) VALUES
	(2, '2020-06-25 05:50:09.858655', '2020-06-25 05:50:09.858655', '', '请假单-2020-06-25-13-42-27-277', 'leave_20200625135009552', '', '[{`customfield`:1,`field_key`:`create_user`,`field_value`:``},{`customfield`:2,`field_key`:`create_time`,`field_value`:``},{`customfield`:3,`field_key`:`group`,`field_value`:``},{`customfield`:4,`field_key`:`id`,`field_value`:``},{`customfield`:5,`field_key`:`start_end_time`,`field_value`:[`2020-06-25`,`2020-07-24`]},{`customfield`:6,`field_key`:`type`,`field_value`:`2`},{`customfield`:7,`field_key`:`memo`,`field_value`:`建设银行`},{`customfield`:8,`field_key`:`leader_radio`,`field_value`:``},{`customfield`:9,`field_key`:`hr_radio`,`field_value`:``}]', 1, 3, 1, 1);
/*!40000 ALTER TABLE `tickets_ticket` ENABLE KEYS */;

-- 导出  表 main.tickets_ticketcustomfield 结构CREATE TABLE `tickets_ticketcustomfield` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`memo` TEXT(65535) NOT NULL COMMENT '备注' COLLATE 'utf8mb4_general_ci',
	`field_value` TEXT(65535) NOT NULL COMMENT '字段值' COLLATE 'utf8mb4_general_ci',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	`customfield_id` BIGINT(20) NOT NULL COMMENT '字段',
	`ticket_id` BIGINT(20) NOT NULL COMMENT '工单',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `tickets_ticketcustomfield_customfield_id` (`customfield_id`) USING BTREE,
	INDEX `tickets_ticketcustomfield_ticket_id` (`ticket_id`) USING BTREE,
	CONSTRAINT `FK_tickets_ticketcustomfield_customfield_id` FOREIGN KEY (`customfield_id`) REFERENCES `goadmindb`.`workflows_customfield` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_tickets_ticketcustomfield_ticket_id` FOREIGN KEY (`ticket_id`) REFERENCES `goadmindb`.`tickets_ticket` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COMMENT='工单自定义字段， 工单自定义字段实际的值。'
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;


-- 正在导出表  main.tickets_ticketcustomfield 的数据：-1 rows
/*!40000 ALTER TABLE `tickets_ticketcustomfield` DISABLE KEYS */;
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(1, '2020-06-13 10:33:37.167798', '2020-06-13 10:33:37.167798', '', 'admin', 1, 1);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(2, '2020-06-13 10:33:37.167798', '2020-06-13 10:33:37.167798', '', '2020-06-13 10:33:36.882952+00:00', 2, 1);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(3, '2020-06-13 10:33:37.167798', '2020-06-13 10:33:37.167798', '', 'top', 3, 1);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(4, '2020-06-13 10:33:37.167798', '2020-06-13 10:33:37.167798', '', '1', 4, 1);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(5, '2020-06-13 10:33:37.167798', '2020-06-13 10:33:37.167798', '', '[''2020-06-16'', ''2020-07-16'']', 5, 1);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(6, '2020-06-13 10:33:37.167798', '2020-06-13 10:33:37.167798', '', '1', 6, 1);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(7, '2020-06-13 10:33:37.167798', '2020-06-13 10:33:37.167798', '', '随便写写', 7, 1);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(8, '2020-06-13 10:33:37.167798', '2020-06-13 10:33:37.167798', '', '', 8, 1);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(9, '2020-06-13 10:33:37.167798', '2020-06-13 10:33:37.167798', '', '', 9, 1);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(10, '2020-06-25 05:50:10.130479', '2020-06-25 05:50:10.130479', '', 'admin', 1, 2);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(11, '2020-06-25 05:50:10.130479', '2020-06-25 05:50:10.130479', '', '2020-06-25 05:50:09.858655+00:00', 2, 2);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(12, '2020-06-25 05:50:10.130479', '2020-06-25 05:50:10.130479', '', 'top', 3, 2);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(13, '2020-06-25 05:50:10.130479', '2020-06-25 05:50:10.130479', '', '1', 4, 2);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(14, '2020-06-25 05:50:10.130479', '2020-06-25 05:50:10.130479', '', '[''2020-06-25'', ''2020-07-24'']', 5, 2);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(15, '2020-06-25 05:50:10.130479', '2020-06-25 05:50:10.130479', '', '2', 6, 2);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(16, '2020-06-25 05:50:10.130479', '2020-06-25 05:50:10.130479', '', '建设银行', 7, 2);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(17, '2020-06-25 05:50:10.131442', '2020-06-25 05:50:10.131442', '', '', 8, 2);
INSERT INTO `tickets_ticketcustomfield` (`id`, `create_time`, `update_time`, `memo`, `field_value`, `customfield_id`, `ticket_id`) VALUES
	(18, '2020-06-25 05:50:10.131442', '2020-06-25 05:50:10.131442', '', '', 9, 2);
/*!40000 ALTER TABLE `tickets_ticketcustomfield` ENABLE KEYS */;

-- 导出  表 main.tickets_ticketflowlog 结构CREATE TABLE `tickets_ticketflowlog` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`memo` TEXT(65535) NOT NULL COMMENT '备注' COLLATE 'utf8mb4_general_ci',
	`suggestion` VARCHAR(140) NOT NULL COMMENT '审批意见' COLLATE 'utf8mb4_general_ci',
	`participant` VARCHAR(50) NOT NULL COMMENT '处理人' COLLATE 'utf8mb4_general_ci',
	`intervene_type` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '干预类型.0: \'转交操作\',1: \'接单操作\',2: \'评论操作\',3: \'删除操作\',4: \'强制关闭操作\',5: \'强制修改状态操作\',6: \'撤回\',',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	`state_id` BIGINT(20) NOT NULL COMMENT '当前状态',
	`ticket_id` BIGINT(20) NOT NULL COMMENT '工单',
	`transition_id` BIGINT(20) NOT NULL COMMENT '流转',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `tickets_ticketflowlog_state_id` (`state_id`) USING BTREE,
	INDEX `tickets_ticketflowlog_ticket_id` (`ticket_id`) USING BTREE,
	INDEX `tickets_ticketflowlog_transition_id` (`transition_id`) USING BTREE,
	CONSTRAINT `FK_tickets_ticketflowlog_state_id` FOREIGN KEY (`state_id`) REFERENCES `goadmindb`.`workflows_state` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_tickets_ticketflowlog_ticket_id` FOREIGN KEY (`ticket_id`) REFERENCES `goadmindb`.`tickets_ticket` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT `FK_tickets_ticketflowlog_transition_id` FOREIGN KEY (`transition_id`) REFERENCES `goadmindb`.`workflows_transition` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COMMENT='工单流转日志'
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;


-- 正在导出表  main.tickets_ticketflowlog 的数据：-1 rows
/*!40000 ALTER TABLE `tickets_ticketflowlog` DISABLE KEYS */;
INSERT INTO `tickets_ticketflowlog` (`id`, `create_time`, `update_time`, `memo`, `suggestion`, `participant`, `intervene_type`, `state_id`, `ticket_id`, `transition_id`) VALUES
	(1, '2020-06-13 10:33:37.015098', '2020-06-13 10:33:37.015098', '', '没啥意见', 'admin', '0', 1, 1, 1);
INSERT INTO `tickets_ticketflowlog` (`id`, `create_time`, `update_time`, `memo`, `suggestion`, `participant`, `intervene_type`, `state_id`, `ticket_id`, `transition_id`) VALUES
	(2, '2020-06-25 05:50:09.997983', '2020-06-25 05:50:09.998992', '', '没啥意见', 'admin', '0', 1, 2, 1);
/*!40000 ALTER TABLE `tickets_ticketflowlog` ENABLE KEYS */;

-- 导出  表 main.tickets_ticketuser 结构CREATE TABLE `tickets_ticketuser` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`memo` TEXT(65535) NOT NULL COMMENT '备注' COLLATE 'utf8mb4_general_ci',
	`username` VARCHAR(100) NOT NULL COMMENT '关系人' COLLATE 'utf8mb4_general_ci',
	`in_process` TINYINT(1) NOT NULL COMMENT '待处理中',
	`worked` TINYINT(1) NOT NULL COMMENT '处理过',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	`deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
	`ticket_id` BIGINT(20) NULL DEFAULT NULL COMMENT '工单',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `tickets_ticketuser_ticket_id` (`ticket_id`) USING BTREE,
	CONSTRAINT `FK_tickets_ticketuser_ticket_id` FOREIGN KEY (`ticket_id`) REFERENCES `goadmindb`.`tickets_ticket` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
)
COMMENT='工单关系人, 用于加速待办工单及关联工单列表查询'
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;


-- 正在导出表  main.tickets_ticketuser 的数据：-1 rows
/*!40000 ALTER TABLE `tickets_ticketuser` DISABLE KEYS */;
INSERT INTO `tickets_ticketuser` (`id`, `create_time`, `update_time`, `memo`, `username`, `in_process`, `worked`, `ticket_id`) VALUES
	(1, '2020-06-13 10:33:37.298875', '2020-06-13 10:33:37.298875', '', 'admin', 0, 1, 1);
INSERT INTO `tickets_ticketuser` (`id`, `create_time`, `update_time`, `memo`, `username`, `in_process`, `worked`, `ticket_id`) VALUES
	(2, '2020-06-13 10:33:37.448840', '2020-06-13 10:33:37.448840', '', '', 1, 0, 1);
INSERT INTO `tickets_ticketuser` (`id`, `create_time`, `update_time`, `memo`, `username`, `in_process`, `worked`, `ticket_id`) VALUES
	(3, '2020-06-25 05:50:10.271901', '2020-06-25 05:50:10.271901', '', 'admin', 0, 1, 2);
INSERT INTO `tickets_ticketuser` (`id`, `create_time`, `update_time`, `memo`, `username`, `in_process`, `worked`, `ticket_id`) VALUES
	(4, '2020-06-25 05:50:10.405216', '2020-06-25 05:50:10.405216', '', '', 1, 0, 2);
/*!40000 ALTER TABLE `tickets_ticketuser` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
