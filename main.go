package main

import (
	"day2/connections"
	"day2/handlers"
)

func main() {
	//connect DB
	connections.Connect()
	//menjalankan router/hosting
	handlers.HandleReq()
}
