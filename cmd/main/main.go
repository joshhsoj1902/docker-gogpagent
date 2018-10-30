package main

import (
    "log"
    "github.com/gorilla/handlers"
    "os"
    "net/http"

	"github.com/joshhsoj1902/docker-gogpagent/internal/xmlrpc"

)

func main() {
    log.Println("STARTING Agent")

    http.HandleFunc("/RPC2", xmlrpc.Handler)

    log.Println("Starting XML-RPC server on localhost:12679/RPC2")
    log.Fatal(http.ListenAndServe(":12679", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
