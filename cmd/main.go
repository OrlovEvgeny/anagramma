package main

import (
	"anagramma"
	"anagramma/bootstrap"
	"anagramma/httplib"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var addr string

//init
func init() {
	flag.StringVar(&addr, "addr", "localhost:8080", "-addr=localhost:8080")
	flag.Parse()
}

//main
func main() {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		fmt.Println("\nReceived CTRL+C an interrupt, stopping all services...\n")

		//TODO тут должен быть какой-то gracefull
		time.Sleep(time.Second * 1)
		cleanupDone <- struct{}{}
		close(cleanupDone)
	}()

	ctx := boot(context.Background())
	go httplib.HttpServer(ctx, addr)

	<-cleanupDone
	fmt.Println("Exit!")
}

func boot(ctx context.Context) context.Context {
	ctx = bootstrap.Register(ctx, bootstrap.HashMapAlgorithm, anagramma.NewHmap())

	return ctx
}
