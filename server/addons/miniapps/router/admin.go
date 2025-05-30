// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/addons/miniapps/controller/admin/sys"
	"hotgo/addons/miniapps/global"
	"hotgo/addons/miniapps/router/genrouter"
	"hotgo/internal/consts"
	"hotgo/internal/library/addons"
	"hotgo/internal/service"
)

// Admin 后台路由
func Admin(ctx context.Context, group *ghttp.RouterGroup) {
	prefix := addons.RouterPrefix(ctx, consts.AppAdmin, global.GetSkeleton().Name)
	group.Group(prefix, func(group *ghttp.RouterGroup) {
		group.Bind(
			sys.Index,
		)
		group.Middleware(service.Middleware().AdminAuth)
		group.Bind(
			sys.Config,
		)
	})

	// 注册生成路由
	genrouter.Register(ctx, group)
}
