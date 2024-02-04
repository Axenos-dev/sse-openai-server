package v1

import (
	"github.com/Axenos-dev/sse-openai-server/internal/controllers/v1/chat"
	"github.com/Axenos-dev/sse-openai-server/internal/controllers/v1/sse"
	"github.com/Axenos-dev/sse-openai-server/internal/llm"
	"github.com/Axenos-dev/sse-openai-server/internal/services"
	"github.com/gin-gonic/gin"
)

func InitV1(r *gin.Engine) {
	v1 := r.Group("/v1/")

	llm := llm.LLM{}
	llm.InitClient()

	chat.RegHandlers(v1, services.ChatService{LLM: llm})
	sse.RegHandlers(v1)
}
