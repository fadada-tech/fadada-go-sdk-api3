package api3req

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"sort"
	"strings"
)

type Req interface {
	Request(appConfig AppConfig, token string, bizContent string, nonce string) ([]byte, string, error)
	UploadRequest(appConfig AppConfig, token string, bizContent string, nonce string, fileMap map[string]*os.File) ([]byte, error)
}

type AppConfig struct {
	AppId   string
	AppKey  string
	UrlPath string
}

/**
  参数排序
*/
func SortParam(paramMap map[string]string) string {
	var keys = make([]string, 0)
	for key := range paramMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var sortParam string
	keysLen := len(keys)
	for i := 0; i < keysLen; i++ {
		sortParam += keys[i] + "=" + paramMap[keys[i]]
		if i < (keysLen - 1) {
			sortParam += "&"
		}
	}
	return sortParam
}

/**
  签名方法，排序后的参数，appKey，时间戳
*/
func Sign(sortParam string, appKey string, timestamp string) string {
	messageSha256 := sha256.New()
	messageSha256.Write([]byte(sortParam))
	message := strings.ToLower(hex.EncodeToString(messageSha256.Sum(nil)))
	keyByes := []byte(appKey)
	h1 := hmac.New(sha256.New, keyByes)
	h1.Write([]byte(timestamp))
	h1Bytes := h1.Sum(nil)
	h2 := hmac.New(sha256.New, h1Bytes)
	h2.Write([]byte(message))
	signature := strings.ToLower(hex.EncodeToString(h2.Sum(nil)))
	return signature
}

/**
  获取文件hash值 sha256 hex string
*/
func GetFileHash(f *os.File) string {
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return ""
	}
	if _, err := f.Seek(0, 0); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}
