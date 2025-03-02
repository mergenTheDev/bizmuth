package bizmuth

type SceneInterface interface {
	Ready()
}

type Scene struct{}

func (def Scene) Ready(callback func()) {
	callback()
}



// Swithes and Runs Scene
func SwithcScene(scene Scene) {
	
}
