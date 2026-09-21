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
)

type AppState struct {
	TargetCPU    int
	DisplayedCPu int
}

func main() {
	go func() {
		window := new(app.Window)
		window.Option(
			app.Title("Day 08 - Animated Progress"),
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

	state := AppState{
		TargetCPU:    100,
		DisplayedCPu: 0,
	}

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			drawProgressBar(gtx, state.DisplayedCPu)

			e.Frame(gtx.Ops)

			if state.DisplayedCPu < state.TargetCPU {
				state.DisplayedCPu++
			}

			if state.DisplayedCPu > state.TargetCPU {
				state.DisplayedCPu--
			}

			window.Invalidate()

			time.Sleep(20 * time.Millisecond)
		}
	}
}

func drawProgressBar(gtx layout.Context, value int) {
	x := 100
	y := 250

	width := 600
	height := 40

	background := clip.Rect{
		Min: image.Pt(x, y),
		Max: image.Pt(x+width, y+height),
	}

	paint.FillShape(
		gtx.Ops,
		color.NRGBA{
			R: 50,
			G: 50,
			B: 50,
			A: 255,
		},

		background.Op(),
	)

	progressWidth := width * value / 100

	progress := clip.Rect{
		Min: image.Pt(x, y),
		Max: image.Pt(x+progressWidth, y+height),
	}

	paint.FillShape(
		gtx.Ops,
		color.NRGBA{
			R: 40,
			G: 180,
			B: 100,
			A: 255,
		},
		progress.Op(),
	)
}
