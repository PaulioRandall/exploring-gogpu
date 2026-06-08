package app

import (
	_ "github.com/gogpu/gg/gpu"

	"github.com/gogpu/gg"
	"github.com/gogpu/gogpu"

	"github.com/gogpu/ui/app"
	"github.com/gogpu/ui/desktop"

	"github.com/PaulioRandall/exploring-gogpu/grid_canvas"
)

// TODO: Rename to something more precise.
//       Everything here is called "app" :(

type Application struct {
	gogpuApp *gogpu.App
	uiApp    *app.App

	drawCtx    *gg.Context
	gridCanvas grid_canvas.GridCanvas
}

func NewApplication() *Application {
	gogpuApp := createGoGPU()
	uiApp := createUiApp(gogpuApp)
	drawCtx := gg.NewContext(100, 100)
	gridCanvas := grid_canvas.Make(drawCtx)

	uiApp.SetRoot(gridCanvas.AsWidget())

	return &Application{
		gogpuApp:   gogpuApp,
		uiApp:      uiApp,
		drawCtx:    drawCtx,
		gridCanvas: gridCanvas,
	}
}

func createGoGPU() *gogpu.App {
	config := gogpu.DefaultConfig().
		WithTitle("My App").
		WithSize(800, 600).
		WithContinuousRender(false)

	return gogpu.NewApp(config)
}

func createUiApp(gogpuApp *gogpu.App) *app.App {
	return app.New(
		app.WithWindowProvider(gogpuApp),
		app.WithPlatformProvider(gogpuApp),
		app.WithEventSource(gogpuApp.EventSource()),
	)
}

func (a *Application) Run() error {
	return desktop.Run(a.gogpuApp, a.uiApp)
}

func (a *Application) Close() {
	a.drawCtx.Close()
}
