package query

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Query() []byte {
	// Define variables

	tokenFilePath := "/go/kubernetes/token.txt"
	apiServer := "https://192.168.49.2:8443"
	namespace := "default"
	url := fmt.Sprintf("%s/api/v1/namespaces/%s/persistentvolumeclaims", apiServer, namespace)

	token, err := os.ReadFile(tokenFilePath)
	if err != nil {
		log.Fatal(err)
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", "Bearer "+string(token))

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	
	return body

}
