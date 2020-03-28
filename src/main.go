package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, _ *http.Request) {
	reason := getRandomOpener()
	fmt.Fprintf(w, "%s", reason.content )
}

