package main

import (
    "log"
    "github.com/gorilla/handlers"
    "os"
    "fmt"
    "net/http"
	"encoding/xml"
	"encoding/base64"

    "github.com/gorilla/rpc"
    gorillaXml "github.com/divan/gorilla-xmlrpc/xml"
	"golang.org/x/net/html/charset"
	"github.com/xxtea/xxtea-go/xxtea"

    // "io/ioutil"
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
    type MethodCall struct {
		XMLName xml.Name `xml:"methodCall"`
		MethodName string	`xml:"methodName"`
		Param string		`xml:"params>param>value>string"`
	}
	
	type Result struct {
		XMLName xml.Name `xml:"methodCall"`
		MethodName string	`xml:"methodName"`
		Param string		`xml:"params>param>value>string"`
    }

	v := MethodCall{MethodName: "", Param: ""}

    decoder := xml.NewDecoder(r.Body)
    decoder.CharsetReader = charset.NewReaderLabel
    err := decoder.Decode(&v)
    fmt.Printf("decoded %s\n", v)


	if err != nil {
		fmt.Printf("error: %v\n", err)
		// return
	}

    fmt.Printf("MethodName %s\n", v.MethodName)
	fmt.Printf("Param %s\n", v.Param)
	
	decodeData1, err := base64.StdEncoding.DecodeString(v.Param)
	fmt.Printf("decodeData1 %s\n", decodeData1)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		// return
	}

	decrypt_data := string(xxtea.Decrypt(decodeData1, []byte(os.Getenv("OGP_KEY"))))
	fmt.Printf("decrypt_data %s\n", decrypt_data)

	decodeData2, err := base64.StdEncoding.DecodeString(decrypt_data)
	fmt.Printf("decodeData2 %s\n", decodeData2)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		// return
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


	// enc := xml.NewEncoder(os.Stdout)
	// enc.Indent("  ", "    ")
	// if err := enc.Encode(v); err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }

	// w.Header().Set("content-type", "application/xml")
	w.Write([]byte("1"))
}










func main() {
    log.Println("STARTING")
    RPC := rpc.NewServer()
    xmlrpcCodec := gorillaXml.NewCodec()
    xmlrpcCodec.RegisterAlias("quick_chk", "HelloService.Say")
    RPC.RegisterCodec(xmlrpcCodec, "text/xml")
    RPC.RegisterService(new(HelloService), "")
    // http.Handle("/RPC2", RPC)
    http.HandleFunc("/RPC2", Cleaner)

    log.Println("Starting XML-RPC server on localhost:12679/RPC2")
    log.Fatal(http.ListenAndServe(":12679", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
