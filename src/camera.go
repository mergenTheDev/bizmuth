package bizmuth

import (
	"sync"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var defVertexShader = `
#version 330 core

layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 aTexCoord;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

out vec2 TexCoord;

void main() {
	gl_Position = projection * view * model * vec4(aPos, 1.0);
	TexCoord = aTexCoord;
}
`

var defFragmentShader = `
#version 330 core

out vec4 FragColor;
in vec2 TexCoord;

uniform sampler2D Texture;

void main() {
	FragColor = texture(Texture, TexCoord);
}`

var (
	shaderProgram  uint32
	initCameraOnce sync.Once
	viewLoc        int32
	projectionLoc  int32
)

type Camera struct {
	Position   Vector2
	projection mgl32.Mat4
}

func CreateCamera(x, y float32) *Camera {
	initCameraOnce.Do(func() {
		vertexShader := compileShader(defVertexShader, gl.VERTEX_SHADER)
		fragmentShader := compileShader(defFragmentShader, gl.FRAGMENT_SHADER)

		shaderProgram = gl.CreateProgram()

		gl.AttachShader(shaderProgram, vertexShader)
		gl.AttachShader(shaderProgram, fragmentShader)
		gl.LinkProgram(shaderProgram)

		projectionLoc = gl.GetUniformLocation(shaderProgram, gl.Str("projection\x00"))
		viewLoc = gl.GetUniformLocation(shaderProgram, gl.Str("view\x00"))

		gl.DeleteShader(vertexShader)
		gl.DeleteShader(fragmentShader)
	})

	return &Camera{
		Position:   Vector2{X: x, Y: y},
		projection: mgl32.Ortho2D(0, float32(windowWidth), float32(windowHeight), 0),
	}
}

func (cam *Camera) Update() {
	gl.UseProgram(shaderProgram)

	windowWidth, windowHeight = glfw.GetCurrentContext().GetSize()
	cam.projection = mgl32.Ortho2D(0, float32(windowWidth), float32(windowHeight), 0)
	gl.Viewport(0, 0, int32(windowWidth), int32(windowHeight))

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
