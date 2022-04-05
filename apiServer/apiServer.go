package main

import (
	"log"
	"net/http"
	"os"

	"apiServer/heartbeat"
	"apiServer/locate"
	"apiServer/objects"
	"apiServer/temp"
	"apiServer/versions"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
