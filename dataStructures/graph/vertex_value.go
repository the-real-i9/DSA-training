package graph

type VertexValue interface {
	EqualTo(VertexValue) bool
	LessThan(VertexValue) bool
	GreaterThan(VertexValue) bool
	Data() any
}

type IntVertexValue int

func (v IntVertexValue) Data() any {
	return int(v)
}

func (a IntVertexValue) EqualTo(b VertexValue) bool {
	return a.Data().(int) == b.(IntVertexValue).Data().(int)
}

func (a IntVertexValue) LessThan(b VertexValue) bool {
	return a.Data().(int) < b.(IntVertexValue).Data().(int)
}

func (a IntVertexValue) GreaterThan(b VertexValue) bool {
	return a.Data().(int) > b.(IntVertexValue).Data().(int)
}
