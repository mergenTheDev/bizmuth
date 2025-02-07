package bizmuth

import "github.com/go-gl/glfw/v3.3/glfw"

func GameLoop(window *glfw.Window) {
	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
