# Structuring Your Go App

The command `cmd` folder is for your application binaries. These are the 'client' for your application's library, `pkg`.

```
# An example of a web app
github.com/lstratta/my-go-app/
    cmd/
        data/ # Initialise database with data
            main.go
        server/ # Run the server 
            main.go 
    pkg/
        


```