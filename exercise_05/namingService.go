package main

import "./namingService"

func main() {
	inv := namingService.NamingServiceInvoker{"localhost",3999}
	inv.Invoke()
}