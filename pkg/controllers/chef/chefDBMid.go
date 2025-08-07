package chef

import (
	"context"
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
)

func DBGetUserRole(w http.ResponseWriter, r *http.Request) *http.Request {
	var role, err = models.GetUserRole(r.Context().Value("UserId").(int))
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database Error") {
		return nil
	}
	r = r.WithContext(context.WithValue(r.Context(), "Role", role))
	return r
}

func DBSetPreparedDish(w http.ResponseWriter, r *http.Request) *http.Request {
	var dishId = r.Context().Value("DishId").(int)
	var prepared = r.Context().Value("Prepared").(int)
	var err = models.SetPreparedDish(dishId, prepared)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database Error") {
		return nil
	}
	return r
}
