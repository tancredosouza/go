package main

import "./namingService"

func main() {
	inv := namingService.Invoker{"localhost",3999}
	inv.Invoke()
}