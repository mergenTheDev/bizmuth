package bizmuth

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func CreateWindow(width int, height int, title string, resizable int) *glfw.Window {
	glfw.WindowHint(glfw.Resizable, resizable)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)

	if err != nil {
		panic(PrefixErr + "Can't create window!")
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(PrefixErr + "Can't initialize OpenGL!")
	}

	gl.Viewport(0, 0, int32(width), int32(height))

	return window
}
