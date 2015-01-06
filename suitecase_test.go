package suitecase

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToHyphenCase(t *testing.T) {
	Convey("to snake case", t, func() {
		Convey("no terminology", func() {
			So(ToHyphenCase("a"), ShouldEqual, "a")
			So(ToHyphenCase("A"), ShouldEqual, "a")
			So(ToHyphenCase("Aa"), ShouldEqual, "aa")
			So(ToHyphenCase("aA"), ShouldEqual, "a-a")
			So(ToHyphenCase("WinterIsComming"), ShouldEqual, "winter-is-comming")
		})

		Convey("terminology", func() {
			So(ToHyphenCase("AA"), ShouldEqual, "aa")
			So(ToHyphenCase("HTTPServer"), ShouldEqual, "http-server")
			So(ToHyphenCase("Yeer Kunst"), ShouldEqual, "yeer-kunst")
			So(ToHyphenCase("Yeer_Kunst"), ShouldEqual, "yeer-kunst")
			So(ToHyphenCase("Yeer        Kunst"), ShouldEqual, "yeer-kunst")
		})
	})
}

func TestToSnakeCase(t *testing.T) {
	Convey("to snake case", t, func() {
		Convey("no terminology", func() {
			So(ToSnakeCase("a"), ShouldEqual, "a")
			So(ToSnakeCase("A"), ShouldEqual, "a")
			So(ToSnakeCase("Aa"), ShouldEqual, "aa")
			So(ToSnakeCase("aA"), ShouldEqual, "a_a")
			So(ToSnakeCase("WinterIsComming"), ShouldEqual, "winter_is_comming")
			So(ToSnakeCase("_camelCase"), ShouldEqual, "_camel_case")
			So(ToSnakeCase("NoHTTPS"), ShouldEqual, "no_https")
			So(ToSnakeCase("Wi_thF"), ShouldEqual, "wi_th_f")
			So(ToSnakeCase("_AnotherTES_TCaseP"), ShouldEqual, "_another_tes_t_case_p")
			So(ToSnakeCase("ALL"), ShouldEqual, "all")
			So(ToSnakeCase("A_B"), ShouldEqual, "a_b")
			So(ToSnakeCase("_A_B_"), ShouldEqual, "_a_b_")
			So(ToSnakeCase("_A___B_"), ShouldEqual, "_a___b_")
			So(ToSnakeCase("A B"), ShouldEqual, "a_b")
			So(ToSnakeCase("A-B"), ShouldEqual, "a_b")
			So(ToSnakeCase("A   B"), ShouldEqual, "a_b")
			So(ToSnakeCase("A b"), ShouldEqual, "a_b")
			So(ToSnakeCase("A _b"), ShouldEqual, "a_b")
			So(ToSnakeCase("A B c"), ShouldEqual, "a_b_c")
			So(ToSnakeCase("A   B C d"), ShouldEqual, "a_b_c_d")
		})

		Convey("terminology", func() {
			So(ToSnakeCase("AA"), ShouldEqual, "aa")
			So(ToSnakeCase("HTTPServer"), ShouldEqual, "http_server")
		})
	})
}

func TestToCamelCase(t *testing.T) {
	Convey("to camel case", t, func() {
		So(ToCamelCase("a"), ShouldEqual, "A")
		So(ToCamelCase("aa"), ShouldEqual, "Aa")
		So(ToCamelCase("a_a"), ShouldEqual, "AA")
		So(ToCamelCase("_a"), ShouldEqual, "_A")
		So(ToCamelCase("AA"), ShouldEqual, "AA")
		So(ToCamelCase("http_server"), ShouldEqual, "HttpServer")
		So(ToCamelCase("no_https"), ShouldEqual, "NoHttps")
		So(ToCamelCase("_left_side"), ShouldEqual, "_LeftSide")
		So(ToCamelCase("right_side_"), ShouldEqual, "RightSide_")
		So(ToCamelCase("_both_side_"), ShouldEqual, "_BothSide_")
		So(ToCamelCase("_complex__case_"), ShouldEqual, "_Complex_Case_")
	})
}
