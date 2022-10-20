# Setting Up Go Project Example

1. Set up basic dependencies:
    + Automated recompiler when files are updated
        - [CompileDaemon](github.com/githubnemo/CompileDaemon)
    + A dotenv package to help manage environment variables
        - [GoDotEnv](github.com/joho/godotenv)
2. Choose your web framework:
    + e.g. The Gin Web Framework
        - [Gin](github.com/gin-gonic/gin)
3. Choose your database driver:
    + SQL (MySQL and PostgreSQL): e.g. GORM
        - [GORM](https://gorm.io/)
    + NoSQL: e.g. MongoDB for Go
        - [MongoDB](https://www.mongodb.com/docs/drivers/go/current/)

Libraries to check out:
* https://github.com/julienschmidt/httprouter



Dependencies:
1. `github.com/githubnemo/CompileDaemon`
2. `github.com/joho/godotenv`
3. `github.com/gin-gonic/gin`
4. `gorm.io/gorm` 


``` 
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon
go get github.com/joho/godotenv
go get github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```