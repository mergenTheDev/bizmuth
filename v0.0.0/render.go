package bizmuth

import "github.com/go-gl/gl/v3.3-core/gl"

type Args struct {
	pos     Vector2
	texture uint32
	scale   float32
}

func BackgroundColor(r float32, g float32, b float32, alpha float32) {
	gl.ClearColor(r, g, b, alpha)
}

func Draw(args Args) {
	//TO-DO
	/*vertices := []float32{}

	indices := []float32{}*/
}
