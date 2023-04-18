package utils

import (
	"nonoDemo/pkg/framework"
	"reflect"
)

func LoadConfig(config framework.Configuration, target interface{}) interface{} {
	t := reflect.TypeOf(config).Elem()
	v := reflect.ValueOf(config).Elem()
	return loadConfig(t, v, target)
}

func loadConfig(t reflect.Type, v reflect.Value, target interface{}) interface{} {
	for i := 0; i < t.NumField(); i++ {
		tName := t.Field(i).Type.String()
		rName := reflect.TypeOf(target).String()
		if tName == rName {
			return v.Field(i).Interface()
		}
		if t.Field(i).Type.Kind() == reflect.Struct {
			result := loadConfig(t.Field(i).Type, v.Field(i), target)
			if result != nil {
				return result
			}
			continue
		}
	}
	return nil
}
