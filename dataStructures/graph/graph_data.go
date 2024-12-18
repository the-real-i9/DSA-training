package graph

type VertexValue interface {
	EqualTo(VertexValue) bool
	LessThan(VertexValue) bool
	GreaterThan(VertexValue) bool
	Data() any
}

type IntVertexVal int

func (v IntVertexVal) Data() any {
	return int(v)
}

func (a IntVertexVal) EqualTo(b VertexValue) bool {
	return a.Data().(int) == b.(IntVertexVal).Data().(int)
}

func (a IntVertexVal) LessThan(b VertexValue) bool {
	return a.Data().(int) < b.(IntVertexVal).Data().(int)
}

func (a IntVertexVal) GreaterThan(b VertexValue) bool {
	return a.Data().(int) > b.(IntVertexVal).Data().(int)
}
