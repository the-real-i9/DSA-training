package graph

type EdgeWeight interface {
	EqualTo(EdgeWeight) bool
	LessThan(EdgeWeight) bool
	GreaterThan(EdgeWeight) bool
	Data() any
}

type IntEdgeWeight int

func (v IntEdgeWeight) Data() any {
	return int(v)
}

func (a IntEdgeWeight) EqualTo(b EdgeWeight) bool {
	return a.Data().(int) == b.(IntEdgeWeight).Data().(int)
}

func (a IntEdgeWeight) LessThan(b EdgeWeight) bool {
	return a.Data().(int) < b.(IntEdgeWeight).Data().(int)
}

func (a IntEdgeWeight) GreaterThan(b EdgeWeight) bool {
	return a.Data().(int) > b.(IntEdgeWeight).Data().(int)
}
