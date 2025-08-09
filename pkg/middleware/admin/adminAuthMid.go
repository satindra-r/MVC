package adminMid

import (
	"mvc/pkg/utils"
	"net/http"
)

func AuthVerifyAdmin(w http.ResponseWriter, r *http.Request) *http.Request {
	if r.Context().Value("Role") != "Admin" {
		utils.RespondFailure(w, http.StatusForbidden, "Forbidden")
		return nil
	}
	return r
}
