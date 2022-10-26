


```Go 
package serverone

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lstratta/simple-api/cmd/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server1 struct {
	Port	uint16
	DB_Client *mongo.Client
}

func (s *Server1) DbConnect() {
	s.DB_Client = database.ConnectDB()
}

func (s *Server1) DbClose() {
	database.DisconnectDB(s.DB_Client)
}

func (s *Server1) ServeHTTP(w http.ResponseWriter, r *http.Request){
	pathHandler1(w, r)
}


func (s *Server1) RunOne() {
	port := fmt.Sprintf(":%d", s.Port)
	log.Fatal(http.ListenAndServe(port, s))
}

```


```Go
package serverone

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
		fmt.Fprintf(w, "SERVER ONE: HOME PAGE")
}

func contactHandler1(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "SERVER ONE: CONTACT PAGE")
}

func notFoundHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SERVER ONE: Page not found!")
}
```