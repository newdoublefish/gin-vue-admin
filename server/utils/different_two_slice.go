package utils

import (
	"errors"
	"reflect"
)

func difference(slice1 interface{}, slice2 interface{}, f func(from interface{}, to interface{}) bool) ([]interface{}, error) {
	s1 := reflect.ValueOf(slice1)
	if s1.Kind() != reflect.Slice {
		return nil, errors.New("slice1 InterfaceSlice() given a non-slice type")
	}
	s2 := reflect.ValueOf(slice2)
	if s2.Kind() != reflect.Slice {
		return nil, errors.New("slice2 InterfaceSlice() given a non-slice type")
	}
	//create slice
	var diff []interface{}
	for i := 0; i < 1; i++ {
		for x := 0; x < s1.Len(); x++ {
			found := false
			for y := 0; y < s2.Len(); y++ {
				if f(s1.Index(x).Interface(), s2.Index(y).Interface()) {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1.Index(x).Interface())
			}
		}
		if i == 0 {
			s1, s2 = s2, s1
		}
	}
	return diff, nil
}
