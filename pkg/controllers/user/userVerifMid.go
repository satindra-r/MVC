package user

import (
	"context"
	"encoding/json"
	"errors"
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
	"regexp"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

var phoneRegex = regexp.MustCompile(`^(?:\+[0-9]{1,2})?[0-9]{10}$`)

type Dishes struct {
	ItemId          int    `json:"itemId"`
	SplInstructions string `json:"splInstructions"`
	Count           int    `json:"count"`
}

type Order struct {
	Items []Dishes `json:"Items"`
}

func VerifyCreateUser(w http.ResponseWriter, r *http.Request) *http.Request {
	var hasErr = false
	var err error
	var user = models.User{}

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

	user.Address, hasErr = utils.GetOrReflect(w, r, "Address")
	if hasErr {
		return nil
	}

	password, hasErr := utils.GetOrReflect(w, r, "Password")
	if hasErr {
		return nil
	}

	if len(password) < 8 {
		utils.RespondFailure(w, http.StatusBadRequest, "Password must be at least 8 characters")
		return nil
	}

	user.Hash, err = GenerateHash(password)

	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		utils.RespondFailure(w, http.StatusBadRequest, "Password too long")
		return nil
	}

	r = r.WithContext(context.WithValue(r.Context(), "User", user))
	return r
}

func VerifyLogin(w http.ResponseWriter, r *http.Request) *http.Request {
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

func VerifyCreateOrder(w http.ResponseWriter, r *http.Request) *http.Request {
	var err error
	var order Order
	err = json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid Order")
		return nil
	}
	if order.Items == nil || len(order.Items) == 0 {
		utils.RespondFailure(w, http.StatusBadRequest, "Empty Order")
		return nil
	}
	var DBDishes []models.Dish
	var DBOrder = models.Order{}
	DBOrder.Price = 0

	var prices []float64
	prices, err = models.GetItemPrices()

	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database Error") {
		return nil
	}

	for _, item := range order.Items {

		if item.Count > 0 {
			var dish = models.Dish{}
			dish.OrderId = DBOrder.OrderId
			dish.ItemId = item.ItemId
			dish.DishCount = item.Count
			dish.SplInstructions = item.SplInstructions
			DBDishes = append(DBDishes, dish)
			DBOrder.Price += prices[item.ItemId] * (float64)(item.Count)
		}
	}

	r = r.WithContext(context.WithValue(r.Context(), "DBDishes", DBDishes))
	r = r.WithContext(context.WithValue(r.Context(), "DBOrder", DBOrder))

	return r
}

func VerifyCountDish(w http.ResponseWriter, r *http.Request) *http.Request {
	var dishIdStr string
	var countStr string
	var dishId int
	var count int
	var hasErr bool
	var err error

	dishIdStr, hasErr = utils.GetOrReflect(w, r, "dishId")
	if hasErr {
		return nil
	}

	dishId, err = strconv.Atoi(dishIdStr)

	if err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid dishId")
		return nil
	}

	countStr, hasErr = utils.GetOrReflect(w, r, "count")
	if hasErr {
		return nil
	}

	count, err = strconv.Atoi(countStr)

	if err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid count")
		return nil
	}

	count = max(-1, min(count, 1))

	r = r.WithContext(context.WithValue(r.Context(), "DishId", dishId))
	r = r.WithContext(context.WithValue(r.Context(), "Count", count))

	return r
}
