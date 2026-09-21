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
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type AppState struct {
	TargetWidth  int
	CurrentWidth int
}

func main() {
	go func() {
		window := new(app.Window)
		window.Option(
			app.Title("Day 08 - Transition"),
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

	var button widget.Clickable

	state := AppState{
		TargetWidth:  200,
		CurrentWidth: 100,
	}

	theme := material.NewTheme()

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			if button.Clicked(gtx) {
				if state.TargetWidth == 200 {
					state.TargetWidth = 500
				} else {
					state.TargetWidth = 200
				}
			}

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(
				gtx,

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					btn := material.Button(
						theme,
						&button,
						"Toggle Panel",
					)

					return btn.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if state.CurrentWidth < state.TargetWidth {
						state.CurrentWidth += 5
					}

					if state.CurrentWidth > state.TargetWidth {
						state.CurrentWidth -= 5
					}

					drawPanel(gtx, state.CurrentWidth)

					return layout.Dimensions{
						Size: image.Pt(state.CurrentWidth, 200),
					}
				}),
			)

			e.Frame(gtx.Ops)

			if state.CurrentWidth != state.TargetWidth {
				window.Invalidate()
			}

			time.Sleep(10 * time.Millisecond)
		}
	}
}

func drawPanel(gtx layout.Context, width int) {
	height := 150

	rect := clip.Rect{
		Min: image.Pt(100, 100),
		Max: image.Pt(100+width, 100+height),
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
}
