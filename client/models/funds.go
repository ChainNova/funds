package models

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Fund struct {
	Id             string  `json:"Id,omitempty"`
	Name           string  `json:"Name,omitempty"`
	CreateTime     string  `json:"CreateTime,omitempty"`
	Quotas         float64 `json:"Quotas,omitempty"`
	MarketValue    float64 `json:"MarketValue,omitempty"`
	NetValue       float64 `json:"NetValue,omitempty"`
	NetDelta       string  `json:"NetDelta,omitempty"`
	ThresholdValue float64 `json:"ThresholdValue,omitempty"`
}

type MyFund struct {
	Fund
	MyQuotas      float64 `json:"MyQuotas,omitempty"`
	MyMarketValue float64 `json:"MyQuotas,omitempty"`
	MyBalance     float64 `json:"MyBalance,omitempty"`
}

type FundMarket struct {
	Index int
	Size  int64
	Type  string
}

// ---------- struct with app ------------
// getMyFundResponse getMyFundResponse
type getMyFundResponse struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Owner  string `protobuf:"bytes,2,opt,name=owner" json:"owner,omitempty"`
	Assets string `protobuf:"bytes,3,opt,name=assets" json:"assets,omitempty"`
	Fund   string `protobuf:"bytes,4,opt,name=fund" json:"fund,omitempty"`
}

// AppFund AppFund
type AppFund struct {
	Name          string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Funds         int    `protobuf:"bytes,1,opt,name=funds" json:"funds,omitempty"`
	Assets        int    `protobuf:"bytes,1,opt,name=assets" json:"assets,omitempty"`
	PartnerAssets int    `protobuf:"bytes,1,opt,name=partnerAssets" json:"partnerAssets,omitempty"`
	PartnerTime   int    `protobuf:"bytes,1,opt,name=partnerTime" json:"partnerTime,omitempty"`
	BuyStart      int    `protobuf:"bytes,1,opt,name=buyStart" json:"buyStart,omitempty"`
	BuyPer        int    `protobuf:"bytes,1,opt,name=buyPer" json:"buyPer,omitempty"`
	BuyAll        int    `protobuf:"bytes,1,opt,name=buyAll" json:"buyAll,omitempty"`
	Net           int    `protobuf:"bytes,1,opt,name=net" json:"net,omitempty"`
	CreateTime    int64  `protobuf:"bytes,1,opt,name=createTime" json:"createTime,omitempty"`
	UpdateTime    int64  `protobuf:"bytes,1,opt,name=updateTime" json:"updateTime,omitempty"`
	LatestTx      string `protobuf:"bytes,1,opt,name=latestTx" json:"latestTx,omitempty"`
}

// AppFundsResponse AppFundsResponse
type AppFundsResponse struct {
	Status string    `json:"status,omitempty"`
	Msg    string    `json:"msg,omitempty"`
	Result []AppFund `json:"result,omitempty"`
}

// AppFundResponse AppFundResponse
type AppFundResponse struct {
	Status string  `json:"status,omitempty"`
	Msg    string  `json:"msg,omitempty"`
	Result AppFund `json:"result,omitempty"`
}

// AppMyFund AppMyFund
type AppMyFund struct {
	Name   string `json:"Name,omitempty"`
	Owner  string `json:"owner,omitempty"`
	Assets int    `json:"assets,omitempty"`
	Fund   int    `json:"fund,omitempty"`
}

// AppMyFundResponse AppMyFundResponse
type AppMyFundResponse struct {
	Status string    `json:"status,omitempty"`
	Msg    string    `json:"msg,omitempty"`
	Result AppMyFund `json:"result,omitempty"`
}

// AppTransfterFundRequest AppTransfterFundRequest
type AppTransfterFundRequest struct {
	EnrollID string `json:"enrollID,omitempty"`
	Name     string `json:"name,omitempty"`
	Funds    int64  `json:"funds,omitempty"`
}

// AppTransfterFundResponse AppTransfterFundResponse
type AppTransfterFundResponse struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
}

// AppCreateFundRequest AppCreateFundRequest
type AppCreateFundRequest struct {
	EnrollID      string `json:"enrollID,omitempty"`
	Name          string `json:"name,omitempty"`
	Funds         int    `json:"funds,omitempty"`
	Assets        int    `json:"assets,omitempty"`
	PartnerAssets int    `json:"partnerAssets,omitempty"`
	PartnerTime   int    `json:"partnerTime,omitempty"`
	BuyStart      int    `json:"buyStart,omitempty"`
	BuyPer        int    `json:"buyPer,omitempty"`
	BuyAll        int    `json:"buyAll,omitempty"`
	Netvalue      int    `json:"net,omitempty"`
}

