package bootstrap

import (
	"github.com/wutongtree/funds/client/models"
	"github.com/wutongtree/funds/client/utils"

	"github.com/astaxie/beego"
)

func InitTemplate() {
	beego.AddFuncMap("getDate", utils.GetDate)
	beego.AddFuncMap("getDateMH", utils.GetDateMH)
	beego.AddFuncMap("getAvatarSource", utils.GetAvatarSource)
	beego.AddFuncMap("getAvatar", utils.GetAvatar)
	beego.AddFuncMap("getAvatarUserid", models.GetAvatarUserid)
}
