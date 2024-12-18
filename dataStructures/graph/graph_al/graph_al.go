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

// newvv - new vertex value || gv - existing graph vertex
func (G *Graph) AddVertex(newvv graph.VertexValue) {
	for _, gv := range G.Vertices {
		if newvv.EqualTo(gv.Value) {
			return
		}
	}

	newVertex := &Vertex{Value: newvv}

	G.Vertices = append(G.Vertices, newVertex)
}
