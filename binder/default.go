package binder

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// Tag Definitions
var (
	QueryTag  = "query"
	BodyTag   = "body"
	ParamTag  = "param"
	HeaderTag = "header"
	IDTag     = "id"
)

// BindService :
func BindService(id string, body map[string]interface{}, query map[string]interface{}, header map[string]interface{}, assign interface{}) error {
	t := reflect.TypeOf(assign)
	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
		return errors.New("binder interface must be addressable struct")
	}

	setter := reflect.ValueOf(assign).Elem()
	npt := t.Elem()
	for i := 0; i < npt.NumField(); i++ {
		field := npt.Field(i)

		htag := string(field.Tag.Get(HeaderTag))
		if htag != "" && header[htag] != nil {
			err := BindValue(setter.Field(i), header[htag])
			if err != nil {
				return err
			}
		}

		qtag := string(field.Tag.Get(QueryTag))
		if qtag != "" && query[qtag] != nil {
			err := BindValue(setter.Field(i), query[qtag])
			if err != nil {
				return err
			}
		}

		btag := string(field.Tag.Get(BodyTag))
		if btag != "" && body[btag] != nil {
			err := BindValue(setter.Field(i), body[btag])
			if err != nil {
				return err
			}
		}

		idtag := string(field.Tag.Get(IDTag))
		if idtag != "" && id != "" {
			err := BindValue(setter.Field(i), id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// BindEndpoint :
func BindEndpoint(header map[string]interface{}, param map[string]string, body map[string]interface{}, query map[string]interface{}, assign interface{}) error {
	t := reflect.TypeOf(assign)
	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
		return errors.New("binder interface must be addressable struct")
	}

	setter := reflect.ValueOf(assign).Elem()
	npt := t.Elem()
	for i := 0; i < npt.NumField(); i++ {
		field := npt.Field(i)

		htag := string(field.Tag.Get(HeaderTag))
		if htag != "" && header[htag] != nil {
			err := BindValue(setter.Field(i), header[htag])
			if err != nil {
				return err
			}
		}

		qtag := string(field.Tag.Get(QueryTag))
		if qtag != "" && query[qtag] != nil {
			err := BindValue(setter.Field(i), query[qtag])
			if err != nil {
				return err
			}
		}

		btag := string(field.Tag.Get(BodyTag))
		if btag != "" && body[btag] != nil {
			err := BindValue(setter.Field(i), body[btag])
			if err != nil {
				return err
			}
		}

		ptag := string(field.Tag.Get(ParamTag))
		if ptag != "" && param[ptag] != "" {
			err := BindValue(setter.Field(i), param[ptag])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// BindValue :
func BindValue(field reflect.Value, value interface{}) error {
	switch field.Type().Kind() {

	case reflect.Float32, reflect.Float64:
		str := fmt.Sprintf("%v", value)
		bit := field.Type().Bits()
		result, err := strconv.ParseFloat(str, bit)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to float%d", value, bit)
		}
		field.SetFloat(result)
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		str := fmt.Sprintf("%v", value)
		bit := field.Type().Bits()
		result, err := strconv.ParseUint(str, 10, bit)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to uint%d", value, bit)
		}
		field.SetUint(result)
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		str := fmt.Sprintf("%v", value)
		bit := field.Type().Bits()
		result, err := strconv.ParseInt(str, 10, bit)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to int%d", value, bit)
		}
		field.SetInt(result)
		break
	case reflect.String:
		result := fmt.Sprintf("%v", value)
		if !field.CanSet() {
			return fmt.Errorf("Trying to BindValue() on unexported field")
		}
		field.SetString(result)
	case reflect.Bool:
		str := fmt.Sprintf("%v", value)
		result, err := strconv.ParseBool(str)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to bool", value)
		}
		if !field.CanSet() {
			return fmt.Errorf("Trying to BindValue() on unexported field")
		}
		field.SetBool(result)
		break
	default:
		qt := reflect.TypeOf(value)
		if !field.Type().AssignableTo(qt) {
			return fmt.Errorf("Trying to convert %v to %v", value, field.Addr().Type())
		}
		if !field.CanSet() {
			return fmt.Errorf("Trying to BindValue() on unexported field")
		}
		field.Set(reflect.ValueOf(value))
		break

	}

	return nil
}
