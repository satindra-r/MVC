package admin

import (
	"mvc/pkg/models"
	"mvc/pkg/utils"
	"net/http"
)

func DBSetPaidOrder(w http.ResponseWriter, r *http.Request) *http.Request {
	var orderId = r.Context().Value("OrderId").(int)
	var paid = r.Context().Value("Paid").(int)
	var err = models.SetPaidOrder(orderId, paid)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database Error") {
		return nil
	}
	return r
}
