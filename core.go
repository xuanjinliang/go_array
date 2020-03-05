package go_array

import (
	"bytes"
	"encoding/gob"
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

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

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

	startArr := data.Slice(0, target)

	start := 0
	end := dataLen
	if len(args) > 0 {
		start = m.minus(args[0])
	}

	if len(args) > 1 {
		end = m.minus(args[1])
	}

	if start >= dataLen || end <= start {
		return m.GetData()
	}

	if end > dataLen {
		end = dataLen
	}

	endArr := data.Slice(start, end)

	if end < dataLen {
		startArr = reflect.AppendSlice(startArr, endArr)
		startLen := startArr.Len()

		if startLen > dataLen {
			startLen = dataLen
		}
		endArr = data.Slice(startLen, dataLen)
	}

	s := reflect.MakeSlice(m.SliceType, dataLen, dataLen)
	reflect.Copy(s, reflect.AppendSlice(startArr, endArr))
	// log.Printf("%v, %v, %v", startArr, endArr, s)
	m.Data = s
	return m.GetData()
}

/*
 * slice every
 */
func (m *manager) Every(f func(interface{}, int) bool) bool {
	data := m.Data
	len := m.Len()

	for i := 0; i < len; i++ {
		o := data.Index(i)
		bool := f(o.Interface(), i)
		if bool == false {
			return false
		}
	}

	return true
}

/*
 * slice fill
 */
func (m *manager) Fill(target interface{}, args ...int) interface{} {
	data := m.Data
	dataLen := m.Len()

	if reflect.TypeOf(target) != m.ElemType {
		return m.GetData()
	}

	start := 0
	end := dataLen

	if len(args) > 0 {
		start = m.minus(args[0])
	}

	if len(args) > 1 {
		end = m.minus(args[1])
	}

	if start > dataLen || end <= start {
		return m.GetData()
	}

	startArr := data.Slice(0, start)
	endArr := data.Slice(start, end)

	if end > dataLen {
		end = dataLen
	}

	endArr = data.Slice(start, end)
	l := endArr.Len()

	for i := 0; i < l; i++ {
		endArr.Index(i).Set(reflect.ValueOf(target))
	}

	if end < dataLen {
		startArr = reflect.AppendSlice(startArr, endArr)
		startLen := startArr.Len()

		if startLen > dataLen {
			startLen = dataLen
		}
		endArr = data.Slice(startLen, dataLen)
	}

	m.Data = reflect.AppendSlice(startArr, endArr)
	return m.GetData()
}

/*
 * slice Filter
 */
func (m *manager) Filter(f func(interface{}, int) bool) interface{} {
	data := m.Data
	len := m.Len()

	s := reflect.MakeSlice(m.SliceType, 0, len)
	for i := 0; i < len; i++ {
		o := data.Index(i)
		if bool := f(o.Interface(), i); bool {
			s = reflect.Append(s, o)
		}
	}

	return s.Interface()
}

/*
 * slice Fine
 */
func (m *manager) Fine(f func(interface{}, int) bool) interface{} {
	data := m.Data
	len := m.Len()

	var v interface{}
	for i := 0; i < len; i++ {
		o := data.Index(i)
		if bool := f(o.Interface(), i); bool {
			v = o.Interface()
			break
		}
	}

	return v
}

/*
 * slice FineIndex
 */
func (m *manager) FineIndex(f func(interface{}, int) bool) int {
	data := m.Data
	len := m.Len()

	index := -1
	for i := 0; i < len; i++ {
		o := data.Index(i)
		if bool := f(o.Interface(), i); bool {
			index = i
			break
		}
	}

	return index
}

/*
 * slice includes
 */
func (m *manager) Includes(v interface{}) (bool, error) {
	data := m.Data
	len := m.Len()

	b, e := getBytes(v)
	if e != nil {
		return false, e
	}

	for i := 0; i < len; i++ {
		o := data.Index(i).Interface()
		a, err := getBytes(o)
		if err != nil {
			return false, err
		}

		if bool := bytes.Equal(a, b); bool {
			return bool, nil
		}
	}
	return false, nil
}
