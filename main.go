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
	}
}