// AppCreateFundResponse AppCreateFundResponse
type AppCreateFundResponse struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
}

// AppSetFundNetvalueRequest AppSetFundNetvalueRequest
type AppSetFundNetvalueRequest struct {
	EnrollID string `json:"enrollID,omitempty"`
	Name     string `json:"name,omitempty"`
	Netvalue int    `json:"net,omitempty"`
}

// AppSetFundNetvalueResponse AppSetFundNetvalueResponse
type AppSetFundNetvalueResponse struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
}

//AppSetFundNewsRequest AppSetFundNewsRequest
type AppSetFundNewsRequest struct {
	EnrollID string `json:"enrollID,omitempty"`
	Name     string `json:"name,omitempty"`
	News     string `json:"news,omitempty"`
}

// AppSetFundNetvalueResponse AppSetFundNetvalueResponse
type AppSetFundNewsResponse struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
}

// AppSetFundThreshholdRequest AppSetFundThreshholdRequest
type AppSetFundThreshholdRequest struct {
	EnrollID      string `json:"enrollID,omitempty"`
	Name          string `json:"name,omitempty"`
	PartnerAssets int    `json:"partnerAssets,omitempty"`
	PartnerTime   int    `json:"partnerTime,omitempty"`
	BuyStart      int    `json:"buyStart,omitempty"`
	BuyPer        int    `json:"buyPer,omitempty"`
	BuyAll        int    `json:"buyAll,omitempty"`
}

// AppSetFundThreshholdResponse AppSetFundThreshholdResponse
type AppSetFundThreshholdResponse struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
}

// FundNetLog FundNetLog
type FundNetLog struct {
	Name string `json:"name"`
	Time int64  `json:"time"`
	Net  int64  `json:"net"`
}

// AppNetLogResponse AppNetLogResponse
type AppNetLogResponse struct {
	Status string       `json:"status,omitempty"`
	Msg    string       `json:"msg,omitempty"`
	Result []FundNetLog `json:"result,omitempty"`
}

type FundNews struct {
	Name string `json:"name"`
	News string `json:"news"`
	Time int64  `json:"time"`
	Date string `json:"date,omitempty"`
}

// AppNetLogResponse AppNetLogResponse
type AppNewsResponse struct {
	Status string     `json:"status,omitempty"`
	Msg    string     `json:"msg,omitempty"`
	Result []FundNews `json:"result,omitempty"`
}

// ListMyFunds ListMyFunds
func ListMyFunds(userId string, page int, offset int) (nums int, funds []Fund, err error) {

	// Get fund
	urlstr := getHTTPURL("funds")
	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("ListMyFunds failed: %v", err)
		return
	}

	logger.Debugf("ListMyFunds: url=%v response=%v", urlstr, string(response))
	logger.Debug(string(response))
	var result AppFundsResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("ListMyFunds failed: %v", err)
		return
	}

	if result.Status != "OK" {
		logger.Errorf("ListMyFunds failed: %v", result.Status)
		return
	}

	// result
	for _, v := range result.Result {
		fund := Fund{
			Id:          v.Name,
			Name:        v.Name,
			CreateTime:  time.Unix(v.CreateTime, 0).Format("2006-01-02"),
			Quotas:      float64(v.Funds),
			MarketValue: float64(v.Funds * v.Net),
			NetValue:    float64(v.Net),
		}

		funds = append(funds, fund)
	}

	nums = len(funds)
	return nums, funds, err
}

// GetMyFund GetMyFund
func GetMyFund(userId string, fundid string) (myfund AppMyFund, err error) {

	// // Get fund
	// urlstr := getHTTPURL("fund/" + fundid)
	// response, err := performHTTPGet(urlstr)
	// if err != nil {
	// 	logger.Errorf("GetMyFund failed: %v", err)
	// 	return
	// }

	// logger.Debugf("GetMyFund: url=%v response=%v", urlstr, string(response))

	// var resultAppFund AppFundResponse
	// err = json.Unmarshal(response, &resultAppFund)
	// if err != nil {
	// 	logger.Errorf("GetMyFund failed: %v", err)
	// 	return
	// }

	// if resultAppFund.Status != "OK" {
	// 	logger.Errorf("GetMyFund failed: %v", resultAppFund.Status)
	// 	return
	// }

	// Get My fund
	urlstr := getHTTPURL("user/" + fundid + "/" + userId)
	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("GetMyFund failed: %v", err)
		return
	}

	logger.Debugf("GetMyFund: url=%v response=%v", urlstr, string(response))

	var resultAppMyFund AppMyFundResponse
	err = json.Unmarshal(response, &resultAppMyFund)
	if err != nil {
		logger.Errorf("GetMyFund failed: %v", err)
		return
	}

	if resultAppMyFund.Status != "OK" {
		logger.Errorf("GetFund failed: %v", resultAppMyFund.Status)
		return
	}

	myfund = resultAppMyFund.Result

	return
}

