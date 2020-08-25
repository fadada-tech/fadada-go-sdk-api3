package test

import (
	"math/rand"
	"time"
)

const (
	TestAppId  = "xxxxx"
	TestAppKey = "xxxxxx"
	TestUrl    = "http://127.0.0.1:8004/api/v3"
	TestToken  = "c3f9aba2a94744cf93877653863048d2"
)

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
