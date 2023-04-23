package datatypes

import "fmt"

type Float64 float64

func (f Float64) ToInt() Int {
	return Int(f)
}

func (f Float64) ToString() String {
	return String(fmt.Sprintf("%f", f))
}
