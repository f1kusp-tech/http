package main

import (
	"fmt"
	"net/http"
)

func handlerNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "404 - Not Found")
}
