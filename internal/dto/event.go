package dto

type InputMessage struct {
	BucketName string `json:"bucketName"`
	Key        string `json:"key"`
	Url        string `json:"url"`
}

type OutputMessage struct {
	FileName string  `json:"fileName"`
	Status   string  `json:"status"` //COMPLETED, ERROR
	Reason   *string `json:"reason"` //FILE_SIZE, PROCESSING_ERROR
}
