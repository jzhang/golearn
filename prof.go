package main

import (
	"os"
	"log"
	"runtime/pprof"
)

func bar() {
	for i := 0; i < 1000000; i++ {
	}
}

func foo() {
	for i := 0; i < 10000000; i++ {
		if i % 1000 == 0 {
			bar()
		}
	}
}

func main() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	foo()
}

