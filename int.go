package datatypes

import "fmt"

type Int int

func (i Int) ToString() String {
	return String(fmt.Sprintf("%d", i))
}

func (i Int) ToFloat32() Float32 {
	return Float32(i)
}

func (i Int) ToFloat64() Float64 {
	return Float64(i)
}
