package main

import (
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
			app.Title("Day 06 - Dynamic Lists"),
			app.Size(600, 500),
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

	var list widget.List

	list.Axis = layout.Vertical

	models := []string{
		"qwen2.5:3b",
		"gemma3:4b",
		"embeddinggemma:300m",
		"llama3.2",
		"phi4",
		"mistral",
		"deepseek-r1",
	}

	for {
		switch e := window.Event().(type) {

		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			title := material.H1(
				theme,
				"Available AI Models",
			)

			modelList := material.List(
				theme,
				&list,
			)

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
						Height: 20,
					}.Layout(gtx)
				}),

				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return modelList.Layout(
						gtx,
						len(models),
						func(gtx layout.Context, index int) layout.Dimensions {

							model := material.Body1(
								theme,
								models[index],
							)

							return layout.Inset{
								Top:    10,
								Bottom: 10,
								Left:   20,
								Right:  20,
							}.Layout(gtx, model.Layout)
						},
					)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
