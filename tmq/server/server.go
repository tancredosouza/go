package main

import "../infrastructure"

func main() {
	srh := infrastructure.ServerRequestHandler{"localhost", 3993}

	srh.StartListening()

	srh.KeepAcceptingNewConnections()
}
