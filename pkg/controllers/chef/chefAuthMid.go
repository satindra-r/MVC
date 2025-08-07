package chef

import (
	"fmt"
	"mvc/pkg/utils"
	"net/http"
)

func AuthVerifyChef(w http.ResponseWriter, r *http.Request) *http.Request {
	if r.Context().Value("Role") != "Chef" && r.Context().Value("Role") != "Admin" {
		utils.RespondFailure(w, http.StatusForbidden, "Forbidden")
		fmt.Println(r.Context().Value("Role"))
		return nil
	}
	return r
}
