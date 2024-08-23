package e

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

// Custom logger initialization
func init() {
	// Disable default log flags for custom formatting
	log.SetFlags(0)
}

// Custom error logging function
func LogError(err error) {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			// Log the error and terminate the program
			log.Fatalf("[%s] %s:%d Error: %v\n", time.Now().Format(time.RFC3339), file, line, err)
		} else {
			// Log the error and terminate the program
			log.Fatalf("Error: %v", err)
		}
	}
}

// Example function that returns an error
func ExampleFunction() error {
	return fmt.Errorf("something went wrong")
}
