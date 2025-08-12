package api

import (
	"fmt"
	"mvc/pkg/config"
	"mvc/pkg/controllers/admin"
	"mvc/pkg/controllers/chef"
	"mvc/pkg/controllers/user"
	adminMid "mvc/pkg/middleware/admin"
	chefMid "mvc/pkg/middleware/chef"
	userMid "mvc/pkg/middleware/user"
	"mvc/pkg/utils"
	"mvc/pkg/views"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	staticFs := http.FileServer(http.Dir("./pkg/views/"))
	router.PathPrefix("/static/").Handler(staticFs)

	utils.SetRoute(router, "GET", "/", userMid.AuthVerifyUser, chefMid.DBGetUserRole, views.RenderHomeScreen)
	utils.SetRoute(router, "GET", "/login", views.RenderLogin)
	utils.SetRoute(router, "GET", "/signUp", views.RenderSignUp)
	utils.SetRoute(router, "GET", "/items", userMid.AuthVerifyUser, chefMid.DBGetUserRole, views.RenderItems)
	utils.SetRoute(router, "GET", "/orders", userMid.AuthVerifyUser, chefMid.DBGetUserRole, views.RenderOrders)
	utils.SetRoute(router, "GET", "/bill", userMid.AuthVerifyUser, chefMid.DBGetUserRole, views.RenderBill)
	utils.SetRoute(router, "GET", "/sections", userMid.AuthVerifyUser, chefMid.DBGetUserRole, views.RenderSections)
	utils.SetRoute(router, "GET", "/users", userMid.AuthVerifyUser, chefMid.DBGetUserRole, views.RenderUsers)

	utils.SetRoute(router, "POST", "/api/user", user.VerifyCreateUser, user.DBCreateUser)
	utils.SetRoute(router, "POST", "/api/user/login", user.VerifyLogin, user.DBGetUserCredentials, user.AuthCheckUserCredentials)
	utils.SetRoute(router, "POST", "/api/order", user.VerifyCreateOrder, userMid.AuthVerifyUser, user.DBCreateOrder)
	utils.SetRoute(router, "PUT", "/api/dish", chef.VerifyPreparedDish, userMid.AuthVerifyUser, chefMid.DBGetUserRole, chefMid.AuthVerifyChef, chef.DBSetPreparedDish)
	utils.SetRoute(router, "PUT", "/api/order", admin.VerifyPaidOrder, userMid.AuthVerifyUser, chefMid.DBGetUserRole, adminMid.AuthVerifyAdmin, admin.DBSetPaidOrder)
	utils.SetRoute(router, "POST", "/api/sections", admin.VerifyCreateSection, userMid.AuthVerifyUser, chefMid.DBGetUserRole, adminMid.AuthVerifyAdmin, admin.DBCreateSection)
	utils.SetRoute(router, "PUT", "/api/sections", admin.VerifySwapSections, userMid.AuthVerifyUser, chefMid.DBGetUserRole, adminMid.AuthVerifyAdmin, admin.DBSwapSections)
	utils.SetRoute(router, "DELETE", "/api/sections", admin.VerifyDeleteSection, userMid.AuthVerifyUser, chefMid.DBGetUserRole, adminMid.AuthVerifyAdmin, admin.DBDeleteSection)
	utils.SetRoute(router, "PUT", "/api/user", admin.VerifySetUserRole, userMid.AuthVerifyUser, chefMid.DBGetUserRole, adminMid.AuthVerifyAdmin, admin.AuthDisallowDemote, admin.DBSetUserRole)
	utils.SetRoute(router, "POST", "/api/item", admin.VerifyCreateItem, userMid.AuthVerifyUser, chefMid.DBGetUserRole, adminMid.AuthVerifyAdmin, admin.DBCreateItem)
	utils.SetRoute(router, "PUT", "/api/item", admin.VerifyEditItem, userMid.AuthVerifyUser, chefMid.DBGetUserRole, adminMid.AuthVerifyAdmin, admin.DBEditItem)

	return router
}

func PrintRoutes() {
	fmt.Println("Server listening on http://localhost:" + config.EnvConfig.ServerPort)
	fmt.Println("Available endpoints:")
	fmt.Println("POST\t/api/user")
	fmt.Println("POST\t/api/user/login")
	fmt.Println("POST\t/api/order")
	fmt.Println("PUT \t/api/dish")
	fmt.Println("PUT \t/api/order")
	fmt.Println("POST \t/api/sections")
	fmt.Println("PUT \t/api/sections")
	fmt.Println("DELETE \t/api/sections")
	fmt.Println("PUT \t/api/user")
	fmt.Println("POST\t/api/item")
	fmt.Println("PUT \t/api/item")
}
