package admin

import (
	"context"
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
		return nil
	}

	orderId, err = strconv.Atoi(orderIdStr)

	if err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid orderId")
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
