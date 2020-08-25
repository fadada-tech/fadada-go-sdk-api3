package bean

type Data interface{}

/**
  通知对象
*/
type Notice struct {
	NotifyWay     int    `json:"notifyWay"`
	NotifyAddress string `json:"notifyAddress"`
}

/**
  法大大api通用返回对象
*/
type FddApiResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

/**
  法大大api通用返回下载文件对象
  如果contentType = application/json 表示下载失败，要将bytes转成 FddApiResponse 查看失败信息
  如果contentType = application/zip 表示返回文件，zip压缩包，bytes 为文件二进制
*/
type FileResponse struct {
	ContentType string
	Bytes       []byte
}
