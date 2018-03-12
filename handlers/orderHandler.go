package handlers

import (
	"net/http"
	"strconv"

	"github.com/laidingqing/feichong/helpers"
	"github.com/laidingqing/feichong/managers"
	"github.com/laidingqing/feichong/models"
)

const orderIDParam = "orderId"
const orderNOParam = "orderNo"
const queryWayParam = "way"

// GetOrders ...
func GetOrders(w http.ResponseWriter, r *http.Request) {
	page := helpers.GetInParam(r, orderPageParam)
	size := helpers.GetInParam(r, orderSizeParam)
	catalog := helpers.GetInParam(r, orderCatalogParam)
	orderNO := helpers.GetInParam(r, orderNoParam)

	pageIndex, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(size)
	catalogQuery, _ := strconv.Atoi(catalog)
	var orders models.Pagination
	var err error
	if orderNO != "" {
		orders, err = managers.GetOrdersByNO(orderNO)
	} else {
		orders, err = managers.GetOrders(pageIndex, pageSize, catalogQuery)
	}

	if err == nil {
		helpers.SetResponse(w, http.StatusOK, orders)
	} else {
		helpers.SetResponse(w, http.StatusNotFound, nil)
	}
}

// GetOrderByID .. include way query string
func GetOrderByID(w http.ResponseWriter, r *http.Request) {

	way := helpers.GetParam(r, queryWayParam)
	orderID := helpers.GetParam(r, orderIDParam)
	userID := helpers.GetInParam(r, userIDParam)
	log := helpers.NewLogger()

	if way == "byId" {
		model, err := managers.GetOrderByID(orderID)
		if err != nil {
			helpers.SetResponse(w, http.StatusNotFound, nil)
		} else {
			helpers.SetResponse(w, http.StatusOK, model)
		}
	} else {
		model, err := managers.GetOrderByNo(orderID)
		if err != nil {
			log.Log("err", err)
			helpers.SetResponse(w, http.StatusNotFound, nil)
		} else {
			log.Log("model", model.ID.Hex(), "userid", userID)
			if userID != "" {
				user := managers.GetUserByID(userID)
				//绑定合同
				model.UserID = userID
				model.Company = user.CompanyName
				order, err := managers.PutOrder(model.ID.Hex(), model)
				if err != nil {
					log.Log("PutOrder err", err)
					helpers.SetResponse(w, http.StatusNotFound, nil)
					return
				}
				helpers.SetResponse(w, http.StatusOK, order)
			} else {
				helpers.SetResponse(w, http.StatusOK, model)
			}
		}
	}
}

// DeleteOrderByID ..
func DeleteOrderByID(w http.ResponseWriter, r *http.Request) {

	orderID := helpers.GetParam(r, orderIDParam)
	err := managers.DeleteOrderByID(orderID)
	log := helpers.NewLogger()
	log.Log("err", err)
	if err != nil {
		helpers.SetResponse(w, http.StatusNotFound, err)
	} else {
		helpers.SetResponse(w, http.StatusOK, orderID)
	}
}

// GetOrdersByUser ...
func GetOrdersByUser(w http.ResponseWriter, r *http.Request) {

	userID := helpers.GetParam(r, userIDParam)

	models, err := managers.GetOrdersByUser(userID)
	log := helpers.NewLogger()
	log.Log("user_id", userID)
	if err != nil {
		helpers.SetResponse(w, http.StatusNotFound, nil)
	} else {
		helpers.SetResponse(w, http.StatusOK, models)
	}
}

// PostOrder create user
func PostOrder(w http.ResponseWriter, r *http.Request) {

	var orderReq models.Order
	helpers.GetOrderBody(w, r, &orderReq)

	user := managers.GetUserByID(orderReq.UserInfo.ID.Hex())
	if user.CompanyName != "" {
		orderReq.Company = user.CompanyName
	}

	order, err := managers.InsertOrder(orderReq)
	if err == nil {
		helpers.SetResponse(w, http.StatusOK, order)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}
}
