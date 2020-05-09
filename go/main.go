package main

import (
	"github.com/carlescere/scheduler"
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {

	scheduler.Every(20).Seconds().NotImmediately().Run(func(){
		log.Println(updateMysqlEventsFromRedis())
	})

	http.HandleFunc("/api/", apiHandler)

	flag.Parse()
	hub := newHub()
	go hub.run()

	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}