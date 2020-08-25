package main

import (
	"fmt"
	"github.com/fadada-go-sdk-apiv3/api3req"
	"github.com/fadada-go-sdk-apiv3/client"
	"github.com/fadada-go-sdk-apiv3/test"
	"testing"
	"time"
)

//初始化ApiV3Client，默认使用
var tokenClient = client.TokenClient{
	Client: &client.ApiV3Client{AppId: test.TestAppId, AppKey: test.TestAppKey,
		Url: test.TestUrl,
		Req: &api3req.ApiV3Request{TimeOut: time.Duration(10) * time.Second}}}

/**
  获取token请求示例
*/
func TestGetAccessToken(t *testing.T) {
	//初始化ApiV3Client，默认使用
	response, err := tokenClient.GetAccessToken(test.GetRandomString(32))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
