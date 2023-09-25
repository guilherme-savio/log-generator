package main

import (
	"io"
	"math/rand"
	"net/http"
	"time"
)

var statusCodesAll = [60]int{
	100, 101, 102, 200, 201, 202, 203, 204, 205, 206,
	207, 208, 226, 300, 301, 302, 303, 304, 305, 306,
	307, 308, 400, 401, 402, 403, 404, 405, 406, 407,
	408, 409, 410, 411, 412, 413, 414, 415, 416, 417,
	418, 422, 423, 424, 426, 428, 429, 431, 451, 500,
	501, 502, 503, 504, 505, 506, 507, 508, 510, 511}

func status(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	foo := statusCodesAll[rand.Intn(len(statusCodesAll))]
	w.WriteHeader(foo)
	io.WriteString(w, "Random HTTP Status Code Generator, written in Go")
}

func main() {
	http.HandleFunc("/", status)
	http.ListenAndServe(":8000", nil)
}
