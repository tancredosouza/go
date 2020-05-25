package result_callback

type ResultCallback struct {
}

var ReceivedMsgs chan string

func (r *ResultCallback) Initialize() {
	ReceivedMsgs = make(chan string, 100)
}

func (r *ResultCallback) PublishReceivedMessage(s string) {
	ReceivedMsgs <- s
}