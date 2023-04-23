package datatypes

import "fmt"

type Int16 int16

func (i Int16) ToString() String {
	return String(fmt.Sprintf("%d", i))
}

func (i Int16) ToFloat32() Float32 {
	return Float32(i)
}

func (i Int16) ToFloat64() Float64 {
	return Float64(i)
}
