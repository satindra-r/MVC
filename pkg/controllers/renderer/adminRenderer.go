package renderer

import (
	"bytes"
	"html/template"
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
	"strconv"
)

var renderTemp = template.Must(template.ParseFiles("pkg/views/adminItems.gohtml"))

func AdminRenderItems(cache *bytes.Buffer, w http.ResponseWriter, page int, filters int, search string) {

	data := map[string]interface{}{
		"Items":    models.GetItems(page, filters, search),
		"Sections": models.GetSections(),
		"Page":     strconv.Itoa(page),
		"Filters":  strconv.Itoa(filters),
		"Search":   search,
	}
	err := renderTemp.Execute(w, data)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error") {
		return
	}
	if cache != nil {
		_ = renderTemp.Execute(cache, data)
	}

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
