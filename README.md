# SuiteCase


A handy lib for converting between cases, from camel case to snake case and hypen case.


## Installation

```bash
	go get "github.com/athom/suitecase"
```

## Usage

###### 1. ToSnakeCase

```go
a := suitecase.ToSnakeCase("HTTPServer")
// a = "http_server"
```

###### 2. ToCamelCase

```go
a := suitecase.ToSnakeCase("find_object_by_id")
// a = "FindObjectById"
```

###### 3. ToHypenCase

```go
a := suitecase.ToSnakeCase("HTTPServer")
// a = "http-server"
```


## License

SuiteCase is released under the [WTFPL License](http://www.wtfpl.net/txt/copying).
