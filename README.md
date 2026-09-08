# 🚀 Gio AI Monitor

A native cross-platform desktop monitoring dashboard built with **Go + Gio**, designed to monitor AI and backend infrastructure using **Prometheus**.

The goal of this project is to learn Gio by building a real-world application instead of isolated GUI examples.

---

## 🎯 Project Goal

Build a native desktop application that can monitor:

* Go / Fiber backend
* Ollama
* Qdrant
* Prometheus metrics
* RAG pipeline
* CPU / Memory
* HTTP requests
* Request latency
* Errors
* AI inference performance

The final application will display real-time metrics using charts, graphs, gauges, and animations.

---

## 🏗️ Final Architecture

```text
                    ┌─────────────────────┐
                    │     Gio Desktop     │
                    │      Dashboard      │
                    └──────────┬──────────┘
                               │
                         PromQL Queries
                               │
                               ▼
                    ┌─────────────────────┐
                    │     Prometheus      │
                    └──────────┬──────────┘
                               │
                ┌──────────────┼──────────────┐
                │              │              │
                ▼              ▼              ▼
             Fiber          Ollama         Qdrant
             API             LLM          Vector DB
                │              │              │
                └──────────────┼──────────────┘
                               │
                               ▼
                         AI / RAG System
```

---

# 🗺️ 14-Day Learning Roadmap

## Week 1 — Gio Fundamentals

### Day 01 — Installation & First Application

* Install Gio
* Create Go module
* Create first Gio window
* Understand basic Gio structure
* Run the application
* Create GitHub project structure

### Day 02 — Windows & Layout

* Window configuration
* Vertical layout
* Horizontal layout
* Spacing
* Alignment
* Constraints

### Day 03 — Widgets

* Buttons
* Labels
* Text fields
* Lists
* Icons
* Basic Material components

### Day 04 — Events

* Mouse events
* Keyboard events
* Button clicks
* Event handling
* Application interaction

### Day 05 — State Management

* UI state
* Application state
* Updating UI
* Reactive rendering concepts

### Day 06 — Lists & Forms

* Dynamic lists
* Forms
* Input validation
* Model selection UI

### Day 07 — Custom Drawing

* Shapes
* Lines
* Paths
* Custom rendering
* Preparing for charts

---

# Week 2 — AI Monitoring Dashboard

### Day 08 — Animation

* Animation loop
* Loading indicators
* Progress animation
* Transitions
* Animated status indicators

### Day 09 — Charts

Build:

* Line charts
* Bar charts
* Metric graphs
* Real-time graph updates

### Day 10 — HTTP & APIs

Connect Gio to backend services.

```text
Gio
 │
 └── HTTP Client
       │
       ├── Fiber
       ├── Ollama
       └── Qdrant
```

### Day 11 — Prometheus

Learn:

* Prometheus architecture
* Metrics
* Counters
* Gauges
* Histograms
* PromQL
* `/metrics`

### Day 12 — Real-Time Metrics

Connect Gio to Prometheus.

```text
Prometheus
     │
     │ PromQL
     ▼
Go HTTP Client
     │
     ▼
Gio Dashboard
```

### Day 13 — Dashboard

Build the complete dashboard:

* Request rate
* Error rate
* Latency
* CPU
* Memory
* Ollama latency
* Qdrant latency
* RAG performance

### Day 14 — Final Project

Complete:

* UI
* Charts
* Animation
* Prometheus integration
* AI metrics
* Configuration
* Error handling
* README
* Screenshots
* Production build

---

# 📊 Final Dashboard

The final application will look conceptually like:

```text
┌─────────────────────────────────────────────────────┐
│                 🤖 Gio AI Monitor                   │
├───────────────┬─────────────────────────────────────┤
│               │                                     │
│ Dashboard     │       HTTP Requests                 │
│               │                                     │
│ Fiber API     │       ╭────╮                       │
│ Ollama        │   ╭───╯    ╰────╮                  │
│ Qdrant        │ ──╯              ╰──               │
│ RAG           │                                     │
│               │                                     │
│               ├─────────────────────────────────────┤
│               │                                     │
│               │       Request Latency               │
│               │                                     │
│               │       ───────────────               │
│               │              124ms                  │
│               │                                     │
└───────────────┴─────────────────────────────────────┘
```

