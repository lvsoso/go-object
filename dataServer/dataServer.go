package main

import (
	"dataServer/heartbeat"
	"log"
	"net/http"
	"os"

	"dataServer/locate"
	"dataServer/objects"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
