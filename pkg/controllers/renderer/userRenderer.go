package renderer

import (
	"html/template"
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
	"strconv"
)

func UserRenderLogin(w http.ResponseWriter) {
	var temp = template.Must(template.ParseFiles("pkg/views/login.gohtml"))
	err := temp.Execute(w, nil)
	utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error")

}

func UserRenderSignUp(w http.ResponseWriter) {
	var temp = template.Must(template.ParseFiles("pkg/views/signUp.gohtml"))
	err := temp.Execute(w, nil)
	utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error")
}

func UserRenderItems(w http.ResponseWriter, page int, filters int, search string) {
	var temp = template.Must(template.ParseFiles("pkg/views/items.gohtml"))
	data := map[string]interface{}{
		"Items":    models.GetItems(page, filters, search),
		"Sections": models.GetSections(),
		"Page":     strconv.Itoa(page),
		"Filters":  strconv.Itoa(filters),
		"Search":   search,
	}
	err := temp.Execute(w, data)
	utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error")

}

func UserRenderOrders(w http.ResponseWriter, userId int, page int) {
	var temp = template.Must(template.ParseFiles("pkg/views/orders.gohtml"))
	data := map[string]interface{}{
		"Orders": models.GetUserOrders(userId, page),
		"Page":   strconv.Itoa(page),
	}
	err := temp.Execute(w, data)
	utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error")

}

func UserRenderBill(w http.ResponseWriter, userId int, orderId int) {
	var temp = template.Must(template.ParseFiles("pkg/views/bill.gohtml"))

	var order, err = models.GetUserOrder(orderId)
	if err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Order Does Not Exist")
		return
	}

	if order.UserId != userId {
		utils.RespondFailure(w, http.StatusForbidden, "Unauthorized Access")
		return
	}

	data := map[string]interface{}{
		"Order": order,
	}
	err = temp.Execute(w, data)
	utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error")

}
