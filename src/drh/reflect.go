package drh

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Username string  `json:"Username"`
	Socre    float64 `json:"Socre"`
}

func GetFieldName(columnName string, info *User) {
	t := reflect.TypeOf(info)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		fmt.Println("Check type error not Struct")
	}
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		fmt.Println(strings.ToUpper(t.Field(i).Name))
		if strings.ToUpper(t.Field(i).Name) == strings.ToUpper(columnName) {
			v := reflect.ValueOf(info).Elem()
			v.FieldByName(t.Field(i).Name).SetString("test")
		}
	}
	fmt.Println(info)
}
