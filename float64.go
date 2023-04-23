package datatypes

import "fmt"

type Float64 float64

func (f Float64) ToInt() Int {
	return Int(f)
}

func (f Float64) ToInt8() Int8 {
	return Int8(f)
}

func (f Float64) ToInt16() Int16 {
	return Int16(f)
}

func (f Float64) ToInt32() Int32 {
	return Int32(f)
}

func (f Float64) ToInt64() Int64 {
	return Int64(f)
}

func (f Float64) ToString() String {
	return String(fmt.Sprintf("%f", f))
}
