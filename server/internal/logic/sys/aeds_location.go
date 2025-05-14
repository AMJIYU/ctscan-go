// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package sys

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysAedsLocation struct{}

func NewSysAedsLocation() *sSysAedsLocation {
	return &sSysAedsLocation{}
}

func init() {
	service.RegisterSysAedsLocation(NewSysAedsLocation())
}

// Model AED定位ORM模型
func (s *sSysAedsLocation) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.AedLocations.Ctx(ctx), option...)
}

// List 获取AED定位列表
func (s *sSysAedsLocation) List(ctx context.Context, in *sysin.AedsLocationListInp) (list []*sysin.AedsLocationListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.AedsLocationListModel{})

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.AedLocations.Columns().Id, in.Id)
	}

	// 查询created_at
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.AedLocations.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 查询地址（模糊搜索）
	if in.Address != "" {
		mod = mod.WhereLike(dao.AedLocations.Columns().Address, "%"+in.Address+"%")
	}

	// 查询经度
	if in.Longitude > 0 {
		mod = mod.Where(dao.AedLocations.Columns().Longitude, in.Longitude)
	}

	// 查询纬度
	if in.Latitude > 0 {
		mod = mod.Where(dao.AedLocations.Columns().Latitude, in.Latitude)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.AedLocations.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取AED定位列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出AED定位
func (s *sSysAedsLocation) Export(ctx context.Context, in *sysin.AedsLocationListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.AedsLocationExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出AED定位-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.AedsLocationExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增AED定位
func (s *sSysAedsLocation) Edit(ctx context.Context, in *sysin.AedsLocationEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.AedsLocationUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改AED定位失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.AedsLocationInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增AED定位失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除AED定位
func (s *sSysAedsLocation) Delete(ctx context.Context, in *sysin.AedsLocationDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除AED定位失败，请稍后重试！")
		return
	}
	return
}

// View 获取AED定位指定信息
func (s *sSysAedsLocation) View(ctx context.Context, in *sysin.AedsLocationViewInp) (res *sysin.AedsLocationViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取AED定位信息，请稍后重试！")
		return
	}
	return
}

