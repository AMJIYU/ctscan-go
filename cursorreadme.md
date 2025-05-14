
aedslocation菜单后端代码路径：   
server/api/admin/aedslocation/aedslocation.go 
server/internal/controller/admin/sys/aeds_location.go  
server/internal/model/input/sysin/aeds_location.go 
server/internal/logic/sys/aeds_location.go
server/internal/router/genrouter/aeds_location.go 

aedslocation菜单前段代码路径
web/src/api/aedsLocation/index.ts 
web/src/views/aedsLocation

借口查询示例
curl 'http://localhost:8001/admin/aedsLocation/list?page=1&pageSize=500' \
                                     -H 'Accept: application/json, text/plain, */*' \
                                     -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7' \
                                     -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwicGlkIjowLCJkZXB0SWQiOjEwMCwiZGVwdFR5cGUiOiJjb21wYW55Iiwicm9sZUlkIjoxLCJyb2xlS2V5Ijoic3VwZXIiLCJ1c2VybmFtZSI6ImFkbWluIiwicmVhbE5hbWUiOiLlrZ_luIUiLCJhdmF0YXIiOiJodHRwOi8vbG9jYWxob3N0OjgwMDAvYXR0YWNobWVudC8yMDI1LTA0LTI0L2Q5ZXdpbnM0czNwc2RvazRqcC5qcGciLCJlbWFpbCI6IjEzMzgxNDI1MEBxcS5jb20iLCJtb2JpbGUiOiIxNTMwMzgzMDU3MSIsImFwcCI6ImFkbWluIiwibG9naW5BdCI6IjIwMjUtMDQtMjUgMTU6MjU6MjQifQ.jn5dBYlR_6TL4f1hcEsNBlqiTl-DoUi_a2mHN-qh6Zo' \
                                     -H 'Cache-Control: no-cache' \
                                     -H 'Connection: keep-alive' \
                                     -b 'gfsessionid=4837d5a1c042391801e2f3753ea3f0f7' \
                                     -H 'Pragma: no-cache' \
                                     -H 'Referer: http://localhost:8001/' \
                                     -H 'Sec-Fetch-Dest: empty' \
                                     -H 'Sec-Fetch-Mode: cors' \
                                     -H 'Sec-Fetch-Site: same-origin' \
                                     -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36' \
                                     -H 'sec-ch-ua: "Google Chrome";v="135", "Not-A.Brand";v="8", "Chromium";v="135"' \
                                     -H 'sec-ch-ua-mobile: ?0' \
                                     -H 'sec-ch-ua-platform: "macOS"'



原来小程序中的接口：
curl "https://api.cloudtorch.cn/api/nearby-aed?latitude=30.2704&longitude=120.6247&radius=5"
