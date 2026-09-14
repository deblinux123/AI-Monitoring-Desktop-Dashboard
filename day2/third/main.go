package main

import (
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)

		window.Option(
			app.Title("Day 02 - Constraints"),
			app.Size(800, 500),
		)

		if err := run(window); err != nil {
			panic(err)
		}

		os.Exit(0)
	}()

	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()

	var ops op.Ops

	for {
		switch e := window.Event().(type) {

		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			title := material.H2(
				theme,
				"AI Server",
			)

			status := material.H3(
				theme,
				"Server Online",
			)

			status.Color = color.NRGBA{
				R: 0,
				G: 180,
				B: 80,
				A: 255,
			}

			layout.Flex{
				Axis:      layout.Vertical,
				Alignment: layout.Middle,
			}.Layout(
				gtx,

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return title.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 30,
					}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {

					// Limit the width of the widget.
					gtx.Constraints.Min.X = 300
					gtx.Constraints.Max.X = 500

					return status.Layout(gtx)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
