package oscrud

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var (
	queryTag      = "query"
	bodyTag       = "body"
	paramTag      = "param"
	headerTag     = "header"
	jsonTag       = "json"
	queryModelTag = "qm"
)

// func (c Context) bindAll(assign interface{}) error {
// 	t := reflect.TypeOf(assign)
// 	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
// 		return errors.New("binder interface must be addressable struct")
// 	}

// 	setter := reflect.ValueOf(assign).Elem()
// 	npt := t.Elem()
// 	for i := 0; i < npt.NumField(); i++ {
// 		field := npt.Field(i)
// 		key := ""
// 		json := field.Tag.Get(jsonTag)
// 		if json != "" {
// 			key = strings.Split(json, ",")[0]
// 		}

// 		qm := field.Tag.Get(queryModelTag)
// 		if qm != "" {
// 			key = qm
// 		}

// 		if key != "" {
// 			if value, ok := c.header[key]; ok {
// 				if err := Bind(setter.Field(i), value, c.oscrud.binder); err != nil {
// 					return err
// 				}
// 				continue
// 			}

// 			if value, ok := c.query[key]; ok {
// 				if err := Bind(setter.Field(i), value, c.oscrud.binder); err != nil {
// 					return err
// 				}
// 				continue
// 			}

// 			if value, ok := c.body[key]; ok {
// 				if err := Bind(setter.Field(i), value, c.oscrud.binder); err != nil {
// 					return err
// 				}
// 				continue
// 			}

// 			if value, ok := c.param[key]; ok {
// 				if err := Bind(setter.Field(i), value, c.oscrud.binder); err != nil {
// 					return err
// 				}
// 				continue
// 			}
// 		}
// 	}
// 	return nil
// }

// func (c Context) bind(assign interface{}) error {
// 	t := reflect.TypeOf(assign)
// 	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
// 		return errors.New("binder interface must be addressable struct")
// 	}

// 	setter := reflect.ValueOf(assign).Elem()
// 	npt := t.Elem()
// 	for i := 0; i < npt.NumField(); i++ {
// 		field := npt.Field(i)

// 		htag := field.Tag.Get(headerTag)
// 		if value, ok := c.header[htag]; ok {
// 			if err := Bind(setter.Field(i), value, c.oscrud.binder); err != nil {
// 				return err
// 			}
// 			continue
// 		}

// 		qtag := field.Tag.Get(queryTag)
// 		if value, ok := c.query[qtag]; ok {
// 			if err := Bind(setter.Field(i), value, c.oscrud.binder); err != nil {
// 				return err
// 			}
// 			continue
// 		}

// 		btag := field.Tag.Get(bodyTag)
// 		if value, ok := c.body[btag]; ok {
// 			if err := Bind(setter.Field(i), value, c.oscrud.binder); err != nil {
// 				return err
// 			}
// 			continue
// 		}

// 		ptag := field.Tag.Get(paramTag)
// 		if value, ok := c.param[ptag]; ok {
// 			if err := Bind(setter.Field(i), value, c.oscrud.binder); err != nil {
// 				return err
// 			}
// 			continue
// 		}
// 	}
// 	return nil
// }

// Binder :
type Binder struct {
	custom map[string]Bind
}

// NewBinder :
func NewBinder() Binder {
	return Binder{
		custom: make(map[string]Bind),
	}
}

// Bind :
type Bind func(string) (interface{}, error)

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
			deserialized, err := binder(fmt.Sprintf("%v", value))
			if err != nil {
				return fmt.Errorf("Trying to deserialize %v to %v, %v", value, field.Type().Name(), err)
			}
			field.Set(reflect.ValueOf(deserialized))
		} else {
			// Check if value is list & is struct -> loop -> deserialize
			qt := reflect.TypeOf(value)
			if !field.Type().AssignableTo(qt) {
				return fmt.Errorf("Trying to convert %v to %v", value, field.Addr().Type())
			}
			field.Set(reflect.ValueOf(value))
		}
		break
	case reflect.Array:
		if binder, ok := b.custom["array$"+field.Type().Elem().Name()]; ok {
			deserialized, err := binder(fmt.Sprintf("%v", value))
			if err != nil {
				return fmt.Errorf("Trying to deserialize %v to %v, %v", value, field.Type().Name(), err)
			}
			field.Set(reflect.ValueOf(deserialized))
		} else {
			// Check if value is list & is struct -> loop -> deserialize
			qt := reflect.TypeOf(value)
			if !field.Type().AssignableTo(qt) {
				return fmt.Errorf("Trying to convert %v to %v", value, field.Addr().Type())
			}
			field.Set(reflect.ValueOf(value))
		}
		break
	case reflect.Struct:
		if binder, ok := b.custom[field.Type().Name()]; ok {
			deserialized, err := binder(fmt.Sprintf("%v", value))
			if err != nil {
				return fmt.Errorf("Trying to deserialize %v to %v, %v", value, field.Type().Name(), err)
			}
			field.Set(reflect.ValueOf(deserialized))
		} else {
			// Check is struct -> loop field -> deserialzie based on field
			qt := reflect.TypeOf(value)
			if !field.Type().AssignableTo(qt) {
				return fmt.Errorf("Trying to convert %v to %v", value, field.Addr().Type())
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
