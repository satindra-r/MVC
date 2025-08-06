package user

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
	"regexp"
)

var phoneRegex = regexp.MustCompile(`^(?:\+[0-9]{1,2})?[0-9]{10}$`)

func VerifCreateUser(w http.ResponseWriter, r *http.Request) *http.Request {
	var hasErr = false
	var err error
	var user = models.User{}

	user.UserId = models.GetNextUserID()

	user.UserName, hasErr = utils.GetOrReflect(w, r, "Username")
	if hasErr {
		return nil
	}

	user.Role, hasErr = utils.GetOrReflect(w, r, "Role")
	if hasErr {
		return nil
	}

	if user.Role != "User" && user.Role != "Chef" && user.Role != "Admin" {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid Role")
		return nil
	}

	user.PhoneNo, hasErr = utils.GetOrReflect(w, r, "PhoneNo")
	if hasErr {
		return nil
	}

	if !phoneRegex.Match([]byte(user.PhoneNo)) {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid PhoneNo")
		return nil
	}

	user.Address = r.FormValue("Address")

	password, hasErr := utils.GetOrReflect(w, r, "Password")
	if hasErr {
		return nil
	}

	if len(password) < 8 {
		utils.RespondFailure(w, http.StatusBadRequest, "Password must be at least 8 characters")
		return nil
	}

	user.Hash, err = GenerateHash(r.FormValue("Password"))

	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		utils.RespondFailure(w, http.StatusBadRequest, "Password too long")
		return nil
	}

	r = r.WithContext(context.WithValue(r.Context(), "User", user))
	return r
}

func VerifLogin(w http.ResponseWriter, r *http.Request) *http.Request {
	var hasErr bool
	var UserName string
	var Password string
	UserName, hasErr = utils.GetOrReflect(w, r, "Username")
	if hasErr {
		return nil
	}

	Password, hasErr = utils.GetOrReflect(w, r, "Password")
	if hasErr {
		return nil
	}

	r = r.WithContext(context.WithValue(r.Context(), "UserName", UserName))
	r = r.WithContext(context.WithValue(r.Context(), "Password", Password))
	return r
}
