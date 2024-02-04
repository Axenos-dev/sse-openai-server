package entity

type StreamEvent string

const (
	StreamEventConnectionEstablished StreamEvent = "connection_established"
	StreamEventMessageComletion      StreamEvent = "message_completion"
	StreamEventEndOfMsg              StreamEvent = "end_of_msg"
	StreamEventError                 StreamEvent = "error"
)

type StreamData struct {
	Content string `json:"content"`
}

type MessageCompletionStream struct {
	Topic string     `json:"-"`
	Data  StreamData `json:"data"`

	// basically describes what kind of data is been recieved on client
	Event StreamEvent `json:"event"`

	Message string `json:"message,omitempty"`
}
