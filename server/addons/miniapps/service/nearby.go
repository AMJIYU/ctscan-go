package service

import (
	"context"
	minapi "hotgo/addons/miniapps/api/api"
	mainsysin "hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

//通过接口抽象服务逻辑，便于替换实现（如测试或扩展）。
// IAedNearby 插件内部接口
type IAedNearby interface {
	Nearby(ctx context.Context, in *minapi.NearbyAedReq) (list []*mainsysin.AedsLocationListModel, err error)
}

//获取接口实例，若未注册则触发 panic（强制初始化检查）。
//设计模式：通过全局变量 + 注册机制实现依赖注入，解耦调用方与具体实现。
//包级变量，存储接口的当前实现。
var localAedNearby IAedNearby
// RegisterAedNearby 注册接口实现，用于在包初始化时设置接口实现。
func RegisterAedNearby(i IAedNearby) {
	localAedNearby = i
}

// AedNearby 获取接口实现
func AedNearby() IAedNearby {
	if localAedNearby == nil {
		panic("未注册 IAedNearby 实现")
	}
	return localAedNearby
}





// aedNearbyImpl 默认实现
type aedNearbyImpl struct{}

func init() {
	RegisterAedNearby(&aedNearbyImpl{})
}

// Nearby 查询附近 AED
func (s *aedNearbyImpl) Nearby(ctx context.Context, in *minapi.NearbyAedReq) (list []*mainsysin.AedsLocationListModel, err error) {
	//  将公里转为米,zans
	//  radiusMeters := in.Radius * 1000
	// 不管前端传入的 radius 参数如何，一律使用默认 5 公里（5000 米）
	const defaultRadiusKm = 5
	radiusMeters := defaultRadiusKm * 1000

	// 获取主模块的模型
	m := service.SysAedsLocation().Model(ctx)
	// 使用 MySQL ST_Distance_Sphere 计算两点距离
	result, err := m.Fields("id, longitude, latitude, address, created_at, updated_at").
		Where("ST_Distance_Sphere(point(longitude, latitude), point(?,?)) <= ?", in.Longitude, in.Latitude, radiusMeters).
		All()
	if err != nil {
		return nil, err
	}
	// 解析到主模块列表模型
	err = result.Structs(&list)
	return
}
