package client

import (
	"github.com/fadada-go-sdk-apiv3/api3req"
	"github.com/fadada-go-sdk-apiv3/bean"
	"github.com/fadada-go-sdk-apiv3/bean/file"
	"os"
)

const (
	uploadFilePath      = "/documents/uploadFile"
	GetBySignFileIdPath = "/documents/getBySignFileId"
)

type FileClient struct {
	Client Client
}

/**
  获取个人unionId地址
*/
func (client *FileClient) UploadFile(token string, nonce string, req file.UploadFileReq, file *os.File) (string, error) {
	req.FileContentHash = api3req.GetFileHash(file)
	var fileMap = make(map[string]*os.File)
	fileMap["fileContent"] = file
	return client.Client.UploadRequest(token, nonce, req, uploadFilePath, fileMap)
}

/**
  获取个人unionId地址
*/
func (client *FileClient) GetBySignFileId(token string, nonce string, req file.GetBySignFileIdReq) (bean.FileResponse, error) {
	return client.Client.DownLoadFileRequest(token, nonce, req, GetBySignFileIdPath)
}
