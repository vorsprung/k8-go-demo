package main

import (
	"fmt"
	"net/http"
	"stock"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", stock.SummaryResponse(r))
}
