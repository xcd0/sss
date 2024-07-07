package main

import (
	"context"
	"fmt"
	"machine"
	"machine/usb"
	"time"

	keyboard "github.com/xcd0/tinygo-keyboard"
)

//go:generate go run ./script/prebuild/prebuild.go ./keymap.hjson
//go:generate gofmt -w ./generated_keymap.go
//go:generate gofmt  ./generated_keymap.go
//go:generate go run ./script/gen-def/main.go ./vial.json

var (
	gpio []machine.Pin
	col  []machine.Pin
	row  []machine.Pin
	i2c  []machine.Pin
)

func main() {
	usb.Product = "sss46g-0.0.4"
	fmt.Printf("Product Name: %v", usb.Product)
	SetPin()
	d := keyboard.New()
	d.AddMatrixKeyboard(
		col, row, GetKeycodes(),
		keyboard.MatrixScanPeriod(760*time.Nanosecond),
	)
	//loadKeyboardDef()
	d.Debug = false
	d.Loop(context.Background())
}

func SetPin() {
	gpio = []machine.Pin{
		machine.GPIO0, machine.GPIO1, machine.GPIO2, machine.GPIO3, machine.GPIO4, machine.GPIO5, machine.GPIO6, machine.GPIO7, machine.GPIO8, machine.GPIO9,
		machine.GPIO10, machine.GPIO11, machine.GPIO12, machine.GPIO13, machine.GPIO14, machine.GPIO15, machine.GPIO16, machine.GPIO17, machine.GPIO18, machine.GPIO19,
		machine.GPIO20, machine.GPIO21, machine.GPIO22, machine.GPIO23, machine.GPIO24, machine.GPIO25, machine.GPIO26, machine.GPIO27, machine.GPIO28, machine.GPIO29,
	}
	for c := range gpio {
		gpio[c].Configure(
			machine.PinConfig{
				//Mode: machine.PinInput,
				//Mode: machine.PinInputPulldown,
				Mode: machine.PinInputPullup,
			},
		)
	}
	row = []machine.Pin{
		gpio[0], // row0
		gpio[1], // row1
		gpio[2], // row2
		gpio[3], // row3
	}
	col = []machine.Pin{
		gpio[4],  // col0
		gpio[5],  // col1
		gpio[6],  // col2
		gpio[7],  // col3
		gpio[8],  // col4
		gpio[17], // col5
	}
	i2c = []machine.Pin{
		// SPI0_SCK_PIN	= GPIO18 // Default Serial Clock Bus 0 for SPI communications
		// SPI0_SDO_PIN	= GPIO19	// Tx Default Serial Out Bus 0 for SPI communications
		// SPI0_SDI_PIN	= GPIO16	// Rx Default Serial In Bus 0 for SPI communications
		// SPI1_SCK_PIN	= GPIO10 // Default Serial Clock Bus 1 for SPI communications
		// SPI1_SDO_PIN	= GPIO11	// Tx Default Serial Out Bus 1 for SPI communications
		// SPI1_SDI_PIN	= GPIO12	// Rx Default Serial In Bus 1 for SPI communications
	}
}
