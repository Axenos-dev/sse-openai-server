package controllers

import (
	v1 "github.com/Axenos-dev/sse-openai-server/internal/controllers/v1"
	"github.com/gin-gonic/gin"
)

func InitControllers(r *gin.Engine) {
	v1.InitV1(r)
}
