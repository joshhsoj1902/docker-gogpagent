// WORKING
package xmlrpc

import (
	"io"
	"fmt"
	"encoding/xml"
)

type WhatOs struct {
	XMLName xml.Name `xml:"methodCall"`
	MethodCall 	string `xml:"methodName"`
	Value 		string `xml:"params>param>value>string"`
}

type WhatOsResult struct {
	XMLName   xml.Name `xml:"methodResponse"`
	Param     string      `xml:"params>param>value>string"`
}

func decode_what_os(body io.Reader) (error, *Check) {
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

//SHOULD WORK
func what_os(body io.Reader) []byte {
	var myResult = "1; Linux x86_64"

	err, what_os := decode_what_os(body)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	
	fmt.Printf("value %s\n", what_os.Value)

	xmlResult := &WhatOsResult{Param: myResult}

	fmt.Printf("encodedResult: %v\n", xmlResult)

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
}