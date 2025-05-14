// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.16.10
package sysin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserWxminiUpdateFields 修改微信小程序用户表字段过滤
type UserWxminiUpdateFields struct {
	Openid     string `json:"openid"     dc:"微信openid"`
	Unionid    string `json:"unionid"    dc:"微信unionid"`
	Nickname   string `json:"nickname"   dc:"微信昵称"`
	Avatar     string `json:"avatar"     dc:"头像"`
	UserId     int64  `json:"userId"     dc:"关联的用户ID"`
	Sessionkey string `json:"sessionkey" dc:"微信sessionkey"`
}

// UserWxminiInsertFields 新增微信小程序用户表字段过滤
type UserWxminiInsertFields struct {
	Openid     string `json:"openid"     dc:"微信openid"`
	Unionid    string `json:"unionid"    dc:"微信unionid"`
	Nickname   string `json:"nickname"   dc:"微信昵称"`
	Avatar     string `json:"avatar"     dc:"头像"`
	UserId     int64  `json:"userId"     dc:"关联的用户ID"`
	Sessionkey string `json:"sessionkey" dc:"微信sessionkey"`
}

// UserWxminiEditInp 修改/新增微信小程序用户表
type UserWxminiEditInp struct {
	entity.UserWxmini
}

func (in *UserWxminiEditInp) Filter(ctx context.Context) (err error) {
	// 验证微信openid
	if err := g.Validator().Rules("required").Data(in.Openid).Messages("微信openid不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证关联的用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("关联的用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type UserWxminiEditModel struct{}

// UserWxminiDeleteInp 删除微信小程序用户表
type UserWxminiDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *UserWxminiDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type UserWxminiDeleteModel struct{}

// UserWxminiViewInp 获取指定微信小程序用户表信息
type UserWxminiViewInp struct {
	Id         int64       `json:"id" v:"required#id不能为空" dc:"id"`
	CreatedAt  *gtime.Time `json:"createdAt" dc:"创建时间"`
	Openid     string      `json:"openid" dc:"微信openid"`
	Unionid    string      `json:"unionid" dc:"微信unionid"`
	Nickname   string      `json:"nickname" dc:"微信昵称"`
	Avatar     string      `json:"avatar" dc:"头像"`
	Sessionkey string      `json:"sessionkey" dc:"微信sessionkey"`
}

func (in *UserWxminiViewInp) Filter(ctx context.Context) (err error) {
	// 验证id
	if err := g.Validator().Rules("required").Data(in.Id).Messages("id不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证创建时间
	if err := g.Validator().Rules("required").Data(in.CreatedAt).Messages("创建时间不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type UserWxminiViewModel struct {
	entity.UserWxmini
}

// UserWxminiListInp 获取微信小程序用户表列表
type UserWxminiListInp struct {
	form.PageReq
	Id         int64         `json:"id"        dc:"id"`
	CreatedAt  []*gtime.Time `json:"createdAt" dc:"创建时间"`
	Openid     string        `json:"openid"    dc:"微信openid"`
	Unionid    string        `json:"unionid"   dc:"微信unionid"`
	Nickname   string        `json:"nickname"  dc:"微信昵称"`
	Avatar     string        `json:"avatar"    dc:"头像"`
	Sessionkey string        `json:"sessionkey" dc:"微信sessionkey"`
}

func (in *UserWxminiListInp) Filter(ctx context.Context) (err error) {

	return
}

type UserWxminiListModel struct {
	Id         int64       `json:"id"         dc:"id"`
	Openid     string      `json:"openid"     dc:"微信openid"`
	Unionid    string      `json:"unionid"    dc:"微信unionid"`
	Nickname   string      `json:"nickname"   dc:"微信昵称"`
	Avatar     string      `json:"avatar"     dc:"头像"`
	UserId     int64       `json:"userId"     dc:"关联的用户ID"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
	Sessionkey string      `json:"sessionkey" dc:"微信sessionkey"`
}

// UserWxminiExportModel 导出微信小程序用户表
type UserWxminiExportModel struct {
	Id         int64       `json:"id"         dc:"id"`
	Openid     string      `json:"openid"     dc:"微信openid"`
	Unionid    string      `json:"unionid"    dc:"微信unionid"`
	Nickname   string      `json:"nickname"   dc:"微信昵称"`
	Avatar     string      `json:"avatar"     dc:"头像"`
	UserId     int64       `json:"userId"     dc:"关联的用户ID"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
	Sessionkey string      `json:"sessionkey" dc:"微信sessionkey"`
}
