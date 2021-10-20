package kgo

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/template"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func inc(i int) int {
	return i + 1
}

func dec(i int) int {
	return i - 1
}

func mod(i, d int) int {
	return i % d
}

func fmul(x, y float64) float64 {
	return x * y
}

func fdiv(x, y float64) float64 {
	return x / y
}

func fadd(x, y float64) float64 {
	return x + y
}

func fsub(x, y float64) float64 {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func ToInt(x interface{}) int {
	i := 0
	switch v := x.(type) {
	case int8:
		i = int(v)
	case int16:
		i = int(v)
	case int32:
		i = int(v)
	case int64:
		i = int(v)
	case int:
		i = int(v)
	case uint:
		i = int(v)
	case uint32:
		i = int(v)
	case uint16:
		i = int(v)
	case uint8:
		i = int(v)
	case uint64:
		i = int(v)
	case float32:
		i = int(v)
	case float64:
		i = int(v)
	case string:
		if len(v) == 0 {
			return 0
		}
		var err error
		i, err = strconv.Atoi(v)
		if err != nil {
			return 0
		}
	}
	return i
}

func ToFloat(x interface{}) float64 {
	var f float64 = 0
	switch v := x.(type) {
	case int8:
		f = float64(v)
	case int16:
		f = float64(v)
	case int32:
		f = float64(v)
	case int64:
		f = float64(v)
	case int:
		f = float64(v)
	case uint:
		f = float64(v)
	case uint32:
		f = float64(v)
	case uint16:
		f = float64(v)
	case uint8:
		f = float64(v)
	case uint64:
		f = float64(v)
	case float32:
		f = float64(v)
	case float64:
		f = float64(v)
	case string:
		var err error
		f, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return 0
		}
	}
	return f
}

func ToBool(x interface{}) bool {
	switch v := x.(type) {
	case bool:
		return v
	case int64:
		return v != 0
	case uint64:
		return v != 0
	case int32:
		return v != 0
	case uint32:
		return v != 0
	case float64:
		return v != 0
	case float32:
		return v != 0
	}
	return false
}

func ToArray(x interface{}) []interface{} {
	a, _ := x.([]interface{})
	return a
}

func noPrint(x ...interface{}) string {
	return ""
}

func CorrectBackSlash(str string) string {
	return strings.ReplaceAll(str, `\`, `\\`)
}

// TempMap ...
type TempMap map[string]interface{}

// Set ...
func (tm TempMap) Set(key string, value interface{}) string {
	tm[key] = value
	return ""
}

// Get ...
func (tm TempMap) Get(key string) interface{} {
	v, ok := tm[key]
	if ok {
		return v
	}
	return ""
}

// Has ...
func (tm TempMap) Has(key string) bool {
	_, ok := tm[key]
	return ok
}

// ToJSON ...
func (tm TempMap) ToJSON() string {
	data, err := json.Marshal(tm)
	if err != nil {
		Error(err)
		return ""
	}
	return string(data)
}

func CreateMap(x ...interface{}) TempMap {
	n := len(x)
	if n == 0 {
		return TempMap{}
	}
	n = n - 1
	m := TempMap{}
	for i := 0; i < n; i = i + 2 {
		k, ok := x[i].(string)
		if ok {
			m[k] = x[i+1]
		}
	}
	return m
}

func CreateList(x ...interface{}) []interface{} {
	if len(x) == 0 {
		return []interface{}{}
	}
	return x
}

func AppendList(ls []interface{}, x ...interface{}) []interface{} {
	return append(ls, x...)
}

func xlog(x ...interface{}) string {
	log.Println(x...)
	return ""
}

func SafeAt(a interface{}, index int) interface{} {
	switch x := a.(type) {
	case []interface{}:
		if index >= len(x) {
			return nil
		}
		return x[index]
	case []string:
		if index >= len(x) {
			return nil
		}
		return x[index]
	case []map[string]interface{}:
		if index >= len(x) {
			return nil
		}
		return x[index]
	case []TempMap:
		if index >= len(x) {
			return nil
		}
		return x[index]
	case []int:
		if index >= len(x) {
			return nil
		}
		return x[index]
	case []float64:
		if index >= len(x) {
			return nil
		}
		return x[index]
	case []int64:
		if index >= len(x) {
			return nil
		}
		return x[index]
	case []int32:
		if index >= len(x) {
			return nil
		}
		return x[index]
	case []bool:
		if index >= len(x) {
			return nil
		}
		return x[index]
	case map[string]interface{}:
		if index == 0 {
			return x
		}
		return nil
	case TempMap:
		if index == 0 {
			return x
		}
		return nil
	}
	return nil
}

func ToMap(x interface{}) TempMap {
	switch v := x.(type) {
	case map[string]interface{}:
		return TempMap(v)
	case TempMap:
		return v
	case bson.M:
		return TempMap(v)
	case map[string]string:
		{
			m := TempMap{}
			for k, s := range v {
				m[k] = s
			}
			return m
		}
	}
	return TempMap{}
}

func Field(x map[string]interface{}, key string) interface{} {
	v, _ := x[key]
	return v
}

func ToString(x interface{}) string {
	s, _ := x.(string)
	return s
}

func ToJSONString(x interface{}) string {
	data, err := json.Marshal(x)
	if err != nil {
		Error(err)
		return ""
	}
	return string(data)
}

func ParseJSONString(str string, x interface{}) error {
	return json.Unmarshal([]byte(str), x)
}

func w3(s string) string {
	return fmt.Sprintf("{{ %s }}", s)
}

func w3x(s string) string {
	return fmt.Sprintf("{{- %s -}}", s)
}

// JoinString ...
func JoinString(sep string, s ...string) string {
	return strings.Join(s, sep)
}

// FMap ...
var FMap = template.FuncMap{
	"w3":               w3,
	"w3x":              w3x,
	"inc":              inc,
	"dec":              dec,
	"mod":              mod,
	"fmul":             fmul,
	"fdiv":             fdiv,
	"fadd":             fadd,
	"fsub":             fsub,
	"mul":              mul,
	"div":              div,
	"add":              add,
	"sub":              sub,
	"toInt":            ToInt,
	"toFloat":          ToFloat,
	"toBool":           ToBool,
	"toArray":          ToArray,
	"toMap":            ToMap,
	"toString":         ToString,
	"newObjectID":      primitive.NewObjectID,
	"hashPassword":     HashPassword,
	"randPassword":     RandPassword,
	"noPrint":          noPrint,
	"correctBackSlash": CorrectBackSlash,
	"now":              time.Now,
	"map":              CreateMap,
	"log":              xlog,
	"list":             CreateList,
	"append":           AppendList,
	"field":            Field,
	"split":            strings.Split,
	"toLower":          strings.ToLower,
	"title":            strings.ToTitle,
	"toUpper":          strings.ToUpper,
	"toJSON":           ToJSONString,
	"reverse":          Reverse,
	"at":               SafeAt,
	"convertDate":      ConvertDate,
	"joinString":       JoinString,
}
