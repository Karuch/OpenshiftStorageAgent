package global

import (
	"fmt"
	"os"

	e "github.com/Karuch/OpenshiftStorageAgent/internal/logs"
)

var (
	TokenFilePath = EnvInit("TOKEN_FILE_PATH", "/var/run/secrets/kubernetes.io/serviceaccount")
	APIServer     = EnvInit("API_SERVER", "")
	Namespace     = EnvInit("NAMESPACE", "")
)

// getEnv retrieves environment variables or returns a default value if not set.
func EnvInit(key, defaultValue string) string {

	// Set environment variables directly  /////////////////
	os.Setenv("TOKEN_FILE_PATH", "/go/kubernetes/token.txt")
	os.Setenv("API_SERVER", "https://192.168.49.2:8443")
	os.Setenv("NAMESPACE", "default")
	// remove before prod ^^^^^^^^^^^^^^^ //////////////////

	value := os.Getenv(key)
	if value == "" {
		if key != "TOKEN_FILE_PATH" {
			e.LogError(fmt.Errorf("environment variable '$%s' is empty, you must declare it.", key))
		}
		return defaultValue
	}
	return value
}
