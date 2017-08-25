package main

import "fmt"
import "context"
import "time"

func add(a, b int) <-chan int {
	sum := make(chan int)
	go func() {
		sum <- a + b
	}()
	return sum
}

func fanIn(sum1, sum2 <-chan int) <-chan int {
	sum := make(chan int)
	go func() {
		for {
			sum <- <-sum1
		}
	}()
	go func() {
		for {
			sum <- <-sum2
		}
	}()
	return sum
}

func fanIn2(sum1, sum2 <-chan int) <-chan int {
	sum := make(chan int)
	go func() {
		for {
			select {
			case c := <-sum1:
				sum <- c
			case c := <-sum2:
				sum <- c
			}
		}
	}()
	return sum
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("监控中...")
		}
	}
}

func main() {
	sum12 := add(1, 2)
	sum49 := add(4, 9)
	fmt.Println(<-sum12)
	fmt.Println(<-sum49)

	sum := fanIn(add(1, 2), add(4, 9))
	fmt.Println(<-sum)
	fmt.Println(<-sum)

	sum2 := fanIn2(add(1, 2), add(4, 9))
	fmt.Println(<-sum2)
	fmt.Println(<-sum2)

	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	go watch(ctx)
	go watch(ctx)
	time.Sleep(1 * time.Second)
	cancel()
	fmt.Println("结束啦。。。")
}
