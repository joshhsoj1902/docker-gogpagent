package main

import (
    "log"
    "github.com/gorilla/handlers"
    "os"
    "fmt"
    "net/http"
    "github.com/gorilla/rpc"
    "github.com/divan/gorilla-xmlrpc/xml"

    "io/ioutil"
)

type HelloArgs struct {
    Who string
}

type HelloReply struct {
    Message string
}

type HelloService struct{}

func (h *HelloService) Say(r *http.Request, args *HelloArgs, reply *HelloReply) error {
    log.Println("Say", args.Who)
    reply.Message = "Hello, " + args.Who + "!"
    return nil
}







// type Message struct {
// 	Id   int64  `json:"id"`
// 	Name string `json:"name"`
// }

// curl localhost:8000 -d '{"name":"Hello"}'
func Cleaner(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
  log.Println("Body", b)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	// var msg Message
	// err = json.Unmarshal(b, &msg)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
  //
	// output, err := json.Marshal(msg)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	// w.Header().Set("content-type", "application/xml")
	// w.Write(output)
}










func main() {
    fmt.Println("STARTING")
    RPC := rpc.NewServer()
    xmlrpcCodec := xml.NewCodec()
    xmlrpcCodec.RegisterAlias("quick_chk", "HelloService.Say")
    RPC.RegisterCodec(xmlrpcCodec, "text/xml")
    RPC.RegisterService(new(HelloService), "")
    // http.Handle("/RPC2", RPC)
    http.HandleFunc("/RPC2", Cleaner)


    log.Println("Starting XML-RPC server on localhost:12679/RPC2")
    log.Fatal(http.ListenAndServe(":12679", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
