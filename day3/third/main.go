package main

import (
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {

	go func() {
		window := new(app.Window)

		window.Option(
			app.Title("Day 03 - List"),
			app.Size(600, 600),
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

	models := []string{
		"gemma3:4b",
		"qwen2.5:3b",
		"embeddinggemma:300m",
		"llama3.2",
		"qwen2.5:7b",
	}

	var list widget.List

	list.Axis = layout.Vertical

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			title := material.H1(
				theme,
				"Ollama Models",
			)

			title.Color = color.NRGBA{
				R: 127,
				G: 0,
				B: 0,
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
						Height: 20,
					}.Layout(gtx)
				}),

				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return material.List(
						theme,
						&list,
					).Layout(
						gtx,
						len(models),

						func(gtx layout.Context, index int) layout.Dimensions {
							item := material.Body1(
								theme,
								models[index],
							)

							return layout.Inset{
								Top:    10,
								Bottom: 10,
								Left:   20,
								Right:  30,
							}.Layout(
								gtx,
								item.Layout,
							)
						},
					)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
