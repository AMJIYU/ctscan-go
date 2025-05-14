// Package api
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package api

import (
	mainsysin "hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"
)

// NearbyAedReq 查询附近 AED 设备
type NearbyAedReq struct {
	g.Meta    `path:"/nearby-aed" method:"get" tags:"MiniApps" summary:"根据经纬度和半径查询附近 AED 设备"`
	Latitude  float64 `json:"latitude"  v:"required#latitude 不能为空"  dc:"纬度"`
	Longitude float64 `json:"longitude" v:"required#longitude 不能为空" dc:"经度"`
	Radius    float64 `json:"radius"    v:"required#radius 不能为空"    dc:"查询半径（公里）"`
}

// NearbyAedRes 返回结果
type NearbyAedRes struct {
	List []*mainsysin.AedsLocationListModel `json:"list" dc:"附近 AED 设备列表"`
}