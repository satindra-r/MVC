package api

import (
	"fmt"
	"mvc/pkg/config"
	"mvc/pkg/controllers/admin"
	"mvc/pkg/controllers/chef"
	"mvc/pkg/controllers/frontend"
	"mvc/pkg/controllers/user"
	"mvc/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	staticFs := http.FileServer(http.Dir("./pkg/views/"))
	router.PathPrefix("/static/").Handler(staticFs)

	utils.SetRoute(router, "GET", "/", user.AuthVerifyUser, chef.DBGetUserRole, frontend.RenderHomeScreen)
	utils.SetRoute(router, "GET", "/login", frontend.RenderLogin)
	utils.SetRoute(router, "GET", "/signUp", frontend.RenderSignUp)
	utils.SetRoute(router, "GET", "/items", user.AuthVerifyUser, chef.DBGetUserRole, frontend.RenderItems)
	utils.SetRoute(router, "GET", "/orders", user.AuthVerifyUser, chef.DBGetUserRole, frontend.RenderOrders)
	utils.SetRoute(router, "GET", "/bill", user.AuthVerifyUser, chef.DBGetUserRole, frontend.RenderBill)
	utils.SetRoute(router, "GET", "/sections", user.AuthVerifyUser, chef.DBGetUserRole, frontend.RenderSections)
	utils.SetRoute(router, "GET", "/users", user.AuthVerifyUser, chef.DBGetUserRole, frontend.RenderUsers)

	utils.SetRoute(router, "POST", "/api/User", user.VerifyCreateUser, user.DBCreateUser)
	utils.SetRoute(router, "POST", "/api/User/login", user.VerifyLogin, user.DBGetUserCredentials, user.AuthCheckUserCredentials)
	utils.SetRoute(router, "POST", "/api/Order", user.VerifyCreateOrder, user.AuthVerifyUser, user.DBCreateOrder)
	utils.SetRoute(router, "PUT", "/api/Dish", chef.VerifyPreparedDish, user.AuthVerifyUser, chef.DBGetUserRole, chef.AuthVerifyChef, chef.DBSetPreparedDish)
	utils.SetRoute(router, "PUT", "/api/Order", admin.VerifyPaidOrder, user.AuthVerifyUser, chef.DBGetUserRole, admin.AuthVerifyAdmin, admin.DBSetPaidOrder)
	utils.SetRoute(router, "PUT", "/api/Sections", admin.VerifySwapSections, user.AuthVerifyUser, chef.DBGetUserRole, admin.AuthVerifyAdmin, admin.DBSwapSections)
	utils.SetRoute(router, "PUT", "/api/User", admin.VerifySetUserRole, user.AuthVerifyUser, chef.DBGetUserRole, admin.AuthVerifyAdmin, admin.AuthDisallowDemote, admin.DBSetUserRole)
	utils.SetRoute(router, "POST", "/api/Item", admin.VerifyCreateItem, user.AuthVerifyUser, chef.DBGetUserRole, admin.AuthVerifyAdmin, admin.DBCreateItem)
	utils.SetRoute(router, "PUT", "/api/Item", admin.VerifyEditItem, user.AuthVerifyUser, chef.DBGetUserRole, admin.AuthVerifyAdmin, admin.DBEditItem)

	return router
}

func PrintRoutes() {
	fmt.Println("Server listening on http://localhost:" + config.EnvConfig.ServerPort)
	fmt.Println("Available endpoints:")
	fmt.Println("POST\t/api/User")
	fmt.Println("POST\t/api/User/login")
	fmt.Println("POST\t/api/Order")
	fmt.Println("PUT \t/api/Dish")
	fmt.Println("PUT \t/api/Order")
	fmt.Println("PUT \t/api/Sections")
	fmt.Println("PUT \t/api/User")
	fmt.Println("POST\t/api/Item")
	fmt.Println("PUT \t/api/Item")
}
