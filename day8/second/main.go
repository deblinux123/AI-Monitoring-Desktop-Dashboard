package main

import (
	"image"
	"image/color"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func main() {
	go func() {
		window := new(app.Window)

		window.Option(
			app.Title("Day 08 - Loading Indicator"),
			app.Size(800, 600),
		)

		if err := run(window); err != nil {
			panic(err)
		}

		os.Exit(0)
	}()

	app.Main()
}

func run(window *app.Window) error {
	var ops op.Ops

	activeDote := 0

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			drawLoading(gtx, activeDote)

			e.Frame(gtx.Ops)

			activeDote++

			if activeDote >= 4 {
				activeDote = 0
			}

			window.Invalidate()

			time.Sleep(150 * time.Millisecond)
		}
	}
}

func drawLoading(gtx layout.Context, activeDote int) {
	startX := 300
	y := 250

	for i := 0; i < 4; i++ {
		x := startX + i*50

		size := 20

		rect := clip.Rect{
			Min: image.Pt(x, y),
			Max: image.Pt(x+size, y+size),
		}

		var c color.NRGBA

		if i == activeDote {
			c = color.NRGBA{
				R: 30,
				G: 120,
				B: 220,
				A: 255,
			}
		} else {
			c = color.NRGBA{
				R: 100,
				G: 100,
				B: 100,
				A: 255,
			}
		}

		paint.FillShape(
			gtx.Ops,
			c,
			rect.Op(),
		)
	}
}
