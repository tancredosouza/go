package main

import "./buffers"
import "./infrastructure"
import "./orchestrator"

func main() {
	InitializeMiddleware()
}

func InitializeMiddleware() {
	orchestrator.Initialize()

	buffers.Initialize()

	srh := infrastructure.ServerRequestHandler{"localhost", 3993}
	srh.Initialize()
}
