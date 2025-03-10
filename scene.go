package bizmuth

import (
	"fmt"

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
	fCount := 0
	pTime := glfw.GetTime()

	for !CurrentContext.ShouldClose() && CurrentScene == scene {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(shaderProgram)
		callback()

		cTime := glfw.GetTime()
		fCount++

		if cTime-pTime >= 1 {
			fmt.Println("FPS:", fCount)
			fCount = 0
			pTime = cTime
		}

		CurrentContext.SwapBuffers()
		glfw.PollEvents()
	}
}
