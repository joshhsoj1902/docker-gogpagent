package xmlrpc

import (
	// "io"
    // "log"
    "fmt"
    "net/http"
	"encoding/xml"
	"bytes"
	// "strings"

	"golang.org/x/net/html/charset"

	"io/ioutil"
)

type StringParam struct {
	Value string `xml:"value>string"`
}

type AgentService struct{}

func (agent *AgentService) Quick_chk(r *http.Request, args *struct{test string}, reply *struct{Message int}) error {
	fmt.Println("==== quick_chk ====")
	var myResult = 0

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
		myResult = 1
	}

	reply.Message = myResult
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) What_os(r *http.Request, args *struct{test string}, reply *struct{Message string}) error {
	fmt.Println("==== What_os ====")
	var myResult = "1; Linux x86_64"

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Discover_ips(r *http.Request, args *struct{test string}, reply *struct{Message string}) error {
	fmt.Println("==== Discover_ips ====")
	var myResult = "0.0.0.0"

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Cpu_count(r *http.Request, args *struct{test string}, reply *struct{Message int}) error {
	fmt.Println("==== Cpu_count ====")
	var myResult = 1

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Exec(r *http.Request, args *struct{test string }, reply *struct{Message string}) error {
	fmt.Println("==== Exec =====")
	var myResult = "foo"

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}


//     cmd1 := exec.Command(functionName, functionArgs)
//     out, err := cmd1.CombinedOutput()
//     if err != nil {
//         log.Fatalf("cmd.Run() failed with %s\n", err)
//     }
//     // fmt.Printf("combined out:\n%s\n", string(out))

// 	myResult = string(out)

	myEncodedResult, err := Encode(myResult)

// 	anotherResult := fmt.Sprintf("1;%v", myEncodedResult)


	reply.Message = fmt.Sprintf("1;%v", myEncodedResult)
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n\n\n", reply.Message)
    return nil
}


func (agent *AgentService) Ftp_mgr(r *http.Request, args *struct{test string}, reply *struct{Message int}) error {
	fmt.Println("==== Ftp_mgr ====")
	var myResult = 1

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

// WORKING
/// \return 1 If is
/// \return 0 If is not
/// \return -1 If agent could not be reached.
func (agent *AgentService) Is_screen_running(r *http.Request, args *struct{test string}, reply *struct{Message int}) error {
	fmt.Println("==== Is_screen_running ====")
	var myResult = 1

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

/// \returns 0 If file exists
/// \returns 1 If file does not exist
/// \returns -1 If server not available.
func (agent *AgentService) Rfile_exists(r *http.Request, args *struct{test string}, reply *struct{Message int}) error {
	fmt.Println("==== Rfile_exist ====")
	var myResult = 1

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Start_server(r *http.Request, args *struct{test string}, reply *struct{Message int}) error {
	fmt.Println("==== Start_server ====")
	var myResult = 1

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Restart_server(r *http.Request, args *struct{test string}, reply *struct{Message int}) error {
	fmt.Println("==== Restart_server ====")
	var myResult = 1

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Stop_server(r *http.Request, args *struct{test string}, reply *struct{Message int}) error {
	fmt.Println("==== Stop_server ====")
	var myResult = 1

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Get_log(r *http.Request, args *struct{test string}, reply *struct{Message int}) error {
	fmt.Println("==== Get_log ====")
	var myResult = 1

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Readfile(r *http.Request, args *struct{test string}, reply *struct{Message int}) error {
	fmt.Println("==== Readfile ====")
	var myResult = 1

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func (agent *AgentService) Writefile(r *http.Request, args *struct{test string}, reply *struct{Message int}) error {
	fmt.Println("==== Writefile ====")
	var myResult = 1

	err := Decode2(&args.test)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}

	reply.Message = myResult
	
	fmt.Printf("decoded Arg: %v\n", args.test)
	fmt.Printf("reply.Message: %v\n", reply.Message)
    return nil
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
    type MethodCall struct {
		XMLName xml.Name `xml:"methodCall"`
		MethodName string	`xml:"methodName"`
	}

	Body, _ := ioutil.ReadAll(r.Body);

	v := MethodCall{MethodName: ""}

	// Decode Method
    decoder := xml.NewDecoder(bytes.NewReader(Body))
    decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&v)
	fmt.Printf("decoded %s\n", v)
	
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

    fmt.Printf("RPC MethodName %s\n", v.MethodName)
    // fmt.Printf("BODY %s\n\n\n", Body)

	// w.Header().Set("content-type", "application/xml")

	// fmt.Printf("MethodName %s Being Called \n Body: %s", v.MethodName, Body)

	switch v.MethodName {

	default:
		fmt.Printf("MethodName %s NOT SUPPORTED \n Body: %s", v.MethodName, Body)
    }
}