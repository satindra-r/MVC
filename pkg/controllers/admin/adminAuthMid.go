package admin

import (
	"mvc/pkg/utils"
	"net/http"
)

func AuthDisallowDemote(w http.ResponseWriter, r *http.Request) *http.Request {
	if r.Context().Value("TargetUserId").(int) == r.Context().Value("UserId").(int) {
		utils.RespondFailure(w, http.StatusForbidden, "Cannot Demote Yourself")
		return nil
	}
	return r
}
