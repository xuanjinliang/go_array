package go_array

import (
	"errors"
	"reflect"
)

type manager struct {
	Data *interface{}
	Type *string
}

func getElemValue(data *interface{}) reflect.Value {
	v := reflect.ValueOf(*data)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

/*
 * Initialize
 */
func Array(array interface{}) (*manager, error) {
	v := getElemValue(&array)
	// log.Printf("%v, %v", v, reflect.ValueOf(array))

	kind := v.Kind()

	if kind == reflect.Array || kind == reflect.Slice {
		k := kind.String()
		return &manager{
				Data: &array,
				Type: &k,
			},
			nil
	}
	return nil, errors.New("the parameter is not array or slice")
}

/*
 * return array or slice length
 */
func (m *manager) Len() int {
	data := m.Data
	v := getElemValue(data)
	return v.Len()
}
