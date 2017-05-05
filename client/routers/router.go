package routers

import (
	"github.com/astaxie/beego"
	"github.com/wutongtree/funds/client/controllers"
)

func init() {
	// 登录
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	// 我的基金
	beego.Router("/my/funds", &controllers.FundsController{}, "get:ListMyFunds")
	// beego.Router("/my/fund", &controllers.FundsController{}, "get:GetMyFund")
	beego.Router("/funds/:id", &controllers.FundsController{}, "get:GetFund")
	beego.Router("/my/buyfund", &controllers.FundsController{}, "post:BuyFund")
	beego.Router("/my/redeemfund", &controllers.FundsController{}, "post:RedeemFund")

	// 基金管理
	beego.Router("/funds/new", &controllers.FundsController{}, "get:GetNewFund")
	beego.Router("/funds/manage", &controllers.FundsController{}, "get:ManageFund")
	beego.Router("/fund/setfundnetvalue/:id", &controllers.FundsController{}, "get:FundNetvalue")
	beego.Router("/fund/setfundthreshhold/:id", &controllers.FundsController{}, "get:FundThreshhold")
	beego.Router("/fund/setfundnews/:id", &controllers.FundsController{}, "get:FundNews")

	beego.Router("/fund/new", &controllers.FundsController{}, "post:CreateNewFund")
	beego.Router("/fund/setfundnetvalue", &controllers.FundsController{}, "post:SetFundNetvalue")
	beego.Router("/fund/setfundthreshhold", &controllers.FundsController{}, "post:SetFundThreshhold")
	beego.Router("/fund/setfundnews", &controllers.FundsController{}, "post:SetFundNews")
}
