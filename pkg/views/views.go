package views

import (
	"bytes"
	"mvc/pkg/controllers/renderer"
	"mvc/pkg/utils"
	"net/http"
	"strconv"
)

var Cache bytes.Buffer

func RenderLogin(w http.ResponseWriter, r *http.Request) *http.Request {
	renderer.UserRenderLogin(w)
	return r
}

func RenderSignUp(w http.ResponseWriter, r *http.Request) *http.Request {
	renderer.UserRenderSignUp(w)
	return r
}

func RenderItems(w http.ResponseWriter, r *http.Request) *http.Request {
	switch r.Context().Value("Role") {
	case "Admin":
		{
			var pageStr = r.URL.Query().Get("page")
			var page, _ = strconv.Atoi(pageStr)
			page = max(1, page)

			var filtersStr = r.URL.Query().Get("filters")
			var filters, _ = strconv.Atoi(filtersStr)
			filters = max(0, filters)

			var search = r.URL.Query().Get("search")
			if page == 1 && filters == 0 && search == "" && Cache.Len() > 0 {
				_, err := w.Write(Cache.Bytes())
				if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error") {
					return nil
				}
			} else if page == 1 && filters == 0 && search == "" {
				renderer.AdminRenderItems(&Cache, w, page, filters, search)

			} else {
				renderer.AdminRenderItems(nil, w, page, filters, search)
			}
			return r
		}
	case "Chef":
		{
			http.Redirect(w, r, "/", http.StatusFound)
			return nil
		}
	default:
		{

			var pageStr = r.URL.Query().Get("page")
			var page, _ = strconv.Atoi(pageStr)
			page = max(1, page)

			var filtersStr = r.URL.Query().Get("filters")
			var filters, _ = strconv.Atoi(filtersStr)
			filters = max(0, filters)

			var search = r.URL.Query().Get("search")
			renderer.UserRenderItems(w, page, filters, search)
			return r
		}
	}
}

func RenderOrders(w http.ResponseWriter, r *http.Request) *http.Request {
	switch r.Context().Value("Role") {
	case "Admin":
		{

			var pageStr = r.URL.Query().Get("page")
			var page, _ = strconv.Atoi(pageStr)
			page = max(1, page)
			renderer.AdminRenderOrders(w, page)
			return r
		}
	case "Chef":
		{
			var pageStr = r.URL.Query().Get("page")
			var page, _ = strconv.Atoi(pageStr)
			page = max(1, page)
			renderer.ChefRenderOrders(w, page)
			return r
		}
	default:
		{
			var pageStr = r.URL.Query().Get("page")
			var page, _ = strconv.Atoi(pageStr)
			page = max(1, page)
			var userId = r.Context().Value("UserId").(int)
			renderer.UserRenderOrders(w, userId, page)
			return r
		}
	}
}

func RenderBill(w http.ResponseWriter, r *http.Request) *http.Request {
	switch r.Context().Value("Role") {
	case "Admin":
		{
			http.Redirect(w, r, "/", http.StatusFound)
			return nil
		}
	case "Chef":
		{
			http.Redirect(w, r, "/", http.StatusFound)
			return nil
		}
	default:
		{
			var orderIdStr = r.URL.Query().Get("order")
			var orderId, _ = strconv.Atoi(orderIdStr)
			var userId = r.Context().Value("UserId").(int)
			renderer.UserRenderBill(w, userId, orderId)
			return r
		}
	}
}

func RenderSections(w http.ResponseWriter, r *http.Request) *http.Request {
	switch r.Context().Value("Role") {
	case "Admin":
		{
			renderer.AdminRenderSections(w)
		}

	default:
		{
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
	return r
}
func RenderUsers(w http.ResponseWriter, r *http.Request) *http.Request {
	switch r.Context().Value("Role") {
	case "Admin":
		{
			var pageStr = r.URL.Query().Get("page")
			var page, _ = strconv.Atoi(pageStr)
			page = max(1, page)
			var userId = r.Context().Value("UserId").(int)
			renderer.AdminRenderUsers(w, userId, page)
			return r
		}
	default:
		{
			http.Redirect(w, r, "/", http.StatusFound)
			return nil
		}
	}
}

func RenderHomeScreen(w http.ResponseWriter, r *http.Request) *http.Request {
	switch r.Context().Value("Role") {
	case "Admin":
		{
			http.Redirect(w, r, "/items", http.StatusFound)
			return nil
		}
	case "Chef":
		{
			http.Redirect(w, r, "/orders", http.StatusFound)
			return nil
		}
	default:
		{
			http.Redirect(w, r, "/items", http.StatusFound)
			return nil
		}
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/404.html", http.StatusFound)
}
