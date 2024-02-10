package reflections

import (
	"reflect"
)

func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// Get the type and value of the struct
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)

	// Iterate through the struct fields
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		fieldValue := dataValue.Field(i).Interface()

		// Check if the field is optional and unset
		if field.Tag.Get("optional") == "true" && reflect.ValueOf(fieldValue).IsNil() {
			continue // Skip optional fields if unset
		}

		// Add the field to the map
		result[field.Name] = fieldValue
	}

	return result
}

func IsString(val interface{}) bool {
	return reflect.ValueOf(val).Type().Kind() == reflect.String
}

func IsStringPtr(val interface{}) bool {
	return reflect.ValueOf(val).Type().Kind() == reflect.Pointer
}
