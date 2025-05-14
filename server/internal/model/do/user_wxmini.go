// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserWxmini is the golang structure of table user_wxmini for DAO operations like Where/Data.
type UserWxmini struct {
	g.Meta     `orm:"table:user_wxmini, do:true"`
	Id         interface{} //
	Openid     interface{} // 微信openid
	Unionid    interface{} // 微信unionid
	Nickname   interface{} // 微信昵称
	Avatar     interface{} // 头像
	UserId     interface{} // 关联的用户ID
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	Sessionkey interface{} // 微信sessionkey
}
