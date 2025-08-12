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
	utils.RespondSuccess(w, http.StatusOK, "Order Paid Set")
	return r
}

func DBCreateSection(w http.ResponseWriter, r *http.Request) *http.Request {
	var sectionName = r.Context().Value("SectionName").(string)
	var sectionId = models.GetNextSectionId()
	var sectionOrder = models.GetNextSectionOrder()
	var err error
	err = models.CreateSection(sectionId, sectionOrder, sectionName)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database error") {
		return nil
	}
	utils.RespondSuccess(w, http.StatusOK, "Section Created")
	return r
}
func DBSwapSections(w http.ResponseWriter, r *http.Request) *http.Request {
	var sectionId1 = r.Context().Value("SectionId1").(int)
	var sectionId2 = r.Context().Value("SectionId2").(int)
	var err error
	err = models.SwapSections(sectionId1, sectionId2)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database error") {
		return nil
	}
	utils.RespondSuccess(w, http.StatusOK, "Sections Swapped")
	return r
}
func DBDeleteSection(w http.ResponseWriter, r *http.Request) *http.Request {
	var sectionId = r.Context().Value("SectionId").(int)
	var err error
	err = models.DeleteSection(sectionId)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database error") {
		return nil
	}
	utils.RespondSuccess(w, http.StatusOK, "Section Deleted")
	return r
}
func DBSetUserRole(w http.ResponseWriter, r *http.Request) *http.Request {
	var userId = r.Context().Value("TargetUserId").(int)
	var role = r.Context().Value("TargetRole").(string)
	var err error
	err = models.SetUserRole(userId, role)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database error") {
		return nil
	}
	utils.RespondSuccess(w, http.StatusOK, "User Role Set")
	return r
}

func DBCreateItem(w http.ResponseWriter, r *http.Request) *http.Request {
	var item = r.Context().Value("Item").(models.Item)
	var err = models.CreateItem(item)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database Error") {
		return nil
	}
	utils.RespondSuccess(w, http.StatusOK, "Item Created")
	return r
}

func DBEditItem(w http.ResponseWriter, r *http.Request) *http.Request {
	var item = r.Context().Value("Item").(models.Item)
	var err = models.EditItem(item)
	if utils.ReflectAndLogErr(w, http.StatusInternalServerError, err, "Database Error") {
		return nil
	}
	utils.RespondSuccess(w, http.StatusOK, "Item Edited")
	return r
}
