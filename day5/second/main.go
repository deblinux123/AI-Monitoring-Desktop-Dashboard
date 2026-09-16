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
	Count     int
	LastEvent string
}

func main() {
	go func() {
		window := new(app.Window)
		window.Option(
			app.Title("Day 05 - State Managment"),
			app.Size(600, 400),
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

	var button widget.Clickable

	newState := AppState{
		Count:     0,
		LastEvent: "Waiting...",
	}

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			if button.Clicked(gtx) {
				newState.Count++
				newState.LastEvent = "Increment clicked"
			}

			title := material.H3(
				theme,
				"Counter",
			)

			counter := material.H3(
				theme,
				fmt.Sprintf("Count: %d", newState.Count),
			)

			event := material.Body1(
				theme,
				"Last Event: "+newState.LastEvent,
			)

			increment := material.Button(
				theme,
				&button,
				"Increment",
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

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return counter.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 10,
					}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return event.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 20,
					}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return increment.Layout(gtx)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
