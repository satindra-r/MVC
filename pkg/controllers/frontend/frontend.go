package frontend

import (
	"html/template"
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
	"strconv"
)

func RenderLogin(w http.ResponseWriter, r *http.Request) *http.Request {
	var temp = template.Must(template.ParseFiles("pkg/views/login.gohtml"))
	err := temp.Execute(w, nil)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error") {
		return nil
	}
	return r
}

func RenderSignUp(w http.ResponseWriter, r *http.Request) *http.Request {
	var temp = template.Must(template.ParseFiles("pkg/views/signUp.gohtml"))
	err := temp.Execute(w, nil)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error") {
		return nil
	}
	return r
}

func RenderItems(w http.ResponseWriter, r *http.Request) *http.Request {

	switch r.Context().Value("Role") {
	case "Admin":
		{
			var temp = template.Must(template.ParseFiles("pkg/views/adminItems.gohtml"))
			var pageStr = r.URL.Query().Get("page")
			var page, _ = strconv.Atoi(pageStr)
			page = max(1, page)
			var filtersStr = r.URL.Query().Get("filters")
			var filters, _ = strconv.Atoi(filtersStr)
			filters = max(0, filters)
			data := map[string]interface{}{
				"Items":    models.GetItems(page, filters),
				"Sections": models.GetSections(),
				"Page":     strconv.Itoa(page),
				"Filters":  strconv.Itoa(filters),
			}
			err := temp.Execute(w, data)
			if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error") {
				return nil
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
			var temp = template.Must(template.ParseFiles("pkg/views/items.gohtml"))
			var pageStr = r.URL.Query().Get("page")
			var page, _ = strconv.Atoi(pageStr)
			page = max(1, page)
			var filtersStr = r.URL.Query().Get("filters")
			var filters, _ = strconv.Atoi(filtersStr)
			filters = max(0, filters)
			data := map[string]interface{}{
				"Items":    models.GetItems(page, filters),
				"Sections": models.GetSections(),
				"Page":     strconv.Itoa(page),
				"Filters":  strconv.Itoa(filters),
			}
			err := temp.Execute(w, data)
			if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error") {
				return nil
			}
			return r
		}
	}
}

func RenderOrders(w http.ResponseWriter, r *http.Request) *http.Request {
	switch r.Context().Value("Role") {
	case "Admin":
		{
			var temp = template.Must(template.ParseFiles("pkg/views/adminOrders.gohtml"))
			var pageStr = r.URL.Query().Get("page")
			var page, _ = strconv.Atoi(pageStr)
			page = max(1, page)
			data := map[string]interface{}{
				"Orders": models.GetAllOrders(page),
				"Page":   strconv.Itoa(page),
			}
			err := temp.Execute(w, data)
			if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error") {
				return nil
			}
			return r
		}
	case "Chef":
		{
			var temp = template.Must(template.ParseFiles("pkg/views/chefOrders.gohtml"))
			var pageStr = r.URL.Query().Get("page")
			var page, _ = strconv.Atoi(pageStr)
			page = max(1, page)
			data := map[string]interface{}{
				"Orders": models.GetAllOrders(page),
				"Page":   strconv.Itoa(page),
			}
			err := temp.Execute(w, data)
			if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error") {
				return nil
			}
			return r
		}
	default:
		{
			var temp = template.Must(template.ParseFiles("pkg/views/orders.gohtml"))
			var pageStr = r.URL.Query().Get("page")
			var page, _ = strconv.Atoi(pageStr)
			page = max(1, page)
			data := map[string]interface{}{
				"Orders": models.GetUserOrders(r.Context().Value("UserId").(int), page),
				"Page":   strconv.Itoa(page),
			}
			err := temp.Execute(w, data)
			if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error") {
				return nil
			}
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
			var temp = template.Must(template.ParseFiles("pkg/views/bill.gohtml"))
			var orderIdStr = r.URL.Query().Get("order")
			var orderId, _ = strconv.Atoi(orderIdStr)
			var order, err = models.GetUserOrder(orderId)
			if err != nil {
				utils.RespondFailure(w, http.StatusBadRequest, "Order Does Not Exist")
				return nil
			}

			if order.UserId != r.Context().Value("UserId").(int) {
				utils.RespondFailure(w, http.StatusForbidden, "Unauthorized Access")
				return nil
			}

			data := map[string]interface{}{
				"Order": order,
			}
			err = temp.Execute(w, data)
			if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error") {
				return nil
			}
			return r
		}
	}
}

func RenderSections(w http.ResponseWriter, r *http.Request) *http.Request {
	switch r.Context().Value("Role") {
	case "Admin":
		{
			var temp = template.Must(template.ParseFiles("pkg/views/sections.gohtml"))
			data := map[string]interface{}{
				"Sections": models.GetSections(),
			}
			err := temp.Execute(w, data)
			if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error") {
				return nil
			}
			return r
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
			var temp = template.Must(template.ParseFiles("pkg/views/users.gohtml"))
			var pageStr = r.URL.Query().Get("page")
			var page, _ = strconv.Atoi(pageStr)
			data := map[string]interface{}{
				"Users":       models.GetUsers(page),
				"AdminUserId": r.Context().Value("UserId"),
				"Page":        strconv.Itoa(page),
			}
			err := temp.Execute(w, data)
			if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error") {
				return nil
			}
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
