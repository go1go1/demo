package goini

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Marshal(data interface{}) (result []byte, err error) {
	typeInfo := reflect.TypeOf(data)
	if typeInfo.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}

	valueInfo := reflect.ValueOf(data)
	var conf []string
	for i := 0; i < typeInfo.NumField(); i++ {
		field := typeInfo.Field(i)
		value := valueInfo.Field(i)

		fieldType := field.Type
		if fieldType.Kind() == reflect.Struct {
			tagVal := field.Name
			if len(tagVal) == 0 {
				tagVal = field.Name
			}
			section := fmt.Sprintf("[%s]\n", tagVal)
			conf = append(conf, section)
		}
		for j := 0; j < fieldType.NumField(); j++ {
			keyField := fieldType.Field(j)
			fieldTagVal := keyField.Tag.Get("ini")
			// 没有设置tag
			if len(fieldTagVal) == 0 {
				fieldTagVal = keyField.Name
			}
			valField := value.Field(j)
			item := fmt.Sprintf("%s=%v\n", fieldTagVal, valField.Interface())
			conf = append(conf, item)
		}
	}
	for _, val := range conf {
		byteVal := []byte(val)
		result = append(result, byteVal...)
	}
	return
}
func UnMarshal(data []byte, result interface{}) (err error) {
	lineArr := strings.Split(string(data), "\n")
	//for _, v := range lineArr {
	//	fmt.Printf("---%s\n", v)
	//}

	resultType := reflect.TypeOf(result)
	if resultType.Kind() != reflect.Ptr {
		err = errors.New("please pass pointer")
		return
	}

	elemType := resultType.Elem()
	if elemType.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}

	var lastSectionName string
	for index, line := range lineArr {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		//忽略注释
		if line[0] == ';' || line[0] == '#' {
			continue
		}
		if line[0] == '[' {
			lastSectionName, err = parseSection(line, elemType)
			if err != nil {
				err = fmt.Errorf("%v, lineNo:%d", err, index+1)
				return
			}
			continue
		}

		err = parseItem(lastSectionName, line, result)
		if err != nil {
			err = fmt.Errorf("%v lineNo %d", err, index+1)
			return
		}
	}

	return
}

// parseSection
func parseSection(line string, elemType reflect.Type) (fieldName string, err error) {
	//核验内容段落，以[xxxx]区分的
	if line[0] == '[' {
		if len(line) <= 2 {
			err = fmt.Errorf("syntax error, invalid section %s", line)
			return
		}

		if line[len(line)-1] != ']' {
			err = fmt.Errorf("syntax error, invalid section %s", line)
			return
		}

		sectionName := line[1 : len(line)-1]
		sectionName = strings.TrimSpace(sectionName)
		if len(sectionName) == 0 {
			err = fmt.Errorf("syntax error, invalid section %s", line)
			return
		}

		for i := 0; i < elemType.NumField(); i++ {
			field := elemType.Field(i)
			tagValue := field.Tag.Get("ini")
			if tagValue == sectionName {
				fieldName = field.Name
				break
			}
		}
	}
	return
}

// parseItem
func parseItem(sectionName, line string, result interface{}) (err error) {
	index := strings.Index(line, "=")
	if index == -1 {
		err = fmt.Errorf("syntax error, line %s", line)
		return
	}

	key := strings.TrimSpace(line[0:index])
	val := strings.TrimSpace(line[index+1:])
	if len(key) == 0 {
		err = fmt.Errorf("syntax error, line %s", line)
		return
	}

	resultValue := reflect.ValueOf(result)
	sectionValue := resultValue.Elem().FieldByName(sectionName)

	sectionType := sectionValue.Type()
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field %s must be struct", sectionName)
		return
	}

	keyFieldName := ""
	for i := 0; i < sectionType.NumField(); i++ {
		field := sectionType.Field(i)
		tagVal := field.Tag.Get("ini")
		if tagVal == key {
			keyFieldName = field.Name
			break
		}
	}
	if len(keyFieldName) == 0 {
		return
	}

	fieldValue := sectionValue.FieldByName(keyFieldName)
	if fieldValue == reflect.ValueOf(nil) {
		return
	}

	fieldKind := fieldValue.Type().Kind()
	switch fieldKind {
	case reflect.String:
		fieldValue.SetString(val)
	case reflect.Int, reflect.Int32, reflect.Int64:
		intVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		fieldValue.SetInt(intVal)
	case reflect.Bool:
		boolVal, err := strconv.ParseBool(val)
		if err != nil {
			return err
		}
		fieldValue.SetBool(boolVal)
	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		fieldValue.SetFloat(floatVal)
	default:
		err = fmt.Errorf("unsupport type:%v", fieldKind)
	}

	return
}
