package entities

import "fmt"

type User struct {
	ID       int64
	Name     string
	Family   string
	UserName string
	Password string
}

func (user User) ToString() string {
	return fmt.Sprintf("id = %d \n name = %s \n family = %s \n username = %s \n password = %s \n", user.ID, user.Name, user.Family, user.UserName, user.Password)
}
