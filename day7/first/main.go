package main

import (
	"image"
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(
			app.Title("Day 07 - Custom Drawing"),
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
	// theme := material.NewTheme()

	var ops op.Ops

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			rect := clip.Rect{
				Min: image.Pt(200, 150),
				Max: image.Pt(400, 250),
			}

			paint.FillShape(
				gtx.Ops,
				color.NRGBA{
					R: 40,
					G: 200,
					B: 100,
					A: 255,
				},
				rect.Op(),
			)
			e.Frame(gtx.Ops)
		}
	}
}
