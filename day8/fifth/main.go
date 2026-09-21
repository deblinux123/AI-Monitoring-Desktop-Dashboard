package main

import (
	"image"
	"image/color"
	"math"
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
			app.Title("Day 08 - Animated Status"),
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

	phase := 0.0

	for {
		switch e := window.Event().(type) {

		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			radius := 20.0 + 10.0*math.Sin(phase)

			drawStatus(gtx, int(radius))

			e.Frame(gtx.Ops)

			phase += 0.08

			window.Invalidate()

			time.Sleep(16 * time.Millisecond)
		}
	}
}

func drawStatus(gtx layout.Context, radius int) {
	centerX := 400
	centerY := 300

	rect := clip.Ellipse{
		Min: image.Pt(
			centerX-radius,
			centerY-radius,
		),
		Max: image.Pt(
			centerX+radius,
			centerY+radius,
		),
	}

	paint.FillShape(
		gtx.Ops,
		color.NRGBA{
			R: 40,
			G: 200,
			B: 100,
			A: 255,
		},
		rect.Op(gtx.Ops),
	)
}
