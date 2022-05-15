// This sample program demonstrates how to use the work package
// to use a pool of goroutines to get work done.
package work

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
	"time"
)

// Worker must be implemented by types that want to use
// the work pool.
type Worker interface {
	GetTaskType() string
	Task()
}

// Pool provides a pool of goroutines that can execute any Worker
// tasks that are submitted.
type WorkerType struct {
	typeSync string
	quantity int
}

type Pool struct {
	work       chan Worker
	workerType WorkerType
}

type PollManager struct {
	Pools []Pool
	Wg    sync.WaitGroup
}

// New creates a new work pool.
func New(maxGoroutines int, workerTypes []WorkerType) *PollManager {
	fmt.Printf("workerTypes: %v, %d\n", workerTypes, len(workerTypes))

	pools := make([]Pool, 0)

	currentGoroutines := maxGoroutines

	for _, wt := range workerTypes {
		currentGoroutines = currentGoroutines - wt.quantity
		pools = append(pools, Pool{
			work:       make(chan Worker),
			workerType: wt,
		})
	}

	pools = append(pools, Pool{
		work: make(chan Worker),
		workerType: WorkerType{
			typeSync: "standard",
			quantity: currentGoroutines,
		},
	})

	pollManager := &PollManager{
		Pools: pools,
	}

	fmt.Printf("PollManager.pools: %v\n", pollManager.Pools)
	pollManager.Wg.Add(maxGoroutines)
	for _, pool := range pollManager.Pools {
		for i := 0; i < pool.workerType.quantity; i++ {
			go func(pool Pool) {
				for w := range pool.work {
					w.Task()
				}
				pollManager.Wg.Done()
			}(pool)
		}
	}

	return pollManager
}

// Run submits work to the pool.
func (p *PollManager) Run(w Worker) {
	for _, pool := range p.Pools {
		if pool.workerType.typeSync == w.GetTaskType() {
			pool.work <- w
		}
	}
}

// Shutdown waits for all the goroutines to shutdown.
func (p *PollManager) Shutdown() {
	for _, pool := range p.Pools {
		close(pool.work)
	}
	p.Wg.Wait()
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	Name     string `json:"name"`
	TaskType string `json:"task_type"`
}

// Task implements the Worker interface.
func (m *namePrinter) Task() {
	log.Println(m.Name)
	time.Sleep(time.Second)
}

func (m *namePrinter) GetTaskType() string {
	return m.TaskType
}

// main is the entry point for all Go programs.
func WorkTests(wgg *sync.WaitGroup) {
	defer wgg.Done()

	path := filepath.Join("data", "name-type.json")
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("jsonFile: %v\n", jsonFile)

	var names []*namePrinter

	err = json.Unmarshal(jsonFile, &names)
	if err != nil {
		log.Fatal(err)
	}

	workerTypes := []WorkerType{{typeSync: "important", quantity: 10}, {typeSync: "most important", quantity: 5}}
	// Create a work pool with 2 goroutines.
	pm := New(30, workerTypes)

	var wg sync.WaitGroup
	wg.Add(len(names))

	// for i := 0; i < 2; i++ {
	// Iterate over the slice of names.
	for _, np := range names {
		// Create a namePrinter and provide the
		go func(np namePrinter) {
			// Submit the task to be worked on. When RunTask
			// returns we know it is being handled.
			pm.Run(&np)
			wg.Done()
		}(*np)
	}
	// }

	wg.Wait()

	// Shutdown the work pool and wait for all existing work
	// to be completed.
	pm.Shutdown()
}
