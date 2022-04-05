package main

import (
	"dataServer/heartbeat"
	"log"
	"net/http"
	"os"

	"dataServer/locate"
	"dataServer/objects"
	"dataServer/temp"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
