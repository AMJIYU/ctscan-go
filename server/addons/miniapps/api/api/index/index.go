// Package index
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package index

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/miniapps/model/input/sysin"
)

// TestReq 测试
type TestReq struct {
	g.Meta `path:"/index/test" method:"get" tags:"miniAPP" summary:"测试前台API"`
	sysin.IndexTestInp
}

type TestRes struct {
	*sysin.IndexTestModel
}
