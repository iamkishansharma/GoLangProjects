package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to RACE CONDITION (Mutex & AwaitGroups)")

	// wait group
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(4)
	// go routines
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One routine")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two routine")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three routine")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Four routine")
		m.RLock()
		fmt.Println(score)
		m.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait() // waits for go routines
	fmt.Println(score)
}
