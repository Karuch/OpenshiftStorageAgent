package createPOD

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Karuch/OpenshiftStorageAgent/internal/global"
	"github.com/Karuch/OpenshiftStorageAgent/internal/logs"
)

func Request(CompletePVCsMap map[string]int64) ([]byte, error){
	// Define variables

	url := fmt.Sprintf("%s/api/v1/namespaces/%s/pods", global.APIServer, global.Namespace)

	token, err := os.ReadFile(global.TokenFilePath)
	if err != nil {
		e.LogError(err)
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Create a new request using http
	podManifest, err := GetPodManifest(CompletePVCsMap)
	if err != nil {
		e.LogError(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(podManifest)))

	// add authorization header to the req
	req.Header.Add("Authorization", "Bearer "+string(token))
	req.Header.Set("Content-Type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	fmt.Println("Sending request to create pod to:", global.APIServer)
	resp, err := client.Do(req)
	if err != nil {
		e.LogError(err)
	}
	defer resp.Body.Close()
	// I have no idea why, kubernetes return empty response on failure in this request
	// Check if the response status code is not 200
	if resp.StatusCode == 409 {
		return nil, fmt.Errorf("server return unexpected status code: 409, agent-pod already exist?")
	} else if resp.StatusCode != 201 && resp.StatusCode != 200 {
		return nil, fmt.Errorf("server return unexpected status code: %d", resp.StatusCode)
	} else {
		fmt.Println("response statuscode of agent-pod creation:", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e.LogError(err)
	}
	
	return body, err

}
