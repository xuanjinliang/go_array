package go_array

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestArray(t *testing.T) {
	Convey("test", t, func() {
		arrayData := [3]int{1, 2, 3}
		slice := arrayData[:] // array to slice
		array, err := Array(slice)
		So(err, ShouldBeNil)
		// change type []int 断言
		getArray := array.GetData().([]int)
		log.Printf("%v", getArray)
		// slice to array
		var arr [3]int
		copy(arr[:], slice[:])
		log.Printf("%v", arr)
	})
}

var sliceData = []int{1, 2, 3}

func TestSlice(t *testing.T) {
	Convey("test", t, func() {
		array, err := Array(sliceData)
		So(err, ShouldBeNil)
		data := array.GetData().([]int)
		So(reflect.DeepEqual(sliceData, data), ShouldBeTrue)
		// log.Printf("%v", array.GetData())
		/*aa := append(arrayData, []int{1, 2, 3}...)
		log.Printf("%v", aa)*/
	})
}

func TestManager_Len(t *testing.T) {
	Convey("test len", t, func() {
		array, err := Array(sliceData)
		So(err, ShouldBeNil)
		l := array.Len()
		So(l, ShouldEqual, len(sliceData))
		//log.Printf("len --> %v", len)
	})
}

func TestManager_ForEach(t *testing.T) {
	Convey("test foreach", t, func() {
		array, err := Array(sliceData)
		So(err, ShouldBeNil)
		array.ForEach(func(v interface{}, i int) {
			//类型转换
			/*o := v.(int)
			log.Println(reflect.TypeOf(o))*/
			num := v.(int) - 1
			So(num, ShouldEqual, i)
			// log.Printf("v --> %v, %v", v, i)
		})
	})
}

func TestManager_Concat(t *testing.T) {
	Convey("test Concat", t, func() {
		array, err := Array(sliceData)
		So(err, ShouldBeNil)

		arr := [3]int{4, 5, 6}
		slice := []int{7, 8, 9}
		v := 5

		newData := array.Concat(arr, slice, v).([]int)
		length := len(sliceData) + len(arr) + len(slice) + len(strconv.Itoa(v))

		So(len(newData), ShouldEqual, length)
		log.Printf("%v", newData)

		// 这里的测试性能
		slice1 := make([]int, 0)
		slice2 := make([]int, 0)
		l := 10000
		t0 := time.Now()
		for i := 1; i < l; i++ {
			slice1 = append(slice1, sliceData...)
		}
		t1 := time.Now()
		newArray, _ := Array(slice2)
		for i := 1; i < l; i++ {
			newArray.Concat(sliceData)
		}
		slice2 = newArray.GetData().([]int)
		t2 := time.Now()
		log.Println("reflect Concat insert:", t2.Sub(t1), "append insert: ", t1.Sub(t0))
		So(reflect.DeepEqual(slice1, slice2), ShouldBeTrue)
	})
}

func TestManager_CopyWithin(t *testing.T) {
	Convey("test CopyWithin", t, func() {
		arr := []int{1, 2, 3, 4, 5, 6}
		array, err := Array(arr)
		So(err, ShouldBeNil)
		getArray := array.CopyWithin(2, 1, 5).([]int)
		target := []int{1, 2, 2, 3, 4, 5}
		So(reflect.DeepEqual(getArray, target), ShouldBeTrue)
	})
}

func TestManager_Every(t *testing.T) {
	Convey("test every", t, func() {
		array, err := Array(sliceData)
		So(err, ShouldBeNil)
		bool := array.Every(func(v interface{}, i int) bool {
			o := v.(int)
			log.Printf("%v --> %v", o, i)
			return o > 3
		})
		So(bool, ShouldBeFalse)
	})
}

func TestManager_Fill(t *testing.T) {
	Convey("test fill", t, func() {
		arr := []string{"Banana", "Orange", "Apple", "Mango", "Pear", "Pineapple"}
		array, err := Array(arr)
		So(err, ShouldBeNil)
		data := array.Fill("Runoob", 1).([]string)
		target := []string{"Banana", "Runoob", "Runoob", "Runoob", "Runoob", "Runoob"}
		// log.Println(data)
		So(reflect.DeepEqual(data, target), ShouldBeTrue)
	})
}

func TestManager_Filter(t *testing.T) {
	Convey("test filter", t, func() {
		array, err := Array(sliceData)
		So(err, ShouldBeNil)
		slice := array.Filter(func(v interface{}, i int) bool {
			o := v.(int)
			return o > 1
		})
		So(len(slice.([]int)), ShouldEqual, 2)
	})
}

func TestManager_Fine(t *testing.T) {
	Convey("test Fine", t, func() {
		array, err := Array(sliceData)
		So(err, ShouldBeNil)
		num := 2
		v := array.Fine(func(v interface{}, i int) bool {
			o := v.(int)
			return o > num
		})

		if num >= len(sliceData) { // not find
			So(v, ShouldBeNil)
		} else {
			log.Println(v)
			So(v.(int), ShouldHaveSameTypeAs, 0)
		}
	})
}

func TestManager_FineIndex(t *testing.T) {
	Convey("test FineIndex", t, func() {
		array, err := Array(sliceData)
		So(err, ShouldBeNil)
		num := 3
		index := array.FineIndex(func(v interface{}, i int) bool {
			o := v.(int)
			return o > num
		})
		if num >= len(sliceData)-1 { // not find
			So(index, ShouldEqual, -1)
		} else {
			log.Println(index)
			So(index, ShouldBeGreaterThanOrEqualTo, 0)
		}
	})
}

func TestManager_Includes(t *testing.T) {
	Convey("test Includes", t, func() {
		array, err := Array(sliceData)
		So(err, ShouldBeNil)
		bool, err := array.Includes("3")
		So(err, ShouldBeNil)
		So(bool, ShouldBeFalse)
	})
}

func TestManager_IndexOf(t *testing.T) {
	Convey("test IndexOf", t, func() {
		array, err := Array(sliceData)
		So(err, ShouldBeNil)
		index, err := array.IndexOf(3)
		So(err, ShouldBeNil)
		So(index, ShouldEqual, 2)
	})
}