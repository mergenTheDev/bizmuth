package bizmuth

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
