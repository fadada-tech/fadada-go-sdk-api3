package account

import "github.com/fadada-go-sdk-apiv3/bean"

/**
  获取个人unionId地址请求对象
*/
type GetPersonUnionIdUrlReq struct {
	ClientId    string      `json:"clientId"`
	Notice      bean.Notice `json:"notice"`
	AllowModify int         `json:"allowModify"`
	RedirectUrl string      `json:"redirectUrl"`
	AuthScope   string      `json:"authScope"`
	AuthScheme  int         `json:"authScheme"`
	Person      PersonReq   `json:"person"`
}

/**
  获取个人unionId地址请求对象 中person信息
*/
type PersonReq struct {
	Name                string `json:"name"`
	IdentType           string `json:"identType"`
	IdentNo             string `json:"identNo"`
	Mobile              string `json:"mobile"`
	IdPhotoOptional     int    `json:"idPhotoOptional"`
	BackIdCardImgBase64 string `json:"backIdCardImgBase64"`
	IdCardImgBase64     string `json:"idCardImgBase64"`
	BankCardNo          string `json:"bankCardNo"`
	IsMiniProgram       int    `json:"isMiniProgram"`
}
