package suitecase

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStack(t *testing.T) {
	Convey("stack init", t, func() {
		var s stack = make([]rune, 0)
		So(s.depth(), ShouldEqual, 0)
		So(s.isEmpty(), ShouldBeTrue)
	})
	Convey("stack push", t, func() {
		var s stack = make([]rune, 0)
		s.push('A')
		s.push('B')
		So(s.depth(), ShouldEqual, 2)
		So(s.top(), ShouldEqual, 'B')

	})
	Convey("stack pop", t, func() {
		var s stack = make([]rune, 0)
		s.push('A')
		s.push('B')
		s.push('C')
		So(s.top(), ShouldEqual, 'C')
		c := s.pop()
		So(c, ShouldEqual, 'C')
		So(s.top(), ShouldEqual, 'B')
		c = s.pop()
		So(c, ShouldEqual, 'B')
		So(s.top(), ShouldEqual, 'A')
		c = s.pop()
		So(c, ShouldEqual, 'A')
		So(s.top(), ShouldEqual, 0)
	})

	Convey("stack String()", t, func() {
		var s stack = make([]rune, 0)
		s.push('A')
		s.push('B')
		s.push('C')
		So(s.String(), ShouldEqual, "ABC")
		s.pop()
		So(s.String(), ShouldEqual, "AB")
	})

	Convey("stack clear", t, func() {
		var s stack = make([]rune, 0)
		s.push('A')
		s.push('B')
		s.push('C')
		s.clear()
		So(s.String(), ShouldEqual, "")
	})
}
