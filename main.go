package main

import (
	"fmt"
	"sync"
	"time"
)

// The Dining Philosopher problem is well known in computer science circles.
// Five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty - beides those of philosophy - is that the dish
// server is a very diffucult kind of spaghetti which has to  be eacten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however this mean that no two neighbours
// may be eating simultaneously, since there are five philosophers and five forks
//
// This is a simple implementation of Dijkstra's solution to the "Dining
// Philosophers" dilema.

// Philosopher is a struct which store information about a philosopher.
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// list of all philosophers
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// define some variables
var hunger = 3                  // how many times does a person eat
var eatTime = 1 * time.Second   // how long it takes to eatTime
var thinkTime = 3 * time.Second // how long a philosopher thinks
var sleepTime = 1 * time.Second // how long to wait when printing things out

// *** added this
var orderMutex sync.Mutex   // a mutex for the slice orderFinished; part of challange
var orederFinished []string // the order in which philosophers dining and leave; part of challange

func main() {
	// print a welcome message
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("---------------------------")
	fmt.Println("The table is empty.")

	// *** added this
	time.Sleep(sleepTime)

	// start the meal
	dine()

	// print out finished message
	fmt.Println("------------------")
	fmt.Println("The table is empty.")
}

func dine() {
	// eatTime = 0 * time.Second
	// sleepTime = 0 * time.Second
	// thinkTime = 0 * time.Second

	// check if all have eaten
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// check if all are seated at the table
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// fire off a go routine for a current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()

	seated.Wait()

	// eat three times(hunger is const for three times)
	for i := hunger; i > 0; i-- {
		// get a lock on both forks
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.name)
	}

	fmt.Println(philosopher.name, " is satisfied")
	fmt.Println(philosopher.name, " left the table")

	// *** added this
	orderMutex.Lock()
	orederFinished = append(orederFinished, philosopher.name)
	orderMutex.Unlock()
}
