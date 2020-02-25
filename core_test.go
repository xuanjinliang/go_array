package go_array

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

var array = []int{1, 2, 3}

func TestNewArray(t *testing.T) {
	Convey("test", t, func() {
		array, err := Array(array)
		So(err, ShouldBeNil)
		log.Printf("%v", *array)
		len := array.Len()
		log.Printf("%v", len)
	})
}
