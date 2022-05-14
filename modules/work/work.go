// This sample program demonstrates how to use the work package
// to use a pool of goroutines to get work done.
package work

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Worker must be implemented by types that want to use
// the work pool.
type Worker interface {
	TaskType() string
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
		if pool.workerType.typeSync == w.TaskType() {
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

// names provides a set of names to display.
var names = []namePrinter{
	{name: "Morten", taskType: "standard"},
	{name: "Abigael", taskType: "standard"},
	{name: "Park", taskType: "standard"},
	{name: "Kalvin", taskType: "standard"},
	{name: "Kleon", taskType: "standard"},
	{name: "Jerry", taskType: "standard"},
	{name: "Emerson", taskType: "standard"},
	{name: "Teddy", taskType: "standard"},
	{name: "Anestassia", taskType: "standard"},
	{name: "Brenn", taskType: "standard"},
	{name: "Loralie", taskType: "standard"},
	{name: "Kelwin", taskType: "standard"},
	{name: "Kimble", taskType: "standard"},
	{name: "Arleta", taskType: "standard"},
	{name: "Juieta", taskType: "standard"},
	{name: "Giustino", taskType: "standard"},
	{name: "Carita", taskType: "standard"},
	{name: "Kelbee", taskType: "standard"},
	{name: "Dosi", taskType: "standard"},
	{name: "Davidson", taskType: "standard"},
	{name: "Orv", taskType: "standard"},
	{name: "Davin", taskType: "standard"},
	{name: "Kandy", taskType: "standard"},
	{name: "Elie", taskType: "standard"},
	{name: "Alaster", taskType: "standard"},
	{name: "Terence", taskType: "standard"},
	{name: "Debee", taskType: "standard"},
	{name: "Neall", taskType: "standard"},
	{name: "Frieda", taskType: "standard"},
	{name: "Shelden", taskType: "standard"},
	{name: "Brina", taskType: "standard"},
	{name: "Thomas", taskType: "standard"},
	{name: "Niall", taskType: "standard"},
	{name: "Vincents", taskType: "standard"},
	{name: "Ewart", taskType: "standard"},
	{name: "Randall", taskType: "standard"},
	{name: "Jandy", taskType: "standard"},
	{name: "Page", taskType: "standard"},
	{name: "Manda", taskType: "standard"},
	{name: "Marcia", taskType: "standard"},
	{name: "Torie", taskType: "standard"},
	{name: "Raymond", taskType: "standard"},
	{name: "Roderic", taskType: "standard"},
	{name: "Vaclav", taskType: "standard"},
	{name: "Vinita", taskType: "standard"},
	{name: "Cy", taskType: "standard"},
	{name: "Rafa", taskType: "standard"},
	{name: "Wynny", taskType: "standard"},
	{name: "Dunstan", taskType: "standard"},
	{name: "Nero", taskType: "standard"},
	{name: "Curt-Most-Important", taskType: "most important"},
	{name: "Chancey", taskType: "standard"},
	{name: "Kinny", taskType: "standard"},
	{name: "York", taskType: "standard"},
	{name: "Charis", taskType: "standard"},
	{name: "Stevana", taskType: "standard"},
	{name: "Gilemette-Important", taskType: "important"},
	{name: "Christyna-Important", taskType: "important"},
	{name: "Galen-Most-Important", taskType: "most important"},
	{name: "Jeanne-Important", taskType: "important"},
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	name     string
	taskType string
}

// Task implements the Worker interface.
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func (m *namePrinter) TaskType() string {
	return m.taskType
}

// main is the entry point for all Go programs.
func WorkTests(wgg *sync.WaitGroup) {
	defer wgg.Done()

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
		}(np)
	}
	// }

	wg.Wait()

	// Shutdown the work pool and wait for all existing work
	// to be completed.
	pm.Shutdown()
}
