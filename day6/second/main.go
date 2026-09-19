package main

import (
	"fmt"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type AppState struct {
	Models        []string
	SelectedModel string
	LastEvent     string
}

func main() {
	go func() {
		window := new(app.Window)

		window.Option(
			app.Title("Day 06 - Model Selection"),
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

	state := AppState{
		Models: []string{
			"qwen2.5:3b",
			"gemma3:4b",
			"embeddinggemma:300m",
			"llama3.2",
			"phi4",
		},
		SelectedModel: "",
		LastEvent:     "No model selected",
	}

	var modelButtons []widget.Clickable

	for len(modelButtons) < len(state.Models) {
		modelButtons = append(
			modelButtons,
			widget.Clickable{},
		)
	}

	for {
		switch e := window.Event().(type) {

		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// =========================
			// Handle Model Selection
			// =========================

			for index := range state.Models {
				if modelButtons[index].Clicked(gtx) {
					state.SelectedModel = state.Models[index]
					state.LastEvent = fmt.Sprintf(
						"Selected: %s",
						state.Models[index],
					)
				}
			}

			// =========================
			// UI
			// =========================

			title := material.H1(
				theme,
				"AI Model Selection",
			)

			selected := material.H3(
				theme,
				fmt.Sprintf(
					"Selected Model: %s",
					state.SelectedModel,
				),
			)

			lastEvent := material.Body1(
				theme,
				"Last Event: "+state.LastEvent,
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
						len(state.Models),
						func(gtx layout.Context, index int) layout.Dimensions {

							button := material.Button(
								theme,
								&modelButtons[index],
								state.Models[index],
							)

							return layout.Inset{
								Top:    5,
								Bottom: 5,
								Left:   20,
								Right:  20,
							}.Layout(
								gtx,
								button.Layout,
							)
						},
					)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 15,
					}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return selected.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 10,
					}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return lastEvent.Layout(gtx)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
