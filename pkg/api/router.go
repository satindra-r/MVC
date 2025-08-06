package api

import (
	"fmt"
	"mvc/pkg/config"
	"mvc/pkg/controllers/user"
	"mvc/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler).Methods("GET")

	utils.SetRoute(router, "POST", "/api/User", user.VerifCreateUser, user.DBCreateUser)
	utils.SetRoute(router, "POST", "/api/User/login", user.VerifLogin, user.DBGetUserCredentials, user.AuthCheckUserCredentials)

	return router
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World")
}

// TODO do this

func PrintRoutes() {
	fmt.Println("Server listening on http://localhost:" + config.ServerPort)
	/*fmt.Println("Available endpoints:")
	fmt.Println("  GET  /")
	fmt.Println("  GET  /users")
	fmt.Println("  GET  /users/{id}")
	fmt.Println("  POST /users/add")
	fmt.Println("  PUT  /users/update/{id}")
	fmt.Println("  DELETE /users/delete/{id}")*/
}
