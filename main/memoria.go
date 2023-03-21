package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var rendereable = []ui.Drawable{}

const width = 14
const height = 6

func principal() {
	var i, j int = 42, 2701
	p := &i

	// Agrego para mostrar
	addVariable("i", &i, 0, 0)
	addVariable("j", &j, 1, 0)
	addVariableP("p", &p, 2, 0)

	actualizarConsola()

	time.Sleep(5 * time.Second)

	*p = *p * 2

	addVariable("i", &i, 0, 1)
	addVariable("j", &j, 1, 1)
	addVariableP("p", &p, 2, 1)

	actualizarConsola()
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// Codigo
	principal()

	actualizarConsola()

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func addVariable(name string, direction *int, column int, row int) {
	pi := widgets.NewParagraph()
	pi.Text = fmt.Sprint(name, " [", reflect.TypeOf(*direction), "](fg:blue)", "\n", "[", *direction, "](fg:red,mod:bold)", "\n", &*direction)
	pi.SetRect(column*width, row*height, (column+1)*width, (row+1)*height)
	//pi.Border = false

	rendereable = append(rendereable, pi)
}

func addVariableP(name string, direction **int, column int, row int) {
	pi := widgets.NewParagraph()
	pi.Text = fmt.Sprint(name, " [", reflect.TypeOf(*direction), "](fg:blue)", "\n", "[", *direction, "](fg:red,mod:bold)", "\n", direction, "\n", "[", **direction, "](fg:yellow,mod:bold)")
	pi.SetRect(column*width, row*height, (column+1)*width, (row+1)*height)
	//pi.Border = false

	rendereable = append(rendereable, pi)
}

func actualizarConsola() {
	ui.Render(rendereable...)
}
