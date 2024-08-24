package queryPVC

import (
	"fmt"
	"strings"
	"strconv"
)

func ConvertStorageToBytes(storage string) (int64, error) {
	// Remove whitespace and convert to lowercase
	storage = strings.TrimSpace(strings.ToLower(storage))

	// Define the multipliers for each unit
	unitMultipliers := map[string]int64{
		"b":  1,
		"kb": 1 << 10,
		"mb": 1 << 20,
		"gb": 1 << 30,
		"tb": 1 << 40,
		"pb": 1 << 50,
		"eb": 1 << 60,
		"gi": 1 << 30,
		"mi": 1 << 20,
		"ki": 1 << 10,
	}

	// Extract the numeric value and unit
	for unit := range unitMultipliers {
		if strings.HasSuffix(storage, unit) {
			valueStr := strings.TrimSuffix(storage, unit)
			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				return 0, err // Return zero and the error if parsing fails
			}
			return int64(value * float64(unitMultipliers[unit])), nil
		}
	}

	return 0, fmt.Errorf("unknown storage unit in %s", storage)
}