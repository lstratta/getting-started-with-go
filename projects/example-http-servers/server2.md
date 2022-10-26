``` Go

package server

import (
	"fmt"
	"log"
	"net/http"

)

type pathHandler2 struct {}

func (p pathHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
		case "/":
			homeHandler2(w, r)
		case "/contact":
			contactHandler2(w, r)
		default:
			notFoundHandler2(w, r)
	}
}

func homeHandler2(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "SERVER TWO: This is home...")
}

func contactHandler2(w http.ResponseWriter, r *http.Request){
	
	fmt.Fprintf(w, "SERVER TWO: This is a contact page!")
}

func notFoundHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SERVER TWO: Page not found!")
}


func RunTwo(){

	// create new multiplexer
	mux := http.NewServeMux() 

	pathHandler := pathHandler2{}
	// Handle a function that does not implement the ServerHTTP method
	mux.Handle("/", pathHandler)

	log.Fatal(http.ListenAndServe(":2222", mux))
}

```