package datatypes

type Bool bool

func (b Bool) ToString() String {
	if b {
		return String("true")
	}
	return String("false")
}

func (b Bool) ToInt() Int {
	if b {
		return Int(1)
	}
	return Int(0)
}

func (b Bool) ToInt8() Int8 {
	if b {
		return Int8(1)
	}
	return Int8(0)
}

func (b Bool) ToInt16() Int16 {
	if b {
		return Int16(1)
	}
	return Int16(0)
}

func (b Bool) ToInt32() Int32 {
	if b {
		return Int32(1)
	}
	return Int32(0)
}

func (b Bool) ToInt64() Int64 {
	if b {
		return Int64(1)
	}
	return Int64(0)
}

func (b Bool) ToFloat32() Float32 {
	if b {
		return Float32(1)
	}
	return Float32(0)
}

func (b Bool) ToFloat64() Float64 {
	if b {
		return Float64(1)
	}
	return Float64(0)
}
