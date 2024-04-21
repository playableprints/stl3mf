package models

import (
	"fmt"

	"github.com/hschendel/stl"
)

func CheckSTL(filepath string) error {
	solid, err := stl.ReadFile(filepath)
	if err != nil {
		return err
	}

	/*
	   type TriangleErrors struct {
	   	// HasEqualVertices is true if some vertices are identical, meaning we are having
	   	// a line, or even a point, as opposed to a triangle.
	   	HasEqualVertices bool

	   	// NormalDoesNotMatch istrue if the normal vector does not match a normal calculated from the
	   	// vertices in the right hand order, even allowing for an angular difference
	   	// of < 90 degree.
	   	NormalDoesNotMatch bool

	   	// EdgeErrors by edge. The edge is indexed by it's first vertex, i.e.
	   	//    0: V0 -> V1
	   	//    1: V1 -> V2
	   	//    2: V2 -> V0
	   	// If the edge has no error its value is nil.
	   	EdgeErrors [3]*EdgeError
	   }

	   type EdgeError struct {
	   	// SameEdgeTriangles are indexes in Solid.Triangles of triangles that contain exactly the same edge.
	   	SameEdgeTriangles []int

	   	// CounterEdgeTriangles are indexes in Solid.Triangles of triangles that contain the edge in the
	   	// opposite direction. If there is exactly one other triangle, this is no
	   	// error.
	   	CounterEdgeTriangles []int
	   }

	*/

	trierrors := solid.Validate()
	verticeerrcount := 0
	normalerrcount := 0
	sameedges := 0
	counteredges := 0

	for _, v := range trierrors {
		if v.HasEqualVertices {
			verticeerrcount++
		}

		if v.NormalDoesNotMatch {
			normalerrcount++
		}

		for _, e := range v.EdgeErrors {
			if e != nil {
				if len(e.SameEdgeTriangles) > 0 {
					sameedges++
				}
				if len(e.CounterEdgeTriangles) != 1 {
					counteredges++
				}
			}
		}
	}

	if len(trierrors) > 0 {

		fmt.Printf(`Found %v triangles that aren't right.
		%v are not triangles.
		%v have inverted normals.
		%v duplicated edges.
		%v edges that incorrectly run counter.
`,
			len(trierrors),
			verticeerrcount,
			normalerrcount,
			sameedges,
			counteredges)

		return fmt.Errorf("invalid STL")
	}

	fmt.Print("No errors found\n")
	return nil
}
