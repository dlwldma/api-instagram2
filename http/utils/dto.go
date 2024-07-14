package utils

import (
	"errors"
	"reflect"
)

func AreAllFieldsSet(s interface{}) (bool, error) {
	// Get the type of the struct
	structType := reflect.TypeOf(s)

	// If s is a pointer, dereference it to get the struct type
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	// Get the value of the struct
	structValue := reflect.ValueOf(s)

	// Iterate through the fields
	for i := 0; i < structType.NumField(); i++ {
		field := structValue.Field(i)

		// Check if the field is zero-initialized
		if reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			return false, errors.New("form-data fields required")
		}
	}

	return true, nil
}
