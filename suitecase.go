package suitecase

import "strings"

func Underscore(s string) (r string) {
	return ToSnakeCase(s)
}

func Camelize(s string) (r string) {
	return ToCamelCase(s)
}

func Dasherize(s string) (r string) {
	return ToHyphenCase(s)
}

func ToHyphenCase(s string) (r string) {
	r = ToSnakeCase(s)
	r = strings.Replace(r, `_`, `-`, -1)
	return
}

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

func ToCamelCase(s string) (r string) {
	array := strings.Split(s, "_")
	var newArray []string

	for _, ele := range array {
		newstr := make([]rune, 0)
		for i, s := range []rune(ele) {
			if i == 0 {
				if !isUpper(s) {
					s += 'A' - 'a'
				}
				newstr = append(newstr, s)
				continue
			}
			newstr = append(newstr, s)
		}
		newArray = append(newArray, string(newstr))
	}

	r = strings.Join(newArray, "")
	return
}

func isUpper(c rune) bool {
	return 'A' <= c && c <= 'Z'
}
