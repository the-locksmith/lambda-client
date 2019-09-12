package primitive

import (
	"github.com/galaco/lambda-client/core/mesh"
)

var cubeVerts = []float32{
	-1.0, -1.0, -1.0,
	1.0, -1.0, -1.0,
	-1.0, -1.0, 1.0,
	1.0, -1.0, -1.0,
	1.0, -1.0, 1.0,
	-1.0, -1.0, 1.0,
	-1.0, 1.0, -1.0,
	-1.0, 1.0, 1.0,
	1.0, 1.0, -1.0,
	1.0, 1.0, -1.0,
	-1.0, 1.0, 1.0,
	1.0, 1.0, 1.0,
	-1.0, -1.0, 1.0,
	1.0, -1.0, 1.0,
	-1.0, 1.0, 1.0,
	1.0, -1.0, 1.0,
	1.0, 1.0, 1.0,
	-1.0, 1.0, 1.0,
	-1.0, -1.0, -1.0,
	-1.0, 1.0, -1.0,
	1.0, -1.0, -1.0,
	1.0, -1.0, -1.0,
	-1.0, 1.0, -1.0,
	1.0, 1.0, -1.0,
	-1.0, -1.0, 1.0,
	-1.0, 1.0, -1.0,
	-1.0, -1.0, -1.0,
	-1.0, -1.0, 1.0,
	-1.0, 1.0, 1.0,
	-1.0, 1.0, -1.0,
	1.0, -1.0, 1.0,
	1.0, -1.0, -1.0,
	1.0, 1.0, -1.0,
	1.0, -1.0, 1.0,
	1.0, 1.0, -1.0,
	1.0, 1.0, 1.0,
}

var cubeNormals = cubeVerts

var cubeUVs = []float32{
	0, 0,
	1, 0,
	0, 1,
	1, 0,
	0, 1,
	1, 1,

	0, 0,
	1, 0,
	0, 1,
	1, 0,
	0, 1,
	1, 1,

	0, 0,
	1, 0,
	0, 1,
	1, 0,
	0, 1,
	1, 1,

	0, 0,
	1, 0,
	0, 1,
	1, 0,
	0, 1,
	1, 1,
	0, 0,
	1, 0,
	0, 1,
	1, 0,
	0, 1,
	1, 1,
	0, 0,
	1, 0,
	0, 1,
	1, 0,
	0, 1,
	1, 1,
}

// Cube
type Cube struct {
	mesh.Mesh
}

// NewCube
func NewCube() *Cube {
	c := &Cube{
		Mesh: *mesh.NewMesh(),
	}
	c.AddVertex(cubeVerts...)
	c.AddNormal(cubeNormals...)
	c.AddUV(cubeUVs...)
	//c.Finish()

	return c
}