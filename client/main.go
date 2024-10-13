package main

import (
	"MTG-test-2/client/ws"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	p := path.Dir("./views/index.html")
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)

}
func TestHandler(w http.ResponseWriter, r *http.Request) {
	num := r.URL.Query().Get("num")
	if num == "" {
		num = "1"
	}
	numInt, err := strconv.Atoi(num)
	if err != nil {
		numInt = 1
	}
	ws.Worker(numInt)
}
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/test", TestHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":3000", nil)

}
