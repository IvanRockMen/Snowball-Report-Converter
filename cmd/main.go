package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/IvanRockMen/Snowball-Report-Converter/internal/config"
	"github.com/IvanRockMen/Snowball-Report-Converter/internal/pluginstorage"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	logger.Debug("debug %v", cfg)

	ps := pluginstorage.InitPluginStorage()

	for k, v := range cfg.Plugins {
		if err := ps.LoadPlugin(k, v); err != nil {
			logger.Error("Error: %v", err)
		}
	}
}
