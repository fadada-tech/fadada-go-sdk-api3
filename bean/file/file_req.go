package file

/**
  上传文件请求对象
*/
type UploadFileReq struct {
	FileType        int    `json:"fileType"`
	FileContentHash string `json:"fileContentHash"`
}

/**
  下载签署文件对象
*/
type GetBySignFileIdReq struct {
	TaskId     string `json:"taskId"`
	SignFileId string `json:"signFileId"`
}
