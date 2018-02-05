package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	log.Printf("hello_svc starting...")

	// start a server to respond on port 3000
	go startHttpServer()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutdown signal received, exiting...")
}

func startHttpServer() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello_svc says 'hello from K8s!'")
	})

	log.Printf("hello_svc running...")

	http.ListenAndServe(":3000", nil)
}
