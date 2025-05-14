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

// AedsLocationUpdateFields 修改AED定位字段过滤
type AedsLocationUpdateFields struct {
	Longitude float64 `json:"longitude" dc:"经度"`
	Latitude  float64 `json:"latitude"  dc:"纬度"`
	Address   string  `json:"address"   dc:"地址"`
}

// AedsLocationInsertFields 新增AED定位字段过滤
type AedsLocationInsertFields struct {
	Longitude float64 `json:"longitude" dc:"经度"`
	Latitude  float64 `json:"latitude"  dc:"纬度"`
	Address   string  `json:"address"   dc:"地址"`
}

// AedsLocationEditInp 修改/新增AED定位
type AedsLocationEditInp struct {
	entity.AedLocations
}

func (in *AedsLocationEditInp) Filter(ctx context.Context) (err error) {
	// 验证经度
	if err := g.Validator().Rules("required").Data(in.Longitude).Messages("经度不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证纬度
	if err := g.Validator().Rules("required").Data(in.Latitude).Messages("纬度不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证地址
	if err := g.Validator().Rules("required").Data(in.Address).Messages("地址不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type AedsLocationEditModel struct{}

// AedsLocationDeleteInp 删除AED定位
type AedsLocationDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *AedsLocationDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type AedsLocationDeleteModel struct{}

// AedsLocationViewInp 获取指定AED定位信息
type AedsLocationViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *AedsLocationViewInp) Filter(ctx context.Context) (err error) {
	return
}

type AedsLocationViewModel struct {
	entity.AedLocations
}

// AedsLocationListInp 获取AED定位列表
type AedsLocationListInp struct {
	form.PageReq
	Id        int           `json:"id"        dc:"id"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"created_at"`
	Address   string        `json:"address"   dc:"地址"`
	Longitude float64       `json:"longitude" dc:"经度"`
	Latitude  float64       `json:"latitude"  dc:"纬度"`
}

func (in *AedsLocationListInp) Filter(ctx context.Context) (err error) {
	return
}

type AedsLocationListModel struct {
	Id        int         `json:"id"        dc:"id"`
	Longitude float64     `json:"longitude" dc:"经度"`
	Latitude  float64     `json:"latitude"  dc:"纬度"`
	Address   string      `json:"address"   dc:"地址"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"created_at"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"updated_at"`
}

// AedsLocationExportModel 导出AED定位
type AedsLocationExportModel struct {
	Id        int         `json:"id"        dc:"id"`
	Longitude float64     `json:"longitude" dc:"经度"`
	Latitude  float64     `json:"latitude"  dc:"纬度"`
	Address   string      `json:"address"   dc:"地址"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"created_at"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"updated_at"`
}
