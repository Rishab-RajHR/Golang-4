package main

import (
	"fmt"

	"github.com/packages/auth"
	"github.com/packages/user"
)

func main() {
	auth.LoginWithCredentials("Alex", "secret")
	session := auth.GetSession()
	fmt.Println("session", session)

	user := user.User{
		Email: "user@gmail.com",
		Name:  "John Doe",
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)

}
