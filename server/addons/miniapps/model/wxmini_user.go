// WxMiniUser 微信小程序用户模型

package model

import (
    "github.com/gogf/gf/v2/os/gtime"
)
type WxMiniUser struct {
    Id        int64       `json:"id"         description:"主键"`
    OpenID    string      `json:"openid"     description:"微信openid"`
    UnionID   string      `json:"unionid"    description:"微信unionid"`
    Nickname  string      `json:"nickname"   description:"微信昵称"`
    Avatar    string      `json:"avatar"     description:"头像"`
    UserID    int64       `json:"userId"     description:"关联的用户ID"`
    CreatedAt *gtime.Time `json:"createdAt"  description:"创建时间"`
    UpdatedAt *gtime.Time `json:"updatedAt"  description:"更新时间"`
}

type SysUser struct {
	Id        int64       `json:"id"         description:"主键"`
	OpenID    string      `json:"openid"     description:"微信openid"`
	UnionID   string      `json:"unionid"    description:"微信unionid"`
	Nickname  string      `json:"nickname"   description:"微信昵称"`
	Avatar    string      `json:"avatar"     description:"头像"`
}