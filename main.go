package main

import "runtime"
import "github.com/go-gl/glfw/v3.3/glfw"
import "github.com/go-gl/gl/v3.3-core/gl"

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()

	if err != nil {
		panic(err)
	}

	defer glfw.Terminate()

	window, err := glfw.CreateWindow(800, 600, "Go OpenGL", nil, nil)

	if err != nil {
		panic(err)
	}

	if err := gl.Init(); err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	gl.ClearColor(0.5, 0.1, 0.6, 1.0)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}