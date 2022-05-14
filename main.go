package main

import (
	"fmt"
	channelsTest "go/test/modules/channels"
	mapTest "go/test/modules/maps"
	pollingTest "go/test/modules/polling"
	routineTest "go/test/modules/routine"
	runnerTest "go/test/modules/runner"
	sliceTest "go/test/modules/slice"
	UnbufferedChannelsTest "go/test/modules/unbufferedChannels"
	userPk "go/test/modules/user"
	"runtime"
	"sync"
)

func main() {

	// fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1)

	tempTestsSlice := []func(*sync.WaitGroup){
		userPk.UseStructs,
		routineTest.Routine,
		sliceTest.RunSliceTest,
		mapTest.MapTests,
		channelsTest.ChannelsTest,
		UnbufferedChannelsTest.UnbufferedChannelsTest,
		runnerTest.RunRunnerTest,
	}

	fmt.Println("tempTestsSlice:", tempTestsSlice)

	testsSlice := []func(*sync.WaitGroup){
		pollingTest.PollingTest,
	}

	// testsSlice = append(testsSlice, tempTestsSlice...)

	var wg sync.WaitGroup
	wg.Add(len(testsSlice))

	for _, myTest := range testsSlice {
		go myTest(&wg)
	}

	wg.Wait()

}
