package suitecase

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToHyphenCase(t *testing.T) {
	Convey("to snake case", t, func() {
		Convey("no terminology", func() {
			So(
				ToHyphenCase("a"),
				ShouldEqual,
				"a",
			)
			So(
				ToHyphenCase("A"),
				ShouldEqual,
				"a",
			)
			So(
				ToHyphenCase("Aa"),
				ShouldEqual,
				"aa",
			)
			So(
				ToHyphenCase("aA"),
				ShouldEqual,
				"a-a",
			)
			So(
				ToHyphenCase("WinterIsComming"),
				ShouldEqual,
				"winter-is-comming",
			)
		})

		Convey("terminology", func() {
			So(
				ToHyphenCase("AA"),
				ShouldEqual,
				"aa",
			)
			So(
				ToHyphenCase("HTTPServer"),
				ShouldEqual,
				"http-server",
			)
		})
	})
}

func TestToSnakeCase(t *testing.T) {
	Convey("to snake case", t, func() {
		Convey("no terminology", func() {
			So(
				ToSnakeCase("a"),
				ShouldEqual,
				"a",
			)
			So(
				ToSnakeCase("A"),
				ShouldEqual,
				"a",
			)
			So(
				ToSnakeCase("Aa"),
				ShouldEqual,
				"aa",
			)
			So(
				ToSnakeCase("aA"),
				ShouldEqual,
				"a_a",
			)
			So(
				ToSnakeCase("WinterIsComming"),
				ShouldEqual,
				"winter_is_comming",
			)
		})

		Convey("terminology", func() {
			So(
				ToSnakeCase("AA"),
				ShouldEqual,
				"aa",
			)
			So(
				ToSnakeCase("HTTPServer"),
				ShouldEqual,
				"http_server",
			)
		})
	})
}

func TestToCamelCase(t *testing.T) {
	Convey("to camel case", t, func() {
		So(
			ToCamelCase("a"),
			ShouldEqual,
			"A",
		)
		So(
			ToCamelCase("aa"),
			ShouldEqual,
			"Aa",
		)
		So(
			ToCamelCase("a_a"),
			ShouldEqual,
			"AA",
		)
		So(
			ToCamelCase("AA"),
			ShouldEqual,
			"AA",
		)
		So(
			ToCamelCase("http_server"),
			ShouldEqual,
			"HttpServer",
		)
	})
}
