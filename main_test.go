package main

import "testing"
import "github.com/google/go-cmp/cmp"
import . "github.com/smartystreets/goconvey/convey"

func TestBDDStyle(t *testing.T) {
	Convey("Given", t, func() {
		t1 := T1{"VAL1", 12}
		t2 := T1{"VAL2", 12}

		Convey("When", func() {
			areEqual := cmp.Equal(t1, t2)

			Convey("Then", func() {
				So(areEqual, ShouldBeTrue)
			})
		})
	})
}
