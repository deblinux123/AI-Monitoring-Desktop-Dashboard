package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"net/http"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type OllamaResponse struct {
	Models []OllamaModel `json:"models"`
}

type OllamaModel struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

func main() {
	go func() {
		window := new(app.Window)

		window.Option(
			app.Title("Ollama Monitor"),
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
	theme := material.NewTheme()

	var ops op.Ops

	var refreshButton widget.Clickable
	var modelList widget.List

	modelList.Axis = layout.Vertical

	var models []OllamaModel

	statusText := "Not Connected"

	models, err := getOllamaModels()

	if err != nil {
		statusText = "Ollama is not available"
	} else {
		statusText = "Conected"
	}

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			title := material.H3(
				theme,
				"Ollama Monitor",
			)

			refresh := material.Button(
				theme,
				&refreshButton,
				"Refresh Models",
			)

			status := material.Body1(
				theme,
				"Status: "+statusText,
			)

			modeltitle := material.H3(
				theme,
				"Installed Models",
			)

			title.Color = color.NRGBA{
				R: 127,
				G: 0,
				B: 0,
				A: 255,
			}

			if refreshButton.Clicked(gtx) {
				newModels, err := getOllamaModels()

				if err != nil {
					statusText = "Ollamam is not available"
					models = nil
				} else {
					statusText = "Connected"
					models = newModels
				}
			}

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

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return refresh.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 20,
					}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return status.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 20,
					}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return modeltitle.Layout(gtx)
				}),

				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return material.List(
						theme,
						&modelList,
					).Layout(
						gtx,

						len(models),
						func(gtx layout.Context, index int) layout.Dimensions {
							model := models[index]

							modelName := material.Body1(
								theme,
								fmt.Sprintf(
									"%s  (%s)",
									model.Name,
									formatSize(model.Size),
								),
							)

							return layout.Inset{
								Top:    12,
								Bottom: 12,
								Left:   30,
								Right:  30,
							}.Layout(
								gtx,
								modelName.Layout,
							)
						},
					)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}

func getOllamaModels() ([]OllamaModel, error) {
	resp, err := http.Get("http://localhost:11434/api/tags")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"Ollama returned status %d",
			resp.StatusCode,
		)
	}

	var result OllamaResponse

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result.Models, nil
}

func formatSize(size int64) string {
	const GB = 1024 * 1024 * 1024
	const MB = 1024 * 1024

	if size >= GB {
		return fmt.Sprintf(
			"%.2f GB",
			float64(size)/GB,
		)
	}

	return fmt.Sprintf(
		"%.2f MB",
		float64(size)/MB,
	)
}
