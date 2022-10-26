package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func homeHandler3() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		
	return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		fmt.Fprintf(w, "SERVER THREE: This is home...")
	}
}


func contactHandler3() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		fmt.Fprintf(w, "SERVER THREE: This is a contact page!")
	}
}

func notFoundHandler3() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		fmt.Fprintf(w, "SERVER THREE: Page not found!")
	}
}

func routes(rtr *httprouter.Router) { 
	rtr.GET("/", homeHandler3())
	rtr.GET("/contact", contactHandler3())

}

func RunThree(){
	// Create multiplexer
	router := httprouter.New()

	routes(router)

	log.Fatal(http.ListenAndServe(":3333", router))
}