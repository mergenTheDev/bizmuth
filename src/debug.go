package bizmuth

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func CheckGLError() {
	if err := gl.GetError(); err != gl.NO_ERROR {
		log.Fatalf("OpenGL error: %v", err)
	}
}
