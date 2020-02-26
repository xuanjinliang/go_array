package go_array

import (
	"errors"
	"reflect"
)

type manager struct {
	Data      reflect.Value
	SliceType reflect.Type
	ElemType  reflect.Type
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
		SliceType: reflect.ValueOf(array).Type(),
		ElemType:  reflect.TypeOf(array).Elem(),
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

/*
 * slice concat, support array, slice and Type value
 */
func (m *manager) Concat(args ...interface{}) interface{} {
	newMData := m.Data
	for _, param := range args {
		v := getElemValue(&param)
		kind := v.Kind()

		if kind == reflect.Array && reflect.TypeOf(param).Elem() == m.ElemType {
			len := v.Len()
			for i := 0; i < len; i++ {
				newMData = reflect.Append(newMData, v.Index(i))
			}
			continue
		}

		if kind == reflect.Slice && reflect.TypeOf(param).Elem() == m.ElemType {
			newMData = reflect.AppendSlice(newMData, v)
			continue
		}

		if reflect.TypeOf(param) == m.ElemType {
			newMData = reflect.Append(newMData, v)
			continue
		}
	}

	m.Data = newMData
	return m.GetData()
}
