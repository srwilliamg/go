package routine

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f(from string) {
	defer wg.Done()
	for i := 0; i < len(from); i++ {
		fmt.Println(from, ":", i)
	}
}

func fRoutine(from string) {
	for i := 0; i < len(from); i++ {
		wg.Add(1)
		go f(fmt.Sprintf("%d%s", i, from))
	}
}

func Routine() {
	wg.Add(3)
	go f("goroutine")
	go f("direct")
	go func(msg string) {
		defer wg.Done()
		fmt.Println("in anonymous function")
		go fRoutine(msg)
	}("going")

	wg.Wait()
	fmt.Println("done")
}
