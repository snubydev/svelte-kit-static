package main

import (
	"fmt"
	"zoo/services"
	"zoo/webserver"
)

func main() {
	port := "3000"
	fmt.Printf("Hello Svelte-Kit-Static App started http://localhost:%s ...\n", port)
	zoo := services.NewZoo()
	webserver.NewWebServer(zoo)
	webserver.Run(port)
}
