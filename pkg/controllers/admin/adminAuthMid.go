package admin

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

func AuthDisallowDemote(w http.ResponseWriter, r *http.Request) *http.Request {
	if r.Context().Value("TargetUserId").(int) == r.Context().Value("UserId").(int) {
		utils.RespondFailure(w, http.StatusForbidden, "Cannot Demote Yourself")
		return nil
	}
	return r
}
