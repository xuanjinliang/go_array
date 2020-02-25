package go_array

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

var array = []int{1, 2, 3}

func TestArray(t *testing.T) {
	Convey("test", t, func() {
		array, err := Array(array)
		So(err, ShouldBeNil)
		log.Printf("%v", *array)
	})
}

func TestManager_Len(t *testing.T) {
	Convey("test len", t, func() {
		array, err := Array(array)
		So(err, ShouldBeNil)
		len := array.Len()
		log.Printf("%v", len)
	})
}

func TestManager_ForEach(t *testing.T) {
	Convey("test foreach", t, func() {
		array, err := Array(array)
		So(err, ShouldBeNil)
		array.ForEach(func(v interface{}, i int) {
			//类型转换
			/*o := v.(int)
			log.Println(reflect.TypeOf(o))*/

			log.Printf("v --> %v, %v", v, i)
		})
	})
}
