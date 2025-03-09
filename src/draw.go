package bizmuth

import (
	"sync"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var genBuffersOnce sync.Once
var vao, vbo, ebo uint32

type DrawArgs struct {
	Position       Vector2
	Texture        uint32
	Scale          float32
	Text           string
	Font           int32
	Physics        int32
	Body           int32
	CollisionShape int32
}

func BackgroundColor(r float32, g float32, b float32, alpha float32) {
	gl.ClearColor(r, g, b, alpha)
}

func Draw(args DrawArgs) {
	gl.UseProgram(shaderProgram)

	modelLoc := gl.GetUniformLocation(shaderProgram, gl.Str("model\x00"))
	model := mgl32.Translate3D(args.Position.X, args.Position.Y, 0).Mul4(mgl32.Scale3D(args.Scale, args.Scale, 1))

	genBuffersOnce.Do(func() {
		vertices := []float32{
			50, 50, 0, 1, 1,
			50, -50, 0, 1, 0,
			-50, -50, 0, 0, 0,
			-50, 50, 0, 0, 1,
		}

		indices := []uint32{
			0, 1, 3,
			1, 2, 3,
		}

		gl.GenVertexArrays(1, &vao)
		gl.GenBuffers(1, &vbo)
		gl.GenBuffers(1, &ebo)

		gl.BindVertexArray(vao)

		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
		gl.EnableVertexAttribArray(0)

		gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
		gl.EnableVertexAttribArray(1)
	})

	gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])

	if args.Texture != 0 {
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, args.Texture)
		gl.Uniform1i(gl.GetUniformLocation(shaderProgram, gl.Str("Texture\x00")), 0)
	}

	gl.BindVertexArray(vao)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.BindVertexArray(0)
}
