package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
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

	go serve(addr)

	<-cleanupDone
	fmt.Println("Exit!")
}

//serve
func serve(addr string) {
	hm := NewHmap()

	http.HandleFunc("/load", hm.LoadHandler)
	http.HandleFunc("/get", hm.GetHandler)

	log.Printf("http_server start by %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

//Load
func (hm *HashMap) LoadHandler(w http.ResponseWriter, r *http.Request) {
	var wslice []string
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&wslice)
	if err != nil {
		msg := map[string]string{
			"error": "body format invalid, must be json array",
		}
		data, _ := json.Marshal(msg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	tsize := hm.Store(wslice...)

	msg := map[string]int{
		"match_anagramm": tsize,
	}
	data, _ := json.Marshal(msg)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//Get
func (hm *HashMap) GetHandler(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")

	w.Header().Set("Content-Type", "application/json")

	matchSlice := hm.Load(strings.TrimSpace(word))
	data, _ := json.Marshal(matchSlice)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
