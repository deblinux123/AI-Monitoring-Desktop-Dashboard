package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type UI struct {
	theme    *material.Theme
	fetchBtn widget.Clickable
	List     widget.List

	mu     sync.Mutex
	todos  []Todo
	status string
}

func (u *UI) fetchTodos(w *app.Window) {
	u.mu.Lock()
	u.status = "Loading ... "
	u.mu.Unlock()

	w.Invalidate()

	go func() {
		resp, err := http.Get("http://localhost:3000/api/todos")
		u.mu.Lock()

		defer u.mu.Unlock()

		if err != nil {
			u.status = "error: " + err.Error()
			w.Invalidate()
			return
		}

		defer resp.Body.Close()

		var result []Todo

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			u.status = "Decode error: " + err.Error()
			w.Invalidate()
			return
		}

		u.todos = result
		u.status = ""
		w.Invalidate()
	}()
}

func (u *UI) Layout(gtx layout.Context) layout.Dimensions {
	u.mu.Lock()
	todos := append([]Todo(nil), u.todos...)
	status := u.status
	u.mu.Unlock()

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(u.theme, &u.fetchBtn, "Getch Todos")

			return layout.UniformInset(unit.Dp(12)).Layout(gtx, btn.Layout)
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if status == "" {
				return layout.Dimensions{}
			}

			return layout.UniformInset(unit.Dp(12)).Layout(gtx, material.Body1(u.theme, status).Layout)
		}),

		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			u.List.Axis = layout.Vertical

			return u.List.Layout(gtx, len(todos), func(gtx layout.Context, index int) layout.Dimensions {
				t := todos[index]
				label := t.Text

				if t.Done {
					label = "✓ " + label
				}

				return layout.UniformInset(unit.Dp(6)).Layout(gtx, material.Body1(u.theme, label).Layout)
			})
		}),
	)
}

func main() {
	go func() {
		window := new(app.Window)
		window.Option(
			app.Title("Gio + Fiber"),
			app.Size(800, 600),
		)

		ui := &UI{
			theme: material.NewTheme(),
		}

		var ops op.Ops

		for {
			switch e := window.Event().(type) {
			case app.DestroyEvent:
				if e.Err != nil {
					log.Fatal(e.Err)
				}
				os.Exit(0)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				if ui.fetchBtn.Clicked(gtx) {
					ui.fetchTodos(window)
				}

				ui.Layout(gtx)

				e.Frame(gtx.Ops)
			}
		}
	}()

	app.Main()
}
