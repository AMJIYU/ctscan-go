// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AedLocations is the golang structure for table aed_locations.
type AedLocations struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Longitude float64     `json:"longitude" orm:"longitude"  description:"经度"`
	Latitude  float64     `json:"latitude"  orm:"latitude"   description:"纬度"`
	Address   string      `json:"address"   orm:"address"    description:"地址"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
