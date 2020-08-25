package main

import (
	"fmt"
	"github.com/fadada-go-sdk-apiv3/api3req"
	"github.com/fadada-go-sdk-apiv3/bean"
	"github.com/fadada-go-sdk-apiv3/bean/account"
	"github.com/fadada-go-sdk-apiv3/client"
	"github.com/fadada-go-sdk-apiv3/test"
	"testing"
	"time"
)

//初始化ApiV3Client，默认使用
var accountClient = client.AccountClient{
	Client: &client.ApiV3Client{AppId: test.TestAppId, AppKey: test.TestAppKey,
		Url: test.TestUrl,
		Req: &api3req.ApiV3Request{TimeOut: time.Duration(10) * time.Second}}}

/**
  获取个人unionId地址 请求示例
*/
func TestPersonUnionIdUrl(t *testing.T) {
	var req = account.GetPersonUnionIdUrlReq{}
	req.ClientId = "15013477347"
	req.AllowModify = 1
	req.AuthScheme = 1
	req.RedirectUrl = "http://www.fadada.com"
	var person = account.PersonReq{}
	person.Name = "胡建"
	person.Mobile = "15013477347"
	person.IsMiniProgram = 1
	req.Person = person
	var notice = bean.Notice{}
	notice.NotifyAddress = "15013477347"
	notice.NotifyWay = 1
	req.Notice = notice

	//初始化ApiV3Client，默认使用
	response, err := accountClient.GetPersonUnionIdUrl(
		test.TestToken, test.GetRandomString(32), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

/**
  测试签名
*/
func TestSign(t *testing.T) {
	var sortParam = "X-FDD-Api-App-Id=FA25444835&X-FDD-Api-Nonce=b06a58d405f6417e985d07b14e9dc580&X-FDD-Api-Sign-Type=HMAC-SHA256&X-FDD-Api-Timestamp=2020-08-15 14:48:03.317&X-FDD-Api-Token=213736ff86a440e999875711471ea99d&bizContent={\"miniProgramSign\":1,\"redirectUrl\":\"http://www.fadada.com\",\"taskId\":\"f8202250301645e2829cb7bc2c363cbd\",\"unionId\":\"897c30d8bb4f4dc999839290e91747a1\"}"
	signature := api3req.Sign(sortParam, "XONZ3L2ADXEEVLA9CC3AXZVVXYAQ8SDD", "2020-08-15 14:48:03.317")
	fmt.Println("signature = " + signature)
}
