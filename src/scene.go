package bizmuth

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Scene struct{}

var CurrentScene *Scene

func CreateScene() *Scene {
	var scene Scene
	return &scene
}

// Swithes and Runs Scene
func SwithcScene(scene *Scene) {
	CurrentScene = scene
}

func (scene *Scene) Ready(callback func()) {
	if callback != nil || scene == CurrentScene {
		callback()
	}
}

func (scene *Scene) Update(callback func()) {
	go func() {
		for {
			gl.Clear(gl.COLOR_BUFFER_BIT)
			gl.UseProgram(shaderProgram)
			callback()
			/*if scene != CurrentScene {
				break
			}*/
			glfw.PollEvents()
		}
	}()
}
