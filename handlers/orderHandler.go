package handlers

import (
	"net/http"
	"strconv"
	"time"

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
		log.Log("description", "账务记账业务记录")
		var year = int(order.CreatedAt.Year())
		var orderMonth = order.StartMonth
		var orderYear = int(order.CreatedAt.Year())
		var month = 0
		var next = false
		for i := 0; i < 12; i++ {
			log.Log("end", orderMonth+month, "orderMonth", orderMonth, "m", month, "orderYear", orderYear)

			business := models.Business{
				OrderID:   order.ID.Hex(),
				Year:      orderYear,
				Month:     orderMonth,
				Catalog:   0, //没有状态变化
				CreatedAt: time.Now(),
			}

			managers.InsertBusinessData(business)

			if orderMonth < 12 {
				orderMonth++
				month = i
			} else {
				orderMonth = 1
				if next == false {
					next = true
				}
				month++
			}

			if next {
				orderYear = year + 1
			}
		}

	}

	if err == nil {
		helpers.SetResponse(w, http.StatusOK, order)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
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