// GetFund GetFund
func GetFund(fundid string) (fund AppFund, err error) {

	// Get fund
	urlstr := getHTTPURL("fund/" + fundid)
	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("GetFund failed: %v", err)
		return
	}

	logger.Debugf("GetFund: url=%v response=%v", urlstr, string(response))

	var resultAppFund AppFundResponse
	err = json.Unmarshal(response, &resultAppFund)
	if err != nil {
		logger.Errorf("GetFund failed: %v", err)
		return
	}

	if resultAppFund.Status != "OK" {
		logger.Errorf("GetFund failed: %v", resultAppFund.Status)
		return
	}
	fund = resultAppFund.Result

	return
}

// GetNetLog GetNetLog
func GetNetLog(fundid string) (history [][]int64, err error) {

	// Get fund
	urlstr := getHTTPURL("netLog/" + fundid)
	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("GetNetLog failed: %v", err)
		return
	}

	logger.Debugf("GetNetLog: url=%v response=%v", urlstr, string(response))
	logger.Debug(string(response))
	var result AppNetLogResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("GetNetLog failed: %v", err)
		return
	}

	if result.Status != "OK" {
		logger.Errorf("GetNetLog failed: %v", result.Status)
		return
	}

	// result
	for _, v := range result.Result {
		history = append(history, []int64{v.Time, v.Net})
	}

	sort.Sort(NetLog(history))

	return
}

// GetFundMarkets GetFundMarkets
func GetFundMarkets(latestTx string) (fundmarkets []FundMarket) {

	latestTxs := strings.Split(latestTx, "|")

	for k, v := range latestTxs {
		tx := strings.Split(v, ",")

		if len(tx) != 3 {
			continue
		}

		size, _ := strconv.ParseInt(tx[1], 10, 64)
		fundmarket := FundMarket{
			Index: k + 1,
			Size:  int64(math.Abs(float64(size))),
		}
		if size > 0 {
			fundmarket.Type = "购买"
		} else if size < 0 {
			fundmarket.Type = "赎回"
		}
		fundmarkets = append(fundmarkets, fundmarket)
	}

	return
}

func GetFundNews(fundid string) (err error, news []FundNews) {
	// Get fund
	urlstr := getHTTPURL("news/" + fundid)
	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("GetNews failed: %v", err)
		return
	}

	logger.Debugf("GetNews: url=%v response=%v", urlstr, string(response))
	logger.Debug(string(response))
	var result AppNewsResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("GetNews failed: %v", err)
		return
	}

	if result.Status != "OK" {
		logger.Errorf("GetNews failed: %v", result.Status)
		return
	}

	news = result.Result
	sort.Sort(NewsByTime(news))
	news = []FundNews(result.Result)

	for k, v := range news {
		news[k].Date = time.Unix(v.Time, 0).Format("2006-01-02 15:04:05")
	}
	return
}

func BuyFund(userId string, fundid string, amount int64) error {
	// Buy fund
	urlstr := getHTTPURL("transfer")
	request := AppTransfterFundRequest{
		EnrollID: userId,
		Name:     fundid,
		Funds:    amount,
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	response, err := performHTTPPost(urlstr, reqBody)
	if err != nil {
		logger.Errorf("BuyFund failed: %v", err)
		return err
	}

	logger.Debugf("BuyFund: url=%v request=%v response=%v", urlstr, request, string(response))

	var result AppTransfterFundResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("BuyFund failed: %v", err)
		return err
	}

	if result.Status != "OK" {
		logger.Errorf("BuyFund failed: %v", result.Status)
		return fmt.Errorf(result.Msg)
	}

	return nil
}

func RedeemFund(userId string, fundid string, quotas int64) error {
	// Redeem fund
	urlstr := getHTTPURL("transfer")
	request := AppTransfterFundRequest{
		EnrollID: userId,
		Name:     fundid,
		Funds:    quotas,
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	response, err := performHTTPPost(urlstr, reqBody)
	if err != nil {
		logger.Errorf("RedeemFund failed: %v", err)
		return err
	}

	logger.Debugf("RedeemFund: url=%v request=%v response=%v", urlstr, request, string(response))

	var result AppTransfterFundResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("RedeemFund failed: %v", err)
		return err
	}

	if result.Status != "OK" {
		logger.Errorf("RedeemFund failed: %v", result.Status)
		return fmt.Errorf(result.Msg)
	}

	return nil
}

