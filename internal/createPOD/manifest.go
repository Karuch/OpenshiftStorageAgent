package createPOD

import (
    "bytes"
    "encoding/json"
    "github.com/itchyny/gojq"
    "os"
    "text/template"
    "github.com/Karuch/OpenshiftStorageAgent/internal/logs" 
	"strings"
)

func counter() func() int {
    i := -1
    return func() int {
        i++
        return i
    }
}

// replaceDash replaces all occurrences of '-' with '_'.
func replaceHyphen(s string) string {
    return strings.ReplaceAll(s, "-", "xxHYPHENxxCHARxx")
}

func GetPodManifest(pvcMap map[string]int64) ([]byte, error) {
    var manifestPath string = "/go/kubernetes/agent-pod.json" // Path to your file in the current directory

    // Read the JSON file
    jsonData, err := os.ReadFile(manifestPath)
    if err != nil {
        e.LogError(err) // Retain e.LogError for logging
        return nil, err
    }

    // Unmarshal the JSON data into a map
    var data map[string]interface{}
    err = json.Unmarshal(jsonData, &data)
    if err != nil {
        e.LogError(err) // Retain e.LogError for logging
        return nil, err
    }

    // Define the template for the query string
    queryTemplate := `
.spec.containers[] |= (
    .volumeMounts += [
        {{$c := counter}}{{- range $name, $storage := . }}{{if call $c}}, {{end}}
        {"mountPath": "/data/{{ $name }}", "name": "{{ $name }}"}
        {{- end }}
    ] |
    .env += [
        {{$c := counter}}{{- range $name, $storage := . }}{{if call $c}}, {{end}}
        {"name": "{{ replaceHyphen $name }}_STORAGE", "value": "{{ $storage }}"}
        {{- end }}
    ]
) |
.spec.volumes += [
    {{$c := counter}}{{- range $name, $storage := . }}{{if call $c}}, {{end}}
    {"name": "{{ $name }}", "persistentVolumeClaim": {"claimName": "{{ $name }}"}}
    {{- end }}
]
`

    // Parse the query template
    t := template.Must(template.New("example").Funcs(template.FuncMap{"counter": counter, "replaceHyphen": replaceHyphen}).Parse(queryTemplate))

    // Create a buffer to hold the generated query string
    var queryBuffer bytes.Buffer

    // Generate the query string from the template
    err = t.Execute(&queryBuffer, pvcMap)
    if err != nil {
        e.LogError(err) // Retain e.LogError for logging
        return nil, err
    }

    queryStr := queryBuffer.String()

    // Parse the generated query
    query, err := gojq.Parse(queryStr)
    if err != nil {
        e.LogError(err) // Retain e.LogError for logging
        return nil, err
    }

    // Execute the query
    iter := query.Run(data)

    // Retrieve the updated result
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
            e.LogError(err) // Retain e.LogError for logging
            return nil, err
        }
        updatedData = v
    }

    // Marshal the updated JSON back to a string
    updatedJSON, err := json.MarshalIndent(updatedData, "", "  ")
    if err != nil {
        e.LogError(err) // Retain e.LogError for logging
        return nil, err
    }

    return updatedJSON, err
}
