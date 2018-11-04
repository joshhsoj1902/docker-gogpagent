package xmlrpc

import (
	"io"
	"fmt"
	"encoding/xml"
	"encoding/base64"
	"github.com/xxtea/xxtea-go/xxtea"
    "os"


	"golang.org/x/net/html/charset"
)

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
	decoder2 := xml.NewDecoder(body)
    decoder2.CharsetReader = charset.NewReaderLabel
    err := decoder2.Decode(&o)

	return err
}