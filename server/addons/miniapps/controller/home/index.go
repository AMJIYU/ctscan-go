// Package home
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package home

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/miniapps/api/home/index"
	"hotgo/addons/miniapps/service"
	"hotgo/internal/model"
	isc "hotgo/internal/service"
)

// Index 基础
var Index = cIndex{}

type cIndex struct{}

func (a *cIndex) Index(ctx context.Context, req *index.TestReq) (res *index.TestRes, err error) {
	data, err := service.SysIndex().Test(ctx, &req.IndexTestInp)
	if err != nil {
		return
	}

	isc.View().RenderTpl(ctx, "home/index.html", model.View{Data: g.Map{
		"name":   data.Name,
		"module": data.Module,
		"time":   data.Time,
	}})
	return
}
