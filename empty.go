package validator

import (
	"reflect"

	"github.com/pkg/errors"
)

var (
	notSupportType = errors.New("not support type")
)

// ValidStringParamsNull fast string match, not use reflect
func ValidParamsNullWithString(ss ...string) bool {
	for _, s := range ss {
		if s == "" {
			return true
		}
	}
	return false
}

// ValidParamsNull use reflect method to check empty vals
func ValidParamsNull(vals ...interface{}) bool {
	for _, val := range vals {
		b := IsEmpty(val)
		if b {
			return true
		}
	}
	return false
}

// IsEmpty check if val is empty
func IsEmpty(val interface{}) bool {
	// get nil case out of the way
	if val == nil {
		return true
	}

	objValue := reflect.ValueOf(val)

	switch objValue.Kind() {
	// collection types are empty when they have no element
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
		// pointers are empty if nil or if the value they point to is empty
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		deref := objValue.Elem().Interface()
		return IsEmpty(deref)
		// for all other types, compare against the zero value
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(val, zero.Interface())
	}
}

// func notEmpty(v reflect.Value, name, param string) error {
// 	switch v.Kind() {
// 	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
// 		if v.Len() == 0 {
// 			return errors.New(name + " must not be empty")
// 		}
// 		return nil
// 	case reflect.Bool:
// 		if !v.Bool() {
// 			return errors.New(name + " must not be false")
// 		}
// 		return nil
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		if v.Int() == 0 {
// 			return errors.New(name + " must not be zero")
// 		}
// 		return nil
// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
// 		if v.Uint() == 0 {
// 			return errors.New(name + " must not be zero")
// 		}
// 		return nil
// 	case reflect.Float32, reflect.Float64:
// 		if v.Float() == 0 {
// 			return errors.New(name + " must not be zero")
// 		}
// 		return nil
// 	case reflect.Interface, reflect.Ptr:
// 		if v.IsNil() {
// 			return errors.New(name + " must not be nil")
// 		}
// 		return nil
// 	default:
// 		return notSupportType
// 	}
// }
