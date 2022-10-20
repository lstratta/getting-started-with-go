# Building a REST API in Go

1. Set up dependency management
`go mod init example/todo-go`

2. Install gin
`go get github.com/gin-gonic/gin`

These two steps will create a `go.mod` and `go.sum` file

**nside the main.go file**
3. Create a todo struct

```
package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

type todo struct {
  ID        string  `json: "id"`
  Item      string  `json: "title"`
  Completed bool    `json: "completed"`
}

var todos = []todo {
  {ID: "1", Item: "Get a job", Completed: false},
  {ID: "2", Item: "Learn Dutch", Completed: false},
  {ID: "3", Item: "Learn Go", Completed: false},
}

```

4. Creating the router inside the main function:

```
func main() {
  router := gin.Default()
  router.Run("localhost:9090")
}
```

5. Creating a HTTP endpoint

Add the `GET` request to `main()`


```
func main() {
  router := gin.Default()
  router.GET("/todos")
  router.Run("localhost:9090")
}
```

Above `main()`, create `getTodos()` function

```
func getTodos(context *gin.Context)
```


