package graph_al

import (
	"dsa/dataStructures/graph"
	"fmt"
)

func Run() {
	myData := graph.IntVertexValue(5)
	yourData := graph.IntVertexValue(4)

	/* fmt.Println(myData.Data())
	fmt.Println(myData.GreaterThan(yourData))
	fmt.Println(myData.LessThan(yourData))
	fmt.Println(myData.EqualTo(yourData))
	fmt.Println(myData.EqualTo(myData)) */

	myGraph := &Graph{Directed: false}
	myGraph.AddVertex(myData)
	myGraph.AddVertex(yourData)
	myGraph.AddEdge(myData, yourData, nil)

	/* graphView, err := json.MarshalIndent(myGraph, "", "  ")
	if err != nil {
		log.Println(err)
		return
	} */

	// fmt.Printf("%s\n", graphView)
	fmt.Println(myGraph.Vertices[0].IncEdges)
	fmt.Println(myGraph.Vertices[1].IncEdges)
	fmt.Println(myGraph.Edges)
}
