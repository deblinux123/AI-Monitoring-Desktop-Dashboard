package main

import (
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)

		window.Option(
			app.Title("Day 02 - Horizontal Layout"),
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

			title := material.H3(
				theme,
				"Sytem Metrics",
			)

			cpu := material.H3(
				theme,
				"CPU\n35%",
			)

			gpu := material.H3(
				theme,
				"GPU\n42%",
			)

			ram := material.H3(
				theme,
				"RAM\n61%",
			)

			disk := material.H3(
				theme,
				"DISK\n28%",
			)

			title.Alignment = text.Middle
			cpu.Alignment = text.Middle
			gpu.Alignment = text.Middle
			ram.Alignment = text.Middle
			disk.Alignment = text.Middle

			cpu.Color = color.NRGBA{
				R: 50,
				G: 150,
				B: 240,
				A: 255,
			}

			gpu.Color = color.NRGBA{
				R: 50,
				G: 200,
				B: 100,
				A: 255,
			}

			ram.Color = color.NRGBA{
				R: 220,
				G: 160,
				B: 50,
				A: 255,
			}

			disk.Color = color.NRGBA{
				R: 180,
				G: 80,
				B: 200,
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
						Height: 30,
					}.Layout(gtx)
				}),

				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis: layout.Horizontal,
					}.Layout(
						gtx,

						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return cpu.Layout(gtx)
						}),

						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return gpu.Layout(gtx)
						}),

						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return ram.Layout(gtx)
						}),

						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return disk.Layout(gtx)
						}),
					)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
