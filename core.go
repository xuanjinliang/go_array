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

/*
 * slice copyWithin
 */
func (m *manager) minus(num int) int {
	if num < 0 {
		return m.Data.Len() + num
	}
	return num
}
func (m *manager) CopyWithin(target int, args ...int) interface{} {
	data := m.Data
	dataLen := data.Len()

	target = m.minus(target)

	if target >= dataLen {
		return m.GetData()
	}

	start := 0
	startArr := data.Slice(0, target)

	end := dataLen
	endArr := data.Slice(start, dataLen)

	if len(args) > 0 {
		start = m.minus(args[0])
		endArr = data.Slice(start, dataLen)
	}

	if len(args) > 1 {
		end = m.minus(args[1])

		if end <= start {
			return m.GetData()
		}

		if end > dataLen {
			end = dataLen
		}

		curArray := data.Slice(start, end)
		index := end + 1
		if index >= dataLen {
			endArr = curArray
		} else {
			endArr = data.Slice(index, dataLen)
			endArr = reflect.AppendSlice(curArray, endArr)
		}
	}

	s := reflect.MakeSlice(m.SliceType, dataLen, dataLen)
	reflect.Copy(s, reflect.AppendSlice(startArr, endArr))
	// log.Printf("%v, %v, %v", startArr, endArr, s)
	m.Data = s
	return m.GetData()
}
