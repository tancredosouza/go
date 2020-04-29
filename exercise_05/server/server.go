package main

import "../distribution"

func main() {
	i := distribution.Invoker{"localhost", 6966}

	i.Invoke()
}