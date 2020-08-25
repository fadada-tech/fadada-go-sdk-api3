package client

import (
	"github.com/fadada-go-sdk-apiv3/bean/account"
)

var getPersonUnionIdUrlPath = "/accounts/getPersonUnionIdUrl"

type AccountClient struct {
	Client Client
}

/**
  获取个人unionId地址
*/
func (client *AccountClient) GetPersonUnionIdUrl(token string, nonce string, req account.GetPersonUnionIdUrlReq) (string, error) {
	return client.Client.Request(token, nonce, req, getPersonUnionIdUrlPath)
}
