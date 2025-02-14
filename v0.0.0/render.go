package bizmuth

import "github.com/go-gl/gl/v3.3-core/gl"

/*var defVertexShader = `
#version 330 core

layout (location = 0) in vec3 aPos;
uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main() {
	gl_Position = projection * view * model * vec4(aPos, 1.0);
}
` + "\x00"*/

func BackgroundColor(r float32, g float32, b float32, alpha float32) {
	gl.ClearColor(r, g, b, alpha)
}

func Add() {

}
