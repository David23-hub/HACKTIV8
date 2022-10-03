package main

import (
	"swagger/app/interface/container"
	"swagger/app/interface/server"
)

func main() {
	server.StartServer(*container.SetUpContainer())
}
