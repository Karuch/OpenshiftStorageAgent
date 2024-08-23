package query

import (
	"encoding/json"
	"github.com/Karuch/OpenshiftStorageAgent/internal/logs"
)

// Define the structure that matches the JSON
type Metadata struct {
	Name string `json:"name"`
}

type Item struct {
	Metadata Metadata `json:"metadata"`
}

type Response struct {
	Items []Item `json:"items"`
}

func FliterJson() ([]string, error) {

	queryResult, err := Query()
	if err != nil {
		e.LogError(err)
	}

	// Decode JSON into the Response struct
	var response Response
	err = json.Unmarshal(queryResult, &response)
	if err != nil {
		e.LogError(err)
	}

	allPVCs := []string{}
	// Access and print the name from each item
	for _, item := range response.Items {
		allPVCs = append(allPVCs, item.Metadata.Name)
	}

	return allPVCs, err

}