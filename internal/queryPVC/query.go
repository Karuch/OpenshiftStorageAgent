package queryPVC

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/Karuch/OpenshiftStorageAgent/internal/logs"
	"github.com/Karuch/OpenshiftStorageAgent/internal/global"
)

func Query() ([]byte, error){
	// Define variables

	url := fmt.Sprintf("%s/api/v1/namespaces/%s/persistentvolumeclaims", global.APIServer, global.Namespace)

	token, err := os.ReadFile(global.TokenFilePath)
	if err != nil {
		e.LogError(err)
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", "Bearer "+string(token))

	// Send req using http Client
	client := &http.Client{}
	fmt.Println("waiting to response from:", global.APIServer)
	resp, err := client.Do(req)
	if err != nil {
		e.LogError(err)
	}
	defer resp.Body.Close()
	
	// Check if the response status code is not 200
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("server return unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e.LogError(err)
	}
	
	return body, err

}
