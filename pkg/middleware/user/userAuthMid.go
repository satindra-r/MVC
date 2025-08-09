package userMid

import (
	"context"
	"mvc/pkg/controllers/user"
	"net/http"
	"strings"
)

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

	var UserId = user.JWTGetUserId(JWT)
	if UserId == -1 {
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil
	}

	r = r.WithContext(context.WithValue(r.Context(), "UserId", UserId))
	return r
}
