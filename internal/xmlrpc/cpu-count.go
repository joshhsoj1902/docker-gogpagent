// Should work
package xmlrpc

import (
	"io"
	"fmt"
	"encoding/xml"
)

type CpuCount struct {
	XMLName xml.Name `xml:"methodCall"`
	MethodCall 	string `xml:"methodName"`
	Value 		string `xml:"params>param>value>string"`
}

type CpuCountResult struct {
	XMLName   xml.Name `xml:"methodResponse"`
	Param     int      `xml:"params>param>value>int"`
}

func decode_cpu_count(body io.Reader) (error, *CpuCount) {
	cpu_count := new(CpuCount)
	err := decode_body(body, cpu_count)
	// fmt.Printf("checks %+v\n", check)
	if err != nil {
		return err, nil
	}

	err = Decode2(&cpu_count.Value)
	if err != nil {
		return err, nil
	}
	return nil, cpu_count
}

func cpu_count(body io.Reader) []byte {
	var myResult = 1

	err, what_os := decode_cpu_count(body)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
	}
	
	fmt.Printf("value %s\n", what_os.Value)

	xmlResult := &CpuCountResult{Param: myResult}

	fmt.Printf("encodedResult: %v\n", xmlResult)

	enc, err := xml.MarshalIndent(xmlResult, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return enc
}