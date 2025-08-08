package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"mvc/pkg/config"
	"mvc/pkg/controllers/admin"
	"mvc/pkg/controllers/chef"
	"mvc/pkg/controllers/frontend"
	"mvc/pkg/controllers/user"
	"mvc/pkg/utils"
	"net/http"
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
	fmt.Println("Server listening on http://localhost:" + config.ServerPort)
	fmt.Println("Available endpoints:")
	fmt.Println("  POST  /api/User")
	fmt.Println("  POST  /api/User/login")
	fmt.Println("  POST  /api/Order")
	fmt.Println("  PUT /api/Dish")
	fmt.Println("  PUT  /api/Order")
	fmt.Println("  PUT /api/Sections")
	fmt.Println("  PUT /api/User")
	fmt.Println("  POST  /api/Item")
	fmt.Println("  PUT /api/Item")
}
