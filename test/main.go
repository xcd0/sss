package main

import (
	"context"
	_ "embed"
	"log"
	"machine"
	"machine/usb"

	keyboard "github.com/sago35/tinygo-keyboard"
	"github.com/sago35/tinygo-keyboard/keycodes/jp"
)

func main() {
	usb.Product = "xxxx-0.1.0"

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	d := keyboard.New()

	col1 := machine.GPIO0
	col2 := machine.GPIO1
	row1 := machine.GPIO2
	row2 := machine.GPIO3

	colPins := []machine.Pin{
		col1,
		col2,
	}

	rowPins := []machine.Pin{
		row1,
		row2,
	}

	d.AddMatrixKeyboard(colPins, rowPins, [][]keyboard.Keycode{
		{
			jp.KeyT, jp.KeyI,
			jp.KeyY, jp.KeyG,
		},
	})

	// for Vial
	//loadKeyboardDef()

	//d.Debug = true
	return d.Loop(context.Background())
}
