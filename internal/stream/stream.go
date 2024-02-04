package stream

import (
	"fmt"

	"github.com/Axenos-dev/sse-openai-server/internal/entity"
)

type messageStream struct {
	msgChan chan entity.MessageCompletionStream

	//endChan will be channel that will send signals for stopping stream
	endChan   chan int
	streaming bool
}

type msgStreams map[string]messageStream

var MessageCompletionStream msgStreams = make(msgStreams)

func (ms msgStreams) InitNewStream(topic string) error {
	if _, ok := ms[topic]; ok {
		return fmt.Errorf("channel with topic '%s' already exists", topic)
	}

	// buffer length is set to 1, for the logic:
	// sending    data -> chan    if chan is empty
	// reading    sse  <- chan
	// data will wait if sse didn't read previous data from chan
	ms[topic] = messageStream{
		msgChan: make(chan entity.MessageCompletionStream, 1),
		endChan: make(chan int, 1),
	}

	return nil
}

// closes stream fully, if client has disconnected
func (ms msgStreams) CloseStream(topic string) {
	if _, ok := ms[topic]; ok {
		close(ms[topic].msgChan)
		close(ms[topic].endChan)
		delete(ms, topic)
	}
}

// ends stream without closing channels
func (ms msgStreams) EndStream(topic string) {
	if stream, ok := ms[topic]; ok {
		if ms[topic].streaming {
			stream.endChan <- 1
			stream.streaming = false
			ms[topic] = stream
		}
	}
}

func (ms msgStreams) Chan(topic string) <-chan entity.MessageCompletionStream {
	return ms[topic].msgChan
}

func (ms msgStreams) EndChan(topic string) <-chan int {
	return ms[topic].endChan
}

func (ms msgStreams) Write(msg entity.MessageCompletionStream, topic string) error {
	if stream, ok := ms[topic]; ok {
		stream.msgChan <- msg
		// if event is an error, or end, so that means that stream will end
		stream.streaming = msg.Event == entity.StreamEventMessageComletion
		ms[topic] = stream
	} else {
		return fmt.Errorf("channel '%s' do not exist", topic)
	}

	return nil
}

func (ms msgStreams) DoesStreamExist(topic string) bool {
	_, ok := ms[topic]
	return ok
}
