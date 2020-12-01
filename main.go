package main

import (
	"fmt"
	"zapood/config"
	"zapood/restapi"
)

func main() {
	Demo1()
}

func Demo1() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		restapi.RunApi("localhost:8383", db)
		// userModel := models.UserModel{
		// 	DB: db,
		// }
		// users, err2 := userModel.FindAll()
		// if err2 != nil {
		// 	fmt.Println(err2)
		// } else {
		// 	for _, user := range users {
		// 		fmt.Println(user.ToString())
		// 		fmt.Println("------------------------")
		// 	}
		// }
	}
}
