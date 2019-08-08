// @APIVersion 1.0.0
// @Title Naddic platform API
// @Description Naddic platform API
// @Contact youngtip@naddic.com
// @TermsOfServiceUrl
// @License
// @LicenseUrl
package routers

import (
	"github.com/YoungsoonLee/backend_datainfra/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/user",
			beego.NSRouter("/confirmEmail/:confirmToken", &controllers.UserController{}, "post:ConfirmEmail"),
			beego.NSRouter("/resendConfirmEmail/:email", &controllers.UserController{}, "post:ResendConfirmEmail"),
			beego.NSRouter("/forgotPassword/:email", &controllers.UserController{}, "post:ForogtPassword"),
			beego.NSRouter("/isValidResetPasswordToken/:resetToken", &controllers.UserController{}, "post:IsValidResetPasswordToken"),
			beego.NSRouter("/resetPassword/", &controllers.UserController{}, "post:ResetPassword"),
			beego.NSRouter("/getProfile", &controllers.UserController{}, "post:GetProfile"),
			beego.NSRouter("/updateProfile/", &controllers.UserController{}, "post:UpdateProfile"),
			beego.NSRouter("/updatePassword/", &controllers.UserController{}, "post:UpdatePassword"),

			beego.NSRouter("/registerIndonesia/", &controllers.UserController{}, "post:RegisterIndonesia"),
		),

		beego.NSNamespace("/auth",
			beego.NSRouter("/checkDisplayName/:displayname", &controllers.AuthController{}, "get:CheckDisplayName"),
			beego.NSRouter("/register", &controllers.AuthController{}, "post:CreateUser"),
			beego.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
			beego.NSRouter("/checkLogin", &controllers.AuthController{}, "get:CheckLogin"),
			beego.NSRouter("/social", &controllers.AuthController{}, "post:Social"),
			//beego.NSRouter("/logout", &controllers.AuthController{}, "post:Logout"),
		),

		beego.NSNamespace("/s3",
			beego.NSRouter("/uploadImage", &controllers.S3Controller{}, "post:UploadImage"),
		),

		beego.NSNamespace("/resource",
			beego.NSRouter("/resigster", &controllers.ResourceController{}, "post:CreateResource"),
			beego.NSRouter("/list", &controllers.ResourceController{}, "get:GetResources"),
			beego.NSRouter("/detail/:id", &controllers.ResourceController{}, "post:GetResourceDetail"),
			beego.NSRouter("/update/:id", &controllers.ResourceController{}, "post:UpdateResource"),
			beego.NSRouter("/delete/:id", &controllers.ResourceController{}, "post:DeleteResource"),
		),

		beego.NSNamespace("/kpi",
			beego.NSRouter("/list", &controllers.KpiController{}, "post:GetKPI"),
			//beego.NSRouter("/test", &controllers.SystemController{}, "get:CheckHealthy"),
			beego.NSRouter("/listUser", &controllers.KpiController{}, "post:GetUserKPI"),
			beego.NSRouter("/listSale", &controllers.KpiController{}, "post:GetSaleKPI"),
		),

		beego.NSNamespace("/system",
			beego.NSRouter("/healthy", &controllers.SystemController{}, "get:CheckHealthy"),
		),
	)
	beego.AddNamespace(ns)
}
