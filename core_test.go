package go_array

import (
	"log"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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
		l := len(sliceData) + len(arr) + len(slice) + len(strconv.Itoa(v))

		So(len(newData), ShouldEqual, l)
		log.Printf("%v, %v", newData)
	})
}

