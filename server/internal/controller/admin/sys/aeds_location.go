// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package sys

import (
	"context"
	"hotgo/api/admin/aedslocation"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	AedsLocation = cAedsLocation{}
)

type cAedsLocation struct{}

// List 查看AED定位列表
func (c *cAedsLocation) List(ctx context.Context, req *aedslocation.ListReq) (res *aedslocation.ListRes, err error) {
	list, totalCount, err := service.SysAedsLocation().List(ctx, &req.AedsLocationListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.AedsLocationListModel{}
	}

	res = new(aedslocation.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出AED定位列表
func (c *cAedsLocation) Export(ctx context.Context, req *aedslocation.ExportReq) (res *aedslocation.ExportRes, err error) {
	err = service.SysAedsLocation().Export(ctx, &req.AedsLocationListInp)
	return
}

// Edit 更新AED定位
func (c *cAedsLocation) Edit(ctx context.Context, req *aedslocation.EditReq) (res *aedslocation.EditRes, err error) {
	err = service.SysAedsLocation().Edit(ctx, &req.AedsLocationEditInp)
	return
}

// View 获取指定AED定位信息
func (c *cAedsLocation) View(ctx context.Context, req *aedslocation.ViewReq) (res *aedslocation.ViewRes, err error) {
	data, err := service.SysAedsLocation().View(ctx, &req.AedsLocationViewInp)
	if err != nil {
		return
	}

	res = new(aedslocation.ViewRes)
	res.AedsLocationViewModel = data
	return
}

// Delete 删除AED定位
func (c *cAedsLocation) Delete(ctx context.Context, req *aedslocation.DeleteReq) (res *aedslocation.DeleteRes, err error) {
	err = service.SysAedsLocation().Delete(ctx, &req.AedsLocationDeleteInp)
	return
}