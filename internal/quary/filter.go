package query

import (
	"encoding/json"
	"fmt"
	"log"
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

func FliterJson() {
	// Decode JSON into the Response struct
	var response Response
	err := json.Unmarshal(Query(), &response)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	all_pvcs := []string{}
	// Access and print the name from each item
	for _, item := range response.Items {
		fmt.Println("PVC Name:", item.Metadata.Name)
		all_pvcs = append(all_pvcs, item.Metadata.Name)
	}

	fmt.Println(all_pvcs)
}