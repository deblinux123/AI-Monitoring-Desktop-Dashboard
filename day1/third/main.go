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

		err := run(window)

		if err != nil {
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
				"AI Monitoring Dahsboard",
			)

			status := material.H3(
				theme,
				"● System Online",
			)

			gpu := material.H3(
				theme,
				"GPU: 42%",
			)

			cpu := material.H3(
				theme,
				"CPU: 35%",
			)

			memory := material.H3(
				theme,
				"Memory: 61%",
			)

			footer := material.Body1(
				theme,
				"Day 01 - Gio",
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

			gpu.Color = color.NRGBA{
				R: 52,
				G: 120,
				B: 235,
				A: 255,
			}

			cpu.Color = color.NRGBA{
				R: 235,
				G: 150,
				B: 52,
				A: 255,
			}

			memory.Color = color.NRGBA{
				R: 180,
				G: 80,
				B: 200,
				A: 255,
			}

			footer.Color = color.NRGBA{
				R: 120,
				G: 120,
				B: 120,
				A: 255,
			}

			title.Alignment = text.Middle
			status.Alignment = text.Middle
			gpu.Alignment = text.Middle
			cpu.Alignment = text.Middle
			memory.Alignment = text.Middle
			footer.Alignment = text.Middle

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
					return status.Layout(gtx)
				}),

				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return gpu.Layout(gtx)
				}),

				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return cpu.Layout(gtx)
				}),

				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return memory.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 20,
					}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return footer.Layout(gtx)
				}),
			)
			e.Frame(gtx.Ops)
		}
	}
}
