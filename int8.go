package datatypes

import "fmt"

type Int8 int8

func (i Int8) ToString() String {
	return String(fmt.Sprintf("%d", i))
}

func (i Int8) ToFloat32() Float32 {
	return Float32(i)
}

func (i Int8) ToFloat64() Float64 {
	return Float64(i)
}
