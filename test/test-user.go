package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	ID       int64
	Name     string
	Family   string
	UserName string `gorm:"column:UserName"`
	Password string
}

func main() {
	user := &user{
		ID:       3,
		Name:     "test1",
		Family:   "Test Family1",
		UserName: "Test1",
		Password: "Password1",
	}

	var data bytes.Buffer
	json.NewEncoder(&data).Encode(user)

	response, err := http.Post("http://localhost:8383/api/User/delete/3", "application/json", nil)
	if err != nil || response.StatusCode != 200 {
		log.Fatalf("ERROR : %v | Status : %s", err, response.Status)
	}
	fmt.Println("Test edit User Is OK!")

}
