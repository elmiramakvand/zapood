package entities

type User struct {
	ID       int64
	Name     string
	Family   string
	UserName string
	Password string
}

func (User) TableName() string {
	return "User"
}
