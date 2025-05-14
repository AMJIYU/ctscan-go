// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserWxmini is the golang structure for table user_wxmini.
type UserWxmini struct {
	Id         int64       `json:"id"         orm:"id"         description:""`
	Openid     string      `json:"openid"     orm:"openid"     description:"微信openid"`
	Unionid    string      `json:"unionid"    orm:"unionid"    description:"微信unionid"`
	Nickname   string      `json:"nickname"   orm:"nickname"   description:"微信昵称"`
	Avatar     string      `json:"avatar"     orm:"avatar"     description:"头像"`
	UserId     int64       `json:"userId"     orm:"user_id"    description:"关联的用户ID"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at" description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at" description:"更新时间"`
	Sessionkey string      `json:"sessionkey" orm:"sessionkey" description:"微信sessionkey"`
}
