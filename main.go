package main

import (
	"errors"
	"fmt"
	routineTest "go/test/modules/routine"
	userPk "go/test/modules/user"
)

const USER_TYPE string = "natural"

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func f1(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("can't work with 42")

	}

	return arg + 3, nil
}

func main() {
	person1 := userPk.User{Name: "Pepe", Age: 26, UserType: USER_TYPE}

	fmt.Println(person1)

	person1.Introduction()
	person1.SetName("Juan")
	person1.Introduction()

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)

	ae, ok := e.(*argError)

	fmt.Println("print: ", ae, ok)

	if ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	routineTest.Routine()

}
