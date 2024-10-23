package main

import (
    "ginDemo/routers"
)

func main() {
    // 在routers/router.go文件中已经创建了路由以及Db.DbConnect，这里执行一下即可
    r := routers.SetupRouter()
    // 在3001端口启动我们的项目（端口号随意）
    r.Run(":8080")
}
