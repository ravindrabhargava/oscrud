package oscrud

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// Binder :
type Binder struct {
	custom map[string]Bind
}

// NewBinder :
func NewBinder() *Binder {
	return &Binder{
		custom: make(map[string]Bind),
	}
}

// Bind :
type Bind func(interface{}) (interface{}, error)

// Register :
func (b *Binder) Register(rtype interface{}, bindFn Bind) *Binder {
	typ := reflect.TypeOf(rtype)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Array {
		b.custom["array$"+typ.Elem().Name()] = bindFn
	} else if typ.Kind() == reflect.Slice {
		b.custom["slice$"+typ.Elem().Name()] = bindFn
	} else {
		b.custom[typ.Name()] = bindFn
	}
	return b
}

// Bind :
func (b Binder) Bind(assign interface{}, value interface{}) error {
	typ := reflect.TypeOf(assign)
	if typ.Kind() != reflect.Ptr {
		return errors.New("binder interface must be addressable struct")
	}

	field := reflect.ValueOf(assign).Elem()
	if !field.CanSet() {
		return fmt.Errorf("Trying to bind() on unexported field")
	}

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
		field.SetString(result)
	case reflect.Bool:
		str := fmt.Sprintf("%v", value)
		result, err := strconv.ParseBool(str)
		if err != nil {
			return fmt.Errorf("Trying to convert %v to bool", value)
		}
		field.SetBool(result)
		break
	case reflect.Slice:
		if binder, ok := b.custom["slice$"+field.Type().Elem().Name()]; ok {
			deserialized, err := binder(value)
			if err != nil {
				return fmt.Errorf("Trying to deserialize %v to %v, %v", value, field.Type().Name(), err)
			}
			field.Set(reflect.ValueOf(deserialized))
		} else {
			qt := reflect.TypeOf(value)
			if !field.Type().AssignableTo(qt) {
				return fmt.Errorf("Trying to convert %v to slice %v", value, field.Type().Elem().Name())
			}
			field.Set(reflect.ValueOf(value))
		}
		break
	case reflect.Array:
		if binder, ok := b.custom["array$"+field.Type().Elem().Name()]; ok {
			deserialized, err := binder(value)
			if err != nil {
				return fmt.Errorf("Trying to deserialize %v to %v, %v", value, field.Type().Name(), err)
			}
			field.Set(reflect.ValueOf(deserialized))
		} else {
			qt := reflect.TypeOf(value)
			if !field.Type().AssignableTo(qt) {
				return fmt.Errorf("Trying to convert %v to array %v", value, field.Type().Elem().Name())
			}
			field.Set(reflect.ValueOf(value))
		}
		break
	case reflect.Struct:
		if binder, ok := b.custom[field.Type().Name()]; ok {
			deserialized, err := binder(value)
			if err != nil {
				return fmt.Errorf("Trying to deserialize %v to %v, %v", value, field.Type().Name(), err)
			}
			field.Set(reflect.ValueOf(deserialized))
		} else {
			qt := reflect.TypeOf(value)
			if !field.Type().AssignableTo(qt) {
				return fmt.Errorf("Trying to convert %v to struct %v", value, field.Type().Name())
			}
			field.Set(reflect.ValueOf(value))
		}
		break
	default:
		qt := reflect.TypeOf(value)
		if !field.Type().AssignableTo(qt) {
			return fmt.Errorf("Trying to convert %v to %v", value, field.Addr().Type())
		}
		field.Set(reflect.ValueOf(value))
		break
	}

	return nil
}
