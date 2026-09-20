package main

import (
	"image"
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type AppState struct {
	CPUUsage    int
	GPUUsage    int
	MemoryUsage int
}

func main() {
	go func() {
		window := new(app.Window)
		window.Option(
			app.Title("Day 07 - Custom Metric Bar"),
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
		CPUUsage:    68,
		GPUUsage:    54,
		MemoryUsage: 93,
	}

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(
				gtx,

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					drawMetricBar(
						gtx,
						"CPU",
						state.CPUUsage,
					)

					return layout.Dimensions{
						Size: image.Pt(gtx.Constraints.Max.X, 100),
					}
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					drawMetricBar(
						gtx,
						"GPU",
						state.GPUUsage,
					)

					return layout.Dimensions{
						Size: image.Pt(gtx.Constraints.Max.X, 100),
					}
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					drawMetricBar(
						gtx,
						"Memory",
						state.MemoryUsage,
					)

					return layout.Dimensions{
						Size: image.Pt(gtx.Constraints.Max.X, 100),
					}
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}

func drawMetricBar(gtx layout.Context, name string, value int) {
	x := 100
	y := 30

	width := 600
	height := 30

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

	valueWidth := width * value / 100
	progress := clip.Rect{
		Min: image.Pt(x, y),
		Max: image.Pt(x+valueWidth, y+height),
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
