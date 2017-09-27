package main

import "fmt"

//ping recv-only
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//pong: recv-only, ping send-only
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

//ic_recv_only <- <-ic_send_only

func consumer(ch <-chan int) int {
	return <-ch
}

func producer(i int, ch chan<- int) {
	ch <- i
}

func f() <-chan int {
	// Create a regular, two-way channel.
	c := make(chan int)

	go func() {
		defer close(c)

		// Do stuff
		c <- 1234
	}()

	// Returning it, implicitely converts it to read-only,
	// as per the function return value.
	return c
}

/*
func produce(l *net.TCPListener, c chan<- net.Conn) {
    for {
        conn, _ := l.Accept()
        c<-conn
    }
}

func consume(c <-chan net.Conn) {
    for conn := range c {
        // do something with conn
    }
}

func main() {
    c := make(chan net.Conn, 10)
    for i := 0; i < 10; i++ {
        go consume(c)
    }

    addr := net.TCPAddr{net.ParseIP("127.0.0.1"), 3000}
    l, _ := net.ListenTCP("tcp", &addr)
    produce(l, c)
}
*/

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	ch := make(chan int)
	go producer(42, ch)
	result := consumer(ch)
	fmt.Println("received", result)

	r2 := <-f()
	fmt.Println("received", r2)

}
