package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// loading config at startup once is fine for now
var appConfig = func() *config {
	c, err := getConfig()
	if err != nil {
		panic(err)
	}
	return c
}()

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!", string(appConfig.AllowedHostname[0]))
}
