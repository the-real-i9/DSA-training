package comparator

import "cmp"

type cmprt struct {
	Compare func(a, b any) int
}

func (c cmprt) Equal(a, b any) bool {
	return c.Compare(a, b) == 0
}
func (c cmprt) LessThan(a, b any) bool {
	return c.Compare(a, b) < 0
}
func (c cmprt) GreaterThan(a, b any) bool {
	return c.Compare(a, b) > 0
}
func (c cmprt) LessThanOrEqual(a, b any) bool {
	return c.LessThan(a, b) || c.Equal(a, b)
}
func (c cmprt) GreaterThanOrEqual(a, b any) bool {
	return c.GreaterThan(a, b) || c.Equal(a, b)
}

func (c *cmprt) Reverse() {
	compareOriginal := c.Compare

	c.Compare = func(a, b any) int {
		return compareOriginal(b, a)
	}
}

func defaultCompareFunction(a, b any) int {
	return cmp.Compare(a.(int), b.(int))
}

func Comparator(compareFunction func(a, b any) int) cmprt {
	if compareFunction == nil {
		return cmprt{Compare: defaultCompareFunction}
	}

	return cmprt{Compare: compareFunction}
}
