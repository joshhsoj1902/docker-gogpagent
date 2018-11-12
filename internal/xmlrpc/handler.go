package xmlrpc

import (
	"io"
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

func rpc_start_server(body io.Reader) []byte {
	type MethodCall struct {
		XMLName xml.Name `xml:"methodCall"`
		Params []StringParam	`xml:"params>param"`
	}
	type Result struct {
		XMLName   xml.Name `xml:"methodResponse"`
		Param     int      `xml:"params>param>value>int"`
	}
	var myResult = 1

	v := MethodCall{Params: nil}

    decoder := xml.NewDecoder(body)
    decoder.CharsetReader = charset.NewReaderLabel
    err := decoder.Decode(&v)
	
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	value, err := Decode(v.Params[0].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 1 value %s\n", value)


	value, err = Decode(v.Params[1].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 2 value %s\n", value)


	value, err = Decode(v.Params[2].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 3 value %s\n", value)


	value, err = Decode(v.Params[3].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 4 value %s\n", value)

	value, err = Decode(v.Params[4].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 5 value %s\n", value)

	value, err = Decode(v.Params[5].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 6 value %s\n", value)

	value, err = Decode(v.Params[6].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 7 value %s\n", value)

	value, err = Decode(v.Params[7].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 8 value %s\n", value)




	xmlResult := &Result{Param: myResult}

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
}

func rpc_restart_server(body io.Reader) []byte {
	type MethodCall struct {
		XMLName xml.Name `xml:"methodCall"`
		Params []StringParam	`xml:"params>param"`
	}
	type Result struct {
		XMLName   xml.Name `xml:"methodResponse"`
		Param     int      `xml:"params>param>value>int"`
	}
	var myResult = 1

	v := MethodCall{Params: nil}

    decoder := xml.NewDecoder(body)
    decoder.CharsetReader = charset.NewReaderLabel
    err := decoder.Decode(&v)
	
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	value, err := Decode(v.Params[0].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 1 value %s\n", value)


	value, err = Decode(v.Params[1].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 2 value %s\n", value)


	value, err = Decode(v.Params[2].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 3 value %s\n", value)


	value, err = Decode(v.Params[3].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 4 value %s\n", value)

	value, err = Decode(v.Params[4].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 5 value %s\n", value)

	value, err = Decode(v.Params[5].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 6 value %s\n", value)

	value, err = Decode(v.Params[6].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 7 value %s\n", value)

	value, err = Decode(v.Params[7].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 8 value %s\n", value)

	value, err = Decode(v.Params[8].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 9 value %s\n", value)

	value, err = Decode(v.Params[9].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 10 value %s\n", value)

	value, err = Decode(v.Params[10].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 11 value %s\n", value)

	value, err = Decode(v.Params[11].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 12 value %s\n", value)

	value, err = Decode(v.Params[12].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 13 value %s\n", value)

	value, err = Decode(v.Params[13].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 14 value %s\n", value)

	value, err = Decode(v.Params[14].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 15 value %s\n", value)

	value, err = Decode(v.Params[15].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 16 value %s\n", value)




	xmlResult := &Result{Param: myResult}

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
}

func rpc_stop_server(body io.Reader) []byte {
	type MethodCall struct {
		XMLName xml.Name `xml:"methodCall"`
		Params []StringParam	`xml:"params>param"`
	}
	type Result struct {
		XMLName   xml.Name `xml:"methodResponse"`
		Param     int      `xml:"params>param>value>int"`
	}
	var myResult = 1

	v := MethodCall{Params: nil}

    decoder := xml.NewDecoder(body)
    decoder.CharsetReader = charset.NewReaderLabel
    err := decoder.Decode(&v)
	
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	value, err := Decode(v.Params[0].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 1 value %s\n", value)


	value, err = Decode(v.Params[1].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 2 value %s\n", value)


	value, err = Decode(v.Params[2].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 3 value %s\n", value)


	value, err = Decode(v.Params[3].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 4 value %s\n", value)

	value, err = Decode(v.Params[4].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 5 value %s\n", value)

	value, err = Decode(v.Params[5].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 6 value %s\n", value)

	value, err = Decode(v.Params[6].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 7 value %s\n", value)

	value, err = Decode(v.Params[7].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 8 value %s\n", value)




	xmlResult := &Result{Param: myResult}

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
}

func rpc_get_log(body io.Reader) []byte {
	type MethodCall struct {
		XMLName xml.Name `xml:"methodCall"`
		Params []StringParam	`xml:"params>param"`
	}
	type Result struct {
		XMLName   xml.Name `xml:"methodResponse"`
		Param     string      `xml:"params>param>value>string"`
	}
	var myResult = "1; blah"

	v := MethodCall{Params: nil}

    decoder := xml.NewDecoder(body)
    decoder.CharsetReader = charset.NewReaderLabel
    err := decoder.Decode(&v)
	
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	value, err := Decode(v.Params[0].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 1 value %s\n", value)


	value, err = Decode(v.Params[1].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 2 value %s\n", value)


	value, err = Decode(v.Params[2].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 3 value %s\n", value)


	value, err = Decode(v.Params[3].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 4 value %s\n", value)

	value, err = Decode(v.Params[4].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 5 value %s\n", value)

	value, err = Decode(v.Params[5].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 6 value %s\n", value)

	// value, err = Decode(v.Params[6].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("stop_server 7 value %s\n", value)

	// value, err = Decode(v.Params[7].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("stop_server 8 value %s\n", value)




	xmlResult := &Result{Param: myResult}

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
}

func rpc_readfile(body io.Reader) []byte {
	type MethodCall struct {
		XMLName xml.Name `xml:"methodCall"`
		Params []StringParam	`xml:"params>param"`
	}
	type Result struct {
		XMLName   xml.Name `xml:"methodResponse"`
		Param     string      `xml:"params>param>value>string"`
	}
	var myResult = "1; blah"

	v := MethodCall{Params: nil}

    decoder := xml.NewDecoder(body)
    decoder.CharsetReader = charset.NewReaderLabel
    err := decoder.Decode(&v)
	
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	value, err := Decode(v.Params[0].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 1 value %s\n", value)


	value, err = Decode(v.Params[1].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 2 value %s\n", value)


	// value, err = Decode(v.Params[2].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("stop_server 3 value %s\n", value)


	// value, err = Decode(v.Params[3].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("stop_server 4 value %s\n", value)

	// value, err = Decode(v.Params[4].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("stop_server 5 value %s\n", value)

	// value, err = Decode(v.Params[5].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("stop_server 6 value %s\n", value)

	// value, err = Decode(v.Params[6].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("stop_server 7 value %s\n", value)

	// value, err = Decode(v.Params[7].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("stop_server 8 value %s\n", value)




	xmlResult := &Result{Param: myResult}

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
}

func rpc_writefile(body io.Reader) []byte {
	type MethodCall struct {
		XMLName xml.Name `xml:"methodCall"`
		Params []StringParam	`xml:"params>param"`
	}
	type Result struct {
		XMLName   xml.Name `xml:"methodResponse"`
		Param     string      `xml:"params>param>value>string"`
	}
	var myResult = "1; blah"

	v := MethodCall{Params: nil}

    decoder := xml.NewDecoder(body)
    decoder.CharsetReader = charset.NewReaderLabel
    err := decoder.Decode(&v)
	
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	value, err := Decode(v.Params[0].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 1 value %s\n", value)


	value, err = Decode(v.Params[1].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 2 value %s\n", value)


	value, err = Decode(v.Params[2].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("stop_server 3 value %s\n", value)


	// value, err = Decode(v.Params[3].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("stop_server 4 value %s\n", value)

	// value, err = Decode(v.Params[4].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("stop_server 5 value %s\n", value)

	// value, err = Decode(v.Params[5].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("stop_server 6 value %s\n", value)

	// value, err = Decode(v.Params[6].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("stop_server 7 value %s\n", value)

	// value, err = Decode(v.Params[7].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("stop_server 8 value %s\n", value)




	xmlResult := &Result{Param: myResult}

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
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
	case "start_server":
		w.Write(rpc_start_server(bytes.NewReader(Body)))
	case "restart_server":
		w.Write(rpc_restart_server(bytes.NewReader(Body)))
	case "stop_server":
		w.Write(rpc_stop_server(bytes.NewReader(Body)))
	case "get_log":
		w.Write(rpc_get_log(bytes.NewReader(Body)))
	case "readfile":
		w.Write(rpc_readfile(bytes.NewReader(Body)))
	case "writefile":
		w.Write(rpc_writefile(bytes.NewReader(Body)))
	default:
		fmt.Printf("MethodName %s NOT SUPPORTED \n Body: %s", v.MethodName, Body)
    }
}