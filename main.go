package main

import (
	"HACKTIVE8/service"
	"fmt"
)

func main() {

	userSvc := service.NewUserService()
	msg := userSvc.Register(&service.User{Name: "david"})
	fmt.Println(msg)
	msg = userSvc.Register(&service.User{Name: "rafi"})
	fmt.Println(msg)

	c := make(chan *service.ListUser, 1)
	go userSvc.GetAll(c)

	res := <-c
	for i, t := range res.Users {
		fmt.Printf("User ke-%d : %v\n", i+1, t.Name)
	}
	fmt.Println("LEWAT")
	close(c)

}
