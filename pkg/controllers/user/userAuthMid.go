package user

import (
	"mvc/pkg/utils"
	"net/http"
)

func AuthCheckUserCredentials(w http.ResponseWriter, r *http.Request) *http.Request {
	var hash = r.Context().Value("Hash").(string)

	var isCorrectPassword = CheckUserPassword(hash, r.Context().Value("Password").(string))
	if !isCorrectPassword {
		utils.RespondFailure(w, http.StatusUnauthorized, "Invalid Username or Password")
		return nil
	}

	var JWT = GenerateJWT(r.Context().Value("UserId").(int))
	utils.RespondSuccess(w, http.StatusOK, JWT)
	return r
}
