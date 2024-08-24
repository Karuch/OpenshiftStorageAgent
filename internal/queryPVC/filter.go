package queryPVC

import (
	"encoding/json"
	"fmt"
	"github.com/Karuch/OpenshiftStorageAgent/internal/logs"
)

// Define the structure that matches the JSON
type Metadata struct {
	Name string `json:"name"`
}

type Capacity struct {
	Storage string `json:"storage"`
}

type Status struct {
	Capacity Capacity `json:"capacity"`
}

type Item struct {
	Metadata Metadata `json:"metadata"`
	Status   Status   `json:"status"`
}

type Response struct {
	Items []Item `json:"items"`
}

func FliterJson() (map[string]string, error) {
	// Query function to get the JSON result
	queryResult, err := Query()
	if err != nil {
		e.LogError(err)
		return nil, err
	}

	// Decode JSON into the Response struct
	var response Response
	err = json.Unmarshal(queryResult, &response)
	if err != nil {
		e.LogError(err)
		return nil, err
	}

	// Initialize a map to store PVC names and their capacities
	allPVCsMap := make(map[string]string)

	// Populate the map with name and capacity if storage is present
	for _, item := range response.Items {
		name := item.Metadata.Name
		capacity := item.Status.Capacity.Storage

		// Check if the capacity is not empty
		if capacity != "" {
			allPVCsMap[name] = capacity
		}
	}

	fmt.Println(allPVCsMap)
	return allPVCsMap, nil
}
