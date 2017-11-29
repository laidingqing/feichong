package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/laidingqing/feichong/helpers"
	"github.com/laidingqing/feichong/managers"
	"github.com/laidingqing/feichong/models"
)

const busOrderIDParam = "orderId"
const businessIDParam = "businessId"
const businessMonthParam = "month"
const businessYearParam = "year"

// GeneratorBusinessEnterPriseData ..
func GeneratorBusinessEnterPriseData(order models.Order) {

}

// GeneratorBusinessFinData ..
func GeneratorBusinessFinData(order models.Order) {
	var year = int(order.CreatedAt.Year())
	var orderMonth = order.StartMonth
	var orderYear = int(order.CreatedAt.Year())
	var month = 0
	var next = false
	for i := 0; i < 12; i++ {
		business := models.Business{
			OrderID:   order.ID.Hex(),
			Year:      orderYear,
			Month:     orderMonth,
			Catalog:   models.BusinessStatusUnknown, //没有状态变化
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

// GetBusinessByOrderID 根据订单编号查询业务情况
func GetBusinessByOrderID(w http.ResponseWriter, r *http.Request) {
	orderID := helpers.GetParam(r, busOrderIDParam)
	year := helpers.GetInParam(r, businessYearParam)
	month := helpers.GetInParam(r, businessMonthParam)

	vYear, _ := strconv.Atoi(year)
	vMonth, _ := strconv.Atoi(month)

	business, err := managers.FindOrderBusiness(orderID, vYear, vMonth)
	if err != nil {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	} else {
		helpers.SetResponse(w, http.StatusOK, business)
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
	month := helpers.GetParam(r, businessMonthParam)
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
