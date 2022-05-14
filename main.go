package main

import (
	"fmt"
	channelsTest "go/test/modules/channels"
	mapTest "go/test/modules/maps"
	routineTest "go/test/modules/routine"
	runnerTest "go/test/modules/runner"
	sliceTest "go/test/modules/slice"
	UnbufferedChannelsTest "go/test/modules/unbufferedChannels"
	userPk "go/test/modules/user"
	"sync"
)

func main() {

	// fmt.Println(runtime.NumCPU())
	// runtime.GOMAXPROCS(4)

	tempTestsSlice := []func(*sync.WaitGroup){
		userPk.UseStructs,
		routineTest.Routine,
		sliceTest.RunSliceTest,
		mapTest.MapTests,
		channelsTest.ChannelsTest,
		UnbufferedChannelsTest.UnbufferedChannelsTest,
	}

	fmt.Println("tempTestsSlice:", tempTestsSlice)

	testsSlice := []func(*sync.WaitGroup){
		runnerTest.RunRunnerTest,
	}

	// testsSlice = append(testsSlice, tempTestsSlice...)

	var wg sync.WaitGroup
	wg.Add(len(testsSlice))

	for _, myTest := range testsSlice {
		go myTest(&wg)
	}

	wg.Wait()

}
