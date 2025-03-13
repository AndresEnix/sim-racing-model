package model

import (
	"reflect"
	"unicode/utf16"
)

func uint16ToString(data []uint16) string {
	cleanData := trimTrailingNulls(data)
	return string(utf16.Decode(cleanData))
}

func trimTrailingNulls(data []uint16) []uint16 {
	for i, v := range data {
		if v == 0 {
			return data[:i]
		}
	}
	return data
}

func getFieldNames(i interface{}) []string {
	v := reflect.ValueOf(i)
	t := v.Type()

	var fieldNames []string
	for i := 0; i < v.NumField(); i++ {
		fieldNames = append(fieldNames, t.Field(i).Name)
	}
	return fieldNames
}