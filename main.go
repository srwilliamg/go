package main

import (
	"errors"
	"fmt"
	routineTest "go/test/modules/routine"
	userPk "go/test/modules/user"
)

func notTheNumber(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

func main() {
	userPk.UseStructs()

	for _, v := range []int{7, 42} {
		if r, e := notTheNumber(v); e != nil {
			fmt.Println("notTheNumber failed:", e)
		} else {
			fmt.Println("notTheNumber worked:", r)
		}
	}

	routineTest.Routine()

}