// CreateNewFund CreateNewFund
func CreateNewFund(
	userId string,
	fundid string,
	quotas float64,
	balance float64,
	tbalance float64,
	ttime int,
	tcount float64,
	tbuyper float64,
	tbuyall float64,
	netvalue float64) error {
	// Create New Fund
	urlstr := getHTTPURL("create")
	request := AppCreateFundRequest{
		Name:          fundid,
		Funds:         int(quotas),
		Assets:        int(balance),
		PartnerAssets: int(tbalance),
		PartnerTime:   ttime,
		BuyStart:      int(tcount),
		BuyPer:        int(tbuyper),
		BuyAll:        int(tbuyall),
		Netvalue:      int(netvalue),
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	response, err := performHTTPPost(urlstr, reqBody)
	if err != nil {
		logger.Errorf("CreateNewFund failed: %v", err)
		return err
	}

	logger.Debugf("CreateNewFund: url=%v request=%v response=%v", urlstr, request, string(response))

	var result AppCreateFundResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("CreateNewFund failed: %v", err)
		return err
	}

	if result.Status != "OK" {
		logger.Errorf("CreateNewFund failed: %v", result.Status)
		return fmt.Errorf(result.Msg)
	}

	return nil
}

// SetFundNetvalue SetFundNetvalue
func SetFundNetvalue(
	userId string,
	fundid string,
	netvalue float64) error {
	// Set Fund netvalue
	urlstr := getHTTPURL("setnet")
	request := AppSetFundNetvalueRequest{
		Name:     fundid,
		Netvalue: int(netvalue),
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	response, err := performHTTPPost(urlstr, reqBody)
	if err != nil {
		logger.Errorf("SetFundNetvalue failed: %v", err)
		return err
	}

	logger.Debugf("SetFundNetvalue: url=%v request=%v response=%v", urlstr, request, string(response))

	var result AppSetFundNetvalueResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("SetFundNetvalue failed: %v", err)
		return err
	}

	if result.Status != "OK" {
		logger.Errorf("SetFundNetvalue failed: %v", result.Status)
		return fmt.Errorf(result.Msg)
	}

	return nil
}

// SetFundThreshhold SetFundThreshhold
func SetFundThreshhold(
	userId string,
	fundid string,
	tbalance float64,
	ttime int,
	tcount float64,
	tbuyper float64,
	tbuyall float64) error {
	// Create New Fund
	urlstr := getHTTPURL("setlimit")
	request := AppSetFundThreshholdRequest{
		Name:          fundid,
		PartnerAssets: int(tbalance),
		PartnerTime:   ttime,
		BuyStart:      int(tcount),
		BuyPer:        int(tbuyper),
		BuyAll:        int(tbuyall),
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	response, err := performHTTPPost(urlstr, reqBody)
	if err != nil {
		logger.Errorf("SetFundThreshhold failed: %v", err)
		return err
	}

	logger.Debugf("SetFundThreshhold: url=%v request=%v response=%v", urlstr, request, string(response))

	var result AppSetFundThreshholdResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("SetFundThreshhold failed: %v", err)
		return err
	}

	if result.Status != "OK" {
		logger.Errorf("SetFundThreshhold failed: %v", result.Status)
		return fmt.Errorf(result.Msg)
	}

	return nil
}

//SetFundNew SetFundNew
func SetFundNews(userId, fundid, news string) error {
	urlstr := getHTTPURL("setnews")
	request := AppSetFundNewsRequest{
		Name: fundid,
		News: news,
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	response, err := performHTTPPost(urlstr, reqBody)
	if err != nil {
		logger.Errorf("SetFundNews failed: %v", err)
		return err
	}

	logger.Debugf("SetFundNews: url=%v request=%v response=%v", urlstr, request, string(response))

	var result AppSetFundNetvalueResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("SetFundNews failed: %v", err)
		return err
	}

	if result.Status != "OK" {
		logger.Errorf("SetFundNews failed: %v", result.Status)
		return fmt.Errorf(result.Msg)
	}

	return nil
}

// 对公告按时间排序
type NewsByTime []FundNews

func (x NewsByTime) Len() int           { return len(x) }
func (x NewsByTime) Less(i, j int) bool { return x[i].Time > x[j].Time }
func (x NewsByTime) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//净值记录排序
type NetLog [][]int64

func (x NetLog) Len() int           { return len(x) }
func (x NetLog) Less(i, j int) bool { return x[i][0] > x[j][0] }
func (x NetLog) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
