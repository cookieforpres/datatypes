package datatypes

import "fmt"

type Float32 float32

func (f Float32) ToInt() Int {
	return Int(f)
}

func (f Float32) ToString() String {
	return String(fmt.Sprintf("%f", f))
}
