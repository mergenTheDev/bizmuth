package bizmuth

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Config struct {
	Debug bool
	VSync bool
}

var engine Config

func init() {
	runtime.LockOSThread()
	runtime.GOMAXPROCS(2)
}

func Init(config Config) {
	if err := glfw.Init(); err != nil {
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

	engine = config
}

func End() {
	glfw.Terminate()
}
