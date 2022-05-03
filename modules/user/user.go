package userPackage

import "fmt"

type User struct {
	Name     string
	Age      uint32
	UserType string
}

func (AnUser User) Introduction() string {
	introduction := fmt.Sprintf("Hello my name is %s", AnUser.Name)
	fmt.Println(introduction)
	return introduction
}

func (AnUser *User) SetName(newName string) {
	AnUser.Name = newName
}
