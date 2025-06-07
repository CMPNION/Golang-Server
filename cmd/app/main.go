package main

import (
	"fmt"
	"net/http"

	"myapp/internal/config"
	"myapp/internal/router"
	"myapp/pkg/logger"
)

func main() {
	r := router.NewRouter()
	logger := logger.NewLogger()
	serverCfg, _ := config.NewServerConfig()

	logger.Info(fmt.Sprintf("Server running at %s", serverCfg.Adress))
	http.ListenAndServe(serverCfg.Adress, r)
}
