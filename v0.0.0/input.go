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
	KeyD     = 68
	KeyE     = 69
	KeyF     = 70
	KeyG     = 71
	KeyH     = 72
	KeyI     = 73
	KeyJ     = 74
	KeyK     = 75
	KeyL     = 76
	KeyM     = 77
	KeyN     = 78
	KeyO     = 79
	KeyP     = 80
	KeyQ     = 81
	KeyR     = 82
	KeyS     = 83
	KeyT     = 84
	KeyU     = 85
	KeyV     = 86
	KeyW     = 87
	KeyX     = 88
	KeyY     = 89
	KeyZ     = 90
	KeySuper = 91
)

//Should call inside of bizmuth.GameLoop()

func (window *Window) OnKeyPress(key glfw.Key, callback func()) {
	if window.GetKey(key) == glfw.Press {
		callback()
	}
}

func (window *Window) OnKeyRelease(key glfw.Key, callback func()) {
	if window.GetKey(key) == glfw.Release {
		callback()
	}
}
