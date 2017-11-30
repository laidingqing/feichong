package handlers

import (
	"net/http"
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
			OrderNO:   order.OrderNO,
			OrderID:   order.ID.Hex(),
			Year:      orderYear,
			Month:     orderMonth,
			Catalog:   models.BusinessStatusUnknown, //没有状态变化
			CreatedAt: time.Now(),
			Seq:       i,
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

// GetBusinessByID 根据订单编号查询业务情况
func GetBusinessByID(w http.ResponseWriter, r *http.Request) {
	businessID := helpers.GetParam(r, businessIDParam)
	log := helpers.NewLogger()
	log.Log("businessID", businessID)
	business, err := managers.FindBusinessByID(businessID)
	if err != nil {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	} else {
		helpers.SetResponse(w, http.StatusOK, business)
	}
}

// GetBusinessByOrderID 根据订单编号查询业务情况
func GetBusinessByOrderID(w http.ResponseWriter, r *http.Request) {
	orderID := helpers.GetParam(r, busOrderIDParam)

	business, err := managers.FindOrderBusiness(orderID)
	if err != nil {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	} else {
		helpers.SetResponse(w, http.StatusOK, business)
	}
}

// PutProfitInfoByOrder 增加订单纳税情况
func PutProfitInfoByOrder(w http.ResponseWriter, r *http.Request) {
	businessID := helpers.GetParam(r, businessIDParam)
	var profit models.ProfitInfo

	profit.BusinessID = businessID

	helpers.GetProfitInfoBody(w, r, &profit)
	res, err := managers.UpdateProfitByBusiness(profit)

	if err == nil {
		helpers.SetResponse(w, http.StatusOK, res)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}
}

// PutCapitalInfoByOrder 增加订单资金情况
func PutCapitalInfoByOrder(w http.ResponseWriter, r *http.Request) {
	businessID := helpers.GetParam(r, businessIDParam)
	var capital models.CapitalInfo
	capital.BusinessID = businessID

	helpers.GetCapitalInfoBody(w, r, &capital)

	log := helpers.NewLogger()
	log.Log("BusinessID", capital.BusinessID)

	res, err := managers.UpdateCapitalByBusiness(capital)

	if err == nil {
		helpers.SetResponse(w, http.StatusOK, res)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}
}

// PutTaxInfoByOrder 增加订单资金情况
func PutTaxInfoByOrder(w http.ResponseWriter, r *http.Request) {
	businessID := helpers.GetParam(r, businessIDParam)
	var tax models.TaxInfo

	tax.BusinessID = businessID

	helpers.GetTaxInfoBody(w, r, &tax)
	res, err := managers.UpdateTaxByBusiness(tax)

	if err == nil {
		helpers.SetResponse(w, http.StatusOK, res)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}
}
