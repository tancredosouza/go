package main

import "github.com/my/repo/mymiddleware/namingService"

func main() {
	inv := namingService.Invoker{"localhost",3999}
	inv.Invoke()
}