package xmlrpc

import (
	"io"
    "log"
    "fmt"
    "net/http"
	"encoding/xml"
	"bytes"
	"strings"

	"golang.org/x/net/html/charset"

	"io/ioutil"
	"os/exec"
)

type StringParam struct {
	Value string `xml:"value>string"`
}

//BUGGY
func rpc_discover_ips(body io.Reader) []byte {
	type MethodCall struct {
		XMLName xml.Name `xml:"methodCall"`
		Params []StringParam	`xml:"params>param"`
	}
	type Result struct {
		XMLName   xml.Name `xml:"methodResponse"`
		Param     string      `xml:"params>param>value>string"`
	}
	var myResult = "0.0.0.0"

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
	fmt.Printf("discover IPS %s\n", value)

	xmlResult := &Result{Param: myResult}

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
}

//NEEDS WORK
func rpc_exec(body io.Reader) []byte {
	type MethodCall struct {
		XMLName xml.Name `xml:"methodCall"`
		Params []StringParam	`xml:"params>param"`
	}
	type Result struct {
		XMLName   xml.Name `xml:"methodResponse"`
		Param     string      `xml:"params>param>value>string"`
	}
	var myResult = "foo"

	v := MethodCall{Params: nil}

    decoder := xml.NewDecoder(body)
    decoder.CharsetReader = charset.NewReaderLabel
    err := decoder.Decode(&v)
	fmt.Printf("decoded PARAMS %#v\n", v)
	
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// Decode Param 1
	value1, err := Decode(v.Params[0].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("value1 %s\n", value1)

	// Decode Param 2
	value2, err := Decode(v.Params[1].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("value2 %s\n", value2)

	execArgs := strings.Split(value1, " ") // Convert func+args into an array
	functionName := execArgs[0]
	execArgs = append(execArgs[:0], execArgs[0+1:]...) // remove the function name from the args list
	functionArgs := strings.Join(execArgs, " ")

	fmt.Printf("functionName: %v\n", functionName)
	fmt.Printf("functionArgs: %v\n", functionArgs)

    cmd1 := exec.Command(functionName, functionArgs)
    out, err := cmd1.CombinedOutput()
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }
    // fmt.Printf("combined out:\n%s\n", string(out))

	myResult = string(out)

	myEncodedResult, err := Encode(myResult)

	anotherResult := fmt.Sprintf("1;%v", myEncodedResult)

	xmlResult := &Result{Param: anotherResult}

	fmt.Printf("encodedResult: %v\n", xmlResult)

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// successfulReturn := []byte("1;")
	// append(s1, 3, 5, 7)

	// TODO: This doesn't seem to be working... Will need to rebug on the PHP side
	// return append([]byte("1;"), enc...)
	return enc

	// return append(successfulReturn[:], enc[:])
	// return fmt.Sprintf("%s%s", "1;", enc)
}

// POC
func rpc_ftp_mgr(body io.Reader) []byte {
	type MethodCall struct {
		XMLName xml.Name `xml:"methodCall"`
		Params []StringParam	`xml:"params>param"`
	}
	type Result struct {
		XMLName   xml.Name `xml:"methodResponse"`
		Param     string      `xml:"params>param>value>string"`
	}
	var myResult = "1; FTP"

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
	fmt.Printf("FTP MGR 1 value %s\n", value)


	value, err = Decode(v.Params[1].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("FTP MGR 2 value %s\n", value)


	value, err = Decode(v.Params[2].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("FTP MGR 3 value %s\n", value)


	value, err = Decode(v.Params[3].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("FTP MGR 4 value %s\n", value)


	value, err = Decode(v.Params[4].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("FTP MGR 5 value %s\n", value)




	xmlResult := &Result{Param: myResult}

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
}




// WORKING
/// \return 1 If is
/// \return 0 If is not
/// \return -1 If agent could not be reached.
func rpc_is_screen_running(body io.Reader) []byte {
	type MethodCall struct {
		XMLName xml.Name `xml:"methodCall"`
		Params []StringParam	`xml:"params>param"`
	}
	type Result struct {
		XMLName   xml.Name `xml:"methodResponse"`
		Param     int      `xml:"params>param>value>int"`
	}
	var myResult = 0 // for now just always ay that it isn't running

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
	fmt.Printf("SCREEN 1 value %s\n", value)


	value, err = Decode(v.Params[1].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("SCREEN 2 value %s\n", value)


	value, err = Decode(v.Params[2].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("SCREEN 3 value %s\n", value)


	// value, err = Decode(v.Params[3].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("SCREEN 4 value %s\n", value)


	// value, err = Decode(v.Params[4].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("SCREEN 5 value %s\n", value)




	xmlResult := &Result{Param: myResult}

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
}

/// \returns 0 If file exists
/// \returns 1 If file does not exist
/// \returns -1 If server not available.
func rpc_rfile_exists(body io.Reader) []byte {
	type MethodCall struct {
		XMLName xml.Name `xml:"methodCall"`
		Params []StringParam	`xml:"params>param"`
	}
	type Result struct {
		XMLName   xml.Name `xml:"methodResponse"`
		Param     int      `xml:"params>param>value>int"`
	}
	var myResult = 0

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
	fmt.Printf("rfile 1 value %s\n", value)


	value, err = Decode(v.Params[1].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("rfile 2 value %s\n", value)


	// value, err = Decode(v.Params[2].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("rfile 3 value %s\n", value)


	// value, err = Decode(v.Params[3].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("rfile 4 value %s\n", value)


	// value, err = Decode(v.Params[4].Value)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }
	// fmt.Printf("rfile 5 value %s\n", value)




	xmlResult := &Result{Param: myResult}

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
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


func Handler(w http.ResponseWriter, r *http.Request) {
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
    case "quick_chk":
		w.Write(quick_chk(bytes.NewReader(Body)))
	case "exec":
		w.Write(rpc_exec(bytes.NewReader(Body)))
	case "what_os":
		w.Write(what_os(bytes.NewReader(Body)))
	case "discover_ips":
		w.Write(rpc_discover_ips(bytes.NewReader(Body)))
	case "ftp_mgr":
		w.Write(rpc_ftp_mgr(bytes.NewReader(Body)))
	case "cpu_count":
		w.Write(cpu_count(bytes.NewReader(Body)))
	case "is_screen_running":
		w.Write(rpc_is_screen_running(bytes.NewReader(Body)))
	case "rfile_exists":
		w.Write(rpc_rfile_exists(bytes.NewReader(Body)))
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