package main

import (
	"fmt"
	"github.com/fadada-go-sdk-apiv3/api3req"
	"github.com/fadada-go-sdk-apiv3/bean/file"
	"github.com/fadada-go-sdk-apiv3/client"
	"github.com/fadada-go-sdk-apiv3/test"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

//初始化ApiV3Client，默认使用
var fileClient = client.FileClient{
	Client: &client.ApiV3Client{AppId: test.TestAppId, AppKey: test.TestAppKey,
		Url: test.TestUrl,
		Req: &api3req.ApiV3Request{TimeOut: time.Duration(10) * time.Second}}}

/**
  上传文件 请求示例
*/
func TestUploadFile(t *testing.T) {
	uploadFile, err := os.Open("C:/Users/huj1/Desktop/temp/授权模板.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file hash = {}", api3req.GetFileHash(uploadFile))
	var req file.UploadFileReq
	req.FileType = 1
	response, err := fileClient.UploadFile(test.TestToken, test.GetRandomString(32), req, uploadFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}

/**
  下载签署文件 请求示例
*/
func TestGetBySignFileIdReq(t *testing.T) {
	var req file.GetBySignFileIdReq
	req.SignFileId = ""
	req.TaskId = "126e9510b66e47068a4b92a16bd84331"
	response, err := fileClient.GetBySignFileId(test.TestToken, test.GetRandomString(32), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ContentType = {}", response.ContentType)
	if response.ContentType == "application/zip" {
		if err := ioutil.WriteFile("C:/Users/huj1/Desktop/temp/126e9510b66e47068a4b92a16bd84331.zip", response.Bytes, 0655); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("response = {}", string(response.Bytes))
	}
}
