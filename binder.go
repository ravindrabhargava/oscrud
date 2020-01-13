package oscrud

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var (
	queryTag  = "query"
	bodyTag   = "body"
	paramTag  = "param"
	headerTag = "header"
)

func (c Context) bindAll(assign interface{}) error {
	t := reflect.TypeOf(assign)
	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
		return errors.New("binder interface must be addressable struct")
	}

	setter := reflect.ValueOf(assign).Elem()
	npt := t.Elem()
	for i := 0; i < npt.NumField(); i++ {
		field := npt.Field(i)
		json := field.Tag.Get("json")
		if json != "" {
			if value, ok := c.header[json]; ok {
				if err := bindValue(setter.Field(i), value); err != nil {
					return err
				}
				continue
			}

			if value, ok := c.query[json]; ok {
				if err := bindValue(setter.Field(i), value); err != nil {
					return err
				}
				continue
			}

			if value, ok := c.body[json]; ok {
				if err := bindValue(setter.Field(i), value); err != nil {
					return err
				}
				continue
			}

			if value, ok := c.param[json]; ok {
				if err := bindValue(setter.Field(i), value); err != nil {
					return err
				}
				continue
			}
		}
	}
	return nil
}

func (c Context) bind(assign interface{}) error {
	t := reflect.TypeOf(assign)
	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
		return errors.New("binder interface must be addressable struct")
	}

	setter := reflect.ValueOf(assign).Elem()
	npt := t.Elem()
	for i := 0; i < npt.NumField(); i++ {
		field := npt.Field(i)

		htag := field.Tag.Get(headerTag)
		if value, ok := c.header[htag]; ok {
			if err := bindValue(setter.Field(i), value); err != nil {
				return err
			}
			continue
		}

		qtag := field.Tag.Get(queryTag)
		if value, ok := c.query[qtag]; ok {
			if err := bindValue(setter.Field(i), value); err != nil {
				return err
			}
			continue
		}

		btag := field.Tag.Get(bodyTag)
		if value, ok := c.body[btag]; ok {
			if err := bindValue(setter.Field(i), value); err != nil {
				return err
			}
			continue
		}

		ptag := field.Tag.Get(paramTag)
		if value, ok := c.param[ptag]; ok {
			if err := bindValue(setter.Field(i), value); err != nil {
				return err
			}
			continue
		}
	}
	return nil
}

func bindValue(field reflect.Value, value interface{}) error {
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
			return fmt.Errorf("Trying to bindValue() on unexported field")
		}
		field.SetString(result)
	case reflect.Bool:
		str := fmt.Sprintf("%v", value)
		result, err := strconv.ParseBool(str)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to bool", value)
		}
		if !field.CanSet() {
			return fmt.Errorf("Trying to bindValue() on unexported field")
		}
		field.SetBool(result)
		break
	default:
		qt := reflect.TypeOf(value)
		if !field.Type().AssignableTo(qt) {
			return fmt.Errorf("Trying to convert %v to %v", value, field.Addr().Type())
		}
		if !field.CanSet() {
			return fmt.Errorf("Trying to bindValue() on unexported field")
		}
		field.Set(reflect.ValueOf(value))
		break

	}

	return nil
}
