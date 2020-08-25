package api3req

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ApiV3Request struct {
	TimeOut time.Duration
}

/**
  返回的报文字节数组
  返回的contentType
  调用方法返回的错误
*/
func (c *ApiV3Request) Request(appConfig AppConfig, token string, bizContent string, notice string) ([]byte, string, error) {
	client := &http.Client{Timeout: c.TimeOut}
	var postString = ""
	if token != "" {
		postValue := url.Values{
			"bizContent": {bizContent},
		}
		postString = postValue.Encode()
	}
	var r, err = http.NewRequest(http.MethodPost, appConfig.UrlPath, strings.NewReader(postString))
	if err != nil {
		return nil, "", err
	}
	r.Header.Add("Content-type", "application/x-www-form-urlencoded")
	headerMap := GetHeader(appConfig, notice, bizContent, token)
	for key := range headerMap {
		r.Header.Add(key, headerMap[key])
	}
	rsp, err := client.Do(r)
	if rsp != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return nil, "", err
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, "", err
	}
	return body, rsp.Header.Get("Content-type"), nil
}

/**
  返回的报文字节数组  json字符串
  调用方法返回的错误
*/
func (c *ApiV3Request) UploadRequest(appConfig AppConfig, token string,
	bizContent string, notice string, fileMap map[string]*os.File) ([]byte, error) {
	client := &http.Client{Timeout: c.TimeOut}

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	for key := range fileMap {
		file := fileMap[key]
		w, _ := writer.CreateFormFile(key, filepath.Base(file.Name()))
		if _, err := io.Copy(w, file); err != nil {
			return nil, err
		}
		if _, err := file.Seek(0, 0); err != nil {
			return nil, err
		}
		if err := file.Close(); err != nil {
			return nil, err
		}
	}
	_ = writer.WriteField("bizContent", bizContent)
	if err := writer.Close(); err != nil {
		return nil, err
	}
	var r, err = http.NewRequest(http.MethodPost, appConfig.UrlPath, buf)
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-type", writer.FormDataContentType())
	headerMap := GetHeader(appConfig, notice, bizContent, token)
	for key := range headerMap {
		r.Header.Add(key, headerMap[key])
	}
	rsp, err := client.Do(r)
	if rsp != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

/**
  获取请求头参数
  token 为空场景表示获取的是 获取token的请求头参数
  token 不为空表示获取通用请求头参数
*/
func GetHeader(appConfig AppConfig, Nonce string, bizContent string, token string) map[string]string {
	var headerMap = make(map[string]string)
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	headerMap["X-FDD-Api-App-Id"] = appConfig.AppId
	headerMap["X-FDD-Api-Sign-Type"] = "HMAC-SHA256"
	headerMap["X-FDD-Api-Timestamp"] = timestamp
	headerMap["X-FDD-Api-Nonce"] = Nonce
	if token != "" {
		headerMap["X-FDD-Api-Token"] = token
		headerMap["bizContent"] = bizContent
	} else {
		headerMap["X-FDD-Api-Grant-Type"] = "client_credential"
	}
	//排序
	sortParam := SortParam(headerMap)
	//签名
	signature := Sign(sortParam, appConfig.AppKey, timestamp)
	headerMap["X-FDD-Api-Sign"] = signature
	log.Printf("request api3 param = %s", headerMap)
	//剔除bizContent
	if token != "" {
		delete(headerMap, "bizContent")
	}
	return headerMap
}
