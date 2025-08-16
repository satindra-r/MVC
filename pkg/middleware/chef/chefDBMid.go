package chefMid

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
