/*
 Navicat Premium Data Transfer

 Target Server Type    : MySQL
 Target Server Version : 80034 (8.0.34)
 File Encoding         : 65001

 Date: 06/04/2024 19:35:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '应用名称',
  `key` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '应用KEY 别名',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态',
  `sort` int NOT NULL DEFAULT 0 COMMENT '排序 ASC',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '应用列表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of app
-- ----------------------------
INSERT INTO `app` VALUES (1, '系统管理', 'system', '主应用，即平台最大的管理平台', 1, 0);
INSERT INTO `app` VALUES (2, '总平台', 'platform', '平台业务管理层', 1, 0);

-- ----------------------------
-- Table structure for device
-- ----------------------------
DROP TABLE IF EXISTS `device`;
CREATE TABLE `device`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `key` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '设备标识id',
  `mac` char(17) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '设备mac地址',
  `product_id` int UNSIGNED NOT NULL COMMENT '产品id',
  `user_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属用户',
  `secret` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密钥',
  `data` varchar(2550) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '自定义属性数据',
  `is_online` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否在线',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `online_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后上线时间',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间 0表示未删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `node_id`(`key` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of device
-- ----------------------------
INSERT INTO `device` VALUES (1, 'test_device_01', '58:BF:25:19:CF:40', 1, 0, 'goiot_device', '', 0, '121', 1708707421, 0, 1712391225, 0);

-- ----------------------------
-- Table structure for log
-- ----------------------------
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `level_id` smallint UNSIGNED NOT NULL DEFAULT 0 COMMENT '日志级别 <10为系统日志',
  `level_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '日志级别名称',
  `title` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '日志内容',
  `value_type` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'text' COMMENT '日志类型  text,json,html',
  `request_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求来源',
  `request_ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求来源IP',
  `request_user_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作人ID',
  `request_user` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作人',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '记录时间',
  `update_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态 0=未处理 1=已查看 2=已处理',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `create_time`(`create_time` ASC) USING BTREE,
  INDEX `log_log_level_id_fk`(`level_id` ASC) USING BTREE,
  CONSTRAINT `log_log_level_id_fk` FOREIGN KEY (`level_id`) REFERENCES `log_level` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 349 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of log
-- ----------------------------

-- ----------------------------
-- Table structure for log_level
-- ----------------------------
DROP TABLE IF EXISTS `log_level`;
CREATE TABLE `log_level`  (
  `id` smallint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '<10为系统日志',
  `name` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '中文名称',
  `label` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '英文别名',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '日志开启状态',
  `color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '日志颜色 #ff0000',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '日志级别' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of log_level
-- ----------------------------
INSERT INTO `log_level` VALUES (1, '调试', 'debug', '调试信息', 1, '#333333');
INSERT INTO `log_level` VALUES (2, '信息', 'info', '一般信息', 1, '#3498db');
INSERT INTO `log_level` VALUES (3, '通知', 'notice', '一般信息，但值得注意', 1, '#f1c40f');
INSERT INTO `log_level` VALUES (4, '警告', 'warning', '警告，非错误的异常情况', 1, '#f1c40f');
INSERT INTO `log_level` VALUES (5, '错误', 'error', '运行时错误，记录并非紧急的问题', 1, '#d63031');
INSERT INTO `log_level` VALUES (6, '关键', 'critical', '关键错误，应用程序的某个组件无法使用', 1, '#c23616');
INSERT INTO `log_level` VALUES (7, '紧急', 'alert', '紧急错误，应立即处理，系统重要部分发生异常', 1, '#ff3333');
INSERT INTO `log_level` VALUES (8, '系统瘫痪', 'emergency', '突发性系统瘫痪，导致完全无法使用', 1, '#ff0000');
INSERT INTO `log_level` VALUES (10, '登录日志', 'login', '用户登录时记录的日志', 1, '#f06595');

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `app_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '应用ID',
  `rule_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '对应权限ID（通过拥有的权限ID查询）',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级菜单ID',
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '别名',
  `title` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '显示名称',
  `type` char(12) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '类型',
  `path` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '组件地址',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否启用',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `meta` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT 'meta路由参数',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `sort`(`sort` ASC) USING BTREE,
  INDEX `menu_app_id_fk`(`app_id` ASC) USING BTREE,
  CONSTRAINT `menu_app_id_fk` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 110 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '菜单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, 1, 0, 0, 'home', '首页', 'menu', '/home', '', 1, 0, '{\"title\":\"首页\", \"type\":\"menu\",\"icon\":\"el-icon-eleme-filled\"}');
INSERT INTO `menu` VALUES (2, 1, 0, 1, 'dashboard', '控制台', 'menu', '/dashboard', 'system/home', 1, 0, '{\"title\":\"控制台\",\"type\":\"menu\",\"affix\":true,\"icon\":\"el-icon-coin\"}');
INSERT INTO `menu` VALUES (3, 1, 0, 0, 'system', '系统管理', 'menu', '/system', '', 1, 0, '{\"title\":\"系统管理\",\"type\":\"menu\",\"icon\":\"el-icon-menu\"}');
INSERT INTO `menu` VALUES (4, 1, 0, 3, 'app', '应用管理', 'menu', '/system/app', 'system/app', 1, 1, '{\"title\":\"应用管理\",\"type\":\"menu\",\"icon\":\"el-icon-message-box\"}');
INSERT INTO `menu` VALUES (5, 1, 0, 3, 'menus', '菜单管理', 'menu', '/system/menu', 'system/menu', 1, 2, '{\"title\":\"菜单管理\",\"type\":\"menu\",\"icon\":\"el-icon-menu\"}');
INSERT INTO `menu` VALUES (6, 1, 0, 3, 'rule', '权限管理', 'menu', '/system/rule', 'system/rule', 1, 3, '{\"title\":\"权限管理\",\"type\":\"menu\",\"icon\":\"el-icon-key\"}');
INSERT INTO `menu` VALUES (7, 1, 0, 3, 'role', '角色管理', 'menu', '/system/role', 'system/role', 1, 4, '{\"title\":\"角色管理\",\"type\":\"menu\",\"icon\":\"el-icon-briefcase\"}');
INSERT INTO `menu` VALUES (8, 1, 0, 3, 'user', '用户管理', 'menu', '/system/user', 'system/user', 1, 5, '{\"title\":\"用户管理\",\"type\":\"menu\",\"icon\":\"el-icon-avatar\"}');
INSERT INTO `menu` VALUES (9, 1, 0, 3, 'setting', '系统日志', 'menu', '/system/log', 'system/log', 1, 6, '{\"title\":\"系统日志\",\"type\":\"menu\",\"icon\":\"el-icon-document\"}');
INSERT INTO `menu` VALUES (100, 2, 18, 0, 'dashboard', '主页', '', '/dashboard', 'platform/home', 1, 0, '{\"title\":\"主页\",\"type\":\"menu\",\"icon\":\"el-icon-home-filled\"}');
INSERT INTO `menu` VALUES (108, 2, 0, 0, 'device', '设备管理', '', '/platform/device', 'platform/device', 1, 3, '{\"title\":\"设备管理\",\"type\":\"menu\",\"icon\":\"el-icon-odometer\"}');
INSERT INTO `menu` VALUES (109, 2, 0, 0, 'product', '产品管理', '', '/platform/product', 'platform/product', 1, 2, '{\"title\":\"产品管理\",\"type\":\"menu\",\"icon\":\"el-icon-cpu\"}');

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '产品id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '产品名称',
  `key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '产品标识名称',
  `property` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '属性设定',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '备注',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间 0表示未删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product
-- ----------------------------
INSERT INTO `product` VALUES (1, '测试设备', 'test', '', 0, '测试用', 0, 1712391249, 0);
INSERT INTO `product` VALUES (2, '测试', '', '', 1, '', 1712389961, 1712389961, 1712390151);
INSERT INTO `product` VALUES (3, '测试2', 'key11', '', 1, 'ds', 1712390118, 1712390118, 0);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `app_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '应用ID',
  `name` char(12) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `rules` varchar(1000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '权限ID ,分割a',
  `rules_checked` varchar(1000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '权限树选中的字节点ID',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '简介',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态',
  `sort` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `is_auto` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否为系统自动创建',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `sort`(`sort` ASC) USING BTREE,
  INDEX `role_app_id_fk`(`app_id` ASC) USING BTREE,
  CONSTRAINT `role_app_id_fk` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '用户角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, 1, '管理员', '', '', 0, 0, '', 1, 0, 0);
INSERT INTO `role` VALUES (2, 2, '平台管理员', '', '', 1673533015, 1673533015, '', 1, 1, 0);

-- ----------------------------
-- Table structure for rule
-- ----------------------------
DROP TABLE IF EXISTS `rule`;
CREATE TABLE `rule`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `app_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '应用ID',
  `name` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '名字',
  `alias` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '英文别名',
  `path` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '规则',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级Id',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态 0:1',
  `sort` tinyint NOT NULL DEFAULT 0 COMMENT '排序',
  `addon_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '插件ID 为0=不验证插件',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `uri`(`path` ASC) USING BTREE,
  INDEX `sort`(`sort` ASC) USING BTREE,
  INDEX `rule_app_id_fk`(`app_id` ASC) USING BTREE,
  CONSTRAINT `rule_app_id_fk` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限规则表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of rule
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `nickname` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `username` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '用户登录名',
  `mobile` char(11) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `avatar` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `sex` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '性别 0=女 1=男',
  `password` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '密码 md5',
  `register_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间',
  `login_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '登录时间',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态',
  `delete_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  `role_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `username`(`username` ASC) USING BTREE,
  INDEX `mobile`(`mobile` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '管理员', 'admin', '15188889999', '', 1, 'e10adc3949ba59abbe56e057f20f883e', 1669968619, 1712387771, 1, 0, 1);
INSERT INTO `user` VALUES (2, '平台管理', 'platform', '15122223333', '', 1, 'e10adc3949ba59abbe56e057f20f883e', 1670223528, 1711787082, 1, 0, 2);
INSERT INTO `user` VALUES (3, '4301', '', '', '', 1, '', 1674737920, 0, 1, 0, 0);
INSERT INTO `user` VALUES (4, '2053', '', '', '', 1, '', 1674908463, 0, 1, 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
