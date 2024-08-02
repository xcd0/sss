package main

import (
	"context"
	"machine"
	"machine/usb"

	// keyboard "github.com/sago35/tinygo-keyboard"
	keyboard "github.com/xcd0/tinygo-keyboard"
	"github.com/xcd0/tinygo-keyboard/keycodes/jp"
)

var (
	gpio []machine.Pin
)

func init() {
	usb.Product = "test-0.0.2"

	// 全てのGPIOをスライスとしてアクセスできるようにしておく。
	gpio = []machine.Pin{
		machine.GPIO0, machine.GPIO1, machine.GPIO2, machine.GPIO3, machine.GPIO4, machine.GPIO5, machine.GPIO6, machine.GPIO7, machine.GPIO8, machine.GPIO9,
		machine.GPIO10, machine.GPIO11, machine.GPIO12, machine.GPIO13, machine.GPIO14, machine.GPIO15, machine.GPIO16, machine.GPIO17, machine.GPIO18, machine.GPIO19,
		machine.GPIO20, machine.GPIO21, machine.GPIO22, machine.GPIO23, machine.GPIO24, machine.GPIO25, machine.GPIO26, machine.GPIO27, machine.GPIO28, machine.GPIO29,
	}
}

func main() {
	d := keyboard.New()

	col := []machine.Pin{gpio[0], gpio[1]}
	row := []machine.Pin{gpio[3], gpio[4]}

	d.AddMatrixKeyboard(col, row, [][]keyboard.Keycode{
		{
			jp.KeyA, jp.KeyB,
		},
		{
			jp.KeyC, jp.KeyD,
		},
	}, keyboard.InvertDiode(false))

	d.Loop(context.Background())
}
