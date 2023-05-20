package backend

import (
	"fmt"
	"http"
	"logs"
)

func helloWorld(w ResponseWriter, r *Request) {
	fmt.Fprintf(w, "Hello World \n")
}

func Run(addr string) {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server started on port 9003")
	log.Fatal(http.ListenAndServe(addr, nil))
}
