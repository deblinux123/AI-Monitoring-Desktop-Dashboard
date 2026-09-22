package main

import (
	"image"
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(
			app.Title("Day 09 - Line Chart"),
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

	data := []int{
		30,
		45,
		40,
		70,
		55,
		80,
		65,
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
					drawLineChart(gtx, data)

					return layout.Dimensions{
						Size: image.Pt(
							gtx.Constraints.Max.X,
							400,
						),
					}
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}

func drawLineChart(gtx layout.Context, data []int) {
	chartX := 100
	chartY := 50

	chartWidth := 600
	chartHeight := 300

	if len(data) < 2 {
		return
	}

	stepX := float32(chartWidth) / float32(len(data)-1)

	var path clip.Path

	path.Begin(gtx.Ops)

	for i, value := range data {
		x := float32(chartX) + float32(i)*stepX

		y := float32(chartY+chartHeight) - (float32(value) / 100.0 * float32(chartHeight))

		point := f32.Pt(x, y)

		if i == 0 {
			path.MoveTo(point)
		} else {
			path.LineTo(point)
		}
	}

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
			A: 225,
		},
		stroke.Op(),
	)
}
