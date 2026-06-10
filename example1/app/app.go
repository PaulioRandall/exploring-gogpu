package app

import (
	// The gogpu graphics suite has a highly decoupled and
	// layered architecture. We must import packages based on
	// our needs to build the graphics stack for an
	// application. There is no God import.
	// github.com/gogpu

	// gg/gpu enables hardware-accelerated rendering if
	// support by your system, else falls back to using CPU.
	_ "github.com/gogpu/gg/gpu"

	_g "github.com/gogpu/gogpu"
	_a "github.com/gogpu/ui/app"
	_d "github.com/gogpu/ui/desktop"
)

type ExampleApp struct {
	// Instances of the gogpu/gogpu and gogpu/ui apps are
	// stored here because we need them to run our app via
	// the Run method.
	gogpuApp *_g.App
	uiApp    *_a.App
}

func NewExampleApp() *ExampleApp {
	// Build a new gogpu App that will power our UI app.
	gogpuApp := createGogpuApp()

	// Build a new UI app powered by the gogpu.App we just
	// created.
	uiApp := createUiApp(gogpuApp)

	// Create an instance of the ClickableArea custom widget.
	clickableArea := NewClickableArea()

	// Set the custom widget as the top-level (root) widget
	// of the UI.
	uiApp.SetRoot(clickableArea.AsWidget())

	return &ExampleApp{
		gogpuApp: gogpuApp,
		uiApp:    uiApp,
	}
}

func createGogpuApp() *_g.App {
	config := _g.DefaultConfig().
		WithTitle("Example App: Clickable Area").
		WithSize(800, 600).
		// By setting to false, the UI will only render when a
		// request to render is triggered by some event or
		// state change in our app.
		WithContinuousRender(false)

	return _g.NewApp(config)
}

func createUiApp(gogpuApp *_g.App) *_a.App {
	return _a.New(
		// Allows our UI app to redraw our widgets to the main
		// UI window managed by the gogpu.App we created.
		_a.WithWindowProvider(gogpuApp),

		// Allows the UI app to pick up on input events
		// detected by the gogpu.App we created. Without it
		// we wouldn't know when the user clicked the mouse
		// button.
		_a.WithEventSource(gogpuApp.EventSource()),
	)
}

// Run runs our built ExampleApp. It blocks so this
// function won't exit until our app has closed.
func (ea *ExampleApp) Run() error {
	return _d.Run(ea.gogpuApp, ea.uiApp)
}
