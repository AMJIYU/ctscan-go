// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package router

import (
	"context"
	"hotgo/addons/miniapps/controller/api"
	"hotgo/addons/miniapps/global"
	"hotgo/internal/consts"
	"hotgo/internal/library/addons"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Api 前台路由
func Api(ctx context.Context, group *ghttp.RouterGroup) {
	prefix := addons.RouterPrefix(ctx, consts.AppApi, global.GetSkeleton().Name)
	group.Group(prefix, func(group *ghttp.RouterGroup) {
		group.Bind(
			// 无需验证的路由
			api.Index,
			api.Aed,
			api.Auth,
		)
		group.Middleware(service.Middleware().ApiAuth)
		group.Bind(
		// 需要验证的路由
		// ...
		)
	})
}
