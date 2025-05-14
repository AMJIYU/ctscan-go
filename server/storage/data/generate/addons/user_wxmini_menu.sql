-- hotgo自动生成菜单权限SQL 通常情况下只在首次生成代码时自动执行一次
-- 如需再次执行请先手动删除生成的菜单权限和SQL文件：/Users/eleven/Documents/GolandProjects/hotgo/server/storage/data/generate/addons/user_wxmini_menu.sql
-- Version: 2.16.10
-- Date: 2025-04-26 15:02:32
-- Link https://github.com/bufanyun/hotgo

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;

--
-- 数据库： `hotgo`
--

-- --------------------------------------------------------

--
-- 插入表中的数据 `hg_admin_menu`
--


SET @now := now();


-- 菜单目录
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, '2227', '微信小程序用户表', 'userWxmini', '/userWxmini', 'MenuOutlined', '1', '/addons/userWxmini/index', '', '', 'ParentLayout', '1', '', '0', '0', '', '0', '0', '0', '2', 'tr_2227 ', '0', '', '1', @now, @now);


SET @dirId = LAST_INSERT_ID();


-- 菜单页面
-- 列表
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @dirId, '微信小程序用户表列表', 'userWxminiIndex', 'index', '', '2', '', '/miniapps/userWxmini/list', '', '/addons/miniapps/userWxmini/index', '1', 'userWxmini', '0', '0', '', '0', '1', '0', '3', CONCAT('tr_2227 tr_', @dirId,' '), '10', '', '1', @now, @now);


SET @listId = LAST_INSERT_ID();

-- 详情
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '微信小程序用户表详情', 'userWxminiView', '', '', '3', '', '/miniapps/userWxmini/view', '', '', '1', '', '0', '0', '', '0', '1', '0', '4', CONCAT('tr_2227 tr_', @dirId, ' tr_', @listId,' '), '10', '', '1', @now, @now);


-- 菜单按钮

-- 编辑
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '编辑/新增微信小程序用户表', 'userWxminiEdit', '', '', '3', '', '/miniapps/userWxmini/edit', '', '', '1', '', '0', '0', '', '0', '1', '0', '4', CONCAT('tr_2227 tr_', @dirId, ' tr_', @listId,' '), '20', '', '1', @now, @now);


SET @editId = LAST_INSERT_ID();


-- 删除
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '删除微信小程序用户表', 'userWxminiDelete', '', '', '3', '', '/miniapps/userWxmini/delete', '', '', '1', '', '0', '0', '', '0', '0', '0', '4', CONCAT('tr_2227 tr_', @dirId, ' tr_', @listId,' '), '40', '', '1', @now, @now);




-- 导出
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '导出微信小程序用户表', 'userWxminiExport', '', '', '3', '', '/miniapps/userWxmini/export', '', '', '1', '', '0', '0', '', '0', '0', '0', '4', CONCAT('tr_2227 tr_', @dirId, ' tr_', @listId,' '), '70', '', '1', @now, @now);


COMMIT;