package bizmuth

type Vector2 struct {
	X, Y float32
}

type Vector3 struct {
	X, Y, Z float32
}

type Vector4 struct {
	X, Y, Z, W float32
}

func (v Vector2) ToVec3() Vector3 {
	return Vector3{v.X, v.Y, 0.0}
}
