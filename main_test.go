package main

import "testing"
import "github.com/google/go-cmp/cmp"
import . "github.com/smartystreets/goconvey/convey"

func TestDump(t *testing.T) {
	t1 := T1{"VAL1", 12}
	Dump(t1)
}

func TestEqual(t *testing.T) {
	t1 := T1{"VAL1", 12}
	t2 := T1{"VAL2", 12}

	result := cmp.Equal(t1, t2)
	if !result {
		t.Error("Expected equality on size")
	}

	t3 := T2{"VAL3", 12}
	t4 := T2{"VAL4", 12}

	opt := cmp.AllowUnexported(t3)
	result = cmp.Equal(t3, t4, opt)
	if result {
		t.Error("Did not expect equality on size")
	}
}

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
