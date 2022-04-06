package main

import (
	"github.com/tobolkin8/go_april_2022/database"
	authModel "github.com/tobolkin8/go_april_2022/models/auth"
	userModel "github.com/tobolkin8/go_april_2022/models/user"
)

func main() {
	repo, err := database.Connect()
	if err != nil {
		panic(err)
	}
	repo.LogMode(true)

	for _, user := range users {
		repo.Create(&user)
	}
	for _, login_info := range login_info {
		login_info.SetPassword(login_info.Password)
		repo.Create(&login_info)
	}

}

var users = []userModel.User{
	{Id: 1, Name: "Test-1", LastName: "Test-1-last-name", Email: "test1@test1.com"},
	{Id: 2, Name: "Test-2", LastName: "Test-2-last-name", Email: "test2@test2.com"},
}
var login_info = []authModel.LoginInfo{
	{Username: "test1_login_username", Password: "somepassword1", UserId: 1},
	{Username: "test2_login_username", Password: "somepassword2", UserId: 2},
}
