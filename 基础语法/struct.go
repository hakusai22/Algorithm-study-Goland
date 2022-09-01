package main

import "fmt"

type user struct {
	name     string
	password string
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func main() {
	u := user{name: "yinpeng", password: "1234356"}
	u.resetPassword("asdasdas")
	fmt.Println(u.checkPassword("2048"))
}
