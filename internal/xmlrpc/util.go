package xmlrpc

import (
	"io"
	"fmt"
	"encoding/xml"
	"encoding/base64"
	"github.com/xxtea/xxtea-go/xxtea"
    "os"
	"reflect"
	"strings"

	"golang.org/x/net/html/charset"

	"gopkg.in/yaml.v1"
	"io/ioutil"
)

type GenericMethodCall struct {
	XMLName xml.Name `xml:"methodCall"`
	Params []interface{}	`xml:"params>param"`
}

func Decode2 (src *string) error {
	if strings.Contains(*src, "<string/>") {
		*src = ""
		return nil
	}

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

type DockerConfig struct {
	Port uint32 `yaml:"port"`
	Namespace string `yaml:"namespace"`
	Image string `yaml:"image"`
	DataVol1 string `yaml:"dataVol1"`
	Maxplayers int `yaml:"maxplayers"`
	Version string `yaml:"version"`
  }


func ParseConfigYaml(file string) (DockerConfig, error) {
	// filename, _ := filepath.Abs("./file.yml")
	fmt.Printf("PARSE file: %+v\n", file)
	yamlFile, err := ioutil.ReadFile(file)
  
	if err != nil {
	  return DockerConfig{}, err
	}
  
	var config DockerConfig
  
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
	  return DockerConfig{}, err
	}
  
	// fmt.Printf("Value: %#v\n", config.Port)
	// fmt.Printf("Value: %#v\n", config.Namespace)

	return config, nil
}

func ParseEnvYaml(file string) ([]string, error) {
	// filename, _ := filepath.Abs("./file.yml")
	fmt.Printf("PARSE file: %+v\n", file)
	yamlFile, err := ioutil.ReadFile(file)

	envs := []string{}
  
	if err != nil {
	  return envs, err
	}
  
  
	err = yaml.Unmarshal(yamlFile, &envs)
	if err != nil {
	  return envs, err
	}
  
	fmt.Printf("envs: %#v\n", envs)

	return envs, nil
}

func GenerateServiceName(gameId string) string{
	return gameId+"_game"
}
