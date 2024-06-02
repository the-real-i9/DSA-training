package tree

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type BTNode struct {
	Left  *BTNode
	Data  any
	Right *BTNode
}
