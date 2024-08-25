package main

import (
	"fmt"

	"github.com/Karuch/OpenshiftStorageAgent/internal/createPOD"
	e "github.com/Karuch/OpenshiftStorageAgent/internal/logs"
	"github.com/Karuch/OpenshiftStorageAgent/internal/queryPVC"
)

func main() {
	fmt.Println("test")

	pvcMap, err := queryPVC.GetPVCs()
	if err != nil {
		e.LogError(err)
	}

	_, err = createPOD.Request(pvcMap)
	if err != nil {
		e.LogError(err)
	}

}
