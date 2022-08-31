package main

import (
	"HACKTIVE8/service"
	"net/http"
)

func main() {
	userSvc := service.NewUserService()

	http.HandleFunc("/register", userSvc.RegisterHandler)
	http.HandleFunc("/alluser", userSvc.UserHandler)

	http.ListenAndServe("localhost:8000", nil)

}
