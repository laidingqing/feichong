package handlers

import (
	"net/http"
	"strconv"

	"github.com/laidingqing/feichong/helpers"
	"github.com/laidingqing/feichong/managers"
	"github.com/laidingqing/feichong/models"
)

const busOrderIDParam = "orderId"
const businessIDParam = "businessId"
const monthParam = "month"

// GetBusinessByOrderID 根据订单编号查询业务情况
func GetBusinessByOrderID(w http.ResponseWriter, r *http.Request) {
	orderID := helpers.GetParam(r, busOrderIDParam)
	if true {
		helpers.SetResponse(w, http.StatusOK, nil)
	} else {
		helpers.SetResponse(w, http.StatusNotFound, orderID)
	}
}

// GetCapitalInfosByOrder 根据订单号查询指定月份资金情况
func GetCapitalInfosByOrder(w http.ResponseWriter, r *http.Request) {
	orderID := helpers.GetParam(r, busOrderIDParam)
	if true {
		helpers.SetResponse(w, http.StatusOK, nil)
	} else {
		helpers.SetResponse(w, http.StatusNotFound, orderID)
	}
}

// GetTaxInfosByOrder 根据订单号查询指定月份纳税情况
func GetTaxInfosByOrder(w http.ResponseWriter, r *http.Request) {
	orderID := helpers.GetParam(r, busOrderIDParam)
	month := helpers.GetParam(r, monthParam)
	vMonth, err := strconv.Atoi(month)

	if err != nil {
		helpers.SetResponse(w, http.StatusNotFound, err)
	}

	tax, err := managers.GetOrderTaxs(orderID, vMonth)

	if err != nil {
		helpers.SetResponse(w, http.StatusNotFound, err)
	} else {
		helpers.SetResponse(w, http.StatusOK, tax)
	}
}

// GetProfitInfosByOrder 根据订单号查询指定月份利润情况
func GetProfitInfosByOrder(w http.ResponseWriter, r *http.Request) {
	orderID := helpers.GetParam(r, busOrderIDParam)
	if true {
		helpers.SetResponse(w, http.StatusOK, nil)
	} else {
		helpers.SetResponse(w, http.StatusNotFound, orderID)
	}
}

// AddProfitInfoByOrder 增加订单纳税情况
func AddProfitInfoByOrder(w http.ResponseWriter, r *http.Request) {
	orderID := helpers.GetParam(r, busOrderIDParam)
	var profit models.ProfitInfo
	helpers.GetProfitInfoBody(w, r, &profit)
	if true {
		helpers.SetResponse(w, http.StatusOK, nil)
	} else {
		helpers.SetResponse(w, http.StatusNotFound, orderID)
	}
}

// AddCapitalInfoByOrder 增加订单资金情况
func AddCapitalInfoByOrder(w http.ResponseWriter, r *http.Request) {
	orderID := helpers.GetParam(r, busOrderIDParam)

	var capital models.CapitalInfo

	helpers.GetCapitalInfoBody(w, r, &capital)
	capital.OrderID = orderID

	if true {
		helpers.SetResponse(w, http.StatusOK, nil)
	} else {
		helpers.SetResponse(w, http.StatusNotFound, orderID)
	}
}

// AddTaxInfoByOrder 增加订单资金情况
func AddTaxInfoByOrder(w http.ResponseWriter, r *http.Request) {
	orderID := helpers.GetParam(r, busOrderIDParam)
	var tax models.TaxInfo
	tax.OrderID = orderID
	helpers.GetTaxInfoBody(w, r, &tax)

	id := managers.InsertOrderTax(tax)

	if true {
		helpers.SetResponse(w, http.StatusOK, id)
	} else {
		helpers.SetResponse(w, http.StatusNotFound, id)
	}
}
