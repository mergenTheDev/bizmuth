package bizmuth

import (
	//"fmt"
	"sync"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Scene struct {
	stopChan chan struct{}
	wg       sync.WaitGroup
	mu       sync.Mutex
}

var CurrentScene *Scene

func CreateScene() *Scene {
	return &Scene{
		stopChan: make(chan struct{}),
	}
}

func SwithcScene(scene *Scene) {
	if CurrentScene != nil {
		CurrentScene.stop()
	}
	CurrentScene = scene
}

func (scene *Scene) stop() {
	close(scene.stopChan)
	scene.wg.Wait()
}

func (scene *Scene) Ready(callback func()) {
	if callback != nil {
		callback()
	}
}

func (scene *Scene) Update(callback func()) {
	scene.wg.Add(1)
	go func() {
		defer scene.wg.Done()
		for {
			select {
			case <-scene.stopChan:
				return
			default:
				scene.mu.Lock()
				gl.Clear(gl.COLOR_BUFFER_BIT)
				gl.UseProgram(shaderProgram)
				callback()
				scene.mu.Unlock()
				//glfw.GetCurrentContext().SwapBuffers()
				glfw.PollEvents()
			}
		}
	}()
}
