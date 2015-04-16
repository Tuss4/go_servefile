package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":5050", "http server port")
var file_location = "/home/tj/Documents/GO_DIR/src/github.com/tuss4/go_servefile/index.html"

func serveIndex(w http.ResponseWriter, req *http.Request) {
	req.Header.Add("Origin", req.Host)
	if req.Method == "GET" {
		http.ServeFile(w, req, file_location)
	}
}
func main() {
	flag.Parse()
	http.HandleFunc("/", serveIndex)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
