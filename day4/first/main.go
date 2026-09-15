package main

import (
	"fmt"
	"image/color"
	"os"
	"strconv"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)

		window.Option(
			app.Title("Day 04 - Events"),
			app.Size(700, 500),
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

	clickCount := 0
	statusText := "Waiting for event..."

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case key.Event:
			fmt.Println("Key: ", e.Name)
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			refresh := material.Button(
				theme,
				&refreshButton,
				"Refresh",
			)

			if refreshButton.Clicked(gtx) {
				clickCount++
				statusText = "Refresh button clicked."
			}

			title := material.H1(
				theme,
				"AI Monitoring - Events",
			)

			status := material.H3(
				theme,
				"Status:  "+statusText,
			)

			counter := material.Body1(
				theme,
				"Click Count: "+strconv.Itoa(clickCount),
			)

			title.Color = color.NRGBA{
				R: 127,
				G: 0,
				B: 0,
				A: 255,
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
						Height: 30,
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
					return refresh.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 20,
					}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return counter.Layout(gtx)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
