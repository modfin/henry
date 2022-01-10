package main

import (
	"fmt"
	"github.com/modfin/go18exp/containerz/heap"
	"math/rand"
	"sync"
)

type Work struct {
	id       int
	priority int
}

func (w Work) Do() {
	fmt.Printf("- Consuming work with prio on %d and id %d\n", w.priority, w.id)
}

func main() {

	wg := sync.WaitGroup{}

	priorityQueue := heap.New[Work](func(a, b Work) bool {
		return a.priority < b.priority
	})

	producer := func() {
		for id := range make([]int, 10) {
			prio := rand.Intn(10)
			fmt.Printf("+ Producing work with prio on %d and id %d\n", prio, id)
			priorityQueue.Push(Work{
				id:       id,
				priority: prio,
			})

		}
		wg.Done()
	}

	consumer := func() {
		for range make([]int, 10) {
			// Spinnlock waiting for reads
			for priorityQueue.Len() == 0 {
			}
			work := priorityQueue.Pop()
			work.Do()

		}
		wg.Done()
	}

	wg.Add(1)
	go producer()
	wg.Add(1)
	go consumer()

	wg.Wait()

}
