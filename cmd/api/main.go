package main

import (
	"github.com/tinrab/go-k8s-background-job/api"
	"log"
	"os"
	"strconv"
)

func main() {
	port := uint16(8080)

	if v, err := strconv.ParseUint(os.Getenv("APP_PORT"), 10, 16); err == nil {
		port = uint16(v)
	}

	t := api.NewTransport(port)
	if err := t.Run(); err != nil {
		log.Fatal(err)
	}
}
