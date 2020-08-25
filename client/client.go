package client

import (
	"github.com/fadada-go-sdk-apiv3/bean"
	"os"
)

type Client interface {
	// 下载文件用到此接口
	DownLoadFileRequest(token string, nonce string, data interface{}, path string) (bean.FileResponse, error)
	Request(token string, nonce string, data interface{}, path string) (string, error)
	UploadRequest(token string, nonce string, data interface{}, path string, fileMap map[string]*os.File) (string, error)
}
