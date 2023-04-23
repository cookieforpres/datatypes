package datatypes

type Byte byte

func (b Byte) ToInt() Int {
	return Int(b)
}

func (b Byte) ToInt8() Int8 {
	return Int8(b)
}

func (b Byte) ToInt16() Int16 {
	return Int16(b)
}

func (b Byte) ToInt32() Int32 {
	return Int32(b)
}

func (b Byte) ToInt64() Int64 {
	return Int64(b)
}

func (b Byte) ToFloat32() Float32 {
	return Float32(b)
}

func (b Byte) ToFloat64() Float64 {
	return Float64(b)
}
