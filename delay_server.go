package main

import (
	"io"
	"net/http"
	"time"
	"strconv"
)

func delay(w http.ResponseWriter, r *http.Request) {
	d := r.URL.Query()["d"][0]
	d_, _ := strconv.Atoi(d)
	time.Sleep(time.Duration(d_) * time.Millisecond)
	io.WriteString(w, "blah")
}

func main() {
	http.HandleFunc("/", delay)
	http.ListenAndServe(":8000", nil)
}
