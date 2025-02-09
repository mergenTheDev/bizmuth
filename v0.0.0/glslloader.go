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

	// Checks for errors I don't need anymore
	/*var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		shaderLog := make([]byte, logLength+1)
		gl.GetShaderInfoLog(shader, logLength, nil, (*uint8)(unsafe.Pointer(&shaderLog[0])))

		log.Fatalf(PrefixErr+" Failed to compile %v: %v", file, string(shaderLog))
	}*/

	return shader
}
