// Package aedslocation
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package aedslocation

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询AED定位列表
type ListReq struct {
	g.Meta `path:"/aedsLocation/list" method:"get" tags:"AED定位" summary:"获取AED定位列表"`
	sysin.AedsLocationListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.AedsLocationListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出AED定位列表
type ExportReq struct {
	g.Meta `path:"/aedsLocation/export" method:"get" tags:"AED定位" summary:"导出AED定位列表"`
	sysin.AedsLocationListInp
}

type ExportRes struct{}

// ViewReq 获取AED定位指定信息
type ViewReq struct {
	g.Meta `path:"/aedsLocation/view" method:"get" tags:"AED定位" summary:"获取AED定位指定信息"`
	sysin.AedsLocationViewInp
}

type ViewRes struct {
	*sysin.AedsLocationViewModel
}

// EditReq 修改/新增AED定位
type EditReq struct {
	g.Meta `path:"/aedsLocation/edit" method:"post" tags:"AED定位" summary:"修改/新增AED定位"`
	sysin.AedsLocationEditInp
}

type EditRes struct{}

// DeleteReq 删除AED定位
type DeleteReq struct {
	g.Meta `path:"/aedsLocation/delete" method:"post" tags:"AED定位" summary:"删除AED定位"`
	sysin.AedsLocationDeleteInp
}

type DeleteRes struct{}