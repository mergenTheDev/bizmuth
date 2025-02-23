package bizmuth

import "github.com/go-gl/glfw/v3.3/glfw"

const (
	KeySpace = 32
	Key0     = 48
	Key1     = 49
	Key2     = 50
	Key3     = 51
	Key4     = 52
	Key5     = 53
	Key6     = 54
	Key7     = 55
	Key8     = 56
	Key9     = 57
	KeyA     = 65
	KeyB     = 66
	KeyC     = 67
	KeySuper = 91
)

func (window *Window) OnKeyPress(key glfw.Key, callback func()) {
	if window.GetKey(key) == glfw.Press {
		callback()
	}
}
