package utils

import (
	"reflect"
)

type ResponseMessage struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type RecordsResponse struct {
	Total           int64       `json:"total"`
	RecordsFiltered int         `json:"records_filtered"`
	Data            interface{} `json:"data"`
}

func UpdateEntity(fetchedData interface{}, requestedData interface{}) bool {
	valCategory := reflect.ValueOf(fetchedData).Elem()
	valRequest := reflect.ValueOf(requestedData)
	// Get the number of fields in the requested data struct
	numRequestFields := valRequest.NumField()
	updated := false // Initialize a boolean to track if any updates occurred

	for i := 0; i < numRequestFields; i++ {
		requestValue := valRequest.Field(i)
		requestType := valRequest.Type().Field(i)

		// Check if the request value is valid and not a zero value
		if requestValue.IsValid() && !isZero(requestValue) {
			// Find the corresponding field in the fetched data by name
			currentCategoryValue := valCategory.FieldByName(requestType.Name)
			// If the field exists in the fetched data struct
			if currentCategoryValue.IsValid() && currentCategoryValue.CanSet() {
				// If types match, compare the values
				if currentCategoryValue.Type() == requestValue.Type() {
					// If the current value is the same as the requested value, skip updating
					if currentCategoryValue.Interface() == requestValue.Interface() {
						continue // No update needed
					}
					// Otherwise, set the new value directly
					currentCategoryValue.Set(requestValue)
					updated = true // Set to true if an update occurs
				}
			}
		}
	}
	return updated // Return true if any updates were made, otherwise false
}

func isZero(v reflect.Value) bool {
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}
