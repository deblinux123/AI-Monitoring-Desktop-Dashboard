package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

type UI struct {
	theme    *material.Theme
	iput     widget.Editor
	askBtn   widget.Clickable
	respList widget.List

	mu      sync.Mutex
	answer  string
	loading bool
}

func (u *UI) ask(w *app.Window, prompt string) {
	if prompt == "" {
		return
	}

	u.mu.Lock()
	u.loading = true
	u.answer = ""
	u.mu.Unlock()
	w.Invalidate()

	go func() {
		body, _ := json.Marshal(OllamaRequest{
			Model:  "qwen2.5:3b",
			Prompt: prompt,
			Stream: false,
		})

		resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewReader(body))

		u.mu.Lock()

		defer u.mu.Unlock()

		u.loading = false

		if err != nil {
			u.answer = "error: " + err.Error()
			w.Invalidate()
			return
		}

		defer resp.Body.Close()

		var out OllamaResponse

		if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
			u.answer = "Decode error: " + err.Error()
			w.Invalidate()
			return
		}

		u.answer = out.Response
		w.Invalidate()
	}()
}

func (u *UI) Layout(gtx layout.Context) layout.Dimensions {
	u.mu.Lock()
	answer := u.answer
	loading := u.loading
	u.mu.Unlock()

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(
		gtx,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			ed := material.Editor(u.theme, &u.iput, "Ask Somthing...")
			u.iput.SingleLine = true
			u.iput.Submit = true
			return layout.UniformInset(unit.Dp(12)).Layout(gtx, ed.Layout)
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			label := "Ask"
			if loading {
				label = "Thinking..."
			}

			btn := material.Button(u.theme, &u.askBtn, label)
			return layout.UniformInset(unit.Dp(12)).Layout(gtx, btn.Layout)
		}),

		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			u.respList.Axis = layout.Vertical

			return u.respList.Layout(gtx, 1, func(gtx layout.Context, index int) layout.Dimensions {
				return layout.UniformInset(unit.Dp(12)).Layout(gtx, material.Body1(u.theme, answer).Layout)
			})
		}),
	)
}

func main() {
	go func() {
		window := new(app.Window)
		window.Option(
			app.Title("Gio + Ollama"),
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
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				for {
					event, ok := ui.iput.Update(gtx)

					if !ok {
						break
					}
					if _, ok := event.(widget.SubmitEvent); ok {
						ui.ask(window, ui.iput.Text())
					}
				}

				if ui.askBtn.Clicked(gtx) {
					ui.ask(window, ui.iput.Text())
				}

				ui.Layout(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()

	app.Main()
}
