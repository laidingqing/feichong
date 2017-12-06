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
	var orderMonth = int(order.StartAt.Month())
	var orderYear = int(order.CreatedAt.Year())

	var month = 0
	var next = false
	for i := 0; i < order.OrderMonth; i++ {
		business := models.Business{
			OrderNO:   order.OrderNO,
			OrderID:   order.ID.Hex(),
			Year:      orderYear,
			Month:     orderMonth,
			CreatedAt: time.Now(),
			Seq:       i,
		}

		managers.InsertBusinessData(business)

		if orderMonth < order.OrderMonth {
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

// CreateBusinessByOrderID ...
func CreateBusinessByOrderID(w http.ResponseWriter, r *http.Request) {
	orderID := helpers.GetParam(r, busOrderIDParam)
	var business models.Business

	helpers.GetBusinessBody(w, r, &business)
	business.OrderID = orderID
	business, err := managers.InsertBusinessData(business)

	if err != nil {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	} else {
		helpers.SetResponse(w, http.StatusOK, business)
	}
}

// DeleteBusinessByOrderID ..
func DeleteBusinessByOrderID(w http.ResponseWriter, r *http.Request) {
	businessID := helpers.GetParam(r, businessIDParam)
	err := managers.DeleteBusiness(businessID)

	if err != nil {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	} else {
		helpers.SetResponse(w, http.StatusOK, nil)
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
	tax.Reported = true

	helpers.GetTaxInfoBody(w, r, &tax)
	res, err := managers.UpdateTaxByBusiness(tax)

	if err == nil {
		helpers.SetResponse(w, http.StatusOK, res)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}
}

// GetBusinessInfoByOrder ...
func GetBusinessInfoByOrder(w http.ResponseWriter, r *http.Request) {
	businessID := helpers.GetParam(r, businessIDParam)
	year := helpers.GetInParam(r, businessYearParam)
	month := helpers.GetInParam(r, businessMonthParam)

	vYear, _ := strconv.Atoi(year)
	vMonth, _ := strconv.Atoi(month)

	res, err := managers.FindOrderBusinessByDate(businessID, vYear, vMonth)
	log := helpers.NewLogger()
	log.Log("err", err)
	if err == nil {
		helpers.SetResponse(w, http.StatusOK, res)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}
}

// PutFeedbackByBusiness ...
func PutFeedbackByBusiness(w http.ResponseWriter, r *http.Request) {
	businessID := helpers.GetParam(r, businessIDParam)
	var fb models.FeedBack
	helpers.GetFeedbackBody(w, r, &fb)

	res, err := managers.UpdateFeedbackBusiness(businessID, fb)

	if err == nil {
		helpers.SetResponse(w, http.StatusOK, res)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}
}

// GetConsults ..
func GetConsults(w http.ResponseWriter, r *http.Request) {
	page := helpers.GetInParam(r, orderPageParam)
	size := helpers.GetInParam(r, orderSizeParam)
	pageIndex, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(size)
	res, err := managers.GetFeedbacks(pageIndex, pageSize)

	if err == nil {
		helpers.SetResponse(w, http.StatusOK, res)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}

}

// PostConsults ...
func PostConsults(w http.ResponseWriter, r *http.Request) {
	var consult models.Consult
	helpers.GetSiteFeedbackBody(w, r, &consult)
	resp, err := managers.PostFeedbacks(consult)
	if err == nil {
		helpers.SetResponse(w, http.StatusOK, resp)
	} else {
		helpers.SetResponse(w, http.StatusBadRequest, err)
	}
}