---

# 🛠️ Technology Stack

## Desktop

* Go
* Gio

## Monitoring

* Prometheus
* PromQL

## Visualization

* Gio custom rendering
* Custom charts

## AI

* Ollama
* Qdrant
* RAG

## Backend

* Go
* Fiber

## Infrastructure

* Docker
* Docker Compose

---

# 📁 Project Structure

```text
gio-ai-monitor/
│
├── cmd/
│   └── monitor/
│       └── main.go
│
├── internal/
│   ├── ui/
│   ├── metrics/
│   ├── prometheus/
│   ├── ollama/
│   ├── qdrant/
│   └── config/
│
├── assets/
│
├── docs/
│
├── go.mod
├── go.sum
├── docker-compose.yml
├── prometheus.yml
└── README.md
```

---

# 🚀 Getting Started

## Requirements

* Go
* Git
* Linux / Windows / macOS
* Gio
* Prometheus
* Docker (later)

---

# 📅 Day 01

## 1. Create the project

```bash
mkdir gio-ai-monitor
cd gio-ai-monitor
```

Initialize Go:

```bash
go mod init github.com/YOUR_USERNAME/gio-ai-monitor
```

Replace `YOUR_USERNAME` with your GitHub username.

---

## 2. Install Gio

```bash
go get gioui.org
```

Then:

```bash
go mod tidy
```

Check:

```bash
go list -m all
```

You should see Gio dependencies in the output.

---

# 🐧 Linux Dependencies

On Ubuntu/Debian, install the basic development libraries:

```bash
sudo apt update
```

```bash
sudo apt install -y \
    gcc \
    pkg-config \
    libwayland-dev \
    libx11-dev \
    libx11-xcb-dev \
    libxkbcommon-dev \
    libxkbcommon-x11-dev \
    libegl1-mesa-dev \
    libgles2-mesa-dev
```

If your Linux desktop uses X11, these dependencies are particularly useful.

---

# 3. Create the first application

Create:

```text
cmd/monitor/main.go
```

Start with a minimal Gio application.

```go
package main

import (
	"gioui.org/app"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(
			app.Title("Gio AI Monitor"),
			app.Size(1000, 700),
		)

		if err := run(w); err != nil {
			panic(err)
		}

		app.Main()
	}()
}

func run(w *app.Window) error {
	return nil
}
```

> We will improve this code during Day 01 as we learn Gio's application and rendering model.

---

# 4. Run the project

```bash
go run ./cmd/monitor
```

If everything is installed correctly, the Gio application should start.

---

# 🧪 Day 01 Exercises

Before moving to Day 02:

* [ ] Create the Go module
* [ ] Install Gio
* [ ] Install Linux dependencies
* [ ] Create `cmd/monitor/main.go`
* [ ] Run the application
* [ ] Change the window title
* [ ] Change the window size
* [ ] Create the first Git commit

---

# 📝 Git

Initialize the repository:

```bash
git init
```

Create `.gitignore`:

```gitignore
# Go
*.exe
*.test
*.out

# Build
/bin/
/build/

# IDE
.idea/
.vscode/

# OS
.DS_Store
```

Then:

```bash
git add .
git commit -m "day01: initialize Gio project"
```

---

# 📈 Future Metrics

The final project will expose and visualize metrics such as:

```text
http_requests_total

http_request_duration_seconds

rag_requests_total

rag_request_duration_seconds

qdrant_search_duration_seconds

ollama_request_duration_seconds

embedding_requests_total

application_memory_bytes

application_goroutines
```

---

# 🎓 Learning Philosophy

This project follows a **build-first learning approach**.

Instead of learning Gio only through isolated examples, every concept will eventually become part of the final monitoring application.

```text
Learn
  ↓
Experiment
  ↓
Build small component
  ↓
Integrate into project
  ↓
Production feature
```

---

# 🚀 Final Goal

The final result should be a professional open-source project demonstrating:

```text
Go
 │
 ├── Gio
 │
 ├── Desktop Development
 │
 ├── Custom UI
 │
 ├── Animation
 │
 ├── Charts
 │
 ├── HTTP
 │
 ├── Prometheus
 │
 ├── Monitoring
 │
 ├── Ollama
 │
 ├── Qdrant
 │
 └── RAG
```

