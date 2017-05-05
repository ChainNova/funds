package controllers

import (
	"fmt"
	"time"

	"github.com/wutongtree/funds/client/models"
)

type FundsController struct {
	BaseController
}

func (c *FundsController) ListMyFunds() {
	_, funds, _ := models.ListMyFunds(c.UserUserId, 1, 100)

	c.Data["funds"] = funds
	c.Data["countFunds"] = len(funds)

	c.TplName = "funds/myfunds.tpl"
}

func (c *FundsController) GetFund() {
	fundid := c.GetString(":id")
	c.Data["fundid"] = fundid

	// 我的基金
	myappfund, _ := models.GetMyFund(c.UserUserId, fundid)
	fund, _ := models.GetFund(fundid)

	myfund := models.MyFund{
		Fund: models.Fund{Id: fund.Name,
			Name:        fund.Name,
			CreateTime:  time.Unix(fund.CreateTime, 0).Format("2006-01-02"),
			Quotas:      float64(fund.Funds),
			MarketValue: float64(fund.Net * fund.Funds),
			NetValue:    float64(fund.Net),
			// NetDelta:       "+0.001|0.94%",
			ThresholdValue: float64(fund.Net * fund.BuyPer / 100),
		},
		MyQuotas:      float64(myappfund.Fund),
		MyMarketValue: float64(myappfund.Fund * fund.Net),
		MyBalance:     float64(myappfund.Assets),
	}

	// 净值走势
	netLog, _ := models.GetNetLog(fundid)
	c.Data["netLog"] = netLog
	fmt.Printf("GetNetLog: %v\n", netLog)

	if len(netLog) <= 1 {
		myfund.Fund.NetDelta = "0|0.00%"
	} else {
		delta := netLog[0][1] - netLog[1][1]

		if delta == 0 {
			myfund.Fund.NetDelta = "0|0.00%"
		} else {
			myfund.Fund.NetDelta = fmt.Sprintf("%+d|%.2f", delta, float64(delta)/float64(netLog[1][1])*100) + "%"
		}
	}

	c.Data["myfund"] = myfund
	fmt.Printf("GetFund: %v\n", myfund)

	// 市场动态
	markets := models.GetFundMarkets(fund.LatestTx)
	c.Data["markets"] = markets

	// 基金公告
	_, notices := models.GetFundNews(fundid)
	c.Data["notices"] = notices

	c.TplName = "funds/showfund.tpl"
}

// func (c *FundsController) GetMyFund() {
// 	fundid := c.GetString("fundid")
// 	fmt.Printf("GetMyFund fundid: %v\n", fundid)

// 	userid := c.UserUserId

// 	// 获取净值
// 	netvalue, err := c.GetFloat("netvalue")
// 	if err != nil {
// 		logger.Errorf("netvalue GetFloat error: %v", err)

// 		netvalue = 1
// 	}

// 	// 我的基金
// 	myfund, _ := models.GetMyFund(userid, fundid)

// 	c.Data["json"] = map[string]interface{}{"code": 0, "myaccount": myfund.MyBalance, "myquotas": myfund.MyQuotas, "mymarketvalue": myfund.MyQuotas * netvalue}
// 	c.ServeJSON()
// }

