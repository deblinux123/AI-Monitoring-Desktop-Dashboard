package main

import (
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func main() {
	go func() {
		window := new(app.Window)

		window.Option(
			app.Title("Day 07 - Drawing Line"),
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

	for {
		switch e := window.Event().(type) {

		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			var path clip.Path

			path.Begin(gtx.Ops)

			path.MoveTo(
				f32.Pt(100, 300),
			)

			path.LineTo(
				f32.Pt(200, 200),
			)

			path.LineTo(
				f32.Pt(300, 250),
			)

			path.LineTo(
				f32.Pt(400, 150),
			)

			path.LineTo(
				f32.Pt(500, 220),
			)

			path.LineTo(
				f32.Pt(600, 100),
			)

			stroke := clip.Stroke{
				Path:  path.End(),
				Width: 5,
			}

			paint.FillShape(
				gtx.Ops,
				color.NRGBA{
					R: 30,
					G: 120,
					B: 220,
					A: 255,
				},
				stroke.Op(),
			)

			e.Frame(gtx.Ops)
		}
	}
}
