package bizmuth

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Scene struct{}

var CurrentScene *Scene

func CreateScene() *Scene {
	return &Scene{}
}

func SwithcScene(scene *Scene) {
	CurrentScene = scene
}

func (scene *Scene) Ready(callback func()) {
	callback()
}

// Optimize in future.
func (scene *Scene) Update(callback func()) {
	for !CurrentContext.ShouldClose() && CurrentScene == scene {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(shaderProgram)
		callback()
		CurrentContext.SwapBuffers()
		glfw.PollEvents()
	}
}
