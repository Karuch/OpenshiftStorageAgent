package createPOD

import (
	"encoding/json"
	"fmt"
	"github.com/Karuch/OpenshiftStorageAgent/internal/logs"
	"github.com/itchyny/gojq"
	"os"
)

func GetPodManifest() (string, error) {
    var manifestPath string = "/go/kubernetes/agent-pod.yaml"

    // Read the JSON file
    jsonData, err := os.ReadFile(manifestPath)
    if err != nil {
        e.LogError(err)
		return "", err
    }

    // Parse the JSON data
    var data interface{}
    if err := json.Unmarshal(jsonData, &data); err != nil {
        e.LogError(err)
        return "", err
    }

    // Define the query to add new volumeMounts to each container
	queryStr := `
	.spec.containers[] |= (
		.volumeMounts += [
			{"mountPath": "/data", "name": "data-volume"},
			{"mountPath": "/logs", "name": "logs-volume"}
		] |
		.env += [
			{"name": "ENV_VAR1", "value": "value1"},
			{"name": "ENV_VAR2", "value": "value2"}
		]
	) |
	.spec.volumes += [
		{"name": "data-volume", "persistentVolumeClaim": {"claimName": "data-pvc"}},
		{"name": "logs-volume", "persistentVolumeClaim": {"claimName": "logs-pvc"}}
	]
	`

    // Parse the query
    query, err := gojq.Parse(queryStr)
    if err != nil {
        e.LogError(err)
        return "", err
    }

    // Execute the query
    iter := query.Run(data)

    // Retrieve and print the updated result
    var updatedData interface{}
    for {
        v, ok := iter.Next()
        if !ok {
            break
        }
        if err, ok := v.(error); ok {
            if err, ok := err.(*gojq.HaltError); ok && err.Value() == nil {
                break
            }
            e.LogError(err)
            return "", err
        }
        updatedData = v
    }

    // Marshal the updated JSON back to a string
    updatedJSON, err := json.MarshalIndent(updatedData, "", "  ")
    if err != nil {
        e.LogError(err)
        return "", err
    }

    // Write the updated JSON back to the file
    err = os.WriteFile("./thing.json", updatedJSON, 0644)
    if err != nil {
        e.LogError(err)
        return "", err
    }

    fmt.Println("Updated JSON written to file:")
    fmt.Println(string(updatedJSON))

	return string(updatedJSON), err
}