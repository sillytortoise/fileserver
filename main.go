package main

import (
	_ "pptserver/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.BConfig.WebConfig.TemplateLeft = "<<<"
	beego.BConfig.WebConfig.TemplateLeft = ">>>"
	beego.Run()
}
