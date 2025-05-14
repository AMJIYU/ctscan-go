// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AedLocations is the golang structure of table aed_locations for DAO operations like Where/Data.
type AedLocations struct {
	g.Meta    `orm:"table:aed_locations, do:true"`
	Id        interface{} //
	Longitude interface{} // 经度
	Latitude  interface{} // 纬度
	Address   interface{} // 地址
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
