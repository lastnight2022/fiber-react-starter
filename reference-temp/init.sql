-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        10.2.11-MariaDB - mariadb.org binary distribution
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 blc 的数据库结构
CREATE DATABASE IF NOT EXISTS `blc` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `blc`;

-- 导出  视图 blc.authority_menu 结构
-- 创建临时表以解决视图依赖性错误
CREATE TABLE `authority_menu` (
	`id` BIGINT(20) UNSIGNED NOT NULL,
	`path` VARCHAR(191) NULL COMMENT '路由path' COLLATE 'utf8mb4_general_ci',
	`icon` VARCHAR(191) NULL COMMENT '附加属性' COLLATE 'utf8mb4_general_ci',
	`name` VARCHAR(191) NULL COMMENT '路由name' COLLATE 'utf8mb4_general_ci',
	`sort` BIGINT(20) NULL COMMENT '排序标记',
	`title` VARCHAR(191) NULL COMMENT '附加属性' COLLATE 'utf8mb4_general_ci',
	`hidden` TINYINT(1) NULL COMMENT '是否在列表隐藏',
	`component` VARCHAR(191) NULL COMMENT '对应前端文件路径' COLLATE 'utf8mb4_general_ci',
	`parent_id` VARCHAR(191) NULL COMMENT '父菜单ID' COLLATE 'utf8mb4_general_ci',
	`created_at` DATETIME(3) NULL,
	`updated_at` DATETIME(3) NULL,
	`deleted_at` DATETIME(3) NULL,
	`keep_alive` TINYINT(1) NULL COMMENT '附加属性',
	`menu_level` BIGINT(20) UNSIGNED NULL,
	`default_menu` TINYINT(1) NULL COMMENT '附加属性',
	`close_tab` TINYINT(1) NULL COMMENT '附加属性',
	`menu_id` BIGINT(20) UNSIGNED NOT NULL,
	`authority_id` VARCHAR(90) NOT NULL COMMENT '角色ID' COLLATE 'utf8mb4_general_ci'
) ENGINE=MyISAM;

