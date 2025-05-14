// Package userwxmini
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package userwxmini

import (
	"hotgo/addons/miniapps/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询微信小程序用户表列表
type ListReq struct {
	g.Meta `path:"/userWxmini/list" method:"get" tags:"微信小程序用户表" summary:"获取微信小程序用户表列表"`
	sysin.UserWxminiListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.UserWxminiListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出微信小程序用户表列表
type ExportReq struct {
	g.Meta `path:"/userWxmini/export" method:"get" tags:"微信小程序用户表" summary:"导出微信小程序用户表列表"`
	sysin.UserWxminiListInp
}

type ExportRes struct{}

// ViewReq 获取微信小程序用户表指定信息
type ViewReq struct {
	g.Meta `path:"/userWxmini/view" method:"get" tags:"微信小程序用户表" summary:"获取微信小程序用户表指定信息"`
	sysin.UserWxminiViewInp
}

type ViewRes struct {
	*sysin.UserWxminiViewModel
}

// EditReq 修改/新增微信小程序用户表
type EditReq struct {
	g.Meta `path:"/userWxmini/edit" method:"post" tags:"微信小程序用户表" summary:"修改/新增微信小程序用户表"`
	sysin.UserWxminiEditInp
}

type EditRes struct{}

// DeleteReq 删除微信小程序用户表
type DeleteReq struct {
	g.Meta `path:"/userWxmini/delete" method:"post" tags:"微信小程序用户表" summary:"删除微信小程序用户表"`
	sysin.UserWxminiDeleteInp
}

type DeleteRes struct{}