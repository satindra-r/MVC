package MVC

import (
	"context"
	"fmt"
	"io"
	"mvc/pkg/api"
	"mvc/pkg/config"
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	config.LoadEnvs()
	models.InitDatabase()
	defer models.CloseDatabase()
	var router = api.SetupRouter()
	server := &http.Server{
		Addr:    ":" + config.EnvConfig.ServerPort,
		Handler: router,
	}
	go func() {
		fmt.Printf("Starting server on %s\n", server.Addr)
		err := server.ListenAndServe()
		utils.QuitIfErr(err, "Server error")
	}()

	time.Sleep(time.Second)

	url := "http://localhost:" + config.EnvConfig.ServerPort + "/api/user/login"

	req, err := http.NewRequest("POST", url, nil)

	utils.PanicIfErr(err, "Error creating request")

	req.Header.Set("Username", "user1")
	req.Header.Set("Password", "password")

	client := &http.Client{}
	resp, err := client.Do(req)
	utils.PanicIfErr(err, "Error Reaching Server")

	body, err := io.ReadAll(resp.Body)
	utils.PanicIfErr(err, "Error reading response body")

	if resp.StatusCode != 200 {
		utils.PanicIfErr(err, "Error logging in")
	}
	if string(body) == "" {
		utils.PanicIfErr(err, "Error logging in")
	}

	err = resp.Body.Close()
	utils.PanicIfErr(err, "Unable to close response body")

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	utils.LogIfErr(err, "Error shutting down server")

	fmt.Println("Server exited gracefully")
}
