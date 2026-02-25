package sse

type SseResponse struct {
	Content    string `json:"content"`
	MessageId  string `json:"messageId"`
	ExtendInfo any    `json:"ExtendInfo"`
}
