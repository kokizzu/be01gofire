package main

import (
	"be01gofire/controller"
	"be01gofire/controller/cGuest"
)

func main() {
	server := controller.InitServer()
	//server.AssignHandler("/guest/login", cGuest.Login)
	//server.AssignHandler("/guest/register", cGuest.Register)
	//server.AssignHandler(`/guest/add-queue`,cGuest.AddQueue)
	//server.AssignHandler(`/guest/update-queue`,cGuest.UpdateQueue)
	//server.AssignHandler(`/guest/remove-queue`,cGuest.RemoveQueue)
	//server.AssignHandler(`/guest/show-queue`,cGuest.ShowQueue)
	server.AssignHandler(`/guest/create-account`,cGuest.CreateAccount)
	server.Listen(":8084")
}
