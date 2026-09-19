package main

import (
	"fmt"
	"os"
	"strings"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type AppState struct {
	Models        []string
	SelectedModel string
	ServerURL     string
	ErrorMessage  string
	Status        string
}

func main() {
	go func() {
		window := new(app.Window)

		window.Option(
			app.Title("Day 06 - Lists & Forms"),
			app.Size(700, 700),
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

	// =========================
	// Widgets
	// =========================

	var list widget.List

	list.Axis = layout.Vertical

	var serverInput widget.Editor

	serverInput.SingleLine = true

	var validateButton widget.Clickable

	// =========================
	// Application State
	// =========================

	state := AppState{
		Models: []string{
			"qwen2.5:3b",
			"gemma3:4b",
			"embeddinggemma:300m",
			"llama3.2",
			"phi4",
		},

		SelectedModel: "",
		ServerURL:     "",
		ErrorMessage:  "",
		Status:        "Waiting for validation...",
	}

	// =========================
	// Model Buttons
	// =========================

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
			// Model Selection
			// =========================

			for index := range state.Models {

				if modelButtons[index].Clicked(gtx) {

					state.SelectedModel = state.Models[index]

					state.Status = fmt.Sprintf(
						"Selected model: %s",
						state.SelectedModel,
					)

					state.ErrorMessage = ""
				}
			}

			// =========================
			// Server Validation
			// =========================

			if validateButton.Clicked(gtx) {

				serverURL := strings.TrimSpace(
					serverInput.Text(),
				)

				state.ServerURL = serverURL

				// Empty input
				if serverURL == "" {

					state.ErrorMessage =
						"Server URL cannot be empty"

					state.Status = "Validation failed"

				} else if !strings.HasPrefix(serverURL, "http://") &&
					!strings.HasPrefix(serverURL, "https://") {

					state.ErrorMessage =
						"URL must start with http:// or https://"

					state.Status = "Validation failed"

				} else {

					state.ErrorMessage = ""

					state.Status =
						"Server URL is valid"
				}
			}

			// =========================
			// UI
			// =========================

			title := material.H1(
				theme,
				"AI Model Settings",
			)

			serverLabel := material.Body1(
				theme,
				"Ollama Server",
			)

			serverField := material.Editor(
				theme,
				&serverInput,
				"http://localhost:11434",
			)

			validate := material.Button(
				theme,
				&validateButton,
				"Validate Server",
			)

			status := material.Body1(
				theme,
				"Status: "+state.Status,
			)

			selected := material.H3(
				theme,
				fmt.Sprintf(
					"Selected Model: %s",
					state.SelectedModel,
				),
			)

			modelTitle := material.H3(
				theme,
				"Available Models",
			)

			// =========================
			// Error Message
			// =========================

			errorLabel := material.Body1(
				theme,
				state.ErrorMessage,
			)

			// =========================
			// Model List
			// =========================

			modelList := material.List(
				theme,
				&list,
			)

			// =========================
			// Main Layout
			// =========================

			layout.Flex{
				Axis:      layout.Vertical,
				Alignment: layout.Middle,
			}.Layout(
				gtx,

				// Title
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return title.Layout(gtx)
					},
				),

				// Spacing
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{
							Height: 20,
						}.Layout(gtx)
					},
				),

				// Server Label
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return serverLabel.Layout(gtx)
					},
				),

				// Server Input
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Top:    10,
							Bottom: 10,
							Left:   40,
							Right:  40,
						}.Layout(
							gtx,
							serverField.Layout,
						)
					},
				),

				// Validate Button
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return validate.Layout(gtx)
					},
				),

				// Status
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Top:    10,
							Bottom: 5,
						}.Layout(
							gtx,
							status.Layout,
						)
					},
				),

				// Error
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return errorLabel.Layout(gtx)
					},
				),

				// Spacing
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{
							Height: 20,
						}.Layout(gtx)
					},
				),

				// Model Title
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return modelTitle.Layout(gtx)
					},
				),

				// Model List
				layout.Flexed(
					1,
					func(gtx layout.Context) layout.Dimensions {

						return modelList.Layout(
							gtx,
							len(state.Models),

							func(
								gtx layout.Context,
								index int,
							) layout.Dimensions {

								button := material.Button(
									theme,
									&modelButtons[index],
									state.Models[index],
								)

								return layout.Inset{
									Top:    5,
									Bottom: 5,
									Left:   40,
									Right:  40,
								}.Layout(
									gtx,
									button.Layout,
								)
							},
						)
					},
				),

				// Selected Model
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return selected.Layout(gtx)
					},
				),

				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{
							Height: 20,
						}.Layout(gtx)
					},
				),
			)

			e.Frame(gtx.Ops)
		}
	}
}
