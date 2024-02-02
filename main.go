package main

import (
	"context"
	"embed"
	"log"

	"github.com/knadh/koanf/v2"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/kmesiab/orion/config"
	gh "github.com/kmesiab/orion/github"
)

//go:embed all:frontend/dist
var assets embed.FS

// k is the global config instance
// This is used to hold the configuration
var k = koanf.New(".")

func main() {
	// Load the configuration file
	config, err := config.LoadConfig(k, config.EnvFileName)
	if err != nil {
		log.Fatal("Error loading configuration: ", err.Error())
	}

	// Create an instance of the app structure
	app := &App{
		ctx:           context.Background(), // Initialize a new background context.
		GithubService: &gh.Client{},         // Set the GitHub service client.
		Config:        config,               // Set the config instance.
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Orion",
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
