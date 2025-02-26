package bizmuth

import (
	"sync"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var defVertexShader = `
#version 330 core

layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 aTexCoord;

uniform mat4 view;
uniform mat4 projection;

out vec2 TexCoord;

void main() {
	gl_Position = projection * view * vec4(aPos, 1.0);
	TexCoord = aTexCoord;
}
`

var emptyFragmentShader = `
#version 330 core

out vec4 FragColor;
in vec2 TexCoord;

uniform sampler2D ourTexture;

void main() {
	FragColor = texture(ourTexture, TexCoord);
}`

var (
	program        uint32
	initCameraOnce sync.Once
	viewLoc        int32
	projectionLoc  int32
)

type Camera struct {
	Position   Vector2
	projection mgl32.Mat4
}

func CreateCamera(x, y float32) (*Camera, uint32) {
	initCameraOnce.Do(func() {
		vertexShader := compileShader(defVertexShader, gl.VERTEX_SHADER)
		fragmentShader := compileShader(emptyFragmentShader, gl.FRAGMENT_SHADER)

		program = gl.CreateProgram()

		gl.AttachShader(program, vertexShader)
		gl.AttachShader(program, fragmentShader)
		gl.LinkProgram(program)

		projectionLoc = gl.GetUniformLocation(program, gl.Str("projection\x00"))
		viewLoc = gl.GetUniformLocation(program, gl.Str("view\x00"))

		gl.DeleteShader(vertexShader)
		gl.DeleteShader(fragmentShader)
	})

	return &Camera{
		Position:   Vector2{X: x, Y: y},
		projection: mgl32.Ortho2D(0, 800, 800, 0),
	}, program
}

func (cam *Camera) Update() {
	gl.UseProgram(program)

	gl.UniformMatrix4fv(projectionLoc, 1, false, &cam.projection[0])

	view := mgl32.Translate3D(-cam.Position.X, cam.Position.Y, 0)
	gl.UniformMatrix4fv(viewLoc, 1, false, &view[0])
}

func compileShader(source string, shaderType uint32) uint32 {
	shader := gl.CreateShader(shaderType)
	shaderSource, free := gl.Strs(source + "\x00")
	defer free()

	gl.ShaderSource(shader, 1, shaderSource, nil)
	gl.CompileShader(shader)

	return shader
}
