package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// custom handler
type ironman int

type messageHandler struct {
	message string
}

func (x *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, x.message)
}

func (x ironman) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. Iron!")
}

type thanos int

func (x thanos) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. Thanos!")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><head><title>index</title></head><body><center><p>index</p></center></body></html>"))
	x := r.URL.Query()
	// url.com?id=123
	fmt.Println(x.Get("id"))

	y := r.URL.Path[1:]
	w.Write([]byte(y))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("naberrrr"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index", indexHandler)

	http.HandleFunc("/about", aboutHandler)

	// call custom handler
	var i ironman
	var t thanos

	mux := http.NewServeMux()
	mux.Handle("/ironman", i)
	mux.Handle("/thanos", t)

	x1 := &messageHandler{
		message: "CoderBig",
	}

	x2 := &messageHandler{
		message: "BigCoder",
	}

	mux.Handle("/message-one", x1)
	mux.Handle("/message-two", x2)

	log.Println("Listening ...")

	http.ListenAndServe(":8080", mux)
}
