package util

import (
	"errors"
	"reflect"
	"strings"
)

func TrimSpace(s interface{}) error {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return errors.New("Input is not a pointer to a struct")
	}
	val = val.Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fieldValue := field.String()
			trimmedValue := strings.TrimSpace(fieldValue)
			field.SetString(trimmedValue)
		}
	}
	return nil
}

func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()
		result[field.Name] = value
	}

	return result
}
