package main

import "net/http"

func main() {
	if err := http.ListenAndServe("127.0.0.1:1211", http.HandlerFunc(handle)); err != nil {
		panic(err)
	}
}

func handle(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("hello http"))
	w.WriteHeader(http.StatusOK)
}
