package service

type ListUser struct {
	Users []User
}

type User struct {
	Name string
}

type UserIface interface {
	Register(user *User) string
	GetAll(c chan *ListUser)
}

func NewUserService() UserIface {
	return &ListUser{}
}

func (u *ListUser) Register(user *User) string {
	u.Users = append(u.Users, *user)
	return user.Name + " Berhasil Ditambahkan!"
}

func (u *ListUser) GetAll(c chan *ListUser) {
	c <- u
}