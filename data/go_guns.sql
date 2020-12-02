/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50728
Source Host           : localhost:3306
Source Database       : go_guns

Target Server Type    : MYSQL
Target Server Version : 50728
File Encoding         : 65001

Date: 2020-12-02 16:58:42
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for casbin_rules
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rules`;
CREATE TABLE `casbin_rules` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of casbin_rules
-- ----------------------------

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `dept_id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) DEFAULT NULL,
  `dept_path` varchar(255) DEFAULT NULL,
  `dept_name` varchar(128) DEFAULT NULL,
  `sort` int(11) DEFAULT NULL,
  `leader` varchar(128) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `status` varchar(4) DEFAULT NULL,
  `create_by` varchar(64) DEFAULT NULL,
  `update_by` varchar(64) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`dept_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES ('1', '0', '/0/1', '爱拓科技', '0', 'aituo', '13782218188', 'atuo@aituo.com', '0', '1', '1', '2020-02-27 15:30:19', '2020-03-10 21:09:21', null);
INSERT INTO `sys_dept` VALUES ('7', '1', '/0/1/7', '研发部', '1', '', '', '', '0', '1', '1', '2020-03-08 23:10:59', '2020-04-08 13:00:03', null);
INSERT INTO `sys_dept` VALUES ('8', '1', '/0/1/8', '运维部', '0', '', '', '', '0', '1', null, '2020-03-08 23:11:08', '2020-03-10 16:50:27', null);
INSERT INTO `sys_dept` VALUES ('9', '1', '/0/1/9', '客服部', '0', '', '', '', '0', '1', null, '2020-03-08 23:11:15', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dept` VALUES ('10', '1', '/0/1/10', '人力资源', '3', '张三', '', '', '1', '1', '1', '2020-04-07 23:48:38', '2020-04-07 23:48:46', null);

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `menu_id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) DEFAULT NULL,
  `title` varchar(128) DEFAULT NULL,
  `icon` varchar(128) DEFAULT NULL,
  `path` varchar(128) DEFAULT NULL,
  `paths` varchar(128) DEFAULT NULL,
  `menu_type` varchar(1) DEFAULT NULL,
  `action` varchar(16) DEFAULT NULL,
  `permission` varchar(255) DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `no_cache` tinyint(1) DEFAULT NULL,
  `breadcrumb` varchar(255) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  `sort` int(11) DEFAULT NULL,
  `visible` varchar(1) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `is_frame` varchar(1) DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=484 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES ('2', 'Upms', '系统管理', 'example', '/upms', '/0/2', 'M', '无', '', '0', '1', '', 'Layout', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('3', 'Sysuser', '用户管理', 'user', 'sysuser', '/0/2/3', 'C', '无', 'system:sysuser:list', '2', null, null, '/sysuser/index', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:10:42', null);
INSERT INTO `sys_menu` VALUES ('43', null, '新增用户', null, '/api/v1/sysuser', '/0/2/3/43', 'F', 'POST', 'system:sysuser:add', '3', null, null, null, '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('44', null, '查询用户', null, '/api/v1/sysuser', '/0/2/3/44', 'F', 'GET', 'system:sysuser:query', '3', null, null, null, '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('45', null, '修改用户', null, '/api/v1/sysuser/', '/0/2/3/45', 'F', 'PUT', 'system:sysuser:edit', '3', null, null, null, '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('46', null, '删除用户', null, '/api/v1/sysuser/', '/0/2/3/46', 'F', 'DELETE', 'system:sysuser:remove', '3', null, null, null, '0', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 15:32:45', null);
INSERT INTO `sys_menu` VALUES ('51', 'Menu', '菜单管理', 'tree-table', 'menu', '/0/2/51', 'C', '无', 'system:sysmenu:list', '2', '1', '', '/menu/index', '3', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:16:30', null);
INSERT INTO `sys_menu` VALUES ('52', 'Role', '角色管理', 'peoples', 'role', '/0/2/52', 'C', '无', 'system:sysrole:list', '2', '1', '', '/role/index', '2', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:16:19', null);
INSERT INTO `sys_menu` VALUES ('56', 'Dept', '部门管理', 'tree', 'dept', '/0/2/56', 'C', '无', 'system:sysdept:list', '2', '0', '', '/dept/index', '4', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:16:47', null);
INSERT INTO `sys_menu` VALUES ('57', 'post', '岗位管理', 'pass', 'post', '/0/2/57', 'C', '无', 'system:syspost:list', '2', '0', '', '/post/index', '5', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:16:53', null);
INSERT INTO `sys_menu` VALUES ('58', 'Dict', '字典管理', 'education', 'dict', '/0/2/58', 'C', '无', 'system:sysdicttype:list', '2', '0', '', '/dict/index', '6', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:17:04', null);
INSERT INTO `sys_menu` VALUES ('59', 'DictData', '字典数据', 'education', 'dict/data/:dictId', '/0/2/59', 'C', '无', 'system:sysdictdata:list', '2', '0', '', '/dict/data', '100', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:17:25', null);
INSERT INTO `sys_menu` VALUES ('60', 'Tools', '系统工具', 'component', '/tools', '/0/60', 'M', '无', '', '0', '0', '', 'Layout', '3', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('61', 'Swagger', '系统接口', 'guide', 'swagger', '/0/60/61', 'C', '无', '', '60', '0', '', '/tools/swagger/index', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('62', 'Config', '参数设置', 'list', '/config', '/0/2/62', 'C', '无', 'system:sysconfig:list', '2', '0', '', '/config/index', '7', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:17:16', null);
INSERT INTO `sys_menu` VALUES ('63', '', '接口权限', 'bug', '', '/0/63', 'M', '', '', '0', '0', '', '', '99', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 16:39:32', null);
INSERT INTO `sys_menu` VALUES ('64', '', '用户管理', 'user', '', '/0/63/64', 'M', '', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('66', '', '菜单管理', 'tree-table', '', '/0/63/66', 'C', '', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('67', '', '菜单列表', 'tree-table', '/api/v1/menulist', '/0/63/66/67', 'A', 'GET', '', '66', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('68', '', '新建菜单', 'tree', '/api/v1/menu', '/0/63/66/68', 'A', 'POST', '', '66', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('69', '', '字典', 'dict', '', '/0/63/69', 'M', '', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('70', '', '类型', 'dict', '', '/0/63/69/70', 'C', '', '', '69', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('71', '', '字典类型获取', 'tree', '/api/v1/dict/databytype/', '/0/63/256/71', 'A', 'GET', '', '256', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('72', '', '修改菜单', 'bug', '/api/v1/menu', '/0/63/66/72', 'A', 'PUT', '', '66', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('73', '', '删除菜单', 'bug', '/api/v1/menu/:id', '/0/63/66/73', 'A', 'DELETE', '', '66', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('74', null, '管理员列表', 'bug', '/api/v1/sysUserList', '/0/63/64/74', 'A', 'GET', null, '64', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('75', null, '根据id获取管理员', 'bug', '/api/v1/sysUser/:id', '/0/63/64/75', 'A', 'GET', null, '64', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('76', null, '获取管理员', 'bug', '/api/v1/sysUser/', '/0/63/256/76', 'A', 'GET', null, '256', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-07-18 14:30:28', null);
INSERT INTO `sys_menu` VALUES ('77', null, '创建管理员', 'bug', '/api/v1/sysUser', '/0/63/64/77', 'A', 'POST', null, '64', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('78', null, '修改管理员', 'bug', '/api/v1/sysUser', '/0/63/64/78', 'A', 'PUT', null, '64', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('79', null, '删除管理员', 'bug', '/api/v1/sysUser/:id', '/0/63/64/79', 'A', 'DELETE', null, '64', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('80', null, '当前用户个人信息', 'bug', '/api/v1/user/profile', '/0/63/256/267/80', 'A', 'GET', null, '267', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-05-03 20:50:40', null);
INSERT INTO `sys_menu` VALUES ('81', null, '角色列表', 'bug', '/api/v1/rolelist', '/0/63/201/81', 'A', 'GET', null, '201', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('82', null, '获取角色信息', 'bug', '/api/v1/role/:id', '/0/63/201/82', 'A', 'GET', null, '201', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('83', null, '创建角色', 'bug', '/api/v1/role', '/0/63/201/83', 'A', 'POST', null, '201', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('84', null, '修改角色', 'bug', '/api/v1/role', '/0/63/201/84', 'A', 'PUT', null, '201', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('85', null, '删除角色', 'bug', '/api/v1/role/:id', '/0/63/201/85', 'A', 'DELETE', null, '201', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('86', null, '参数列表', 'bug', '/api/v1/configList', '/0/63/202/86', 'A', 'GET', null, '202', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('87', null, '根据id获取参数', 'bug', '/api/v1/config/:id', '/0/63/202/87', 'A', 'GET', null, '202', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('89', null, '创建参数', 'bug', '/api/v1/config', '/0/63/202/89', 'A', 'POST', null, '202', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('90', null, '修改参数', 'bug', '/api/v1/config', '/0/63/202/90', 'A', 'PUT', null, '202', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('91', null, '删除参数', 'bug', '/api/v1/config/:id', '/0/63/202/91', 'A', 'DELETE', null, '202', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('92', null, '获取角色菜单', 'bug', '/api/v1/menurole', '/0/63/256/92', 'A', 'GET', null, '256', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-07-10 20:47:39', null);
INSERT INTO `sys_menu` VALUES ('93', null, '根据角色id获取角色', 'bug', '/api/v1/roleMenuTreeselect/:id', '/0/63/256/93', 'A', 'GET', null, '256', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-07-10 20:47:52', null);
INSERT INTO `sys_menu` VALUES ('94', null, '获取菜单树', 'bug', '/api/v1/menuTreeselect', '/0/63/256/94', 'A', 'GET', null, '256', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-05-03 20:52:11', null);
INSERT INTO `sys_menu` VALUES ('95', null, '获取角色菜单', 'bug', '/api/v1/rolemenu', '/0/63/205/95', 'A', 'GET', null, '205', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('96', null, '创建角色菜单', 'bug', '/api/v1/rolemenu', '/0/63/205/96', 'A', 'POST', null, '205', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('97', null, '删除用户菜单数据', 'bug', '/api/v1/rolemenu/:id', '/0/63/205/97', 'A', 'DELETE', null, '205', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('103', null, '部门菜单列表', 'bug', '/api/v1/deptList', '/0/63/203/103', 'A', 'GET', null, '203', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('104', null, '根据id获取部门信息', 'bug', '/api/v1/dept/:id', '/0/63/203/104', 'A', 'GET', null, '203', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('105', null, '创建部门', 'bug', '/api/v1/dept', '/0/63/203/105', 'A', 'POST', null, '203', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('106', null, '修改部门', 'bug', '/api/v1/dept', '/0/63/203/106', 'A', 'PUT', null, '203', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('107', null, '删除部门', 'bug', '/api/v1/dept/:id', '/0/63/203/107', 'A', 'DELETE', null, '203', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('108', null, '字典数据列表', 'bug', '/api/v1/dict/datalist', '/0/63/69/206/108', 'A', 'GET', null, '206', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('109', null, '通过编码获取字典数据', 'bug', '/api/v1/dict/data/:id', '/0/63/69/206/109', 'A', 'GET', null, '206', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('110', null, '通过字典类型获取字典数据', 'bug', '/api/v1/dict/databytype/:id', '/0/63/256/110', 'A', 'GET', null, '256', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('111', null, '创建字典数据', 'bug', '/api/v1/dict/data', '/0/63/69/206/111', 'A', 'POST', null, '206', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('112', null, '修改字典数据', 'bug', '/api/v1/dict/data/', '/0/63/69/206/112', 'A', 'PUT', null, '206', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('113', null, '删除字典数据', 'bug', '/api/v1/dict/data/:id', '/0/63/69/206/113', 'A', 'DELETE', null, '206', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('114', null, '字典类型列表', 'bug', '/api/v1/dict/typelist', '/0/63/69/70/114', 'A', 'GET', null, '70', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('115', null, '通过id获取字典类型', 'bug', '/api/v1/dict/type/:id', '/0/63/69/70/115', 'A', 'GET', null, '70', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('116', null, '创建字典类型', 'bug', '/api/v1/dict/type', '/0/63/69/70/116', 'A', 'POST', null, '70', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('117', null, '修改字典类型', 'bug', '/api/v1/dict/type', '/0/63/69/70/117', 'A', 'PUT', null, '70', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('118', null, '删除字典类型', 'bug', '/api/v1/dict/type/:id', '/0/63/69/70/118', 'A', 'DELETE', null, '70', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('119', null, '获取岗位列表', 'bug', '/api/v1/postlist', '/0/63/204/119', 'A', 'GET', null, '204', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('120', null, '通过id获取岗位信息', 'bug', '/api/v1/post/:id', '/0/63/204/120', 'A', 'GET', null, '204', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('121', null, '创建岗位', 'bug', '/api/v1/post', '/0/63/204/121', 'A', 'POST', null, '204', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('122', null, '修改岗位', 'bug', '/api/v1/post', '/0/63/204/122', 'A', 'PUT', null, '204', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('123', null, '删除岗位', 'bug', '/api/v1/post/:id', '/0/63/204/123', 'A', 'DELETE', null, '204', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('138', null, '获取根据id菜单信息', 'bug', '/api/v1/menu/:id', '/0/63/66/138', 'A', 'GET', null, '66', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('142', null, '获取角色对应的菜单id数组', 'bug', '/api/v1/menuids', '/0/63/256/142', 'A', 'GET', null, '256', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('201', '', '角色管理', 'peoples', '', '/0/63/201', 'C', 'GET', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('202', '', '参数设置', 'list', '', '/0/63/202', 'C', 'DELETE', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('203', '', '部门管理', 'tree', '', '/0/63/203', 'C', 'POST', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('204', '', '岗位管理', 'pass', '', '/0/63/204', 'C', '', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('205', '', '角色菜单管理', 'nested', '', '/0/63/205', 'C', '', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('206', '', '数据', '', '', '/0/63/69/206', 'C', 'PUT', '', '69', '0', '', '', '2', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('211', 'Log', '日志管理', 'log', '/log', '/0/2/211', 'M', '', '', '2', '0', '', '/log/index', '8', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:15:32', null);
INSERT INTO `sys_menu` VALUES ('212', 'LoginLog', '登录日志', 'logininfor', '/loginlog', '/0/2/211/212', 'C', '', 'system:sysloginlog:list', '211', '0', '', '/loginlog/index', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('213', null, '获取登录日志', 'bug', '/api/v1/loginloglist', '/0/63/214/213', 'A', 'GET', null, '214', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('214', '', '日志管理', 'log', '', '/0/63/214', 'M', 'GET', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('215', '', '删除日志', 'bug', '/api/v1/loginlog/:id', '/0/63/214/215', 'A', 'DELETE', '', '214', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('216', 'OperLog', '操作日志', 'skill', '/operlog', '/0/2/211/216', 'C', '', 'system:sysoperlog:list', '211', '0', '', '/operlog/index', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('217', '', '获取操作日志', 'bug', '/api/v1/operloglist', '/0/63/214/217', 'A', 'GET', '', '214', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('220', '', '新增菜单', '', '', '/0/2/51/220', 'F', '', 'system:sysmenu:add', '51', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('221', '', '修改菜单', 'edit', '', '/0/2/51/221', 'F', '', 'system:sysmenu:edit', '51', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('222', '', '查询菜单', 'search', '', '/0/2/51/222', 'F', '', 'system:sysmenu:query', '51', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('223', '', '删除菜单', '', '', '/0/2/51/223', 'F', '', 'system:sysmenu:remove', '51', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('224', '', '新增角色', '', '', '/0/2/52/224', 'F', '', 'system:sysrole:add', '52', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('225', '', '查询角色', '', '', '/0/2/52/225', 'F', '', 'system:sysrole:query', '52', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('226', '', '修改角色', '', '', '/0/2/52/226', 'F', '', 'system:sysrole:edit', '52', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('227', '', '删除角色', '', '', '/0/2/52/227', 'F', '', 'system:sysrole:remove', '52', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('228', '', '查询部门', '', '', '/0/2/56/228', 'F', '', 'system:sysdept:query', '56', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('229', '', '新增部门', '', '', '/0/2/56/229', 'F', '', 'system:sysdept:add', '56', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('230', '', '修改部门', '', '', '/0/2/56/230', 'F', '', 'system:sysdept:edit', '56', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('231', '', '删除部门', '', '', '/0/2/56/231', 'F', '', 'system:sysdept:remove', '56', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('232', '', '查询岗位', '', '', '/0/2/57/232', 'F', '', 'system:syspost:query', '57', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('233', '', '新增岗位', '', '', '/0/2/57/233', 'F', '', 'system:syspost:add', '57', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('234', '', '修改岗位', '', '', '/0/2/57/234', 'F', '', 'system:syspost:edit', '57', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('235', '', '删除岗位', '', '', '/0/2/57/235', 'F', '', 'system:syspost:remove', '57', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('236', '', '字典查询', '', '', '/0/2/58/236', 'F', '', 'system:sysdicttype:query', '58', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('237', '', '新增类型', '', '', '/0/2/58/237', 'F', '', 'system:sysdicttype:add', '58', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('238', '', '修改类型', '', '', '/0/2/58/238', 'F', '', 'system:sysdicttype:edit', '58', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('239', '', '删除类型', '', '', '/0/2/58/239', 'F', '', 'system:sysdicttype:remove', '58', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('240', '', '查询数据', '', '', '/0/2/59/240', 'F', '', 'system:sysdictdata:query', '59', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('241', '', '新增数据', '', '', '/0/2/59/241', 'F', '', 'system:sysdictdata:add', '59', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('242', '', '修改数据', '', '', '/0/2/59/242', 'F', '', 'system:sysdictdata:edit', '59', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('243', '', '删除数据', '', '', '/0/2/59/243', 'F', '', 'system:sysdictdata:remove', '59', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('244', '', '查询参数', '', '', '/0/2/62/244', 'F', '', 'system:sysconfig:query', '62', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('245', '', '新增参数', '', '', '/0/2/62/245', 'F', '', 'system:sysconfig:add', '62', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('246', '', '修改参数', '', '', '/0/2/62/246', 'F', '', 'system:sysconfig:edit', '62', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('247', '', '删除参数', '', '', '/0/2/62/247', 'F', '', 'system:sysconfig:remove', '62', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('248', '', '查询登录日志', '', '', '/0/2/211/212/248', 'F', '', 'system:sysloginlog:query', '212', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('249', '', '删除登录日志', '', '', '/0/2/211/212/249', 'F', '', 'system:sysloginlog:remove', '212', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('250', '', '查询操作日志', '', '', '/0/2/211/216/250', 'F', '', 'system:sysoperlog:query', '216', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('251', '', '删除操作日志', '', '', '/0/2/211/216/251', 'F', '', 'system:sysoperlog:remove', '216', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('252', '', '获取登录用户信息', '', '/api/v1/getinfo', '/0/63/256/252', 'A', 'GET', '', '256', '0', '', '', '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('253', '', '角色数据权限', '', '/api/v1/roledatascope', '/0/63/256/253', 'A', 'PUT', '', '256', '0', '', '', '0', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-07-10 20:47:58', null);
INSERT INTO `sys_menu` VALUES ('254', '', '部门树接口【数据权限】', '', '/api/v1/roleDeptTreeselect/:id', '/0/63/256/254', 'A', 'GET', '', '256', '0', '', '', '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('255', '', '部门树【用户列表】', '', '/api/v1/deptTree', '/0/63/256/255', 'A', 'GET', '', '256', '0', '', '', '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('256', '', '必开接口', '', '', '/0/63/256', 'M', 'GET', '', '63', '0', '', '', '0', '1', '1', '', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('257', '', '通过key获取参数', 'bug', '/api/v1/configKey/:id', '/0/63/256/257', 'A', 'GET', '', '256', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('258', '', '退出登录', '', '/api/v1/logout', '/0/63/256/258', 'A', 'POST', '', '256', '0', '', '', '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('259', '', '头像上传', '', '/api/v1/user/avatar', '/0/63/256/267/259', 'A', 'POST', '', '267', '0', '', '', '0', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-05-03 20:50:05', null);
INSERT INTO `sys_menu` VALUES ('260', '', '修改密码', '', '/api/v1/user/pwd', '/0/63/256/260', 'A', 'PUT', '', '256', '0', '', '', '0', '1', '1', '', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('261', 'Gen', '代码生成', 'code', 'gen', '/0/60/261', 'C', '', '', '60', '0', '', '/tools/gen/index', '2', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:18:12', null);
INSERT INTO `sys_menu` VALUES ('262', 'EditTable', '代码生成修改', 'build', 'editTable', '/0/60/262', 'C', '', '', '60', '0', '', '/tools/gen/editTable', '100', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-05-03 20:38:36', null);
INSERT INTO `sys_menu` VALUES ('263', '', '字典类型下拉框【生成功能】', '', '/api/v1/dict/typeoptionselect', '/0/63/256/263', 'A', 'GET', '', '256', '0', '', '', '0', '1', '1', '', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('264', 'Build', '表单构建', 'build', 'build', '/0/60/264', 'C', '', '', '60', '0', '', '/tools/build/index', '1', '0', '1', '1', '1', '2020-04-11 15:52:48', '2020-07-18 13:51:47', null);
INSERT INTO `sys_menu` VALUES ('267', '', '个人中心', '', '', '/0/63/256/267', 'M', '', '', '256', '0', '', '', '0', '1', '1', '', '1', '2020-05-03 20:49:39', '2020-05-03 20:49:39', null);
INSERT INTO `sys_menu` VALUES ('269', 'Server', '服务监控', 'druid', 'system/monitor', '/0/60/269', 'C', '', 'monitor:server:list', '60', '0', '', '/system/monitor', '0', '0', '1', '1', '1', '2020-04-14 00:28:19', '2020-08-09 02:07:53', null);
INSERT INTO `sys_menu` VALUES ('459', 'sys_job管理', '定时任务', 'time-range', '/sys_job', '/0/459', 'M', '无', '', '0', '0', '', 'Layout', '2', '0', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-09 01:27:11', null);
INSERT INTO `sys_menu` VALUES ('460', 'sys_job管理', '定时任务', 'tool', 'sys_job', '/0/459/460', 'C', '无', 'sysjob:sysjob:list', '459', '0', '', '/sysjob/index', '0', '0', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-04 15:18:32', null);
INSERT INTO `sys_menu` VALUES ('461', 'sys_job', '分页获取定时任务', 'pass', 'sys_job', '/0/459/460/461', 'F', '无', 'sysjob:sysjob:query', '460', '0', '', '', '0', '0', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', null);
INSERT INTO `sys_menu` VALUES ('462', 'sys_job', '创建定时任务', 'pass', 'sys_job', '/0/459/460/462', 'F', '无', 'sysjob:sysjob:add', '460', '0', '', '', '0', '0', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', null);
INSERT INTO `sys_menu` VALUES ('463', 'sys_job', '修改定时任务', 'pass', 'sys_job', '/0/459/460/463', 'F', '无', 'sysjob:sysjob:edit', '460', '0', '', '', '0', '0', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', null);
INSERT INTO `sys_menu` VALUES ('464', 'sys_job', '删除定时任务', 'pass', 'sys_job', '/0/459/460/464', 'F', '无', 'sysjob:sysjob:remove', '460', '0', '', '', '0', '0', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', null);
INSERT INTO `sys_menu` VALUES ('465', 'sys_job', '定时任务', 'bug', 'sys_job', '/0/63/465', 'M', '无', '', '63', '0', '', '', '0', '1', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', null);
INSERT INTO `sys_menu` VALUES ('466', 'sys_job', '分页获取定时任务', 'bug', '/api/v1/sysjobList', '/0/63/465/466', 'A', 'GET', '', '465', '0', '', '', '0', '1', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', null);
INSERT INTO `sys_menu` VALUES ('467', 'sys_job', '根据id获取定时任务', 'bug', '/api/v1/sysjob/:id', '/0/63/465/467', 'A', 'GET', '', '465', '0', '', '', '0', '1', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', null);
INSERT INTO `sys_menu` VALUES ('468', 'sys_job', '创建定时任务', 'bug', '/api/v1/sysjob', '/0/63/465/468', 'A', 'POST', '', '465', '0', '', '', '0', '1', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', null);
INSERT INTO `sys_menu` VALUES ('469', 'sys_job', '修改定时任务', 'bug', '/api/v1/sysjob', '/0/63/465/469', 'A', 'PUT', '', '465', '0', '', '', '0', '1', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', null);
INSERT INTO `sys_menu` VALUES ('470', 'sys_job', '删除定时任务', 'bug', '/api/v1/sysjob/:id', '/0/63/465/470', 'A', 'DELETE', '', '465', '0', '', '', '0', '1', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', null);
INSERT INTO `sys_menu` VALUES ('471', 'job_log', '日志', 'bug', 'job_log', '/0/459/471', 'C', '', '', '459', '0', '', '/sysjob/log', '0', '1', '1', '1', '1', '2020-08-05 21:24:46', '2020-08-06 00:02:20', null);
INSERT INTO `sys_menu` VALUES ('473', 'sysSetting', '系统配置', 'form', 'syssettings', '/0/60/473', 'C', '无', 'syssetting:syssetting:list', '60', '0', '', '/system/settings', '0', '0', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 02:17:10', null);
INSERT INTO `sys_menu` VALUES ('474', 'sys_setting', '分页获取SysSetting', 'pass', 'sys_setting', '/0/60/473/474', 'F', '无', 'syssetting:syssetting:query', '473', '0', '', '', '0', '0', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', null);
INSERT INTO `sys_menu` VALUES ('475', 'sys_setting', '创建SysSetting', 'pass', 'sys_setting', '/0/60/473/475', 'F', '无', 'syssetting:syssetting:add', '473', '0', '', '', '0', '0', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', null);
INSERT INTO `sys_menu` VALUES ('476', 'sys_setting', '修改SysSetting', 'pass', 'sys_setting', '/0/60/473/476', 'F', '无', 'syssetting:syssetting:edit', '473', '0', '', '', '0', '0', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', null);
INSERT INTO `sys_menu` VALUES ('477', 'sys_setting', '删除SysSetting', 'pass', 'sys_setting', '/0/60/473/477', 'F', '无', 'syssetting:syssetting:remove', '473', '0', '', '', '0', '0', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', null);
INSERT INTO `sys_menu` VALUES ('478', 'sys_setting', 'SysSetting', 'bug', 'sys_setting', '/0/63/478', 'M', '无', '', '63', '0', '', '', '0', '1', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', null);
INSERT INTO `sys_menu` VALUES ('479', 'sys_setting', '分页获取SysSetting', 'bug', '/api/v1/syssettingList', '/0/63/478/479', 'A', 'GET', '', '478', '0', '', '', '0', '1', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', null);
INSERT INTO `sys_menu` VALUES ('480', 'sys_setting', '根据id获取SysSetting', 'bug', '/api/v1/syssetting/:id', '/0/63/478/480', 'A', 'GET', '', '478', '0', '', '', '0', '1', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', null);
INSERT INTO `sys_menu` VALUES ('481', 'sys_setting', '创建SysSetting', 'bug', '/api/v1/syssetting', '/0/63/478/481', 'A', 'POST', '', '478', '0', '', '', '0', '1', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', null);
INSERT INTO `sys_menu` VALUES ('482', 'sys_setting', '修改SysSetting', 'bug', '/api/v1/syssetting', '/0/63/478/482', 'A', 'PUT', '', '478', '0', '', '', '0', '1', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', null);
INSERT INTO `sys_menu` VALUES ('483', 'sys_setting', '删除SysSetting', 'bug', '/api/v1/syssetting/:id', '/0/63/478/483', 'A', 'DELETE', '', '478', '0', '', '', '0', '1', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', null);

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `post_id` int(11) NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) DEFAULT NULL,
  `post_code` varchar(128) DEFAULT NULL,
  `sort` int(11) DEFAULT NULL,
  `status` varchar(4) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`post_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES ('1', '首席执行官', 'CEO', '0', '0', '首席执行官', '1', '2020-03-08 23:11:15', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_post` VALUES ('2', '首席技术执行官', 'CTO', '2', '0', '首席技术执行官', '1', '1', '2020-04-11 15:52:48', '2020-05-03 20:58:01', null);
INSERT INTO `sys_post` VALUES ('3', '首席运营官', 'COO', '3', '0', '测试工程师', '1', '1', '2020-04-11 15:52:48', null, null);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` int(11) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) DEFAULT NULL,
  `status` varchar(4) DEFAULT NULL,
  `role_key` varchar(128) DEFAULT NULL,
  `role_sort` int(11) DEFAULT NULL,
  `flag` varchar(128) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `admin` tinyint(1) DEFAULT NULL,
  `data_scope` varchar(128) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES ('1', '系统管理员', '0', 'admin', '1', null, '1', null, null, '0', null, '2020-04-11 15:52:48', '2020-05-03 21:35:17', null);
