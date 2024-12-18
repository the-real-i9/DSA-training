package graph_al

import (
	"dsa/dataStructures/graph"
	"fmt"
)

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

/*
gv1v and gv2v - The respective values of the vertex pair to be connected by this new edge

newew - The weight of the edge to be added (optional, if unweighted)
*/
func (G *Graph) AddEdge(gv1v, gv2v graph.VertexValue, newew graph.EdgeWeight) error {
	var vertex1 *Vertex
	var vertex2 *Vertex

	// retrieve the corresponding vertices for the specified vertex values
	for _, gv := range G.Vertices {
		if gv1v.EqualTo(gv.Value) {
			vertex1 = gv
		} else if gv2v.EqualTo(gv.Value) {
			vertex2 = gv
		} else {
			continue
		}

		if vertex1 != nil && vertex2 != nil {
			break
		}
	}

	// check if vertices specified exists
	if vertex1 == nil || vertex2 == nil {
		return fmt.Errorf("unexisting vertex specified")
	}

	// check if this edge already exists
	for _, ge := range G.Edges {
		if G.Directed {
			if ge.AdjVert[vertex1] == vertex2 {
				return nil
			}
		} else {
			if ge.AdjVert[vertex1] == vertex2 || ge.AdjVert[vertex2] == vertex1 {
				return nil
			}
		}
	}

	newAdjVert := map[*Vertex]*Vertex{vertex1: vertex2}
	newEdge := &Edge{AdjVert: newAdjVert, Weight: newew}

	vertex1.IncEdges = append(vertex1.IncEdges, newEdge)
	vertex2.IncEdges = append(vertex2.IncEdges, newEdge)

	G.Edges = append(G.Edges, newEdge)

	return nil
}
