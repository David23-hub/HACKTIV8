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

	res := userSvc.GetAll()

	for i, t := range res {
		fmt.Printf("User ke-%d : %v\n", i+1, t.Name)
	}

}
