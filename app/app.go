package app

import (
	"github.com/gogpu/gogpu"

	_a "github.com/gogpu/ui/app"
	_d "github.com/gogpu/ui/desktop"

	"github.com/PaulioRandall/exploring-gogpu/grid_canvas"
)

// TODO: Rename to something more precise.
//       Everything here is called "app" :(

type Application struct {
	gogpuApp *gogpu.App
	uiApp    *_a.App
}

func NewApplication() *Application {
	gogpuApp := createGoGPU()
	uiApp := createUiApp(gogpuApp)

	gridCanvas := grid_canvas.NewGridCanvas()

	uiApp.SetRoot(gridCanvas.AsWidget())

	gogpuApp.OnDraw(func(ctx *gogpu.Context) {
		println("HERE!!")
		ctx.Clear(0.5, 0.5, 0.5, 1.0)
		uiApp.Frame()
	})

	return &Application{
		gogpuApp: gogpuApp,
		uiApp:    uiApp,
	}
}

func createGoGPU() *gogpu.App {
	config := gogpu.DefaultConfig().
		WithTitle("Grid Canvas").
		WithSize(800, 600).
		WithContinuousRender(false)

	return gogpu.NewApp(config)
}

func createUiApp(gogpuApp *gogpu.App) *_a.App {
	return _a.New(
		_a.WithWindowProvider(gogpuApp),
		_a.WithPlatformProvider(gogpuApp),
		_a.WithEventSource(gogpuApp.EventSource()),
	)
}

func (a *Application) Run() error {
	return _d.Run(a.gogpuApp, a.uiApp)
}

func (a *Application) Close() {
	// TODO
}
