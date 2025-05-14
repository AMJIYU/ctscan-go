// Package api
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package api

import (
	// "hotgo/internal/model/input/adminin"

	"github.com/gogf/gf/v2/frame/g"
)

// UserLoginReq 用户登录入参
// type UserLoginReq struct {
// 	g.Meta `path:"/login" method:"post" tags:"MiniApps" summary:"小程序用户登录获取Token"`
// 	adminin.AccountLoginInp
// }

// type UserLoginRes struct {
// 	*adminin.LoginModel
// }

// WxMiniLoginReq 微信小程序登录入参
type WxMiniLoginReq struct {
	g.Meta `path:"/wxlogin" method:"post" tags:"MiniApps" summary:"微信小程序登录"`
	Code   string `json:"code" v:"required#临时登录凭证code不能为空" dc:"临时登录凭证code"`
}

type WxMiniLoginRes struct {
	Token  string `json:"token" dc:"登录凭证"`
	// *adminin.LoginModel
	// OpenID     string `json:"openId"`     // 用户唯一标识
	// SessionKey string `json:"sessionKey"` // 会话密钥
	// UnionID    string `json:"unionId"`    // 用户在开放平台的唯一标识符(如果有)
}
