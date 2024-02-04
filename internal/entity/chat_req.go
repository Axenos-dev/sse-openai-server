package entity

import (
	"fmt"
	"strings"
)

type SendChatMessageRequest struct {
	Message string `json:"message"`
}

type SendMessageResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (req *SendChatMessageRequest) Validate() error {
	req.Message = strings.TrimSpace(req.Message)
	if len(req.Message) == 0 {
		return fmt.Errorf("empty message")
	}

	if len(req.Message) > 200 {
		return fmt.Errorf("message could not be over 200 chars")
	}

	return nil
}
