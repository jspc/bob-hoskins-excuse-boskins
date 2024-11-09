package main

import "machine/usb/hid/keyboard"

var keysToPresses = map[rune]keyboard.Keycode{
	'a': keyboard.KeyA,
	'b': keyboard.KeyB,
	'c': keyboard.KeyC,
	'd': keyboard.KeyD,
	'e': keyboard.KeyE,
	'f': keyboard.KeyF,
	'g': keyboard.KeyG,
	'h': keyboard.KeyH,
	'i': keyboard.KeyI,
	'j': keyboard.KeyJ,
	'k': keyboard.KeyK,
	'l': keyboard.KeyL,
	'm': keyboard.KeyM,
	'n': keyboard.KeyN,
	'o': keyboard.KeyO,
	'p': keyboard.KeyP,
	'q': keyboard.KeyQ,
	'r': keyboard.KeyR,
	's': keyboard.KeyS,
	't': keyboard.KeyT,
	'u': keyboard.KeyU,
	'v': keyboard.KeyV,
	'w': keyboard.KeyW,
	'x': keyboard.KeyX,
	'y': keyboard.KeyY,
	'z': keyboard.KeyZ,
}
