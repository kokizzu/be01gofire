package main

import (
	"be01gofire/controller"
	"be01gofire/controller/cGuest"
)

func main() {
	server := controller.InitServer()
	server.AssignHandler("/guest/login", cGuest.Login)
	server.Listen(":8084")
}
