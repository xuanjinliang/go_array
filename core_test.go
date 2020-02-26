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
		/*aa := append(arrayData, []int{1, 2, 3}...)
		log.Printf("%v", aa)*/
	})
}

var sliceData = []int{1, 2, 3}

func TestSlice(t *testing.T) {
	Convey("test", t, func() {
		array, err := Array(sliceData)
		So(err, ShouldBeNil)
		log.Printf("%v", array.GetData())
		/*aa := append(arrayData, []int{1, 2, 3}...)
		log.Printf("%v", aa)*/
	})
}

func TestManager_Len(t *testing.T) {
	Convey("test len", t, func() {
		array, err := Array(sliceData)
		So(err, ShouldBeNil)
		len := array.Len()
		log.Printf("len --> %v", len)
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

			log.Printf("v --> %v, %v", v, i)
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
		for i := 1; i < l; i++ {
			array, _ = Array(slice2)
			slice2 = array.Concat(sliceData).([]int)
		}
		t2 := time.Now()
		log.Println("reflect Concat insert:", t2.Sub(t1), "append insert: ", t1.Sub(t0))
		So(reflect.DeepEqual(slice1, slice2), ShouldBeTrue)
	})
}

func TestManager_CopyWithin(t *testing.T) {
	Convey("test Concat", t, func() {
		arr := []int{1, 2, 3, 4, 5, 6}
		array, err := Array(arr)
		So(err, ShouldBeNil)
		getArray := array.CopyWithin(2, 1, 8).([]int)
		log.Printf("%v", getArray)
	})
}
