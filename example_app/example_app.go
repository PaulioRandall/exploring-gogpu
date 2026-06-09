package example_app

import (
	// The gogpu "eco system?" has a highly decoupled and
	// layered architecture. We must pick imports based on
	// our needs to build the graphics stack for an
	// application. There is no God import.
	// github.com/gogpu

	// gg/gpu enables hardware-accelerated rendering, if
	// supports it. Falls back to using CPU if not.
	_ "github.com/gogpu/gg/gpu"

	// gogpu/gogpu high-level API for graphics including
	// window management. The UI toolkit this example app is
	// built on uses this API.
	_g "github.com/gogpu/gogpu"

	// gogpu/ui/app provides framework structures for us to
	// configure and populate with our widgets and buisness
	// logic.
	_a "github.com/gogpu/ui/app"

	// gogpu/ui/desktop provides an environment to run the
	// gogpu/ui/app App we configured and populated with our
	// widgets and business logic.
	_d "github.com/gogpu/ui/desktop"

	// clickable_area is a custom widget I created that
	// detects mouse clicks and places boxes containing the
	// mouse click location at the mosue click location.
	"github.com/PaulioRandall/exploring-gogpu/clickable_area"
)

type ExampleApp struct {
	// gogpuApp is stored here because we need it to run
	// uiApp.
	gogpuApp *_g.App

	// uiApp contains our custom widget, it's stored here so
	// we can run it when the Run method is called.
	uiApp *_a.App
}

// NewExampleApp builds and returns our app, it does not
// run it.
func NewExampleApp() *ExampleApp {
	// Create a gogpu.App that will power our ui/app.App.
	gogpuApp := createGogpuApp()

	// Create a ui/app.App powered by the gogpu.App we
	// created above.
	uiApp := createUiApp(gogpuApp)

	// Create an instance of the custom widget I created.
	clickableArea := clickable_area.NewClickableArea()

	// Set the custom widget as the top-level (root) widget.
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
		WithContinuousRender(false)

	return _g.NewApp(config)
}

func createUiApp(gogpuApp *_g.App) *_a.App {
	return _a.New(
		_a.WithWindowProvider(gogpuApp),
		_a.WithPlatformProvider(gogpuApp),
		_a.WithEventSource(gogpuApp.EventSource()),
	)
}

// Run runs our built ExampleApp. It blocks so this
// function won't exit until our app has closed.
func (ea *ExampleApp) Run() error {
	return _d.Run(ea.gogpuApp, ea.uiApp)
}
