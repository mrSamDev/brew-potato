# brew-potato

A terminal UI for managing Homebrew packages, built with [Bubble Tea](https://github.com/charmbracelet/bubbletea).

## Why

I use Homebrew constantly on my Mac. `brew list` dumps everything — your packages, their dependencies, transitive dependencies, all tangled together. There's no easy way to just see what *you* actually installed.

I knew about alternatives. Nix exists. There's even a Homebrew rewrite in Rust (zerobrew). I tried them, forgot about them, went back to `brew`.

So I built this to scratch my own itch: show only user-installed packages, nothing else.

Also I was learning Go and needed a real project to practice on. This is that project.

## Features

- Full-screen TUI with rounded table border and styled header/footer
- Lists all user-installed Homebrew formulae with install date
- Async uninstall with a live spinner — row turns red in-place, no view flicker
- Graceful error display if `brew` is unavailable

## Requirements

- [Go](https://go.dev/) 1.21+
- [Homebrew](https://brew.sh/)

## Run

```sh
go mod tidy
go run .
```

## Build

```sh
go build -o brew-potato .
./brew-potato
```

## Project Structure

```
brew-ui/
├── main.go                 # entry point
├── internal/
│   ├── brew/
│   │   └── brew.go         # brew CLI wrapper & data types
│   └── ui/
│       ├── model.go        # Bubble Tea model, Init, Update
│       ├── view.go         # View rendering
│       └── styles.go       # lipgloss styles & table theme
├── go.mod
└── go.sum
```

## Keybindings

| Key       | Action             |
|-----------|--------------------|
| `↑` / `↓` | Navigate packages  |
| `d`       | Uninstall selected |
| `q`       | Quit               |
