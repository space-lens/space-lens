package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/space-lens/space-lens/desktop"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func runDesktop() {
	// Create an instance of the app structure
	app := desktop.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Space-lens 🔭",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func runCLI() {
	fmt.Println("CLI soon")
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "cli" {
			runCLI()
			os.Exit(1)
		}
	}
	runDesktop()

}
