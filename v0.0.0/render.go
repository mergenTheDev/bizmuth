package bizmuth

import "github.com/go-gl/gl/v3.3-core/gl"

type DrawArgs struct {
	Pos            Vector2
	Texture        uint32
	Scale          float32
	Text           string
	Font           int32
	Physics        int32
	Body           int32
	CollisionShape int32
}

func BackgroundColor(r float32, g float32, b float32, alpha float32) {
	gl.ClearColor(r, g, b, alpha)
}

func Draw(args DrawArgs) DrawArgs {
	//TO-DO
	/*vertices := []float32{}

	indices := []float32{}*/
	return DrawArgs{}
}