func (c *FundsController) BuyFund() {
	fundid := c.GetString("fundid")

	// 获取购买金额
	buycount, err := c.GetInt64("buycount")
	if err != nil {
		logger.Errorf("buycount GetInt64 error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "购买失败：" + err.Error()}
		c.ServeJSON()

		return
	}

	// 购买基金
	err = models.BuyFund(c.UserUserId, fundid, buycount)
	if err != nil {
		logger.Errorf("BuyFund error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "购买失败：" + err.Error()}
		c.ServeJSON()

		return
	}

	c.Data["json"] = map[string]interface{}{"code": 1, "message": "购买成功：" + fmt.Sprintf("%v", buycount)}
	c.ServeJSON()
}

func (c *FundsController) RedeemFund() {
	fundid := c.GetString("fundid")

	redeemcount, err := c.GetInt64("redeemcount")
	if err != nil {
		logger.Errorf("RedeemFund GetInt64 error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "赎回失败：" + err.Error()}
		c.ServeJSON()

		return
	}

	// 赎回基金
	err = models.RedeemFund(c.UserUserId, fundid, redeemcount*-1)
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "赎回失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "赎回成功：" + fmt.Sprintf("%v", redeemcount)}
	c.ServeJSON()
}

func (c *FundsController) GetNewFund() {

	c.TplName = "funds/createfund.tpl"
}

func (c *FundsController) CreateNewFund() {

	fundname := c.GetString("fundname")
	quotas, err := c.GetFloat("quotas")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "新建基金失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	balance, err := c.GetFloat("balance")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "新建基金失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	tbalance, err := c.GetFloat("tbalance")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "新建基金失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	ttime, err := c.GetInt("ttime")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "新建基金失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	tcount, err := c.GetFloat("tcount")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "新建基金失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	tbuyper, err := c.GetFloat("tbuyper")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "新建基金失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	tbuyall, err := c.GetFloat("tbuyall")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "新建基金失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	netvalue, err := c.GetFloat("netvalue")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "新建基金失败：" + err.Error()}
		c.ServeJSON()

		return
	}

	// 新建基金
	err = models.CreateNewFund(c.UserUserId, fundname, quotas, balance, tbalance, ttime, tcount, tbuyper, tbuyall, netvalue)
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "新建基金失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "新建基金成功：" + fundname}
	c.ServeJSON()
}

func (c *FundsController) ManageFund() {
	_, funds, _ := models.ListMyFunds(c.UserUserId, 1, 100)
	c.Data["funds"] = funds
	c.Data["countFunds"] = len(funds)

	c.TplName = "funds/managefund.tpl"
}

func (c *FundsController) FundNetvalue() {
	fundname := c.GetString(":id")
	fund, _ := models.GetFund(fundname)
	c.Data["fund"] = fund

	c.TplName = "funds/setfundnetvalue.tpl"
}

func (c *FundsController) FundThreshhold() {
	fundname := c.GetString(":id")

	fund, _ := models.GetFund(fundname)
	c.Data["fund"] = fund

	c.TplName = "funds/setfundthreshhold.tpl"
}

func (c *FundsController) FundNews() {
	fundname := c.GetString(":id")
	c.Data["fundname"] = fundname

	c.TplName = "funds/setfundnews.tpl"
}

func (c *FundsController) SetFundNetvalue() {
	fundname := c.GetString("fundname")
	netvalue, err := c.GetFloat("netvalue")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "设置基金净值失败：" + err.Error()}
		c.ServeJSON()

		return
	}

	// 设置基金净值
	err = models.SetFundNetvalue(c.UserUserId, fundname, netvalue)
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "设置基金净值失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "设置基金净值成功：" + fundname}
	c.ServeJSON()
}

func (c *FundsController) SetFundThreshhold() {
	fundname := c.GetString("fundname")
	tbalance, err := c.GetFloat("tbalance")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "设置基金限制失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	ttime, err := c.GetInt("ttime")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "设置基金限制失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	tcount, err := c.GetFloat("tcount")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "设置基金限制失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	tbuyper, err := c.GetFloat("tbuyper")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "设置基金限制失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	tbuyall, err := c.GetFloat("tbuyall")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "设置基金限制失败：" + err.Error()}
		c.ServeJSON()

		return
	}

	// 设置基金限制
	err = models.SetFundThreshhold(c.UserUserId, fundname, tbalance, ttime, tcount, tbuyper, tbuyall)
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "设置基金限制失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "设置基金限制成功：" + fundname}
	c.ServeJSON()
}

func (c *FundsController) SetFundNews() {
	fundname := c.GetString("fundname")
	news := c.GetString("news")

	// 设置基金净值
	err := models.SetFundNews(c.UserUserId, fundname, news)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "设置基金公告失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "设置基金公告成功：" + fundname}
	c.ServeJSON()
}
