package user

import (
	"context"
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
	"strconv"
)

func DBCreateUser(w http.ResponseWriter, r *http.Request) *http.Request {
	var user models.User
	var err error

	user = (r.Context().Value("User")).(models.User)

	var hash string
	hash, _ = models.GetUserCredentials(user.UserName)

	if len(hash) != 0 {
		utils.RespondFailure(w, http.StatusForbidden, "Username already taken")
		return nil
	}

	err = models.CreateUser(user)
	if err != nil {
		return nil
	}

	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database Error") {
		return nil
	}

	utils.RespondSuccess(w, http.StatusCreated, strconv.Itoa(user.UserId))
	return r
}

func DBGetUserCredentials(w http.ResponseWriter, r *http.Request) *http.Request {
	var hash, userId = models.GetUserCredentials(r.Context().Value("UserName").(string))

	r = r.WithContext(context.WithValue(r.Context(), "Hash", hash))
	r = r.WithContext(context.WithValue(r.Context(), "UserId", userId))

	return r
}

func DBCreateOrder(w http.ResponseWriter, r *http.Request) *http.Request {

	var DBDishes = r.Context().Value("DBDishes").([]models.Dish)
	var DBOrder = r.Context().Value("DBOrder").(models.Order)

	DBOrder.UserId = r.Context().Value("UserId").(int)
	var err = models.CreateOrder(DBOrder)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database Error") {
		return nil
	}

	for _, dish := range DBDishes {
		var err = models.CreateDish(dish)
		if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database Error") {
			return nil
		}
	}

	utils.RespondSuccess(w, http.StatusCreated, "Order Created")
	return r
}

func DBSetCountDish(w http.ResponseWriter, r *http.Request) *http.Request {
	var dishId = r.Context().Value("DishId").(int)
	var count = r.Context().Value("Count").(int)
	err := models.EditDishCount(dishId, count)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database Error") {
		return nil
	}
	utils.RespondSuccess(w, http.StatusOK, "Dish Count Set")
	return r
}
