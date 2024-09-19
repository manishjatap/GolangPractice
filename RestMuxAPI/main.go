package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/temp", &aa{})
	http.ListenAndServe(":8080", mux)
}

type aa struct{}

func (a aa) ServeHTTP(res http.ResponseWriter, req *http.Request) {

}

// import "net/http"

// func main() {
// 	mux := http.NewServeMux()
// 	mux.Handle("/temp", &handle{})
// 	http.ListenAndServe(":8080", mux)
// }

// type handle struct{}

// func (h *handle) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	res.Write([]byte("dssd"))
// }
