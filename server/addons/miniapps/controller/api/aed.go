package api

import (
	"context"
	"hotgo/addons/miniapps/api/api"
	"hotgo/addons/miniapps/service"

	"github.com/gogf/gf/v2/errors/gerror"
)


type cAed struct{}

// 在包级别声明了一个变量 Aed，它的类型就是上面定义的 cAed，并用空结构体字面量 cAed{} 进行初始化。
// 这样在其他地方就可以直接用 api.Aed（或按照包路径引用的 Aed）来调用它的方法。
var Aed = cAed{}



// (c *cAed) 是方法接收者，表示该方法属于 *cAed，即指针类型的 cAed。
// 使用指针接收者可以在方法内部修改接收者数据（本例中 struct 没字段，影响不大），也避免了值拷贝。
// Nearby 查询附近AED设备
func (c *cAed) Nearby(ctx context.Context, req *api.NearbyAedReq) (res *api.NearbyAedRes, err error) {
	// 内置函数 new(T) 会分配一个类型为 T 的指针并返回，且其内部值是该类型的零值。
	// 相当于 var tmp api.NearbyAedRes; res = &tmp，在这里用来构造返回的响应对象。
	res = new(api.NearbyAedRes)
	result, err := service.AedNearby().Nearby(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询附近AED设备失败")
	}
	res.List = result
	return
}
