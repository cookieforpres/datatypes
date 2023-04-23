package datatypes

type ByteArray []Byte

func (b ByteArray) ToString() String {
	return String(b)
}

func (b ByteArray) Concat(other ByteArray) ByteArray {
	return append(b, other...)
}
