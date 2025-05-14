package api

import (
	"context"
	"hotgo/addons/miniapps/api/api"
	"hotgo/addons/miniapps/service"
	// internalService "hotgo/internal/service"
)

var Auth = cAuth{}

type cAuth struct{}

// Login 账号登录并返回 Token
// func (c *cAuth) Login(ctx context.Context, req *api.UserLoginReq) (res *api.UserLoginRes, err error) {
// 	// 调用主模块登录逻辑
// 	data, err := internalService.AdminSite().AccountLogin(ctx, &req.AccountLoginInp)
// 	if err != nil {
// 		return
// 	}
// 	res = &api.UserLoginRes{
// 		LoginModel: data,
// 	}
// 	return
// }

// WxMiniLogin 微信小程序登录
func (c *cAuth) WxMiniLogin(ctx context.Context, req *api.WxMiniLoginReq) (res *api.WxMiniLoginRes, err error) {
	return service.WxMini().Login(ctx, req)
}
