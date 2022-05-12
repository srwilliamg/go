package userPackage

import (
	"fmt"
	"sync"
)

const USER_TYPE string = "natural"

type User struct {
	Name     string
	Lastname string
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

func UseStructs(wg *sync.WaitGroup) {
	defer wg.Done()
	person1 := User{Name: "Pepe", Lastname: "Perez", UserType: USER_TYPE}

	fmt.Println(person1)

	person1.Introduction()
	person1.SetName("Juan")
	person1.Introduction()
}
