package main

import (
	"fmt"
	"github.com/Karuch/OpenshiftStorageAgent/internal/createPOD"
	// "github.com/Karuch/OpenshiftStorageAgent/internal/queryPVC"
)

func main() {
	// queryPVC.GetPVCs()
	fmt.Println("test")
	createPOD.GetPodManifest()
}
