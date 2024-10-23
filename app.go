package main

import (
	"log"
    "fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gopkg.in/ini.v1"
	"ginDemo/model"
	"ginDemo/router"
)

var (
	g errgroup.Group
)

func main() {
    fmt.Println("AAAAAAAAAAAAAAA")
	conf, err := ini.Load("config/app.ini")
	if err != nil {
		return
	}
	model.InitDB(conf)
	r := gin.Default()
	router.Init(r)
	// 监听并在 0.0.0.0:8080 上启动服务
	g.Go(func() error {
		return r.Run(conf.Section("server").Key("addr").String())
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}