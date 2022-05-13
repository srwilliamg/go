package main

import (
	"errors"
	"fmt"
	channelsTest "go/test/modules/channels"
	mapTest "go/test/modules/maps"
	routineTest "go/test/modules/routine"
	sliceTest "go/test/modules/slice"
	UnbufferedChannelsTest "go/test/modules/unbufferedChannels"
	userPk "go/test/modules/user"
	"runtime"
	"sync"
)

func notTheNumber(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

func main() {

	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(4)

	for _, v := range []int{7, 42} {
		if r, e := notTheNumber(v); e != nil {
			fmt.Println("notTheNumber failed:", e)
		} else {
			fmt.Println("notTheNumber worked:", r)
		}
	}

	tempTestsSlice := []func(*sync.WaitGroup){
		userPk.UseStructs,
		routineTest.Routine,
		sliceTest.RunSliceTest,
		mapTest.MapTests,
		channelsTest.ChannelsTest,
	}

	fmt.Println("tempTestsSlice:", tempTestsSlice)

	testsSlice := []func(*sync.WaitGroup){
		UnbufferedChannelsTest.UnbufferedChannelsTest,
	}

	// testsSlice = append(testsSlice, tempTestsSlice...)

	var wg sync.WaitGroup
	wg.Add(len(testsSlice))

	for _, myTest := range testsSlice {
		go myTest(&wg)
	}

	wg.Wait()

}
