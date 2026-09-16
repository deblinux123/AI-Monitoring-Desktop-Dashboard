package main

import (
	"fmt"
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type AppState struct {
	CPUUsage    int
	GPUUsage    int
	MemoryUsage int
	ClickCount  int
	LastEvent   string
}

func main() {
	go func() {
		window := new(app.Window)

		window.Option(
			app.Title("Day 05 - State Management"),
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

	var resetButton widget.Clickable

	state := AppState{
		CPUUsage:    35,
		GPUUsage:    42,
		MemoryUsage: 60,
		ClickCount:  0,
		LastEvent:   "Waiting...",
	}

	for {
		switch e := window.Event().(type) {

		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// =========================
			// Event → Update State
			// =========================

			if refreshButton.Clicked(gtx) {
				state.CPUUsage = 72
				state.GPUUsage = 51
				state.MemoryUsage = 68

				state.ClickCount++

				state.LastEvent = "Metrics refreshed"
			}

			if resetButton.Clicked(gtx) {
				state.CPUUsage = 0
				state.GPUUsage = 0
				state.MemoryUsage = 0

				state.ClickCount = 0

				state.LastEvent = "All Metrics Reset"
			}

			// =========================
			// State → UI
			// =========================

			title := material.H1(
				theme,
				"AI Monitoring",
			)

			cpu := material.H3(
				theme,
				fmt.Sprintf("CPU: %d%%", state.CPUUsage),
			)

			gpu := material.H3(
				theme,
				fmt.Sprintf("GPU: %d%%", state.GPUUsage),
			)

			memory := material.H3(
				theme,
				fmt.Sprintf("Memory: %d%%", state.MemoryUsage),
			)

			clicks := material.Body1(
				theme,
				fmt.Sprintf(
					"Refresh Count: %d",
					state.ClickCount,
				),
			)

			lastEvent := material.Body1(
				theme,
				"Last Event: "+state.LastEvent,
			)

			refresh := material.Button(
				theme,
				&refreshButton,
				"Refresh Metrics",
			)

			reset := material.Button(
				theme,
				&resetButton,
				"Reset All",
			)

			refresh.Color = color.NRGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			}

			refresh.Background = color.NRGBA{
				R: 30,
				G: 120,
				B: 220,
				A: 255,
			}

			reset.Background = color.NRGBA{
				R: 180,
				G: 40,
				B: 40,
				A: 255,
			}

			layout.Flex{
				Axis:      layout.Vertical,
				Alignment: layout.Middle,
			}.Layout(
				gtx,

				// Title
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return title.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 30,
					}.Layout(gtx)
				}),

				// Metrics
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis: layout.Horizontal,
					}.Layout(
						gtx,

						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Left:  10,
								Right: 10,
							}.Layout(gtx, cpu.Layout)
						}),

						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Left:  10,
								Right: 10,
							}.Layout(gtx, gpu.Layout)
						}),

						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Left:  10,
								Right: 10,
							}.Layout(gtx, memory.Layout)
						}),
					)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 30,
					}.Layout(gtx)
				}),

				// Last event
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return lastEvent.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 10,
					}.Layout(gtx)
				}),

				// Click count
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return clicks.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{
						Height: 30,
					}.Layout(gtx)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis:      layout.Horizontal,
						Alignment: layout.Middle,
					}.Layout(
						gtx,

						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Left:  10,
								Right: 10,
							}.Layout(gtx, refresh.Layout)
						}),

						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Left:  10,
								Right: 10,
							}.Layout(gtx, reset.Layout)
						}),
					)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
