
## 法大大OpenApi3.0 go sdk 
go version go1.13.1 

### 调用示例   
#### 获取token示例  
参考oauth2_test.go中的TestGetAccessToken测试方法

#### 获取个人unionId地址示例    
参考account_test.go中的TestPersonUnionIdUrl测试方法

#### 上传文件示例
参考file_test.go中的TestUploadFile测试方法

#### 下载签署文件示例   
参考file_test.go中的TestGetBySignFileIdReq测试方法

### 初始化client
```go
// 拿初始化文件file_test.go中的fileClient示例
// 需要先构建默认的ApiV3Client，实现了Client接口，具体接口的方法参考Client接口
// 构建Client的时候需要初始化appId，appKey， 服务端地址等参数，还需要传入http请求超时时间
var fileClient = client.FileClient{
	Client: &client.ApiV3Client{AppId: test.TestAppId, AppKey: test.TestAppKey,
		Url: test.TestUrl,
		Req: &request.ApiV3Request{TimeOut: time.Duration(10) * time.Second}}}
```
