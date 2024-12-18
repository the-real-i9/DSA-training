package graph_al

import (
	"dsa/dataStructures/graph"
	"encoding/json"
	"fmt"
)

func Run() {
	myData := graph.IntVertexVal(5)
	yourData := graph.IntVertexVal(4)

	/* fmt.Println(myData.Data())
	fmt.Println(myData.GreaterThan(yourData))
	fmt.Println(myData.LessThan(yourData))
	fmt.Println(myData.EqualTo(yourData))
	fmt.Println(myData.EqualTo(myData)) */

	myGraph := &Graph{}
	myGraph.AddVertex(myData)
	myGraph.AddVertex(yourData)

	graphView, _ := json.MarshalIndent(myGraph, "", "  ")

	fmt.Printf("%s\n", graphView)
}
