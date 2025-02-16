//go:build goexperiment.arenas

/*
Shortlink application

API-service
*/
package main

import (
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/handle_signal"
	_ "github.com/batazor/shortlink/internal/pkg/i18n"
	"github.com/batazor/shortlink/internal/services/api/di"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-api")

	// Init a new service
	service, cleanup, err := api_di.InitializeAPIService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			service.Logger.Error(r.(string))
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	handle_signal.WaitExitSignal()

	cleanup()
}
