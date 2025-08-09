package chefMid

import (
	"mvc/pkg/utils"
	"net/http"
)

func AuthVerifyChef(w http.ResponseWriter, r *http.Request) *http.Request {
	if r.Context().Value("Role") != "Chef" && r.Context().Value("Role") != "Admin" {
		utils.RespondFailure(w, http.StatusForbidden, "Forbidden")
		return nil
	}
	return r
}
