-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.7.17 - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 goworkflow 的数据库结构
-- CREATE DATABASE IF NOT EXISTS `goworkflow` /*!40100 DEFAULT CHARACTER SET utf8 */;
-- USE `goworkflow`;

CREATE DATABASE IF NOT EXISTS `goadmindb` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `goadmindb`;

-- 导出  表 goworkflow.execution 结构
CREATE TABLE IF NOT EXISTS `execution` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rev` int(11) DEFAULT NULL,
  `proc_def_id` int(11) DEFAULT NULL,
  `proc_def_name` varchar(255) DEFAULT NULL,
  `node_infos` varchar(4000) DEFAULT NULL,
  `is_active` tinyint(4) DEFAULT NULL,
  `start_time` varchar(255) DEFAULT NULL,
  `proc_inst_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `execution_proc_inst_id_proc_inst_id_foreign` (`proc_inst_id`),
  KEY `idx_id` (`id`),
  CONSTRAINT `execution_proc_inst_id_proc_inst_id_foreign` FOREIGN KEY (`proc_inst_id`) REFERENCES `proc_inst` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  goworkflow.execution 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `execution` DISABLE KEYS */;
/*!40000 ALTER TABLE `execution` ENABLE KEYS */;

-- 导出  表 goworkflow.execution_history 结构
CREATE TABLE IF NOT EXISTS `execution_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rev` int(11) DEFAULT NULL,
  `proc_def_id` int(11) DEFAULT NULL,
  `proc_def_name` varchar(255) DEFAULT NULL,
  `node_infos` varchar(4000) DEFAULT NULL,
  `is_active` tinyint(4) DEFAULT NULL,
  `start_time` varchar(255) DEFAULT NULL,
  `proc_inst_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `execution_history_proc_inst_id_proc_inst_history_id_foreign` (`proc_inst_id`),
  KEY `idx_id` (`id`),
  CONSTRAINT `execution_history_proc_inst_id_proc_inst_history_id_foreign` FOREIGN KEY (`proc_inst_id`) REFERENCES `proc_inst_history` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  goworkflow.execution_history 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `execution_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `execution_history` ENABLE KEYS */;

-- 导出  表 goworkflow.identitylink 结构
CREATE TABLE IF NOT EXISTS `identitylink` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `user_id` varchar(255) DEFAULT NULL,
  `user_name` varchar(255) DEFAULT NULL,
  `task_id` int(11) DEFAULT NULL,
  `step` int(11) DEFAULT NULL,
  `proc_inst_id` int(11) DEFAULT NULL,
  `company` varchar(255) DEFAULT NULL,
  `comment` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `identitylink_proc_inst_id_proc_inst_id_foreign` (`proc_inst_id`),
  KEY `idx_id` (`id`),
  CONSTRAINT `identitylink_proc_inst_id_proc_inst_id_foreign` FOREIGN KEY (`proc_inst_id`) REFERENCES `proc_inst` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  goworkflow.identitylink 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `identitylink` DISABLE KEYS */;
/*!40000 ALTER TABLE `identitylink` ENABLE KEYS */;

-- 导出  表 goworkflow.identitylink_history 结构
CREATE TABLE IF NOT EXISTS `identitylink_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `user_id` varchar(255) DEFAULT NULL,
  `user_name` varchar(255) DEFAULT NULL,
  `task_id` int(11) DEFAULT NULL,
  `step` int(11) DEFAULT NULL,
  `proc_inst_id` int(11) DEFAULT NULL,
  `company` varchar(255) DEFAULT NULL,
  `comment` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `identitylink_history_proc_inst_id_proc_inst_history_id_foreign` (`proc_inst_id`),
  KEY `idx_id` (`id`),
  CONSTRAINT `identitylink_history_proc_inst_id_proc_inst_history_id_foreign` FOREIGN KEY (`proc_inst_id`) REFERENCES `proc_inst_history` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  goworkflow.identitylink_history 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `identitylink_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `identitylink_history` ENABLE KEYS */;

-- 导出  表 goworkflow.procdef 结构
CREATE TABLE IF NOT EXISTS `procdef` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `resource` varchar(10000) DEFAULT NULL,
  `userid` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `company` varchar(255) DEFAULT NULL,
  `deploy_time` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  goworkflow.procdef 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `procdef` DISABLE KEYS */;
