package handlers

import (
	"net/http"
	"strconv"

	"github.com/laidingqing/feichong/helpers"
	"github.com/laidingqing/feichong/managers"
	"github.com/laidingqing/feichong/models"
)

const orderIDParam = "orderId"

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

// GetOrderByID ..
func GetOrderByID(w http.ResponseWriter, r *http.Request) {

	orderID := helpers.GetParam(r, orderIDParam)

	model, err := managers.GetOrderByID(orderID)

	if err != nil {
		helpers.SetResponse(w, http.StatusNotFound, nil)
	} else {
		helpers.SetResponse(w, http.StatusOK, model)
	}
}

// DeleteOrderByID ..
func DeleteOrderByID(w http.ResponseWriter, r *http.Request) {

	orderID := helpers.GetParam(r, orderIDParam)
	err := managers.DeleteOrderByID(orderID)

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
	log := helpers.NewLogger()
	log.Log("catalog", order.Catalog)
	if order.Catalog == 1 {
		//账务记账
		GeneratorBusinessFinData(order)
	}

	if order.Catalog == 2 {
		GeneratorBusinessEnterPriseData(order)
	}

	if err == nil {
		helpers.SetResponse(w, http.StatusOK, order)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}
}