-- 导出  表 blc.auto_code_examples 结构
CREATE TABLE IF NOT EXISTS `auto_code_examples` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `auto_code_example_field` varchar(191) DEFAULT NULL COMMENT '仅作示例条目无实际作用',
  PRIMARY KEY (`id`),
  KEY `idx_auto_code_examples_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.auto_code_examples 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `auto_code_examples` DISABLE KEYS */;
/*!40000 ALTER TABLE `auto_code_examples` ENABLE KEYS */;

-- 导出  表 blc.casbin_rule 结构
CREATE TABLE IF NOT EXISTS `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.casbin_rule 的数据：~164 rows (大约)
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
	('p', '888', '/base/login', 'POST', '', '', ''),
	('p', '888', '/user/register', 'POST', '', '', ''),
	('p', '888', '/api/createApi', 'POST', '', '', ''),
	('p', '888', '/api/getApiList', 'POST', '', '', ''),
	('p', '888', '/api/getApiById', 'POST', '', '', ''),
	('p', '888', '/api/deleteApi', 'POST', '', '', ''),
	('p', '888', '/api/updateApi', 'POST', '', '', ''),
	('p', '888', '/api/getAllApis', 'POST', '', '', ''),
	('p', '888', '/api/deleteApisByIds', 'DELETE', '', '', ''),
	('p', '888', '/authority/copyAuthority', 'POST', '', '', ''),
	('p', '888', '/authority/updateAuthority', 'PUT', '', '', ''),
	('p', '888', '/authority/createAuthority', 'POST', '', '', ''),
	('p', '888', '/authority/deleteAuthority', 'POST', '', '', ''),
	('p', '888', '/authority/getAuthorityList', 'POST', '', '', ''),
	('p', '888', '/authority/setDataAuthority', 'POST', '', '', ''),
	('p', '888', '/menu/getMenu', 'POST', '', '', ''),
	('p', '888', '/menu/getMenuList', 'POST', '', '', ''),
	('p', '888', '/menu/addBaseMenu', 'POST', '', '', ''),
	('p', '888', '/menu/getBaseMenuTree', 'POST', '', '', ''),
	('p', '888', '/menu/addMenuAuthority', 'POST', '', '', ''),
	('p', '888', '/menu/getMenuAuthority', 'POST', '', '', ''),
	('p', '888', '/menu/deleteBaseMenu', 'POST', '', '', ''),
	('p', '888', '/menu/updateBaseMenu', 'POST', '', '', ''),
	('p', '888', '/menu/getBaseMenuById', 'POST', '', '', ''),
	('p', '888', '/user/getUserInfo', 'GET', '', '', ''),
	('p', '888', '/user/setUserInfo', 'PUT', '', '', ''),
	('p', '888', '/user/setSelfInfo', 'PUT', '', '', ''),
	('p', '888', '/user/getUserList', 'POST', '', '', ''),
	('p', '888', '/user/deleteUser', 'DELETE', '', '', ''),
	('p', '888', '/user/changePassword', 'POST', '', '', ''),
	('p', '888', '/user/setUserAuthority', 'POST', '', '', ''),
	('p', '888', '/user/setUserAuthorities', 'POST', '', '', ''),
	('p', '888', '/user/resetPassword', 'POST', '', '', ''),
	('p', '888', '/fileUploadAndDownload/findFile', 'GET', '', '', ''),
	('p', '888', '/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', ''),
	('p', '888', '/fileUploadAndDownload/breakpointContinue', 'POST', '', '', ''),
	('p', '888', '/fileUploadAndDownload/removeChunk', 'POST', '', '', ''),
	('p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', ''),
	('p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', ''),
	('p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', ''),
	('p', '888', '/casbin/updateCasbin', 'POST', '', '', ''),
	('p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', ''),
	('p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', ''),
	('p', '888', '/system/getSystemConfig', 'POST', '', '', ''),
	('p', '888', '/system/setSystemConfig', 'POST', '', '', ''),
	('p', '888', '/system/getServerInfo', 'POST', '', '', ''),
	('p', '888', '/customer/customer', 'GET', '', '', ''),
	('p', '888', '/customer/customer', 'PUT', '', '', ''),
	('p', '888', '/customer/customer', 'POST', '', '', ''),
	('p', '888', '/customer/customer', 'DELETE', '', '', ''),
	('p', '888', '/customer/customerList', 'GET', '', '', ''),
	('p', '888', '/autoCode/getDB', 'GET', '', '', ''),
	('p', '888', '/autoCode/getMeta', 'POST', '', '', ''),
	('p', '888', '/autoCode/preview', 'POST', '', '', ''),
	('p', '888', '/autoCode/getTables', 'GET', '', '', ''),
	('p', '888', '/autoCode/getColumn', 'GET', '', '', ''),
	('p', '888', '/autoCode/rollback', 'POST', '', '', ''),
	('p', '888', '/autoCode/createTemp', 'POST', '', '', ''),
	('p', '888', '/autoCode/delSysHistory', 'POST', '', '', ''),
	('p', '888', '/autoCode/getSysHistory', 'POST', '', '', ''),
	('p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', ''),
	('p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', ''),
	('p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', ''),
	('p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', ''),
	('p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', ''),
	('p', '888', '/sysDictionary/findSysDictionary', 'GET', '', '', ''),
	('p', '888', '/sysDictionary/updateSysDictionary', 'PUT', '', '', ''),
	('p', '888', '/sysDictionary/getSysDictionaryList', 'GET', '', '', ''),
	('p', '888', '/sysDictionary/createSysDictionary', 'POST', '', '', ''),
	('p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', ''),
	('p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', ''),
	('p', '888', '/sysOperationRecord/updateSysOperationRecord', 'PUT', '', '', ''),
	('p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', ''),
	('p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', ''),
	('p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', ''),
	('p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', ''),
	('p', '888', '/email/emailTest', 'POST', '', '', ''),
	('p', '888', '/simpleUploader/upload', 'POST', '', '', ''),
	('p', '888', '/simpleUploader/checkFileMd5', 'GET', '', '', ''),
	('p', '888', '/simpleUploader/mergeFileMd5', 'GET', '', '', ''),
	('p', '888', '/excel/importExcel', 'POST', '', '', ''),
	('p', '888', '/excel/loadExcel', 'GET', '', '', ''),
	('p', '888', '/excel/exportExcel', 'POST', '', '', ''),
	('p', '888', '/excel/downloadTemplate', 'GET', '', '', ''),
	('p', '888', '/authorityBtn/setAuthorityBtn', 'POST', '', '', ''),
	('p', '888', '/authorityBtn/getAuthorityBtn', 'POST', '', '', ''),
	('p', '888', '/authorityBtn/canRemoveAuthorityBtn', 'POST', '', '', ''),
	('p', '8881', '/base/login', 'POST', '', '', ''),
	('p', '8881', '/user/register', 'POST', '', '', ''),
	('p', '8881', '/api/createApi', 'POST', '', '', ''),
	('p', '8881', '/api/getApiList', 'POST', '', '', ''),
	('p', '8881', '/api/getApiById', 'POST', '', '', ''),
	('p', '8881', '/api/deleteApi', 'POST', '', '', ''),
	('p', '8881', '/api/updateApi', 'POST', '', '', ''),
	('p', '8881', '/api/getAllApis', 'POST', '', '', ''),
	('p', '8881', '/authority/createAuthority', 'POST', '', '', ''),
	('p', '8881', '/authority/deleteAuthority', 'POST', '', '', ''),
	('p', '8881', '/authority/getAuthorityList', 'POST', '', '', ''),
	('p', '8881', '/authority/setDataAuthority', 'POST', '', '', ''),
	('p', '8881', '/menu/getMenu', 'POST', '', '', ''),
	('p', '8881', '/menu/getMenuList', 'POST', '', '', ''),
	('p', '8881', '/menu/addBaseMenu', 'POST', '', '', ''),
	('p', '8881', '/menu/getBaseMenuTree', 'POST', '', '', ''),
	('p', '8881', '/menu/addMenuAuthority', 'POST', '', '', ''),
	('p', '8881', '/menu/getMenuAuthority', 'POST', '', '', ''),
	('p', '8881', '/menu/deleteBaseMenu', 'POST', '', '', ''),
	('p', '8881', '/menu/updateBaseMenu', 'POST', '', '', ''),
	('p', '8881', '/menu/getBaseMenuById', 'POST', '', '', ''),
	('p', '8881', '/user/changePassword', 'POST', '', '', ''),
	('p', '8881', '/user/getUserList', 'POST', '', '', ''),
	('p', '8881', '/user/setUserAuthority', 'POST', '', '', ''),
	('p', '8881', '/fileUploadAndDownload/upload', 'POST', '', '', ''),
	('p', '8881', '/fileUploadAndDownload/getFileList', 'POST', '', '', ''),
	('p', '8881', '/fileUploadAndDownload/deleteFile', 'POST', '', '', ''),
	('p', '8881', '/casbin/updateCasbin', 'POST', '', '', ''),
	('p', '8881', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', ''),
	('p', '8881', '/jwt/jsonInBlacklist', 'POST', '', '', ''),
	('p', '8881', '/system/getSystemConfig', 'POST', '', '', ''),
	('p', '8881', '/system/setSystemConfig', 'POST', '', '', ''),
	('p', '8881', '/customer/customer', 'POST', '', '', ''),
	('p', '8881', '/customer/customer', 'PUT', '', '', ''),
	('p', '8881', '/customer/customer', 'DELETE', '', '', ''),
	('p', '8881', '/customer/customer', 'GET', '', '', ''),
	('p', '8881', '/customer/customerList', 'GET', '', '', ''),
	('p', '8881', '/user/getUserInfo', 'GET', '', '', ''),
	('p', '9528', '/base/login', 'POST', '', '', ''),
	('p', '9528', '/user/register', 'POST', '', '', ''),
	('p', '9528', '/api/createApi', 'POST', '', '', ''),
	('p', '9528', '/api/getApiList', 'POST', '', '', ''),
	('p', '9528', '/api/getApiById', 'POST', '', '', ''),
	('p', '9528', '/api/deleteApi', 'POST', '', '', ''),
	('p', '9528', '/api/updateApi', 'POST', '', '', ''),
	('p', '9528', '/api/getAllApis', 'POST', '', '', ''),
	('p', '9528', '/authority/createAuthority', 'POST', '', '', ''),
	('p', '9528', '/authority/deleteAuthority', 'POST', '', '', ''),
	('p', '9528', '/authority/getAuthorityList', 'POST', '', '', ''),
	('p', '9528', '/authority/setDataAuthority', 'POST', '', '', ''),
	('p', '9528', '/menu/getMenu', 'POST', '', '', ''),
	('p', '9528', '/menu/getMenuList', 'POST', '', '', ''),
	('p', '9528', '/menu/addBaseMenu', 'POST', '', '', ''),
	('p', '9528', '/menu/getBaseMenuTree', 'POST', '', '', ''),
	('p', '9528', '/menu/addMenuAuthority', 'POST', '', '', ''),
	('p', '9528', '/menu/getMenuAuthority', 'POST', '', '', ''),
	('p', '9528', '/menu/deleteBaseMenu', 'POST', '', '', ''),
	('p', '9528', '/menu/updateBaseMenu', 'POST', '', '', ''),
	('p', '9528', '/menu/getBaseMenuById', 'POST', '', '', ''),
	('p', '9528', '/user/changePassword', 'POST', '', '', ''),
	('p', '9528', '/user/getUserList', 'POST', '', '', ''),
	('p', '9528', '/user/setUserAuthority', 'POST', '', '', ''),
	('p', '9528', '/fileUploadAndDownload/upload', 'POST', '', '', ''),
	('p', '9528', '/fileUploadAndDownload/getFileList', 'POST', '', '', ''),
	('p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', '', '', ''),
	('p', '9528', '/casbin/updateCasbin', 'POST', '', '', ''),
	('p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', ''),
	('p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', ''),
	('p', '9528', '/system/getSystemConfig', 'POST', '', '', ''),
	('p', '9528', '/system/setSystemConfig', 'POST', '', '', ''),
	('p', '9528', '/customer/customer', 'PUT', '', '', ''),
	('p', '9528', '/customer/customer', 'GET', '', '', ''),
	('p', '9528', '/customer/customer', 'POST', '', '', ''),
	('p', '9528', '/customer/customer', 'DELETE', '', '', ''),
	('p', '9528', '/customer/customerList', 'GET', '', '', ''),
	('p', '9528', '/autoCode/createTemp', 'POST', '', '', ''),
	('p', '9528', '/user/getUserInfo', 'GET', '', '', '');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;

-- 导出  表 blc.exa_customers 结构
CREATE TABLE IF NOT EXISTS `exa_customers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_name` varchar(191) DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint(20) unsigned DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` varchar(191) DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.exa_customers 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `exa_customers` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_customers` ENABLE KEYS */;

-- 导出  表 blc.exa_files 结构
CREATE TABLE IF NOT EXISTS `exa_files` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `file_name` varchar(191) DEFAULT NULL,
  `file_md5` varchar(191) DEFAULT NULL,
  `file_path` varchar(191) DEFAULT NULL,
  `chunk_total` bigint(20) DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_files_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.exa_files 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `exa_files` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_files` ENABLE KEYS */;

-- 导出  表 blc.exa_file_chunks 结构
CREATE TABLE IF NOT EXISTS `exa_file_chunks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `exa_file_id` bigint(20) unsigned DEFAULT NULL,
  `file_chunk_number` bigint(20) DEFAULT NULL,
  `file_chunk_path` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.exa_file_chunks 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `exa_file_chunks` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_file_chunks` ENABLE KEYS */;

-- 导出  表 blc.exa_file_upload_and_downloads 结构
CREATE TABLE IF NOT EXISTS `exa_file_upload_and_downloads` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.exa_file_upload_and_downloads 的数据：~2 rows (大约)
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` DISABLE KEYS */;
INSERT INTO `exa_file_upload_and_downloads` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `url`, `tag`, `key`) VALUES
	(1, '2023-07-05 16:57:10.018', '2023-07-05 16:57:10.018', NULL, '10.png', 'https://qmplusimg.henrongyi.top/gvalogo.png', 'png', '158787308910.png'),
	(2, '2023-07-05 16:57:10.018', '2023-07-05 16:57:10.018', NULL, 'logo.png', 'https://qmplusimg.henrongyi.top/1576554439myAvatar.png', 'png', '1587973709logo.png');
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` ENABLE KEYS */;

-- 导出  表 blc.jwt_blacklists 结构
CREATE TABLE IF NOT EXISTS `jwt_blacklists` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text DEFAULT NULL COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.jwt_blacklists 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `jwt_blacklists` DISABLE KEYS */;
/*!40000 ALTER TABLE `jwt_blacklists` ENABLE KEYS */;

-- 导出  表 blc.sys_apis 结构
CREATE TABLE IF NOT EXISTS `sys_apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=88 DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_apis 的数据：~87 rows (大约)
/*!40000 ALTER TABLE `sys_apis` DISABLE KEYS */;
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES
	(1, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/base/login', '用户登录(必选)', 'base', 'POST'),
	(2, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwt', 'POST'),
	(3, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/user/deleteUser', '删除用户', '系统用户', 'DELETE'),
	(4, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/user/register', '用户注册', '系统用户', 'POST'),
	(5, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/user/getUserList', '获取用户列表', '系统用户', 'POST'),
	(6, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/user/setUserInfo', '设置用户信息', '系统用户', 'PUT'),
	(7, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/user/setSelfInfo', '设置自身信息(必选)', '系统用户', 'PUT'),
	(8, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/user/getUserInfo', '获取自身信息(必选)', '系统用户', 'GET'),
	(9, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/user/setUserAuthorities', '设置权限组', '系统用户', 'POST'),
	(10, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/user/changePassword', '修改密码（建议选择)', '系统用户', 'POST'),
	(11, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/user/setUserAuthority', '修改用户角色(必选)', '系统用户', 'POST'),
	(12, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/user/resetPassword', '重置用户密码', '系统用户', 'POST'),
	(13, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/api/createApi', '创建api', 'api', 'POST'),
	(14, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/api/deleteApi', '删除Api', 'api', 'POST'),
	(15, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/api/updateApi', '更新Api', 'api', 'POST'),
	(16, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/api/getApiList', '获取api列表', 'api', 'POST'),
	(17, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/api/getAllApis', '获取所有api', 'api', 'POST'),
	(18, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST'),
	(19, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/api/deleteApisByIds', '批量删除api', 'api', 'DELETE'),
	(20, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/authority/copyAuthority', '拷贝角色', '角色', 'POST'),
	(21, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/authority/createAuthority', '创建角色', '角色', 'POST'),
	(22, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/authority/deleteAuthority', '删除角色', '角色', 'POST'),
	(23, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/authority/updateAuthority', '更新角色信息', '角色', 'PUT'),
	(24, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/authority/getAuthorityList', '获取角色列表', '角色', 'POST'),
	(25, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/authority/setDataAuthority', '设置角色资源权限', '角色', 'POST'),
	(26, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST'),
	(27, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST'),
	(28, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/menu/addBaseMenu', '新增菜单', '菜单', 'POST'),
	(29, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/menu/getMenu', '获取菜单树(必选)', '菜单', 'POST'),
	(30, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/menu/deleteBaseMenu', '删除菜单', '菜单', 'POST'),
	(31, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/menu/updateBaseMenu', '更新菜单', '菜单', 'POST'),
	(32, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/menu/getBaseMenuById', '根据id获取菜单', '菜单', 'POST'),
	(33, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/menu/getMenuList', '分页获取基础menu列表', '菜单', 'POST'),
	(34, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', '菜单', 'POST'),
	(35, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/menu/getMenuAuthority', '获取指定角色menu', '菜单', 'POST'),
	(36, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', '菜单', 'POST'),
	(37, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/fileUploadAndDownload/findFile', '寻找目标文件（秒传）', '分片上传', 'POST'),
	(38, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/fileUploadAndDownload/breakpointContinue', '断点续传', '分片上传', 'POST'),
	(39, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/fileUploadAndDownload/breakpointContinueFinish', '断点续传完成', '分片上传', 'POST'),
	(40, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/fileUploadAndDownload/removeChunk', '上传完成移除文件', '分片上传', 'POST'),
	(41, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/fileUploadAndDownload/upload', '文件上传示例', '文件上传与下载', 'POST'),
	(42, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/fileUploadAndDownload/deleteFile', '删除文件', '文件上传与下载', 'POST'),
	(43, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', '文件上传与下载', 'POST'),
	(44, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/system/getServerInfo', '获取服务器信息', '系统服务', 'POST'),
	(45, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/system/getSystemConfig', '获取配置文件内容', '系统服务', 'POST'),
	(46, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/system/setSystemConfig', '设置配置文件内容', '系统服务', 'POST'),
	(47, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/customer/customer', '更新客户', '客户', 'PUT'),
	(48, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/customer/customer', '创建客户', '客户', 'POST'),
	(49, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/customer/customer', '删除客户', '客户', 'DELETE'),
	(50, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/customer/customer', '获取单一客户', '客户', 'GET'),
	(51, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/customer/customerList', '获取客户列表', '客户', 'GET'),
	(52, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/autoCode/getDB', '获取所有数据库', '代码生成器', 'GET'),
	(53, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/autoCode/getTables', '获取数据库表', '代码生成器', 'GET'),
	(54, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/autoCode/createTemp', '自动化代码', '代码生成器', 'POST'),
	(55, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/autoCode/preview', '预览自动化代码', '代码生成器', 'POST'),
	(56, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/autoCode/getColumn', '获取所选table的所有字段', '代码生成器', 'GET'),
	(57, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/autoCode/getMeta', '获取meta信息', '代码生成器历史', 'POST'),
	(58, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/autoCode/rollback', '回滚自动生成代码', '代码生成器历史', 'POST'),
	(59, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/autoCode/getSysHistory', '查询回滚记录', '代码生成器历史', 'POST'),
	(60, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/autoCode/delSysHistory', '删除回滚记录', '代码生成器历史', 'POST'),
	(61, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', '系统字典详情', 'PUT'),
	(62, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', '系统字典详情', 'POST'),
	(63, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', '系统字典详情', 'DELETE'),
	(64, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', '系统字典详情', 'GET'),
	(65, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', '系统字典详情', 'GET'),
	(66, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysDictionary/createSysDictionary', '新增字典', '系统字典', 'POST'),
	(67, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysDictionary/deleteSysDictionary', '删除字典', '系统字典', 'DELETE'),
	(68, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysDictionary/updateSysDictionary', '更新字典', '系统字典', 'PUT'),
	(69, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典', '系统字典', 'GET'),
	(70, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', '系统字典', 'GET'),
	(71, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', '操作记录', 'POST'),
	(72, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', '操作记录', 'GET'),
	(73, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', '操作记录', 'GET'),
	(74, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', '操作记录', 'DELETE'),
	(75, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', '操作记录', 'DELETE'),
	(76, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/simpleUploader/upload', '插件版分片上传', '断点续传(插件版)', 'POST'),
	(77, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/simpleUploader/checkFileMd5', '文件完整度验证', '断点续传(插件版)', 'GET'),
	(78, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/simpleUploader/mergeFileMd5', '上传完成合并文件', '断点续传(插件版)', 'GET'),
	(79, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/email/emailTest', '发送测试邮件', 'email', 'POST'),
	(80, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/email/emailSend', '发送邮件示例', 'email', 'POST'),
	(81, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/excel/importExcel', '导入excel', 'excel', 'POST'),
	(82, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/excel/loadExcel', '下载excel', 'excel', 'GET'),
	(83, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/excel/exportExcel', '导出excel', 'excel', 'POST'),
	(84, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/excel/downloadTemplate', '下载excel模板', 'excel', 'GET'),
	(85, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/authorityBtn/setAuthorityBtn', '设置按钮权限', '按钮权限', 'POST'),
	(86, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/authorityBtn/getAuthorityBtn', '获取已有按钮权限', '按钮权限', 'POST'),
	(87, '2023-07-05 16:57:09.970', '2023-07-05 16:57:09.970', NULL, '/authorityBtn/canRemoveAuthorityBtn', '删除按钮', '按钮权限', 'POST');
/*!40000 ALTER TABLE `sys_apis` ENABLE KEYS */;

-- 导出  表 blc.sys_authorities 结构
CREATE TABLE IF NOT EXISTS `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  `authority_name` varchar(191) DEFAULT NULL COMMENT '角色名',
  `parent_id` varchar(191) DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `authority_id` (`authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_authorities 的数据：~3 rows (大约)
/*!40000 ALTER TABLE `sys_authorities` DISABLE KEYS */;
INSERT INTO `sys_authorities` (`created_at`, `updated_at`, `deleted_at`, `authority_id`, `authority_name`, `parent_id`, `default_router`) VALUES
	('2023-07-05 16:57:09.990', '2023-07-05 16:57:09.990', NULL, '888', '普通用户', '0', 'dashboard'),
	('2023-07-05 16:57:09.990', '2023-07-05 16:57:09.990', NULL, '8881', '普通用户子角色', '888', 'dashboard'),
	('2023-07-05 16:57:09.990', '2023-07-05 16:57:09.990', NULL, '9528', '测试角色', '0', 'dashboard');
/*!40000 ALTER TABLE `sys_authorities` ENABLE KEYS */;

-- 导出  表 blc.sys_authority_btns 结构
CREATE TABLE IF NOT EXISTS `sys_authority_btns` (
  `authority_id` varchar(191) DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_authority_btns 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `sys_authority_btns` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_authority_btns` ENABLE KEYS */;

-- 导出  表 blc.sys_authority_menus 结构
CREATE TABLE IF NOT EXISTS `sys_authority_menus` (
  `sys_base_menu_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_authority_menus 的数据：~43 rows (大约)
/*!40000 ALTER TABLE `sys_authority_menus` DISABLE KEYS */;
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES
	(1, '888'),
	(1, '8881'),
	(1, '9528'),
	(2, '888'),
	(2, '8881'),
	(2, '9528'),
	(3, '888'),
	(3, '9528'),
	(4, '888'),
	(4, '9528'),
	(5, '888'),
	(5, '9528'),
	(6, '888'),
	(6, '9528'),
	(7, '888'),
	(7, '9528'),
	(8, '888'),
	(8, '8881'),
	(8, '9528'),
	(9, '888'),
	(9, '9528'),
	(10, '888'),
	(10, '9528'),
	(11, '888'),
	(11, '9528'),
	(12, '888'),
	(12, '9528'),
	(13, '888'),
	(14, '888'),
	(14, '9528'),
	(15, '888'),
	(15, '9528'),
	(16, '888'),
	(16, '9528'),
	(17, '888'),
	(17, '9528'),
	(18, '888'),
	(19, '888'),
	(20, '888'),
	(22, '888'),
	(23, '888'),
	(24, '888'),
	(25, '888');
/*!40000 ALTER TABLE `sys_authority_menus` ENABLE KEYS */;

-- 导出  表 blc.sys_auto_code_histories 结构
CREATE TABLE IF NOT EXISTS `sys_auto_code_histories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `table_name` varchar(191) DEFAULT NULL,
  `request_meta` text DEFAULT NULL,
  `auto_code_path` text DEFAULT NULL,
  `injection_meta` text DEFAULT NULL,
  `struct_name` varchar(191) DEFAULT NULL,
  `struct_cn_name` varchar(191) DEFAULT NULL,
  `api_ids` varchar(191) DEFAULT NULL,
  `flag` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_histories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_auto_code_histories 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `sys_auto_code_histories` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_auto_code_histories` ENABLE KEYS */;

-- 导出  表 blc.sys_base_menus 结构
CREATE TABLE IF NOT EXISTS `sys_base_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint(20) unsigned DEFAULT NULL,
  `parent_id` varchar(191) DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_base_menus 的数据：~25 rows (大约)
/*!40000 ALTER TABLE `sys_base_menus` DISABLE KEYS */;
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES
	(1, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '0', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, 0, 0, '仪表盘', 'odometer', 0),
	(2, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '0', 'about', 'about', 0, 'view/about/index.vue', 7, 0, 0, '关于我们', 'info-filled', 0),
	(3, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '0', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, 0, 0, '超级管理员', 'user', 0),
	(4, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '3', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, 0, 0, '角色管理', 'avatar', 0),
	(5, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '3', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, 1, 0, '菜单管理', 'tickets', 0),
	(6, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '3', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, 1, 0, 'api管理', 'platform', 0),
	(7, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '3', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, 0, 0, '用户管理', 'coordinate', 0),
	(8, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '0', 'person', 'person', 1, 'view/person/person.vue', 4, 0, 0, '个人信息', 'message', 0),
	(9, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '0', 'example', 'example', 0, 'view/example/index.vue', 6, 0, 0, '示例文件', 'management', 0),
	(10, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '9', 'excel', 'excel', 0, 'view/example/excel/excel.vue', 4, 0, 0, 'excel导入导出', 'takeaway-box', 0),
	(11, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '9', 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, 0, 0, '媒体库（上传下载）', 'upload', 0),
	(12, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '9', 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, 0, 0, '断点续传', 'upload-filled', 0),
	(13, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '9', 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, 0, 0, '客户列表（资源示例）', 'avatar', 0),
	(14, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '0', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, 0, 0, '系统工具', 'tools', 0),
	(15, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '14', 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, 1, 0, '代码生成器', 'cpu', 0),
	(16, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '14', 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, 1, 0, '表单生成器', 'magic-stick', 0),
	(17, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '14', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, 0, 0, '系统配置', 'operation', 0),
	(18, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '3', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, 0, 0, '字典管理', 'notebook', 0),
	(19, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '3', 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/sysDictionaryDetail.vue', 1, 0, 0, '字典详情', 'order', 0),
	(20, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '3', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, 0, 0, '操作历史', 'pie-chart', 0),
	(21, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '9', 'simpleUploader', 'simpleUploader', 0, 'view/example/simpleUploader/simpleUploader', 6, 0, 0, '断点续传（插件版）', 'upload', 0),
	(22, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '0', 'https://www.gin-vue-admin.com', 'https://www.gin-vue-admin.com', 0, '/', 0, 0, 0, '官方网站', 'home-filled', 0),
	(23, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '0', 'state', 'state', 0, 'view/system/state.vue', 6, 0, 0, '服务器状态', 'cloudy', 0),
	(24, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '14', 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 1, 0, 0, '自动化代码管理', 'magic-stick', 0),
	(25, '2023-07-05 16:57:09.987', '2023-07-05 16:57:09.987', NULL, 0, '14', 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, 0, 0, '自动化代码（复用）', 'magic-stick', 0);
/*!40000 ALTER TABLE `sys_base_menus` ENABLE KEYS */;

-- 导出  表 blc.sys_base_menu_btns 结构
CREATE TABLE IF NOT EXISTS `sys_base_menu_btns` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_btns_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_base_menu_btns 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `sys_base_menu_btns` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_base_menu_btns` ENABLE KEYS */;

-- 导出  表 blc.sys_base_menu_parameters 结构
CREATE TABLE IF NOT EXISTS `sys_base_menu_parameters` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL,
  `type` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_base_menu_parameters 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `sys_base_menu_parameters` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_base_menu_parameters` ENABLE KEYS */;

-- 导出  表 blc.sys_data_authority_id 结构
CREATE TABLE IF NOT EXISTS `sys_data_authority_id` (
  `sys_authority_authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_data_authority_id 的数据：~5 rows (大约)
/*!40000 ALTER TABLE `sys_data_authority_id` DISABLE KEYS */;
INSERT INTO `sys_data_authority_id` (`sys_authority_authority_id`, `data_authority_id_authority_id`) VALUES
	('888', '888'),
	('888', '8881'),
	('888', '9528'),
	('9528', '8881'),
	('9528', '9528');
/*!40000 ALTER TABLE `sys_data_authority_id` ENABLE KEYS */;

-- 导出  表 blc.sys_dictionaries 结构
CREATE TABLE IF NOT EXISTS `sys_dictionaries` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_dictionaries 的数据：~6 rows (大约)
/*!40000 ALTER TABLE `sys_dictionaries` DISABLE KEYS */;
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES
	(1, '2023-07-05 16:57:09.995', '2023-07-05 16:57:09.995', NULL, '性别', 'gender', 1, '性别字典'),
	(2, '2023-07-05 16:57:09.995', '2023-07-05 16:57:09.995', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型'),
	(3, '2023-07-05 16:57:09.995', '2023-07-05 16:57:09.995', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型'),
	(4, '2023-07-05 16:57:09.995', '2023-07-05 16:57:09.995', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型'),
	(5, '2023-07-05 16:57:09.995', '2023-07-05 16:57:09.995', NULL, '数据库字符串', 'string', 1, '数据库字符串'),
	(6, '2023-07-05 16:57:09.995', '2023-07-05 16:57:09.995', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');
/*!40000 ALTER TABLE `sys_dictionaries` ENABLE KEYS */;

-- 导出  表 blc.sys_dictionary_details 结构
CREATE TABLE IF NOT EXISTS `sys_dictionary_details` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) DEFAULT NULL COMMENT '展示值',
  `value` bigint(20) DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint(20) unsigned DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_dictionary_details 的数据：~25 rows (大约)
/*!40000 ALTER TABLE `sys_dictionary_details` DISABLE KEYS */;
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES
	(1, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, '男', 1, 1, 1, 1),
	(2, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, '女', 2, 1, 2, 1),
	(3, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'smallint', 1, 1, 1, 2),
	(4, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'mediumint', 2, 1, 2, 2),
	(5, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'int', 3, 1, 3, 2),
	(6, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'bigint', 4, 1, 4, 2),
	(7, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'date', 0, 1, 0, 3),
	(8, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'time', 1, 1, 1, 3),
	(9, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'year', 2, 1, 2, 3),
	(10, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'datetime', 3, 1, 3, 3),
	(11, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'timestamp', 5, 1, 5, 3),
	(12, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'float', 0, 1, 0, 4),
	(13, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'double', 1, 1, 1, 4),
	(14, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'decimal', 2, 1, 2, 4),
	(15, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'char', 0, 1, 0, 5),
	(16, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'varchar', 1, 1, 1, 5),
	(17, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'tinyblob', 2, 1, 2, 5),
	(18, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'tinytext', 3, 1, 3, 5),
	(19, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'text', 4, 1, 4, 5),
	(20, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'blob', 5, 1, 5, 5),
	(21, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'mediumblob', 6, 1, 6, 5),
	(22, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'mediumtext', 7, 1, 7, 5),
	(23, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'longblob', 8, 1, 8, 5),
	(24, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'longtext', 9, 1, 9, 5),
	(25, '2023-07-05 16:57:10.010', '2023-07-05 16:57:10.010', NULL, 'tinyint', 0, 1, 0, 6);
/*!40000 ALTER TABLE `sys_dictionary_details` ENABLE KEYS */;

-- 导出  表 blc.sys_operation_records 结构
CREATE TABLE IF NOT EXISTS `sys_operation_records` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) DEFAULT NULL COMMENT '请求路径',
  `status` bigint(20) DEFAULT NULL COMMENT '请求状态',
  `latency` bigint(20) DEFAULT NULL COMMENT '延迟',
  `agent` varchar(191) DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) DEFAULT NULL COMMENT '错误信息',
  `body` text DEFAULT NULL COMMENT '请求Body',
  `resp` text DEFAULT NULL COMMENT '响应Body',
  `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_operation_records 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `sys_operation_records` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_operation_records` ENABLE KEYS */;

-- 导出  表 blc.sys_users 结构
CREATE TABLE IF NOT EXISTS `sys_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `base_color` varchar(191) DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(191) DEFAULT '#1890ff' COMMENT '活跃颜色',
  `authority_id` varchar(191) DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) DEFAULT NULL COMMENT '用户邮箱',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_users 的数据：~2 rows (大约)
/*!40000 ALTER TABLE `sys_users` DISABLE KEYS */;
INSERT INTO `sys_users` (`id`, `created_at`, `updated_at`, `deleted_at`, `uuid`, `username`, `password`, `nick_name`, `side_mode`, `header_img`, `base_color`, `active_color`, `authority_id`, `phone`, `email`) VALUES
	(1, '2023-07-05 16:57:09.976', '2023-07-05 16:57:09.976', NULL, '09ff30dc-0aac-4c01-9776-f6e4614f23bb', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', '888', '17611111111', '333333333@qq.com'),
	(2, '2023-07-05 16:57:09.976', '2023-07-05 16:57:09.976', NULL, '4aeb916a-4777-4bc4-bb39-935cc4ad1327', 'a303176530', '3ec063004a6f31642261936a379fde3d', 'QMPlusUser', 'dark', 'https:///qmplusimg.henrongyi.top/1572075907logo.png', '#fff', '#1890ff', '9528', '17611111111', '333333333@qq.com');
/*!40000 ALTER TABLE `sys_users` ENABLE KEYS */;

-- 导出  表 blc.sys_user_authority 结构
CREATE TABLE IF NOT EXISTS `sys_user_authority` (
  `sys_user_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` varchar(90) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  blc.sys_user_authority 的数据：~4 rows (大约)
/*!40000 ALTER TABLE `sys_user_authority` DISABLE KEYS */;
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES
	(1, '888'),
	(1, '8881'),
	(1, '9528'),
	(2, '888');
/*!40000 ALTER TABLE `sys_user_authority` ENABLE KEYS */;

-- 导出  视图 blc.authority_menu 结构
-- 移除临时表并创建最终视图结构
DROP TABLE IF EXISTS `authority_menu`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` VIEW `authority_menu` AS select sys_base_menus.id                AS id,
		   sys_base_menus.path              AS path,
		   sys_base_menus.icon              AS icon,
		   sys_base_menus.name              AS name,
		   sys_base_menus.sort              AS sort,
		   sys_base_menus.title             AS title,
		   sys_base_menus.hidden            AS hidden,
		   sys_base_menus.component         AS component,
		   sys_base_menus.parent_id         AS parent_id,
		   sys_base_menus.created_at        AS created_at,
		   sys_base_menus.updated_at        AS updated_at,
		   sys_base_menus.deleted_at        AS deleted_at,
		   sys_base_menus.keep_alive        AS keep_alive,
		   sys_base_menus.menu_level        AS menu_level,
		   sys_base_menus.default_menu      AS default_menu,
		   sys_base_menus.close_tab      	AS close_tab,
		   sys_authority_menus.sys_base_menu_id      AS menu_id,
		   sys_authority_menus.sys_authority_authority_id AS authority_id
	from (sys_authority_menus
			 join sys_base_menus on ((sys_authority_menus.sys_base_menu_id = sys_base_menus.id))) ;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
