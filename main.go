package main

import (
	"io"
    "log"
    "github.com/gorilla/handlers"
    "os"
    "fmt"
    "net/http"
	"encoding/xml"
	"encoding/base64"
	"bytes"
	"strings"

	"golang.org/x/net/html/charset"
	"github.com/xxtea/xxtea-go/xxtea"

	"io/ioutil"
	"os/exec"
)

type StringParam struct {
	Value string `xml:"value>string"`
}

func Decode (src string) (string, error) {
	// Decode and Decrypt params
	decodeData1, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}
	decrypt_data := string(xxtea.Decrypt(decodeData1, []byte(os.Getenv("OGP_KEY"))))
	decodeData2, err := base64.StdEncoding.DecodeString(decrypt_data)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}
	return string(decodeData2), nil
}

func Encode (src string) (string, error) {
	// Decode and Decrypt params
	encodeData1 := base64.StdEncoding.EncodeToString([]byte(src))
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// 	return "", err
	// }
	encode_data := string(xxtea.Encrypt([]byte(encodeData1), []byte(os.Getenv("OGP_KEY"))))
	encodeData2 := base64.StdEncoding.EncodeToString([]byte(encode_data))

	return string(encodeData2), nil
}

func quick_chk(body io.Reader) []byte {
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
	fmt.Printf("decoded PARAMS %s\n", v)
	
	if err != nil {
		fmt.Printf("error: %v\n", err)
		myResult = 1
	}

	value, err := Decode(v.Params[0].Value)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		myResult = 1
	}
	fmt.Printf("value %s\n", value)

	xmlResult := &Result{Param: myResult}

	fmt.Printf("encodedResult: %v\n", xmlResult)

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
}

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

    fmt.Printf("MethodName %s\n", v.MethodName)
    fmt.Printf("BODY %s\n\n\n", Body)

	// w.Header().Set("content-type", "application/xml")

	switch v.MethodName {
    case "quick_chk":
		w.Write(quick_chk(bytes.NewReader(Body)))
	case "exec":
		w.Write(rpc_exec(bytes.NewReader(Body)))
	default:
		fmt.Printf("MethodName %s NOT SUPPORTED \n Body: %s", v.MethodName, Body)
    }
}





func main() {
    log.Println("STARTING")

    http.HandleFunc("/RPC2", Handler)

    log.Println("Starting XML-RPC server on localhost:12679/RPC2")
    log.Fatal(http.ListenAndServe(":12679", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
