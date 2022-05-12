package main

import (
	"errors"
	"fmt"
	mapTest "go/test/modules/maps"
	routineTest "go/test/modules/routine"
	sliceTest "go/test/modules/slice"
	userPk "go/test/modules/user"
	"sync"
)

func notTheNumber(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

func main() {
	for _, v := range []int{7, 42} {
		if r, e := notTheNumber(v); e != nil {
			fmt.Println("notTheNumber failed:", e)
		} else {
			fmt.Println("notTheNumber worked:", r)
		}
	}

	testsSlice := []func(*sync.WaitGroup){
		userPk.UseStructs,
		routineTest.Routine,
		sliceTest.RunSliceTest,
		mapTest.MapTests,
	}

	var wg sync.WaitGroup
	wg.Add(len(testsSlice))

	for _, myTest := range testsSlice {
		go myTest(&wg)
	}

	wg.Wait()

}