/*!40000 ALTER TABLE `procdef` ENABLE KEYS */;

-- 导出  表 goworkflow.procdef_history 结构
CREATE TABLE IF NOT EXISTS `procdef_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  `resource` varchar(10000) DEFAULT NULL,
  `userid` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `company` varchar(255) DEFAULT NULL,
  `deploy_time` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  goworkflow.procdef_history 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `procdef_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `procdef_history` ENABLE KEYS */;

-- 导出  表 goworkflow.proc_inst 结构
CREATE TABLE IF NOT EXISTS `proc_inst` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `proc_def_id` int(11) DEFAULT NULL,
  `proc_def_name` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `department` varchar(255) DEFAULT NULL,
  `company` varchar(255) DEFAULT NULL,
  `node_id` varchar(255) DEFAULT NULL,
  `candidate` varchar(255) DEFAULT NULL,
  `task_id` int(11) DEFAULT NULL,
  `start_time` varchar(255) DEFAULT NULL,
  `end_time` varchar(255) DEFAULT NULL,
  `duration` bigint(20) DEFAULT NULL,
  `start_user_id` varchar(255) DEFAULT NULL,
  `start_user_name` varchar(255) DEFAULT NULL,
  `is_finished` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  goworkflow.proc_inst 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `proc_inst` DISABLE KEYS */;
/*!40000 ALTER TABLE `proc_inst` ENABLE KEYS */;

-- 导出  表 goworkflow.proc_inst_history 结构
CREATE TABLE IF NOT EXISTS `proc_inst_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `proc_def_id` int(11) DEFAULT NULL,
  `proc_def_name` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `department` varchar(255) DEFAULT NULL,
  `company` varchar(255) DEFAULT NULL,
  `node_id` varchar(255) DEFAULT NULL,
  `candidate` varchar(255) DEFAULT NULL,
  `task_id` int(11) DEFAULT NULL,
  `start_time` varchar(255) DEFAULT NULL,
  `end_time` varchar(255) DEFAULT NULL,
  `duration` bigint(20) DEFAULT NULL,
  `start_user_id` varchar(255) DEFAULT NULL,
  `start_user_name` varchar(255) DEFAULT NULL,
  `is_finished` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  goworkflow.proc_inst_history 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `proc_inst_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `proc_inst_history` ENABLE KEYS */;

-- 导出  表 goworkflow.task 结构
CREATE TABLE IF NOT EXISTS `task` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `node_id` varchar(255) DEFAULT NULL,
  `step` int(11) DEFAULT NULL,
  `proc_inst_id` int(11) DEFAULT NULL,
  `assignee` varchar(255) DEFAULT NULL,
  `create_time` varchar(255) DEFAULT NULL,
  `claim_time` varchar(255) DEFAULT NULL,
  `member_count` tinyint(4) DEFAULT '1',
  `un_complete_num` tinyint(4) DEFAULT '1',
  `agree_num` tinyint(4) DEFAULT NULL,
  `act_type` varchar(255) DEFAULT 'or',
  `is_finished` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `task_proc_inst_id_proc_inst_id_foreign` (`proc_inst_id`),
  KEY `idx_id` (`id`),
  CONSTRAINT `task_proc_inst_id_proc_inst_id_foreign` FOREIGN KEY (`proc_inst_id`) REFERENCES `proc_inst` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  goworkflow.task 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `task` DISABLE KEYS */;
/*!40000 ALTER TABLE `task` ENABLE KEYS */;

-- 导出  表 goworkflow.task_history 结构
CREATE TABLE IF NOT EXISTS `task_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `node_id` varchar(255) DEFAULT NULL,
  `step` int(11) DEFAULT NULL,
  `proc_inst_id` int(11) DEFAULT NULL,
  `assignee` varchar(255) DEFAULT NULL,
  `create_time` varchar(255) DEFAULT NULL,
  `claim_time` varchar(255) DEFAULT NULL,
  `member_count` tinyint(4) DEFAULT '1',
  `un_complete_num` tinyint(4) DEFAULT '1',
  `agree_num` tinyint(4) DEFAULT NULL,
  `act_type` varchar(255) DEFAULT 'or',
  `is_finished` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  goworkflow.task_history 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `task_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `task_history` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
