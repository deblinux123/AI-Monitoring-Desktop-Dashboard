package main

import (
	"fmt"
	"image/color"
	"os"
	"strings"

	"gioui.org/app"
	"gioui.org/io/event" // Added for event.Op
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
			app.Title("Day 04 - Interactive AI Monitor"), // Fixed typo: Intaractive -> Interactive
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

	var cpuCard widget.Clickable
	var gpuCard widget.Clickable
	var refreshButton widget.Clickable

	lastEvent := "Waiting for an event..."
	keyPressed := "None"
	clickCount := 0

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			event.Op(gtx.Ops, window)
			for {
				evt, ok := gtx.Event(key.Filter{})
				if !ok {
					break
				}
				switch evt := evt.(type) {
				case key.Event:
					if evt.State == key.Press {
						keyPressed = string(evt.Name)
						lastEvent = "Keyboard: " + keyPressed
						if strings.EqualFold(keyPressed, "R") {
							lastEvent = "Keyboard: R → Refresh"
						}
					}
				}
			}

			cpu := material.Button(
				theme,
				&cpuCard,
				"CPU\n35%",
			)

			gpu := material.Button(
				theme,
				&gpuCard,
				"GPU\n42%",
			)

			refresh := material.Button(
				theme,
				&refreshButton,
				"Refresh",
			)

			if cpuCard.Clicked(gtx) {
				clickCount++
				lastEvent = "CPU card clicked"
			}

			if gpuCard.Clicked(gtx) {
				clickCount++
				lastEvent = "GPU card clicked"
			}

			if refreshButton.Clicked(gtx) {
				clickCount++
				lastEvent = "Refresh clicked"
			}

			title := material.H1(
				theme,
				"AI Monitoring",
			)

			status := material.H3(
				theme,
				"● System Online",
			)

			eventLabel := material.Body1(
				theme,
				"Last Event: "+lastEvent,
			)

			keyLabel := material.Body1(
				theme,
				"Keyboard: "+keyPressed,
			)

			clickLabel := material.Body1(
				theme,
				fmt.Sprintf(
					"Total Clicks: %d",
					clickCount,
				),
			)

			title.Color = color.NRGBA{
				R: 127,
				G: 0,
				B: 0,
				A: 255,
			}

			status.Color = color.NRGBA{
				R: 0,
				G: 180,
				B: 80,
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
					return layout.Spacer{Height: 20}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return status.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{Height: 30}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal}.Layout(
						gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{Left: 10, Right: 10}.Layout(gtx, cpu.Layout)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{Left: 10, Right: 10}.Layout(gtx, gpu.Layout)
						}),
					)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{Height: 30}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return eventLabel.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{Height: 10}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return keyLabel.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{Height: 30}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return clickLabel.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{Height: 30}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return refresh.Layout(gtx)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
