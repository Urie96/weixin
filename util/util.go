package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func PrintStruct(in interface{}) {
	b, err := json.Marshal(in)
	if err != nil {
		fmt.Printf("%+v", in)
		fmt.Println()
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		fmt.Printf("%+v", in)
		fmt.Println()
		return
	}
	fmt.Println(out.String())
}

func DecodeURLParamsToStruct(req *http.Request, ptr interface{}) interface{} {
	if err := req.ParseForm(); err != nil {
		panic(err)
	}
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		fields[strings.ToLower(fieldInfo.Name)] = v.Field(i)
	}
	for name, values := range req.Form {
		f := fields[strings.ToLower(name)]
		if _, exist := fields[strings.ToLower(name)]; !exist || !f.IsValid() {
			continue
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					panic(fmt.Errorf("%s: %v", name, err))
				}
				f.Set(reflect.Append(f, elem))
			} else if err := populate(f, value); err != nil {
				panic(fmt.Errorf("%s: %v", name, err))
			}
		}
	}
	return ptr
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int, reflect.Int32:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
