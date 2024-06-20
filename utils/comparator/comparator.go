package comparator

import "cmp"

type Comparator struct {
	Compare func(a, b any) int
}

func (c Comparator) Equal(a, b any) bool {
	if a == nil || b == nil {
		return a == b
	}

	return c.Compare(a, b) == 0
}

func (c Comparator) LessThan(a, b any) bool {
	if a == nil || b == nil {
		return false
	}

	return c.Compare(a, b) < 0
}

func (c Comparator) GreaterThan(a, b any) bool {
	if a == nil || b == nil {
		return false
	}

	return c.Compare(a, b) > 0
}

func (c Comparator) LessThanOrEqual(a, b any) bool {
	return c.LessThan(a, b) || c.Equal(a, b)
}

func (c Comparator) GreaterThanOrEqual(a, b any) bool {
	return c.GreaterThan(a, b) || c.Equal(a, b)
}

func (c *Comparator) Reverse() {
	compareOriginal := c.Compare

	c.Compare = func(a, b any) int {
		return compareOriginal(b, a)
	}
}

func DefaultCompareFunc(a, b any) int {
	return cmp.Compare(a.(int), b.(int))
}
