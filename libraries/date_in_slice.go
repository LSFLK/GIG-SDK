package libraries

import "time"

/*
DateInSlice - check if a given date exists in a given date slice
*/
func DateInSlice(slice []time.Time, element time.Time) bool {
	for _, existingElement := range slice {
		if existingElement == element {
			return true
		}
	}
	return false
}
