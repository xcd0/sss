package main

import (
	"context"
	"machine"
	"machine/usb"
	"time"

	//keyboard "github.com/sago35/tinygo-keyboard"
	//"github.com/sago35/tinygo-keyboard/keycodes/jp"
	keyboard "github.com/xcd0/tinygo-keyboard"
	"github.com/xcd0/tinygo-keyboard/keycodes/jp"
)

var (
	gpioPin []machine.Pin
)

func init() {
	usb.Product = "test-0.0.1"

	// 全てのGPIOをスライスとしてアクセスできるようにしておく。
	gpioPin = []machine.Pin{
		machine.GPIO0, machine.GPIO1, machine.GPIO2, machine.GPIO3, machine.GPIO4, machine.GPIO5, machine.GPIO6, machine.GPIO7, machine.GPIO8, machine.GPIO9,
		machine.GPIO10, machine.GPIO11, machine.GPIO12, machine.GPIO13, machine.GPIO14, machine.GPIO15, machine.GPIO16, machine.GPIO17, machine.GPIO18, machine.GPIO19,
		machine.GPIO20, machine.GPIO21, machine.GPIO22, machine.GPIO23, machine.GPIO24, machine.GPIO25, machine.GPIO26, machine.GPIO27, machine.GPIO28, machine.GPIO29,
	}
}

func main() {
	d := keyboard.New()

	col := gpioPin[0:2]
	row := gpioPin[2:4]

	d.AddMatrixKeyboard(
		col, row, [][]keyboard.Keycode{
			{
				jp.KeyA, jp.KeyB,
				jp.KeyC, jp.KeyD,
			},
		},
		//keyboard.MatrixScanPeriod(750*time.Microsecond),
		keyboard.MatrixScanPeriod(760*time.Nanosecond),
	)

	d.Loop(context.Background())
}
