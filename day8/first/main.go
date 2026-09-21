package main

import (
	"image"
	"image/color"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(
			app.Title("Day 08 - Animation"),
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

	x := 100
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			rect := clip.Rect{
				Min: image.Pt(x, 250),
				Max: image.Pt(x+50, 300),
			}

			paint.FillShape(
				gtx.Ops,
				color.NRGBA{
					R: 30,
					G: 120,
					B: 220,
					A: 255,
				},
				rect.Op(),
			)

			e.Frame(gtx.Ops)

			x += 5

			if x > 750 {
				x = 0
			}

			window.Invalidate()

			time.Sleep(16 * time.Millisecond)
		}
	}
}
