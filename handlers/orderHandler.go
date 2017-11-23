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

	pageIndex, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(size)
	catalogQuery, _ := strconv.Atoi(catalog)

	orders, err := managers.GetOrders(pageIndex, pageSize, catalogQuery)
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

// PostOrder create user
func PostOrder(w http.ResponseWriter, r *http.Request) {

	var order models.Order
	helpers.GetOrderBody(w, r, &order)

	id := managers.InsertOrder(order)

	if id != "" {
		helpers.SetResponse(w, http.StatusCreated, nil)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, nil)
	}
}

// PutOrderByViews 更新订单可见用户编号
func PutOrderByViews(w http.ResponseWriter, r *http.Request) {
	orderID := helpers.GetParam(r, orderIDParam)
	var order models.Order
	helpers.GetOrderBody(w, r, &order)

	oldOrder, err := managers.GetOrderByID(orderID)
	oldOrder.Views = order.Views
	if err != nil {
		helpers.SetResponse(w, http.StatusBadRequest, nil)
	}
	order, err = managers.PutOrder(orderID, oldOrder)
	if err != nil {
		helpers.SetResponse(w, http.StatusBadRequest, nil)
	}
	helpers.SetResponse(w, http.StatusCreated, nil)
}
