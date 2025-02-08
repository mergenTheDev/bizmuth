package bizmuth

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func Init() {
	runtime.LockOSThread()

	err := glfw.Init()

	if err != nil {
		log.Fatal(PrefixErr + "Can't initialize GLFW!")
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}

func End() {
	glfw.Terminate()
}
