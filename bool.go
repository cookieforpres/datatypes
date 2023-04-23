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
