package main

import "../requestHandlers"

func main() {
	srh := requestHandlers.ServerRequestHandler{"localhost", 6966}
	srh.StartListening()

	for {
		srh.Receive()
		ans := []byte("alive")
		srh.Send(ans)
	}
}