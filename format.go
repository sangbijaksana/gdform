package gdform

import (
	"fmt"
	"time"
)

// Format function
func Format(dateString string) (time.Time, error) {

	for _, layout := range layoutList {
		time, err := time.Parse(layout, dateString)
		if err != nil {
			continue
		}

		return time, nil
	}

	return time.Time{}, fmt.Errorf("no layout matching found for %s date", dateString)
}
