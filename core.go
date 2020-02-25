package go_array

import (
	"errors"
	"log"
	"reflect"
)

type manager struct {
	Data      reflect.Value
	sliceType reflect.Type
	elemType  reflect.Type
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
	kind := v.Kind()

	// log.Printf("%v, %v, %v", reflect.TypeOf(array), reflect.ValueOf(array), kind)

	if kind != reflect.Slice {
		return nil, errors.New("the parameter is not slice")
	}

	// log.Printf("%v, %v", reflect.TypeOf(array).Elem(), reflect.ValueOf(array).Type())

	return &manager{
		Data:      v,
		sliceType: reflect.ValueOf(array).Type(),
		elemType:  reflect.TypeOf(array).Elem(),
	}, nil

}

func (m *manager) GetData() interface{} {
	return m.Data.Interface()
}

func (m *manager) Len() int {
	return m.Data.Len()
}

/*
 * Only support return value and index
 */
func (m *manager) ForEach(f func(interface{}, int)) {
	data := m.Data
	len := m.Len()

	for i := 0; i < len; i++ {
		o := data.Index(i)
		f(o.Interface(), i)
	}
}

func (m *manager) Concat(arr interface{}) {
	// data := m.Data
	// v := getElemValue(data)
	/*v := getElemValue(data)
	vv := getElemValue(&arr)
	newArray := reflect.Append(reflect.ValueOf(*data), vv)*/
	// log.Printf("%v, %v", reflect.TypeOf(*data), reflect.TypeOf(arr))
	// log.Printf("%v, %v", reflect.ValueOf(*data), reflect.ValueOf(arr))
	newArray := reflect.Append(m.Data, reflect.ValueOf(arr))
	log.Printf("%v", newArray)
}
