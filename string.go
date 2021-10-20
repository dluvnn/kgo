package kgo

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"unicode"
)

// URLSplit2 splits url into two paths
func URLSplit2(url string) (string, string) {
	b := []byte(url)
	n := len(b)
	for i := 1; i < n; i++ {
		if b[i] == '/' {
			return url[:i], url[i:]
		}
	}
	return url, "/"
}

// StringsToInterfaces ...
func StringsToInterfaces(s []string) []interface{} {
	var a []interface{}
	n := len(s)
	if n > 0 {
		a = make([]interface{}, n)
		for i := 0; i < n; i++ {
			a[i] = s[i]
		}
	}
	return a
}

// URLSplit3 splits url into three paths
func URLSplit3(url string) (string, string, string) {
	b := []byte(url)
	n := len(b)
	var i1 int
	for i1 = 1; i1 < n; i1++ {
		if b[i1] == '/' {
			break
		}
	}
	if i1 == n {
		return "/", "/", "/"
	}
	var i2 int
	for i2 = i1 + 1; i2 < n; i2++ {
		if b[i2] == '/' {
			break
		}
	}
	if i2 == i1 {
		return url[:i1], url[i1:], "/"
	}

	return url[:i1], url[i1:i2], url[i2:]
}

// URLFirstTwoPath ...
func URLFirstTwoPath(url string) (string, string) {
	b := []byte(url)
	n := len(b)
	var i1 int
	for i1 = 1; i1 < n; i1++ {
		if b[i1] == '/' {
			break
		}
	}
	if i1 == n {
		if n > 0 {
			return url[1:], ""
		}
		return "", ""
	}
	var i2 int
	for i2 = i1 + 1; i2 < n; i2++ {
		if b[i2] == '/' {
			break
		}
	}
	if i2 == i1 {
		return url[1:i1], url[i1+1:]
	}

	return url[1:i1], url[i1+1 : i2]
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	n := len(r) / 2
	for i, j := 0, len(r)-1; i < n; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isSlashRune(r rune) bool { return r == '/' || r == '\\' }

// ContainsDotDot ...
func ContainsDotDot(v string) bool {
	if !strings.Contains(v, "..") {
		return false
	}
	for _, ent := range strings.FieldsFunc(v, isSlashRune) {
		if ent == ".." {
			return true
		}
	}
	return false
}

// ListContainString ...
func ListContainString(v []string, s string) bool {
	for _, x := range v {
		if x == s {
			return true
		}
	}
	return false
}

// RemoveSpace ...
func RemoveSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// ToXLink converts path to style 'a/b/c'
func ToXLink(path string) string {
	name := path
	n := len(name)
	if n == 0 {
		return name
	}
	if name[0] == '/' {
		name = name[1:]
		n = n - 1
	}
	n = n - 1
	if n > 0 && name[n] == '/' {
		name = name[:n]
	}
	return name
}

// FuncName ...
func FuncName(x interface{}) string {
	full := runtime.FuncForPC(reflect.ValueOf(x).Pointer()).Name()
	v := strings.Split(full, ".")
	n := len(v)
	if n == 0 {
		return ""
	}
	return v[n-1]
}

// PrintString converts an interface object to string
func PrintString(x interface{}) string {
	return fmt.Sprintf("%v", x)
}
