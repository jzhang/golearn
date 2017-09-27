package main

import (
	"fmt"
	"time"
)

type work struct {
	Fn        func()
	Completed bool
}

func newwork(fn func()) *work {
	return &work{fn, false}
}

func worker(in chan *work, out chan *work) {
	for {
		t := <-in
		t.Fn()
		t.Completed = true
		out <- t
	}
}

func masterrun(fn func(), jobs int, concurrency int) {
	pending := make(chan *work)
	done := make(chan *work)

	go func() {
		for i := 0; i < jobs; i++ {
			pending <- newwork(fn)
		}
	}()

	for i := 0; i < concurrency; i++ {
		go worker(pending, done)
	}

	for i := 0; i < jobs; i++ {
		<-done
	}
}

func main() {

	printTime := func() {
		now := time.Now()
		fmt.Println(now.String())
	}

	masterrun(printTime, 100, 10)

}
