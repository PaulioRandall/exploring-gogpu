package example_app

import (
	_g "github.com/gogpu/gogpu"
	_a "github.com/gogpu/ui/app"
	_d "github.com/gogpu/ui/desktop"

	"github.com/PaulioRandall/exploring-gogpu/clickable_area"
)

type ExampleApp struct {
	gogpuApp *_g.App
	uiApp    *_a.App
}

func NewExampleApp() *ExampleApp {
	gogpuApp := createGogpuApp()
	uiApp := createUiApp(gogpuApp)

	clickableArea := clickable_area.NewClickableArea()
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

func (ea *ExampleApp) Run() error {
	// Blocks thread.
	return _d.Run(ea.gogpuApp, ea.uiApp)
}
