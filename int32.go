package datatypes

import "fmt"

type Int32 int32

func (i Int32) ToString() String {
	return String(fmt.Sprintf("%d", i))
}

func (i Int32) ToFloat32() Float32 {
	return Float32(i)
}

func (i Int32) ToFloat64() Float64 {
	return Float64(i)
}
