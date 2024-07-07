// Code generated generated_keymap.go DO NOT EDIT.

package main

import (
	//keyboard "github.com/sago35/tinygo-keyboard"
	//"github.com/sago35/tinygo-keyboard/keycodes/jp"
	keyboard "github.com/xcd0/tinygo-keyboard"
	"github.com/xcd0/tinygo-keyboard/keycodes/jp"
)

func GetKeycodes() [][]keyboard.Keycode {
	return [][]keyboard.Keycode{
		{
			jp.KeyQ, jp.KeyW, jp.KeyE, jp.KeyR, jp.KeyT, jp.KeyTab, jp.KeyAt, jp.KeyY, jp.KeyU, jp.KeyI, jp.KeyO, jp.KeyP,
			jp.KeyA, jp.KeyS, jp.KeyD, jp.KeyF, jp.KeyG, jp.KeyMinus, jp.KeyColon, jp.KeyH, jp.KeyJ, jp.KeyK, jp.KeyL, jp.KeySemicolon,
			jp.KeyZ, jp.KeyX, jp.KeyC, jp.KeyV, jp.KeyB, jp.KeyBackslash | jp.KeyLeftShift, jp.KeyMod4, jp.KeyMod4, jp.KeyBackslash2, jp.KeyN, jp.KeyM, jp.KeyComma, jp.KeyPeriod, jp.KeySlash,
			jp.KeyMuhenkan, jp.KeyMod1, jp.KeySpace, jp.KeyLeftShift, jp.KeyEsc, jp.KeyDelete, jp.KeyBackspace, jp.KeyEnter, jp.KeyHenkan, 0,
		},
		{
			jp.Key1, jp.Key2, jp.Key3, jp.Key4, jp.Key5, jp.KeyTab, jp.KeyAt, jp.Key6, jp.Key7, jp.Key8, jp.Key9, jp.Key0,
			jp.Key1 | jp.KeyLeftShift, jp.Key2 | jp.KeyLeftShift, jp.Key3 | jp.KeyLeftShift, jp.Key4 | jp.KeyLeftShift, jp.Key5 | jp.KeyLeftShift, jp.KeyMinus, jp.KeyColon, jp.Key6 | jp.KeyLeftShift, jp.Key7 | jp.KeyLeftShift, jp.Key8 | jp.KeyLeftShift, jp.Key9 | jp.KeyLeftShift, 0,
			jp.KeyLeftBrace, jp.KeyRightBrace, jp.KeyInsert, jp.KeyInsert | jp.KeyLeftCtrl, jp.KeyInsert | jp.KeyLeftShift, jp.KeyBackslash | jp.KeyLeftShift, 0, 0, jp.KeyBackslash2, jp.KeyLeftBrace, jp.KeyRightBrace, jp.KeyComma, jp.KeyPeriod, jp.KeySlash,
			jp.KeyMuhenkan, jp.KeyMod1, jp.KeySpace, jp.KeyLeftShift, jp.KeyEsc, jp.KeyDelete, jp.KeyBackspace, jp.KeyEnter, jp.KeyHenkan, 0,
		},
		{
			0, 0, 0, 0, 0, 0, jp.KeyAt, jp.Key6, jp.Key7, jp.Key8, jp.Key9, jp.Key0,
			jp.KeyF1, jp.KeyF2, jp.KeyF3, jp.KeyF4, jp.KeyF5, jp.KeyF6, jp.KeyHome, jp.KeyRight, jp.KeyDown, jp.KeyUp, jp.KeyLeft, jp.KeyEnd,
			jp.KeyF7, jp.KeyF8, jp.KeyF9, jp.KeyF10, jp.KeyF11, jp.KeyF12, 0, 0, jp.MouseLeft, jp.MouseRight, jp.MouseMiddle, jp.WheelUp, jp.WheelDown, 0,
			jp.KeyMuhenkan, jp.KeyMod1, jp.KeySpace, jp.KeyLeftShift, jp.KeyEsc, jp.KeyDelete, jp.KeyBackspace, jp.KeyEnter, jp.KeyHenkan, 0,
		},
		{
			0, jp.KeyV | jp.KeyWindows, jp.KeyPeriod | jp.KeyWindows, jp.KeyLeftShift, jp.KeyWindows | jp.KeyLeftShift, 0, 0, jp.KeySlash, jp.Key7, jp.Key8, jp.Key9, jp.KeyColon | jp.KeyLeftShift,
			0, jp.KeyBackspace, jp.KeyDelete, 0, 0, 0, 0, jp.KeyPeriod, jp.Key4, jp.Key5, jp.Key6, jp.KeyMinus,
			0, 0, 0, 0, 0, 0, 0, 0, 0, jp.Key0, jp.Key1, jp.Key2, jp.Key3, jp.KeySemicolon | jp.KeyLeftShift,
			jp.KeyMuhenkan, jp.KeyMod1, jp.KeySpace, jp.KeyLeftShift, jp.KeyEsc, jp.KeyDelete, jp.KeyBackspace, jp.KeyEnter, jp.KeyHenkan, 0,
		},
	}
}
