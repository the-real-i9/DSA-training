package tree

type User struct {
	Id       int
	Username string
	Name     string
}

type BTNode struct {
	Left  *BTNode
	Data  any
	Right *BTNode
}
