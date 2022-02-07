package main

import (
	"leclerc/pkg"
	"log"
	"net/http"
)

func main() {

	// handle received requests for following API endpoints (/RequestWebsite /Halt /CheckService )
	// start server
	http.HandleFunc("/", pkg.RequestWebsite)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
