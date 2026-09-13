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
			app.Title("Day 02 - Vertical Layout"),
			app.Size(600, 400),
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

			title := material.H1(
				theme,
				"AI Monitoring Dashboard",
			)

			status := material.H3(
				theme,
				"System Online",
			)

			footer := material.Body1(
				theme,
				"Day 02- Gio",
			)

			title.Color = color.NRGBA{
				R: 127,
				G: 0,
				B: 0,
				A: 255,
			}

			status.Color = color.NRGBA{
				R: 0,
				G: 180,
				B: 80,
				A: 255,
			}

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(
				gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return title.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 40,
					}.Layout(gtx)
				}),

				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return status.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return footer.Layout(gtx)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
