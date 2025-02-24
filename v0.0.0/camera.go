package bizmuth

import "github.com/go-gl/gl/v3.3-core/gl"

var defVertexShader = `
#version 330 core

layout (location = 0) in vec3 aPos;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main() {
	gl_Position = projection * view * model * vec4(aPos, 1.0);
}
` + "\x00"

var emptyFragmentShader = `
#version 330 core
out vec4 FragColor;
void main() {
	FragColor = vec4(1.0);
}` + "\x00"

type Camera struct {
	Position Vector2
	Program  uint32
}

func CreateCamera(x, y float32) *Camera {
	program := cameraProgram(defVertexShader, emptyFragmentShader)

	return &Camera{
		Position: Vector2{x, y},
		Program:  program,
	}
}

func (cam *Camera) UpdateCamera() {
	gl.UseProgram(cam.Program)

	//TO-DO
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
