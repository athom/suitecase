# SuiteCase


A handy lib for converting between cases, from camel case to snake case and hyphen case.

[![Build Status](https://api.travis-ci.org/athom/suitecase.png?branch=master)](https://travis-ci.org/athom/suitecase)
[![GoDoc](https://godoc.org/github.com/athom/suitecase?status.png)](http://godoc.org/github.com/athom/suitecase)

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
