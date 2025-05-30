// Package genrouter
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package genrouter

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/addons/miniapps/global"
	"hotgo/internal/consts"
	"hotgo/internal/library/addons"
	"hotgo/internal/service"
)

var (
	NoLoginRouter       []interface{} // 无需登录
	LoginRequiredRouter []interface{} // 需要登录
)

// Register 注册通过代码生成的后台路由
func Register(ctx context.Context, group *ghttp.RouterGroup) {
	prefix := addons.RouterPrefix(ctx, consts.AppAdmin, global.GetSkeleton().Name)
	group.Group(prefix, func(group *ghttp.RouterGroup) {
		if len(NoLoginRouter) > 0 {
			group.Bind(NoLoginRouter...)
		}
		group.Middleware(service.Middleware().AdminAuth)
		if len(LoginRequiredRouter) > 0 {
			group.Bind(LoginRequiredRouter...)
		}
	})
}
