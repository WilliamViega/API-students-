package main

import (
    "log"

    "github.com/WilliamViega/API-students/api"
   
)

func main() {
    server := api.NewServer()
    server.ConfigureRoute()
    if err := server.Start(); err != nil {
        log.Fatal(err)
    } 
}
