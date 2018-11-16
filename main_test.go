package main

import "testing"
import "github.com/google/go-cmp/cmp"
import . "github.com/smartystreets/goconvey/convey"

func TestBDDStyle(t *testing.T) {
	Convey("Given three structs", t, func() {
		t1 := T1{"VAL1", 12}
		t2 := T1{"VAL2", 12}
		t3 := T1{"VAL3", 1}

		Convey("When t1 and t2 are compared", func() {
			areEqual := cmp.Equal(t1, t2)

			Convey("Then they should be equal", func() {
				So(areEqual, ShouldBeTrue)
			})
		})

		Convey("When t1 and t3 is added and compared", func() {

			areEqual := cmp.Equal(t1, t3)

			Convey("Then they should not be equal", func() {
				So(areEqual, ShouldBeFalse)
			})
		})
	})
}
