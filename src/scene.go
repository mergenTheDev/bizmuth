package bizmuth

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
	if callback != nil {
		callback()
	}
}

func (scene *Scene) Update(callback func()) {
	go func() {
		for {
			callback()
			if scene != CurrentScene {
				break
			}
		}
	}()
}
