package user

import (
	"context"
	"mvc/pkg/utils"
	"net/http"
	"strings"
)

func AuthCheckUserCredentials(w http.ResponseWriter, r *http.Request) *http.Request {
	var hash = r.Context().Value("Hash").(string)
	var userId = r.Context().Value("UserId").(int)
	if userId == -1 {
		utils.RespondFailure(w, http.StatusUnauthorized, "Invalid Username or Password")
		return nil
	}

	var isCorrectPassword = CheckUserPassword(hash, r.Context().Value("Password").(string))
	if !isCorrectPassword {
		utils.RespondFailure(w, http.StatusUnauthorized, "Invalid Username or Password")
		return nil
	}

	var JWT = GenerateJWT(r.Context().Value("UserId").(int))
	utils.RespondSuccess(w, http.StatusOK, JWT)
	return r
}

func AuthVerifyUser(w http.ResponseWriter, r *http.Request) *http.Request {

	var cookies = r.Cookies()

	var JWT = ""
	for _, cookie := range cookies {
		var found bool
		JWT, found = strings.CutPrefix(cookie.Value, "JWT=")
		if found {
			break
		}

	}

	var UserId = JWTGetUserId(JWT)
	if UserId == -1 {
		//utils.RespondFailure(w, http.StatusUnauthorized, "Invalid or Expired JWT")
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil
	}

	r = r.WithContext(context.WithValue(r.Context(), "UserId", UserId))
	return r
}
