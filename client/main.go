package main

import (
	"MTG-test-2/client/ws"
	"fmt"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	p := path.Dir("./views/index.html")
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)

}
func TestHandler(w http.ResponseWriter, r *http.Request) {
	//num := r.URL.Query().Get("num")
	ws.Worker()
}
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/test", TestHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":3000", nil)

}
