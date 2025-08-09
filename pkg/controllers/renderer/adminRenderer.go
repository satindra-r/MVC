package renderer

import (
	"html/template"
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
	"strconv"
)

func AdminRenderItems(w http.ResponseWriter, page int, filters int) {
	var temp = template.Must(template.ParseFiles("pkg/views/adminItems.gohtml"))

	data := map[string]interface{}{
		"Items":    models.GetItems(page, filters),
		"Sections": models.GetSections(),
		"Page":     strconv.Itoa(page),
		"Filters":  strconv.Itoa(filters),
	}
	err := temp.Execute(w, data)
	utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error")

}
func AdminRenderOrders(w http.ResponseWriter, page int) {
	var temp = template.Must(template.ParseFiles("pkg/views/adminOrders.gohtml"))

	data := map[string]interface{}{
		"Orders": models.GetAllOrders(page),
		"Page":   strconv.Itoa(page),
	}
	err := temp.Execute(w, data)
	utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error")

}
func AdminRenderSections(w http.ResponseWriter) {

	var temp = template.Must(template.ParseFiles("pkg/views/sections.gohtml"))
	data := map[string]interface{}{
		"Sections": models.GetSections(),
	}
	err := temp.Execute(w, data)
	utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error")

}
func AdminRenderUsers(w http.ResponseWriter, userId int, page int) {

	var temp = template.Must(template.ParseFiles("pkg/views/users.gohtml"))

	data := map[string]interface{}{
		"Users":       models.GetUsers(page),
		"AdminUserId": userId,
		"Page":        strconv.Itoa(page),
	}
	err := temp.Execute(w, data)
	utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error")
}
