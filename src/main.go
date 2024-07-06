package main

import (
	"context"
	"machine"
	"machine/usb"

	keyboard "github.com/sago35/tinygo-keyboard"
)

//go:generate go run ./script/prebuild/prebuild.go ./keymap.hjson
//go:generate go run ./script/gen-def/main.go ./vial.json

var (
	gpioPin []machine.Pin
)

func init() {
	usb.Product = "sss46g-0.0.2"
	// 全てのGPIOにjスライスとしてアクセスできるようにしておく。
	gpioPin = []machine.Pin{
		machine.GPIO0, machine.GPIO1, machine.GPIO2, machine.GPIO3, machine.GPIO4, machine.GPIO5, machine.GPIO6, machine.GPIO7, machine.GPIO8, machine.GPIO9,
		machine.GPIO10, machine.GPIO11, machine.GPIO12, machine.GPIO13, machine.GPIO14, machine.GPIO15, machine.GPIO16, machine.GPIO17, machine.GPIO18, machine.GPIO19,
		machine.GPIO20, machine.GPIO21, machine.GPIO22, machine.GPIO23, machine.GPIO24, machine.GPIO25, machine.GPIO26, machine.GPIO27, machine.GPIO28, machine.GPIO29,
	}
	// すべてのGPIOをプルアップしておく。
	for i, _ := range gpioPin {
		gpioPin[i].Configure(
			machine.PinConfig{
				Mode: machine.PinInputPullup,
			},
		)
	}
}

func main() {
	d := keyboard.New()
	row := gpioPin[0:4]                      // 0,1,2,3
	col := append(gpioPin[4:9], gpioPin[17]) // 4,5,6,7,8,17

	d.AddMatrixKeyboard(col, row, GetKeycodes())
	loadKeyboardDef()
	d.Loop(context.Background())
}

/*
var (
	gpioPin []machine.Pin
	i2cPin  []machine.Pin
	rowPin  []machine.Pin
	colPin  []machine.Pin
)

func run() {
	SetPin()

	var d *keyboard.Device
	d = keyboard.New()

	//d.AddGpioKeyboard(gpioPin, GetKeycodes())
	//d.AddMatrixKeyboard(colPin, rowPin , GetKeycodes(), opt ...Option)
	d.AddMatrixKeyboard(colPin, rowPin, GetKeycodes())

	d.Loop(context.Background())
}

func SetPin() {
	gpioPin = []machine.Pin{
		machine.GPIO0, machine.GPIO1, machine.GPIO2, machine.GPIO3, machine.GPIO4, machine.GPIO5, machine.GPIO6, machine.GPIO7, machine.GPIO8, machine.GPIO9,
		machine.GPIO10, machine.GPIO11, machine.GPIO12, machine.GPIO13, machine.GPIO14, machine.GPIO15, machine.GPIO16, machine.GPIO17, machine.GPIO18, machine.GPIO19,
		machine.GPIO20, machine.GPIO21, machine.GPIO22, machine.GPIO23, machine.GPIO24, machine.GPIO25, machine.GPIO26, machine.GPIO27, machine.GPIO28, machine.GPIO29,
	}
	for c := range gpioPin {
		gpioPin[c].Configure(
			machine.PinConfig{
				//Mode: machine.PinInput,
				//Mode: machine.PinInputPullup,
				Mode: machine.PinInputPulldown,
			},
		)
	}

	i2cPin = []machine.Pin{
		// SPI0_SCK_PIN	= GPIO18 // Default Serial Clock Bus 0 for SPI communications
		// SPI0_SDO_PIN	= GPIO19	// Tx Default Serial Out Bus 0 for SPI communications
		// SPI0_SDI_PIN	= GPIO16	// Rx Default Serial In Bus 0 for SPI communications
		// SPI1_SCK_PIN	= GPIO10 // Default Serial Clock Bus 1 for SPI communications
		// SPI1_SDO_PIN	= GPIO11	// Tx Default Serial Out Bus 1 for SPI communications
		// SPI1_SDI_PIN	= GPIO12	// Rx Default Serial In Bus 1 for SPI communications
	}

	rowPin = []machine.Pin{
		gpioPin[0], // row0
		gpioPin[1], // row1
		gpioPin[2], // row2
		gpioPin[3], // row3
	}
	colPin = []machine.Pin{
		gpioPin[4],  // col0
		gpioPin[5],  // col1
		gpioPin[6],  // col2
		gpioPin[7],  // col3
		gpioPin[8],  // col4
		gpioPin[17], // col5
	}
}
*/
