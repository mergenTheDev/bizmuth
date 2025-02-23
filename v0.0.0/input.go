package bizmuth

import "github.com/go-gl/glfw/v3.3/glfw"

type Key int

const (
	Key0     Key = 48
	Key1     Key = 49
	Key2     Key = 50
	Key3     Key = 51
	Key4     Key = 52
	Key5     Key = 53
	Key6     Key = 54
	Key7     Key = 55
	Key8     Key = 56
	Key9     Key = 57
	KeyA     Key = 65
	KeyB     Key = 66
	KeyC     Key = 67
	KeySuper Key = 91
)

func (window *Window) OnKeyPress(key glfw.Key, callback func()) {
	if window.GetKey(key) == glfw.Press {
		callback()
	}
}
