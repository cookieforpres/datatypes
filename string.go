package datatypes

import (
	"fmt"
	"strconv"
	"strings"
)

type String string

func (s String) ToInt() (Int, error) {
	i, err := strconv.Atoi(string(s))
	return Int(i), err
}

func (s String) ToInt8() (Int8, error) {
	i, err := strconv.Atoi(string(s))
	return Int8(i), err
}

func (s String) ToInt16() (Int16, error) {
	i, err := strconv.Atoi(string(s))
	return Int16(i), err
}

func (s String) ToInt32() (Int32, error) {
	i, err := strconv.Atoi(string(s))
	return Int32(i), err
}

func (s String) ToInt64() (Int64, error) {
	i, err := strconv.Atoi(string(s))
	return Int64(i), err
}

func (s String) ToFloat32() (Float32, error) {
	f, err := strconv.ParseFloat(string(s), 64)
	return Float32(f), err
}

func (s String) ToFloat64() (Float64, error) {
	f, err := strconv.ParseFloat(string(s), 64)
	return Float64(f), err
}

func (s String) ToBool() (Bool, error) {
	b, err := strconv.ParseBool(string(s))
	return Bool(b), err
}

func (s String) ToByteArray() ByteArray {
	return ByteArray(s)
}

func (s String) Split(delimiter string) []String {
	split := strings.Split(string(s), delimiter)
	var result []String
	for _, s := range split {
		result = append(result, String(s))
	}

	return result
}

func (s String) Trim() String {
	return String(strings.TrimSpace(string(s)))
}

func (s String) TrimLeft() String {
	return String(strings.TrimLeft(string(s), " "))
}

func (s String) TrimRight() String {
	return String(strings.TrimRight(string(s), " "))
}

func (s String) TrimPrefix(prefix string) String {
	return String(strings.TrimPrefix(string(s), prefix))
}

func (s String) TrimSuffix(suffix string) String {
	return String(strings.TrimSuffix(string(s), suffix))
}

func (s String) Replace(old, new string, n int) String {
	return String(strings.Replace(string(s), old, new, n))
}

func (s String) ReplaceAll(old, new string) String {
	return String(strings.ReplaceAll(string(s), old, new))
}

func (s String) Contains(substr string) bool {
	return strings.Contains(string(s), substr)
}

func (s String) ContainsAny(chars string) bool {
	return strings.ContainsAny(string(s), chars)
}

func (s String) Sprintf(args ...interface{}) String {
	return String(fmt.Sprintf(string(s), args...))
}
