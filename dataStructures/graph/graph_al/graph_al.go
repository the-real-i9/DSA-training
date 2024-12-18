// Graph: Adjacency List Representation
package graph_al

import (
	"dsa/dataStructures/graph"
)

type Vertex struct {
	Value    graph.VertexValue `json:"value,omitempty"`
	IncEdges []*Edge           `json:"incidentEdges,omitempty"` // Incident edges
}

type Edge struct {
	AdjVert map[*Vertex]*Vertex `json:"adjVertexPair,omitempty"` // Adjacent vertex pair
	Weight  graph.EdgeWeight    `json:"weight,omitempty"`        // Weight of edge
}

type Graph struct {
	Edges    []*Edge   `json:"edges,omitempty"`
	Vertices []*Vertex `json:"vertices,omitempty"`
}
