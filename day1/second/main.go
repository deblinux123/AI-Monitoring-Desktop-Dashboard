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

			title := material.H1(theme, "Hello, Gophers")

			subTitle := material.H3(theme, "This is testing for day 1")

			maroon := color.NRGBA{
				R: 127,
				G: 0,
				B: 0,
				A: 255,
			}

			subTitleColor := color.NRGBA{
				R: 52,
				G: 235,
				B: 0,
				A: 255,
			}

			title.Color = maroon
			subTitle.Color = subTitleColor

			title.Alignment = text.Middle
			subTitle.Alignment = text.Start

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(
				gtx,

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return title.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 190,
					}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return subTitle.Layout(gtx)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
