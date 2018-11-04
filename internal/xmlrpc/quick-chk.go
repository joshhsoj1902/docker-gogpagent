// WORKING
package xmlrpc

import (
	"io"
	"fmt"
	"encoding/xml"
)

type Check struct {
	XMLName xml.Name `xml:"methodCall"`
	MethodCall 	string `xml:"methodName"`
	Value 		string `xml:"params>param>value>string"`
}

type CheckResult struct {
	XMLName   xml.Name `xml:"methodResponse"`
	Param     int      `xml:"params>param>value>int"`
}

func decode_check(body io.Reader) (error, *Check) {
	check := new(Check)
	err := decode_body(body, check)
	// fmt.Printf("checks %+v\n", check)
	if err != nil {
		return err, nil
	}

	err = Decode2(&check.Value)
	if err != nil {
		return err, nil
	}
	return nil, check
}

func quick_chk(body io.Reader) []byte {
	var myResult = 0

	err, check := decode_check(body)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
		myResult = 1
	}
	
	fmt.Printf("value %s\n", check.Value)

	xmlResult := &CheckResult{Param: myResult}

	fmt.Printf("encodedResult: %v\n", xmlResult)

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
}