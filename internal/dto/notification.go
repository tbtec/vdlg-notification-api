package dto

type CreateNotification struct {
	InputMessage  *InputMessage
	OutputMessage *OutputMessage
}

type SendNotification struct {
	Email   string
	Message string
	VideoId string
}
