package oscrud

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// Tag Definitions
var (
	QueryTag = "query"
)

// GetStruct :
func GetStruct(query map[string]interface{}, assign interface{}) error {
	t := reflect.TypeOf(assign)
	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
		return errors.New("Query interface must be addressable struct")
	}

	setter := reflect.ValueOf(assign).Elem()
	npt := t.Elem()
	for i := 0; i < npt.NumField(); i++ {
		field := npt.Field(i)
		tag := string(field.Tag.Get(QueryTag))
		if tag != "" && query[tag] != nil {
			err := ParseValue(query[tag], setter.Field(i))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ParseValue :
func ParseValue(query interface{}, field reflect.Value) error {
	switch field.Type().Kind() {
	case reflect.Float32, reflect.Float64:
		str, ok := query.(string)
		if !ok {
			return fmt.Errorf("Trying to convert %v to string", query)
		}
		bit := field.Type().Bits()
		result, err := strconv.ParseFloat(str, bit)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to float%d", query, bit)
		}
		field.SetFloat(result)
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		str, ok := query.(string)
		if !ok {
			return fmt.Errorf("Trying to convert %v to string", query)
		}
		bit := field.Type().Bits()
		result, err := strconv.ParseUint(str, 10, bit)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to uint%d", query, bit)
		}
		field.SetUint(result)
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		str, ok := query.(string)
		if !ok {
			return fmt.Errorf("Trying to convert %v to string", query)
		}
		bit := field.Type().Bits()
		result, err := strconv.ParseInt(str, 10, bit)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to int%d", query, bit)
		}
		field.SetInt(result)
		break
	case reflect.String:
		result, ok := query.(string)
		if !ok {
			return fmt.Errorf("Trying to convert %v to string", query)
		}
		if !field.CanSet() {
			return fmt.Errorf("Trying to SetValue() on unexported field")
		}
		field.SetString(result)
	case reflect.Bool:
		str, ok := query.(string)
		if !ok {
			return fmt.Errorf("Trying to convert %v to string", query)
		}
		result, err := strconv.ParseBool(str)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to bool", query)
		}
		if !field.CanSet() {
			return fmt.Errorf("Trying to SetValue() on unexported field")
		}
		field.SetBool(result)
		break
	default:
		qt := reflect.TypeOf(query)
		if !field.Type().AssignableTo(qt) {
			return fmt.Errorf("Trying to convert %v to %v", query, field.Addr().Type())
		}
		if !field.CanSet() {
			return fmt.Errorf("Trying to SetValue() on unexported field")
		}
		field.Set(reflect.ValueOf(query))
		break
	}

	return nil
}
