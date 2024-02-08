package sse

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/Axenos-dev/sse-openai-server/internal/entity"
	"github.com/Axenos-dev/sse-openai-server/internal/stream"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type sse struct{}

func RegHandlers(r fiber.Router) {
	sse := sse{}

	r.Get("/sse/:topic", sse.serverSentEvents)
}

func (sse) serverSentEvents(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, "text/event-stream")
	c.Set(fiber.HeaderCacheControl, "no-cache")
	c.Set(fiber.HeaderConnection, "keep-alive")
	c.Set(fiber.HeaderTransferEncoding, "chunked")

	topic := c.Params("topic")

	streamErr := stream.MessageCompletionStream.InitNewStream(topic)

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		defer stream.MessageCompletionStream.CloseStream(topic)

		if streamErr != nil {
			json.NewEncoder(w).Encode(entity.MessageCompletionStream{
				Event:   entity.StreamEventError,
				Topic:   topic,
				Message: fmt.Sprintf("can not create listening channel: %v", streamErr),
			})
			w.Flush()
			return
		} else {
			json.NewEncoder(w).Encode(entity.MessageCompletionStream{
				Event:   entity.StreamEventConnectionEstablished,
				Topic:   topic,
				Message: "Connection established!",
			})

			if err := w.Flush(); err != nil {
				return
			}
		}

		for {
			msg := <-stream.MessageCompletionStream.Chan(topic)
			json.NewEncoder(w).Encode(msg)
			if err := w.Flush(); err != nil {
				return
			}
		}
	}))

	return nil
}
