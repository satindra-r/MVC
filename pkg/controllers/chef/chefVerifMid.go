package chef

import (
	"context"
	"mvc/pkg/utils"
	"net/http"
	"strconv"
)

func VerifyPreparedDish(w http.ResponseWriter, r *http.Request) *http.Request {
	var dishIdStr string
	var preparedStr string
	var dishId int
	var prepared int
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

	preparedStr, hasErr = utils.GetOrReflect(w, r, "prepared")
	if hasErr {
		return nil
	}

	prepared, err = strconv.Atoi(preparedStr)

	if err != nil {
		utils.RespondFailure(w, http.StatusBadRequest, "Invalid prepared")
		return nil
	}

	prepared = max(0, min(prepared, 1))

	r = r.WithContext(context.WithValue(r.Context(), "DishId", dishId))
	r = r.WithContext(context.WithValue(r.Context(), "Prepared", prepared))

	return r
}
