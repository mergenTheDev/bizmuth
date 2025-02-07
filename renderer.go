package bizmuth

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

/*func glslLoader(file string, shaderType uint32) uint32 {
	buffer, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(PrefixErr + "Can't read GLSL Shader")
	}

	shader := gl.CreateShader(shaderType)


}*/

func BackgroundColor(r float32, g float32, b float32, alpha float32) {
	gl.ClearColor(r, g, b, alpha)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
