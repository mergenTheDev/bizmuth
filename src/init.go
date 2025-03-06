package bizmuth

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func Init() {
	err := glfw.Init()

	if err != nil {
		log.Fatal(PrefixErr + "Can't initialize GLFW!")
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.AlphaBits, 8)

	if err := gl.Init(); err != nil {
		log.Fatal(PrefixErr + "Can't initialize OpenGL!")
	}
}

func End() {
	glfw.Terminate()
}
