package bizmuth

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func CheckRenderError() {
	if err := gl.GetError(); err != gl.NO_ERROR {
		log.Fatalf(PrefixErr+"OpenGL Error: %v", err)
	}
}
