package bizmuth

import (
	"log"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func LoadShader(file string, shaderType uint32) uint32 {
	buffer, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(PrefixErr + " Can't read file!")
	}

	var str byte
	makeStr := string(append(buffer, str))
	makeStr = string(makeStr) + "\x00"

	shader := gl.CreateShader(shaderType)
	source, free := gl.Strs(makeStr)

	gl.ShaderSource(shader, 1, source, nil)
	free()
	gl.CompileShader(shader)

	return shader
}
