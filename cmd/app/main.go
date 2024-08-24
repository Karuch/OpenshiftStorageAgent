package main

import (
	"fmt"
	"github.com/Karuch/OpenshiftStorageAgent/internal/query"
)

func main() {
	query.GetPVCs()
	fmt.Println("test")
}
