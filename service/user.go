package service

import (
	"encoding/json"
	"net/http"
)

type ListUser struct {
	Users []User
}

type User struct {
	Name string
}

type UserIface interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	UserHandler(w http.ResponseWriter, r *http.Request)
	Register(user *User) string
	GetAll() []User
}

func NewUserService() UserIface {
	return &ListUser{}
}

func (u *ListUser) Register(user *User) string {
	u.Users = append(u.Users, *user)
	return user.Name + " Berhasil Ditambahkan!"
}

func (u *ListUser) GetAll() []User {

	return u.Users
}

func (u *ListUser) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "david"}
	u.Users = append(u.Users, user)
	res, _ := json.Marshal(user)
	w.Write(res)

}

func (u *ListUser) UserHandler(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(u.Users)
	w.Write(res)
}
