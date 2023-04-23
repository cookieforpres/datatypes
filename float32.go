package datatypes

import "fmt"

type Float32 float32

func (f Float32) ToInt() Int {
	return Int(f)
}

func (f Float32) ToInt8() Int8 {
	return Int8(f)
}

func (f Float32) ToInt16() Int16 {
	return Int16(f)
}

func (f Float32) ToInt32() Int32 {
	return Int32(f)
}

func (f Float32) ToInt64() Int64 {
	return Int64(f)
}

func (f Float32) ToString() String {
	return String(fmt.Sprintf("%f", f))
}
