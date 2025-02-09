package bizmuth

import "github.com/go-gl/gl/v3.3-core/gl"

func BackgroundColor(r float32, g float32, b float32, alpha float32) {
	gl.ClearColor(r, g, b, alpha)
}
