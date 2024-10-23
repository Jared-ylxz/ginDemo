package routers // routers包

import (
    "database/sql"
    "fmt"
    "ginDemo/utils"  // 工具包里做了数据库的连接，下面会详细说

    "github.com/gin-gonic/gin"
)

var (
    Db *sql.DB
)

// 路由中间件
func MiddleWare() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("中间件执行完毕")
        c.Next()
    }
}

func SetupRouter() *gin.Engine {
    Db = utils.DbConnect()
    r := gin.Default()
    r.Use(MiddleWare())
    r.POST("/login", LoginHandler) // 子路由单独抽离出来在login.go中
    user := r.Group("/user")  // 二级路由设置
    {
        user.GET("/list", UserHandler)
    }
    return r
}