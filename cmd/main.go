package main

import (
	"context"
	"fmt"
	"mvc/pkg/api"
	"mvc/pkg/config"
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.LoadEnvs()

	models.InitDatabase()

	defer models.CloseDatabase()

	var router = api.SetupRouter()

	api.PrintRoutes()
	server := &http.Server{
		Addr:    ":" + config.ServerPort,
		Handler: router,
	}
	go func() {
		fmt.Printf("Starting server on %s\n", server.Addr)
		err := server.ListenAndServe()
		utils.QuitIfErr(err, "Server error")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var err = server.Shutdown(ctx)
	utils.LogIfErr(err, "Error shutting down server")

	fmt.Println("Server exited gracefully")
}
