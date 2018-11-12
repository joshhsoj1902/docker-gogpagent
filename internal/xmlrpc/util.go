package xmlrpc

import (
	"io"
	"fmt"
	"encoding/xml"
	"encoding/base64"
	"github.com/xxtea/xxtea-go/xxtea"
    "os"
	"reflect"

	"golang.org/x/net/html/charset"
)

type GenericMethodCall struct {
	XMLName xml.Name `xml:"methodCall"`
	Params []interface{}	`xml:"params>param"`
}

func Decode2 (src *string) error {
	// Decode and Decrypt params
	decodeData1, err := base64.StdEncoding.DecodeString(*src)
	if err != nil {
		return err
	}
	decrypt_data := string(xxtea.Decrypt(decodeData1, []byte(os.Getenv("OGP_KEY"))))
	decodeData2, err := base64.StdEncoding.DecodeString(decrypt_data)
	if err != nil {
		return err
	}
	*src = string(decodeData2)
	return nil
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

func decode_body(body io.Reader, o interface{}) error {
	decoder := xml.NewDecoder(body)
    decoder.CharsetReader = charset.NewReaderLabel
    err := decoder.Decode(&o)

	return err
}

func decode_body2(body io.Reader, o interface{}) error {
	gmc := new(GenericMethodCall)

	decoder := xml.NewDecoder(body)
    decoder.CharsetReader = charset.NewReaderLabel
    err := decoder.Decode(&gmc)

	fmt.Printf("gmc %+v\n", gmc)	

	v := reflect.ValueOf(o)

    values := make([]interface{}, v.NumField())

    for i := 0; i < v.NumField(); i++ {
        values[i] = v.Field(i).Interface()
    }

    fmt.Println(values)

	fmt.Printf("FOR DONE %+v\n", gmc.XMLName)	


	return err
}