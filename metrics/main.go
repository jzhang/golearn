//main.go
package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/bmhatfield/go-runtime-metrics"
)

func main() {
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    ":8000", // Normally ":443"
		Handler: http.FileServer(http.Dir(cwd)),
	}
	log.Fatal(srv.ListenAndServe())
}
