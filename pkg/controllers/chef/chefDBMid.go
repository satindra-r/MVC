package chef

import (
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
)

func DBSetPreparedDish(w http.ResponseWriter, r *http.Request) *http.Request {
	var dishId = r.Context().Value("DishId").(int)
	var prepared = r.Context().Value("Prepared").(int)
	var err = models.SetPreparedDish(dishId, prepared)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database Error") {
		return nil
	}
	utils.RespondSuccess(w, http.StatusOK, "Dish Prepared Set")
	return r
}
