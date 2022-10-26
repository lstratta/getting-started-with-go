# HTTP Server using a user defined server struct

In this example, we create a struct that defines some details about the server (port number and holds details of the active MongoDB connection client) and implement the `ServeHTTP` method to make it a http.handler.

```Go 
// server.go

package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lstratta/simple-api/cmd/database"
	"go.mongodb.org/mongo-driver/mongo"
)

// Our server struct
// This struct takes a port number which can be assigned by an environment variable or
// by hard coding it when creating the struct 
type Server struct {
	Port	uint16
	DB_Client *mongo.Client
}

// ServeHTTP is the method we need to include so that Server interfaces
// with the http.Handler interface. It contains the handler that handles the
// paths 
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	pathHandler1(w, r)
}

// DbConnect is an access function to the database.go file which is below
// This might be better off structured in a different way where there is a separate 
// set of functions that can be used to access the database functions.
func (s *Server) DbConnect() {
	s.DB_Client = database.ConnectDB()
}

func (s *Server) DbClose() {
	database.DisconnectDB(s.DB_Client)
	s.DB_Client = nil
}

// Run is called by the main method of the application to start the server
func (s *Server) Run() {
	port := fmt.Sprintf(":%d", s.Port)
	log.Fatal(http.ListenAndServe(port, s))
}

```


```Go
package server

import (
	"fmt"
	"net/http"
)

func pathHandler1(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path {
		case "/":
			homeHandler1(w, r)
		case "/contact":
			contactHandler1(w, r)
		default:
			notFoundHandler1(w, r)
	}
}

func homeHandler1(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "SERVER: Home Page")
}

func contactHandler1(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "SERVER: Contact Page")
}

func notFoundHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SERVER: Page not found!")
}
```

```Go
package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


func ConnectDB() *mongo.Client {

	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set a MONGODB_URI environment variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil{
		log.Panicln("ConnectDB Error:",err)
	}

	return client

}

func DisconnectDB(client *mongo.Client) {

	if err := client.Disconnect(context.TODO()); err != nil {
		log.Panicln("DisconnectDB Error: ", err)
	}

}

func PingDB(client *mongo.Client){

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Panicln("PingDB Error:", err)
	}

	log.Println("Database connected succesfully")
}
```