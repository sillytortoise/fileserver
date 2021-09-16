package routers

import (
	"pptserver/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/allfiles", &controllers.MainController{}, "get:GetFiles")
	beego.Router("/delete/*", &controllers.MainController{}, "get:Delete")
	beego.Router("/posts/*", &controllers.MainController{}, "post:Upload")
	beego.Router("/download/*", &controllers.MainController{}, "get:Download")
	beego.Router("/folder/*", &controllers.MainController{}, "get:GetFolderFiles")
	beego.Router("/isempty/*", &controllers.MainController{}, "get:IsEmpty")
	beego.Router("/create/*", &controllers.MainController{}, "get:CreateFolder")
	beego.Router("/movefile", &controllers.MainController{}, "get:Movefile")
	beego.Router("/getdir", &controllers.MainController{}, "get:GetDir")
	beego.Router("/:name([a-z]+.vue)", &controllers.MainController{}, "get:GetStaticFile")
	beego.Router("/*", &controllers.MainController{})

}
