// Package model
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package model

// BasicConfig 基础配置
type BasicConfig struct {
	Test string `json:"basicTest"`
}

// WxMiniAppConfig 微信小程序配置
type WxMiniAppConfig struct {
	AppID     string `json:"wxMiniAppId"`     // 小程序AppID
	AppSecret string `json:"wxMiniAppSecret"` // 小程序AppSecret
}
