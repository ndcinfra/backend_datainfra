package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/YoungsoonLee/backend_datainfra/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/YoungsoonLee/backend_datainfra/controllers:AuthController"],
		beego.ControllerComments{
			Method: "CheckDisplayName",
			Router: `/checkDisplayName/:displayname`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
