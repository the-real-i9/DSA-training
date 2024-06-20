package comparator

import "cmp"

func NewComparator(compareFunction func(a, b any) int) Comparator {
	if compareFunction == nil {
		return Comparator{compare: defaultCompareFunction}
	}

	return Comparator{compare: compareFunction}
}

type Comparator struct {
	compare func(a, b any) int
}

func (c Comparator) Equal(a, b any) bool {
	if a == nil || b == nil {
		return a == b
	}

	return c.compare(a, b) == 0
}

func (c Comparator) LessThan(a, b any) bool {
	if a == nil || b == nil {
		return false
	}

	return c.compare(a, b) < 0
}

func (c Comparator) GreaterThan(a, b any) bool {
	if a == nil || b == nil {
		return false
	}

	return c.compare(a, b) > 0
}

func (c Comparator) LessThanOrEqual(a, b any) bool {
	return c.LessThan(a, b) || c.Equal(a, b)
}

func (c Comparator) GreaterThanOrEqual(a, b any) bool {
	return c.GreaterThan(a, b) || c.Equal(a, b)
}

func (c *Comparator) Reverse() {
	compareOriginal := c.compare

	c.compare = func(a, b any) int {
		return compareOriginal(b, a)
	}
}

func defaultCompareFunction(a, b any) int {
	return cmp.Compare(a.(int), b.(int))
}
