package sse

import (
	"fmt"

	"github.com/Axenos-dev/sse-openai-server/internal/entity"
	"github.com/Axenos-dev/sse-openai-server/internal/stream"
	"github.com/gin-gonic/gin"
)

type sse struct{}

func RegHandlers(r *gin.RouterGroup) {
	sse := sse{}

	r.GET("/sse/:topic", sse.serverSentEvents)
}

func (sse) serverSentEvents(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Transfer-Encoding", "chunked")

	topic := c.Param("topic")
	if len(topic) == 0 {
		c.SSEvent(topic, entity.MessageCompletionStream{
			Event:   entity.StreamEventError,
			Topic:   topic,
			Message: "empty topic",
		})
		c.Writer.Flush()
		return
	}

	if err := stream.MessageCompletionStream.InitNewStream(topic); err != nil {
		c.SSEvent(topic, entity.MessageCompletionStream{
			Event:   entity.StreamEventError,
			Topic:   topic,
			Message: fmt.Sprintf("can not create listening channel: %v", err),
		})
		c.Writer.Flush()
		return
	} else {
		c.SSEvent(topic, entity.MessageCompletionStream{
			Event:   entity.StreamEventConnectionEstablished,
			Topic:   topic,
			Message: "Connection established!",
		})
		c.Writer.Flush()
	}
	defer stream.MessageCompletionStream.CloseStream(topic)

	for {
		select {
		case msg := <-stream.MessageCompletionStream.Chan(topic):
			c.SSEvent(topic, msg)
			c.Writer.Flush()

		case <-c.Writer.CloseNotify():
			return
		}
	}
}
