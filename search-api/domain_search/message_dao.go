package domain_search

type QueueMessageDto struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type QueueMessagesDto []QueueMessageDto
