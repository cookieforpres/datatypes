package datatypes

import "fmt"

type Int64 int64

func (i Int64) ToString() String {
	return String(fmt.Sprintf("%d", i))
}

func (i Int64) ToFloat32() Float32 {
	return Float32(i)
}

func (i Int64) ToFloat64() Float64 {
	return Float64(i)
}
