package main

import (
	"fmt"

	"github.com/Karuch/OpenshiftStorageAgent/internal/createPOD"
	e "github.com/Karuch/OpenshiftStorageAgent/internal/logs"
	"github.com/Karuch/OpenshiftStorageAgent/internal/queryPVC"
)

func main() {
	// queryPVC.GetPVCs()
	fmt.Println("test")
	pvcMap, err := queryPVC.GetPVCs()
	if err != nil {
		e.LogError(err)
	}

	respone, err := createPOD.Request(pvcMap)
	if err != nil {
		e.LogError(err)
	}
	
	fmt.Println(string(respone))
}
