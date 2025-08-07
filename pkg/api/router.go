package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"mvc/pkg/config"
	"mvc/pkg/controllers/chef"
	"mvc/pkg/controllers/user"
	"mvc/pkg/utils"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	utils.SetRoute(router, "POST", "/api/User", user.VerifyCreateUser, user.DBCreateUser)
	utils.SetRoute(router, "POST", "/api/User/login", user.VerifyLogin, user.DBGetUserCredentials, user.AuthCheckUserCredentials)
	utils.SetRoute(router, "POST", "/api/Order", user.VerifyCreateOrder, user.AuthVerifyUser, user.DBCreateOrder)
	utils.SetRoute(router, "PUT", "/api/Dish", chef.VerifyPreparedDish, user.AuthVerifyUser, chef.DBGetUserRole, chef.AuthVerifyChef, chef.DBSetPreparedDish)

	return router
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
