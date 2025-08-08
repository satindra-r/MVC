package admin

import (
	"context"
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
	"strconv"
)

func VerifyPaidOrder(w http.ResponseWriter, r *http.Request) *http.Request {
	var orderIdStr string
	var paidStr string
	var orderId int
	var paid int
	var hasErr bool
	var err error

	orderIdStr, hasErr = utils.GetOrReflect(w, r, "orderId")
	if hasErr {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid orderId")
		return nil
	}

	orderId, err = strconv.Atoi(orderIdStr)

	if err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid orderId")
		return nil
	}

	paidStr, hasErr = utils.GetOrReflect(w, r, "paid")
	if hasErr {
		return nil
	}

	paid, err = strconv.Atoi(paidStr)

	if err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid paid")
	}

	paid = max(0, min(paid, 1))

	r = r.WithContext(context.WithValue(r.Context(), "OrderId", orderId))
	r = r.WithContext(context.WithValue(r.Context(), "Paid", paid))

	return r
}

func VerifySwapSections(w http.ResponseWriter, r *http.Request) *http.Request {
	var sectionId1Str string
	var sectionId2Str string
	var sectionId1 int
	var sectionId2 int
	var hasErr bool
	var err error

	sectionId1Str, hasErr = utils.GetOrReflect(w, r, "sectionId1")
	if hasErr {
		return nil
	}

	sectionId1, err = strconv.Atoi(sectionId1Str)

	if err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid sectionId1")
		return nil
	}

	sectionId2Str, hasErr = utils.GetOrReflect(w, r, "sectionId2")
	if hasErr {
		return nil
	}

	sectionId2, err = strconv.Atoi(sectionId2Str)

	if err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid sectionId2")
		return nil
	}

	r = r.WithContext(context.WithValue(r.Context(), "SectionId1", sectionId1))
	r = r.WithContext(context.WithValue(r.Context(), "SectionId2", sectionId2))

	return r
}

func VerifySetUserRole(w http.ResponseWriter, r *http.Request) *http.Request {
	var userIdStr string
	var role string
	var userId int
	var hasErr bool
	var err error

	userIdStr, hasErr = utils.GetOrReflect(w, r, "userId")
	if hasErr {
		return nil
	}

	userId, err = strconv.Atoi(userIdStr)

	if err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid userId")
		return nil
	}

	role, hasErr = utils.GetOrReflect(w, r, "role")
	if hasErr {
		return nil
	}

	if role != "User" && role != "Chef" && role != "Admin" {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid role")
		return nil
	}

	r = r.WithContext(context.WithValue(r.Context(), "TargetUserId", userId))
	r = r.WithContext(context.WithValue(r.Context(), "TargetRole", role))

	return r
}

func VerifyCreateItem(w http.ResponseWriter, r *http.Request) *http.Request {
	var item models.Item
	var sectionIdStr string
	var priceStr string

	var hasErr bool
	var err error

	item.ItemId = models.GetNextItemID()

	item.ItemName, hasErr = utils.GetOrReflect(w, r, "itemName")
	if hasErr {
		return nil
	}

	sectionIdStr, hasErr = utils.GetOrReflect(w, r, "sectionId")
	if hasErr {
		return nil
	}

	item.SectionId, err = strconv.Atoi(sectionIdStr)

	if err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid sectionId")
		return nil
	}

	priceStr, hasErr = utils.GetOrReflect(w, r, "price")
	if hasErr {
		return nil
	}
	item.Price, err = strconv.ParseFloat(priceStr, 64)
	if item.Price < 0 || err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid price")
		return nil
	}
	r = r.WithContext(context.WithValue(r.Context(), "Item", item))

	return r
}

func VerifyEditItem(w http.ResponseWriter, r *http.Request) *http.Request {
	var item models.Item
	var itemIdStr string
	var err error
	var hasErr bool

	itemIdStr, hasErr = utils.GetOrReflect(w, r, "itemId")
	if hasErr {
		return nil
	}

	item.ItemId, err = strconv.Atoi(itemIdStr)

	if err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid itemId")
		return nil
	}

	item.ItemName = r.Header.Get("itemName")

	item.SectionId, err = strconv.Atoi(r.Header.Get("sectionId"))
	if err != nil {
		item.SectionId = -1
	}

	item.Price, err = strconv.ParseFloat(r.Header.Get("price"), 64)
	if err != nil {
		item.Price = -1
	}

	r = r.WithContext(context.WithValue(r.Context(), "Item", item))

	return r
}
