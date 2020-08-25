package client

var accessTokenPath = "/oauth2/accessToken"

type TokenClient struct {
	Client Client
}

/**
  获取token
*/
func (client *TokenClient) GetAccessToken(nonce string) (string, error) {
	return client.Client.Request("", nonce, nil, accessTokenPath)
}
