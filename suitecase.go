package suitecase

import "strings"

// Underscore can convert all upper case characters in a string to
// underscore format.
//
// Alias to ToSnakeCase
func Underscore(s string) (r string) {
	return ToSnakeCase(s)
}

// Camelize can convert all upper case characters in a string to
// camel format.
//
// Alias to ToCamelCase
func Camelize(s string) (r string) {
	return ToCamelCase(s)
}

// Dasherize can convert all upper case characters in a string to
// hyphen format.
//
// Alias to ToHyphenCase
func Dasherize(s string) (r string) {
	return ToHyphenCase(s)
}

// ToHyphenCase can convert all upper case characters in a string to
// hyphen format.
//
// Some samples.
//     "FirstName"  => "first-name"
//     "HTTPServer" => "http-server"
//     "NoHTTPS"    => "no-https"
func ToHyphenCase(s string) (r string) {
	r = ToSnakeCase(s)
	r = strings.Replace(r, `_`, `-`, -1)
	return
}

// ToSnakeCase can convert all upper case characters in a string to
// underscore format.
//
// Some samples.
//     "FirstName"  => "first_name"
//     "HTTPServer" => "http_server"
//     "NoHTTPS"    => "no_https"
func ToSnakeCase(s string) (r string) {
	var visited stack = make([]rune, 0)
	var array []string

	for _, chr := range s {
		if isUpper(chr) {
			if !visited.isEmpty() && !isUpper(visited.top()) {
				array = append(array, strings.ToLower(visited.String()))
				visited.clear()
			}
			visited.push(chr)
		} else {
			if visited.isEmpty() || !isUpper(visited.top()) {
				visited.push(chr)
				continue
			}

			if visited.depth() > 1 {
				a := visited.pop()
				array = append(array, strings.ToLower(visited.String()))
				visited.clear()
				visited.push(a)
			}
			visited.push(chr)
		}
	}

	if !visited.isEmpty() {
		array = append(array, strings.ToLower(visited.String()))
	}

	r = strings.Join(array, "_")
	return
}

// ToCamelCase can convert all lower case characters behind underscores
// to upper case character.
// Underscore character will be removed in result except following cases.
//     * More than 1 underscore, e.g. "a__b" => "A_B"
//     * At the beginning of string, e.g. "_a" => "_A"
//     * At the end of string, e.g. "ab_" => "Ab_"
//     * At the beginning and end of string, e.g. "_ab_" => "_Ab_"
func ToCamelCase(s string) (r string) {
	array := strings.Split(s, "_")
	var newArray []string

	for _, ele := range array {
		newstr := make([]rune, 0)
		for i, s := range []rune(ele) {
			if i == 0 {
				if isLower(s) {
					s += 'A' - 'a'
				}
				newstr = append(newstr, s)
				continue
			}
			newstr = append(newstr, s)
		}
		ns := string(newstr)
		if ns == "" {
			ns = `_`
		}
		newArray = append(newArray, ns)
	}

	r = strings.Join(newArray, "")
	return
}

func isUpper(c rune) bool {
	return 'A' <= c && c <= 'Z'
}
func isLower(c rune) bool {
	return 'a' <= c && c <= 'z'
}
