package bizmuth

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var defVertexShader = `
#version 330 core

layout (location = 0) in vec3 aPos;

uniform mat4 view;
uniform mat4 projection;

void main() {
	gl_Position = projection * model * vec4(aPos, 1.0);
}
` + "\x00"

var emptyFragmentShader = `
#version 330 core
out vec4 FragColor;
void main() {
	FragColor = vec4(1.0);
}` + "\x00"

type Camera struct {
	Position          Vector2
	Program           uint32
	viewUniform       int32
	projectionUniform int32
	projection        mgl32.Mat4
}

func CreateCamera(x, y float32) *Camera {
	program := cameraProgram(defVertexShader, emptyFragmentShader)

	/*projectionUniform = gl.GetUniformLocation(program, gl.Str("projection\x00"))
	viewUniform = gl.GetUniformLocation(program, gl.Str("view\x00"))

	projection = mgl32.Ortho2D(0, 800, 800, 0)*/

	return &Camera{
		Position:          Vector2{X: x, Y: y},
		Program:           program,
		projectionUniform: gl.GetUniformLocation(program, gl.Str("projection\x00")),
		viewUniform:       gl.GetUniformLocation(program, gl.Str("view\x00")),
		projection:        mgl32.Ortho2D(0, 800, 800, 0),
	}
}

func (cam *Camera) UpdateCamera() {
	gl.UseProgram(cam.Program)

	gl.UniformMatrix4fv(cam.projectionUniform, 1, false, &cam.projection[0])

	view := mgl32.Translate3D(-cam.Position.X, -cam.Position.Y, 0)
	gl.UniformMatrix4fv(cam.viewUniform, 1, false, &view[0])

}

func cameraProgram(vertex string, fragment string) uint32 {
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)

	vertexSource, freeV := gl.Strs(vertex)
	fragmentSource, freeF := gl.Strs(fragment)

	gl.ShaderSource(vertexShader, 1, vertexSource, nil)
	gl.ShaderSource(fragmentShader, 1, fragmentSource, nil)

	freeV()
	freeF()

	gl.CompileShader(vertexShader)
	gl.CompileShader(fragmentShader)

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	gl.DetachShader(program, vertexShader)
	gl.DetachShader(program, fragmentShader)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program
}
