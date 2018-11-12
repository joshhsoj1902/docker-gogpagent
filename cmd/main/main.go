package main

import (
    "log"
    "github.com/gorilla/handlers"
    "os"
    "net/http"

    "github.com/joshhsoj1902/docker-gogpagent/internal/xmlrpc"
    "github.com/gorilla/rpc"
    "github.com/divan/gorilla-xmlrpc/xml"

)
func RpcHandler() {


    log.Println("Starting XML-RPC server on localhost:1234/RPC2")
    log.Fatal(http.ListenAndServe(":1234", nil))
}

func main() {
    log.Println("STARTING Agent")

    RPC := rpc.NewServer()
    xmlrpcCodec := xml.NewCodec()
    
    RPC.RegisterCodec(xmlrpcCodec, "text/xml")
    RPC.RegisterService(new(xmlrpc.AgentService), "agent")
    http.Handle("/RPC2", RPC)


    // http.HandleFunc("/RPC2", xmlrpc.HttpHandler)

    log.Println("Starting XML-RPC server on localhost:12679/RPC2")
    log.Fatal(http.ListenAndServe(":12679", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
