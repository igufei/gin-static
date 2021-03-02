### gin框架的内嵌静态资源中间件，需要go版本v1.16
####代码例子如下
```bigquery
package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	. "github.com/igufei/gin-static"
)
// 以下需要go1.16支持，可以把资源文件和主程序整合成一个可执行文件 
//go:embed static/*
var f embed.FS

func main() {
	// 设置静态资源
	engine := gin.Default()
	{
		engine.Use(StaticEmbed("static", f))
	}
	engine.Run()
}
```