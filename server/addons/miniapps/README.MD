## miniAPP

### 简介

小程序插件


### 使用说明

AED 小程序插件


### 迁移或安装

1、安装 HotGo (2.16.10及以上)

项目介绍：https://github.com/bufanyun/hotgo

2、将当前插件项目拷贝进 HotGo 根目录的 server/addons 目录下

3、在 HotGo 根目录的 server/addons/modules 目录下创建go文件:miniapps.go，内容如下：
```go
package modules

import _ "hotgo/addons/miniapps"
```

4、HotGo 后台进入 开发工具->插件管理->找到 miniAPP (miniapps) 进行安装

5、重启服务即可生效


### 常用命令行

```shell
# 接口维护-gen service
gf gen service -s=addons/miniapps/logic -d=addons/miniapps/service

```



## 测试过程
### 登入接口

```shell
需要在微信开发者工具中打开开发者工具，打开调试器输入
wx.login({success:(res)=>console.log(res)});

{errMsg: "login:ok", code: "0f3sk31w3l0IM43FXx3w3A86aF3sk31a"}
code: "0f3sk31w3l0IM43FXx3w3A86aF3sk31a"
errMsg: "login:ok"
constructor: (...)
hasOwnProperty: (...)
isPrototypeOf: (...)
propertyIsEnumerable: (...)
toLocaleString: (...)
toString: (...)
valueOf: (...)
__defineGetter__: (...)
__defineSetter__: (...)
__lookupGetter__: (...)
__lookupSetter__: (...)
__proto__: Object
```


```shell
 curl -X POST \
                                     http://localhost:8000/api/miniapps/wxMiniLogin \
                                     -H 'Content-Type: application/json' \
                                     -d '{
                                   "code": "0f3sk31w3l0IM43FXx3w3A86aF3sk31a"
                                 }'
{"code":0,"message":"操作成功","data":{"id":1,"username":"admin","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicGlkIjowLCJkZXB0SWQiOjEwMCwiZGVwdFR5cGUiOiJjb21wYW55Iiwicm9sZUlkIjoxLCJyb2xlS2V5Ijoic3VwZXIiLCJ1c2VybmFtZSI6ImFkbWluIiwicmVhbE5hbWUiOiLlrZ_luIUiLCJhdmF0YXIiOiJodHRwOi8vbG9jYWxob3N0OjgwMDAvYXR0YWNobWVudC8yMDI1LTA0LTI0L2Q5ZXdpbnM0czNwc2RvazRqcC5qcGciLCJlbWFpbCI6IjEzMzgxNDI1MEBxcS5jb20iLCJtb2JpbGUiOiIxNTMwMzgzMDU3MSIsImFwcCI6ImFkbWluIiwibG9naW5BdCI6IjIwMjUtMDQtMjYgMDg6NTA6MjcifQ.sqVo_Ex3QvO_GJst4f7eJud2VvrYQCTCse0FmfODpVI","expires":604800,"openId":"oi7D16x03RsiWfR8f_XqHco7J8E8","sessionKey":"6ZjM0jQSQlZq//NxGPaskA==","unionId":""},"timestamp":1745628627,"traceID":"08e2b4c8fcb73918f117a822a52a9bb8"} 
```


## 附近5公里接口，不需要登入
```shell
 curl 'http://localhost:8000/api/miniapps/nearby-aed?latitude=31.238633&longitude=120.41958&radius=100'
{"code":0,"message":"操作成功","data":{"list":[{"id":5529,"longitude":120.450268,"latitude":31.224513,"address":"江苏省苏州市吴中区香山街道孙武路2015号，苏州太湖游客中心","createdAt":"2025-04-25 10:07:27","updatedAt":"2025-04-25 10:07:27"},{"id":43881,"longitude":120.431816,"latitude":31.229643,"address":"详细地址未知","createdAt":"2025-04-25 10:07:27","updatedAt":"2025-04-25 10:07:27"}]},"timestamp":1745590719,"traceID":"383f65aa8295391834a72b05994244ee"}
```