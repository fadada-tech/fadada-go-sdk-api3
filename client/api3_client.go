package client

import (
	"encoding/json"
	"fmt"
	"github.com/fadada-go-sdk-apiv3/api3req"
	"github.com/fadada-go-sdk-apiv3/bean"
	"os"
)

/**
  默认的api 3.0调用客户端
*/
type ApiV3Client struct {
	AppId  string
	AppKey string
	Url    string
	Req    api3req.Req
}

/**
  通用调用请求，
*/
func (client *ApiV3Client) Request(token string, nonce string, data interface{}, path string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常 = {}", err)
		}
	}()
	err := client.checkClientParam()
	if err != nil {
		return "", err
	}
	bizContent, err := getBizContent(data)
	if err != nil {
		return "", err
	}
	bytes, _, err := client.Req.Request(client.getAppConfig(path), token, bizContent, nonce)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

/**
  通用调用请求，
*/
func (client *ApiV3Client) UploadRequest(token string, nonce string, data interface{}, path string, fileMap map[string]*os.File) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常 = {}", err)
		}
	}()
	err := client.checkClientParam()
	if err != nil {
		return "", err
	}
	bizContent, err := getBizContent(data)
	if err != nil {
		return "", err
	}
	bytes, err := client.Req.UploadRequest(client.getAppConfig(path), token, bizContent, nonce, fileMap)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (client *ApiV3Client) getAppConfig(path string) api3req.AppConfig {
	return api3req.AppConfig{client.AppId, client.AppKey, client.Url + path}
}

func getBizContent(data interface{}) (string, error) {
	if data == nil {
		return "", nil
	}
	datas, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(datas), nil
}

func (client ApiV3Client) checkClientParam() error {
	if client.AppId == "" || client.AppKey == "" || client.Url == "" {
		return fmt.Errorf("平台请求参数为空")
	}
	return nil
}

/**
  下载文件接口
*/
func (client *ApiV3Client) DownLoadFileRequest(token string, nonce string, data interface{}, path string) (bean.FileResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常 = {}", err)
		}
	}()
	rsp := bean.FileResponse{}
	err := client.checkClientParam()
	if err != nil {
		return rsp, err
	}
	bizContent, err := getBizContent(data)
	if err != nil {
		return rsp, err
	}
	bytes, contentType, err := client.Req.Request(client.getAppConfig(path), token, bizContent, nonce)
	if err != nil {
		return rsp, err
	}
	rsp.Bytes = bytes
	rsp.ContentType = contentType
	return rsp, nil
}