INSERT INTO `sys_role` VALUES ('2', '普通角色', '0', 'common', '2', null, '1', null, null, '0', '2', '2020-04-11 15:52:48', '2020-05-03 21:32:59', null);
INSERT INTO `sys_role` VALUES ('3', '测试角色', '0', 'Tester', '3', '', '1', null, null, '0', null, '2020-04-11 15:52:48', '2020-04-12 14:10:52', null);

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `role_id` int(11) DEFAULT NULL,
  `dept_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` int(11) DEFAULT NULL,
  `menu_id` int(11) DEFAULT NULL,
  `role_name` longtext,
  `create_by` longtext,
  `update_by` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES ('1', '2', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '3', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '43', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '44', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '45', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '46', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '51', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '52', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '56', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '57', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '58', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '59', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '60', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '61', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '62', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '63', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '64', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '66', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '67', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '68', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '69', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '70', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '71', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '72', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '73', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '74', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '75', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '76', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '77', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '78', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '79', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '80', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '81', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '82', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '83', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '84', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '85', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '86', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '87', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '89', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '90', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '91', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '92', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '93', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '94', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '95', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '96', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '97', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '103', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '104', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '105', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '106', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '107', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '108', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '109', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '110', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '111', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '112', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '113', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '114', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '115', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '116', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '117', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '118', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '119', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '120', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '121', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '122', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '123', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '138', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '142', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '201', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '202', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '203', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '204', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '205', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '206', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '211', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '212', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '213', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '214', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '215', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '216', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '217', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '220', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '221', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '222', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '223', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '224', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '225', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '226', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '227', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '228', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '229', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '230', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '231', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '232', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '233', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '234', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '235', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '236', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '237', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '238', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '239', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '240', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '241', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '242', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '243', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '244', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '245', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '246', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '247', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '248', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '249', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '250', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '251', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '252', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '253', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '254', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '255', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '256', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '257', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '258', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '259', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '260', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '261', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '262', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '263', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '264', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '267', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '269', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '459', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '460', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '461', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '462', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '463', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '464', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '465', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '466', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '467', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '468', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '469', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '470', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '471', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '473', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '474', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '475', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '476', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '477', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '478', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '479', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '480', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '481', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '482', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '483', 'admin', null, null);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `nick_name` varchar(128) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `salt` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `sex` varchar(255) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `dept_id` int(11) DEFAULT NULL,
  `post_id` int(11) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `status` varchar(4) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `username` varchar(64) DEFAULT NULL,
  `password` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('1', 'zhangwj', '13818888888', '1', null, '', '0', '1@qq.com', '1', '1', '1', '1', null, '0', '2019-11-10 14:05:55', '2020-05-03 20:45:59', null, 'admin', '$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2');
INSERT INTO `sys_user` VALUES ('2', 'zhangwj', '13211111111', '3', null, null, '0', 'q@q.com', '8', '2', '1', '1', null, '0', '2019-11-12 18:28:27', '2020-03-14 20:08:43', null, 'zhangwj', '$2a$10$CqMwHahA3cNrNv16CoSxmeD4XMPU.BiKHPEAeaG5oXMavOKrjInXi');
