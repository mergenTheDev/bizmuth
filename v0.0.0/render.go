package bizmuth

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

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
	modelLoc := gl.GetUniformLocation(shaderProgram, gl.Str("model\x00"))
	model := mgl32.Translate3D(args.Pos.X, args.Pos.Y, 0)
	gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])

	//vertices := []float32{}

	//indices := []float32{}

	return DrawArgs{}
}
