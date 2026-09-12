package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

func main() {
	go func() {
		w := new(app.Window)

		w.Option(
			app.Title("Gio AI Monitor"),
			app.Size(unit.Dp(1000), unit.Dp(700)),
		)

		if err := run(w); err != nil {
			panic(err)
		}
	}()

	app.Main()
}

func run(w *app.Window) error {
	var ops op.Ops

	for {
		e := w.Event()

		switch e := e.(type) {

		case app.FrameEvent:
			gtx := layout.Context{
				Now:    e.Now,
				Metric: e.Metric,
			}

			layout.Stack{}.Layout(gtx)

			e.Frame(&ops)

		case app.DestroyEvent:
			return e.Err
		}
	}
}
