package bizmuth

type SceneInterface interface {
	Ready()
	Update()
}

type Scene struct{}

var CurrentScene Scene

func (scene Scene) Ready(callback func()) {
	callback()

}

func (scene Scene) Update(callback func()) {
	go func() {
		for {
			callback()
			if scene != CurrentScene {
				break
			}
		}
	}()
}

// Swithes and Runs Scene
func SwithcScene(scene Scene) {
	CurrentScene = scene
}
