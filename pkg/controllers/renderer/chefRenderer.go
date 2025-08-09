package renderer

import (
	"html/template"
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
	"strconv"
)

func ChefRenderOrders(w http.ResponseWriter, page int) {
	var temp = template.Must(template.ParseFiles("pkg/views/chefOrders.gohtml"))

	data := map[string]interface{}{
		"Orders": models.GetAllOrders(page),
		"Page":   strconv.Itoa(page),
	}
	err := temp.Execute(w, data)
	utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Connection Error")

}
