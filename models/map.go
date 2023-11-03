package models

import (
	"fmt"
	"reflect"
	"strings"
)

func MapStructFields(source any, target any) any {
	sourceVal := reflect.ValueOf(source)
	targetVal := reflect.ValueOf(target)

	for i := 0; i < sourceVal.NumField(); i++ {
		sourceFieldName := sourceVal.Type().Field(i).Name
		targetField := targetVal.FieldByNameFunc(func(s string) bool {
			return strings.EqualFold(s, sourceFieldName)
		})
		fmt.Println(sourceFieldName)
		fmt.Println(targetField)
		fmt.Println("=======================")
		fmt.Println()

		if targetField.IsValid() && targetField.CanSet() {
			targetField.Set(sourceVal.Field(i))
		}
	}

	return target
}
