package dto

import (
	"fmt"
	"reflect"
)

type CarRequestBody struct {
	CarTypeID       uint   `json:"carTypeID"`
	CarTypeName     string `json:"carTypeName"`
	CarMarkID       uint   `json:"carMarkID"`
	CarMarkName     string `json:"carMarkName"`
	CarMarkNameRus  string `json:"carMarkNameRus"`
	CarModelID      uint   `json:"carModelID"`
	CarModelName    string `json:"carModelName"`
	CarModelNameRus string `json:"carModelNameRus"`
}

type CarType struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToDTO(model interface{}) interface{} {
	// fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(model)
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		fmt.Println(fieldInfo)
	}
	return v
}
