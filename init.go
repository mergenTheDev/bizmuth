package bizmuth

import (
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func Init() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(prefixErr + "Can't initialize GLFW!" + reset)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}
