package utils

import (
	"fmt"
	"reflect"
)

type Field struct {
	Name string
	Type string
	// Add other fields if necessary
}

// GetListOfFields generates a string representation of a list of fields with "Name" and "Type" properties.
func GetListOfFields(fields interface{}) string {
	val := reflect.ValueOf(fields)
	if val.Kind() != reflect.Slice {
		return ""
	}

	var listOfFields string
	for i := 0; i < val.Len(); i++ {
		fieldVal := val.Index(i)
		if fieldVal.Kind() != reflect.Struct {
			continue
		}
		nameField := fieldVal.FieldByName("Name")
		typeField := fieldVal.FieldByName("Type")
		if !nameField.IsValid() || !typeField.IsValid() || nameField.Kind() != reflect.String || typeField.Kind() != reflect.String {
			continue
		}
		fieldString := fmt.Sprintf("%s:%s", nameField.String(), typeField.String())
		if listOfFields == "" {
			listOfFields = fieldString
		} else {
			listOfFields = fmt.Sprintf("%s %s", listOfFields, fieldString)
		}
	}
	return listOfFields
}
