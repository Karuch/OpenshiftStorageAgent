package query

import (
	"fmt"
	"strings"
	"os"
	"github.com/Karuch/OpenshiftStorageAgent/internal/logs"
)

func GetPVCs() (map[string]string, error) {

	var ENABLE_ONLY string = os.Getenv("ENABLE_ONLY")
	var DISABLE_ONLY string = os.Getenv("DISABLE_ONLY")

	if ENABLE_ONLY != "" && DISABLE_ONLY != "" {
		e.LogError(fmt.Errorf("Cannot use ENABLE_ONLY and DISABLE_ONLY at the same time, only one can contain characters. $ENABLE_ONLY: '%s', $DISABLE_ONLY: '%s'", ENABLE_ONLY, DISABLE_ONLY))
	}

	desiredPVCs := map[string]string{}

	allPVCsMap, err := FliterJson()
	if err != nil {
		e.LogError(err)
	}

	if ENABLE_ONLY != "" || DISABLE_ONLY != "" {

		if ENABLE_ONLY != "" {
			fmt.Printf("found ENABLE_ONLY: '%s'\n", ENABLE_ONLY)
			enableOnlylist := strings.Fields(ENABLE_ONLY)

			for _, only := range enableOnlylist {
				for pvc, storage := range allPVCsMap {
					if pvc == only {
						desiredPVCs[pvc] = storage
					}
				}
			}
		}

		if DISABLE_ONLY != "" {
			fmt.Printf("found DISABLE_ONLY: '%s'\n", DISABLE_ONLY)
			disableOnlylist := strings.Fields(DISABLE_ONLY)

			for pvc, storage := range allPVCsMap {
				found := false
				for _, disabledPVC := range disableOnlylist {
					if pvc == disabledPVC {
						found = true
						break
					}
				}
				if !found {
					desiredPVCs[pvc] = storage
				}
			}
		}

	} else {
		fmt.Println("Didn't find any use of ENABLE_ONLY or DISABLE_ONLY")
		desiredPVCs = allPVCsMap
	}

	if len(desiredPVCs) < 1 {
		e.LogError(fmt.Errorf("Didn't found any PVCs to monitor"))
	}

	fmt.Println("found and will monitor those:")
	for pvc, storage := range desiredPVCs {
		fmt.Printf("['%s' %s] ", pvc, storage)
	}

	return desiredPVCs, err

}
