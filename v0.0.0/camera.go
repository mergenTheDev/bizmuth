package bizmuth

import "github.com/go-gl/mathgl/mgl32"

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

func CreateCamera(pos Vector2) {
	projection := mgl32.Ortho2D(0, 800, 0, 800)
	//view :=
}
