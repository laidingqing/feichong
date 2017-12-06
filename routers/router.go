package routers

import (
	"net/http"

	"github.com/laidingqing/feichong/handlers"

	"github.com/gorilla/mux"
)

// Route ..
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes []Route

func init() {
	// users routes.
	routes = append(routes, Route{
		Name:        "UsersGet",
		Method:      "GET",
		Pattern:     "/api/users",
		HandlerFunc: handlers.GetUsers,
	})

	routes = append(routes, Route{
		Name:        "UserGetById",
		Method:      "GET",
		Pattern:     "/api/users/{userId}/",
		HandlerFunc: handlers.GetUserByID,
	})

	routes = append(routes, Route{
		Name:        "UserGetById",
		Method:      "GET",
		Pattern:     "/api/users/{userId}/orders",
		HandlerFunc: handlers.GetOrdersByUser,
	})

	routes = append(routes, Route{
		Name:        "PutUserSecurity",
		Method:      "PUT",
		Pattern:     "/api/users/{userId}/security",
		HandlerFunc: handlers.PutUserSecurity,
	})

	routes = append(routes, Route{
		Name:        "UserGetById",
		Method:      "GET",
		Pattern:     "/api/checkname/",
		HandlerFunc: handlers.CheckUserName,
	})

	routes = append(routes, Route{
		Name:        "UpdateUser",
		Method:      "PUT",
		Pattern:     "/api/users/{userId}/",
		HandlerFunc: handlers.PutUserByID,
	})

	routes = append(routes, Route{
		Name:        "UserCreate",
		Method:      "POST",
		Pattern:     "/api/users",
		HandlerFunc: handlers.PostUser,
	})

	routes = append(routes, Route{
		Name:        "UserCreate",
		Method:      "POST",
		Pattern:     "/api/session",
		HandlerFunc: handlers.LoginUser,
	})

	routes = append(routes, Route{
		Name:        "UserSelf",
		Method:      "GET",
		Pattern:     "/api/operators",
		HandlerFunc: handlers.SelfUsers,
	})

	routes = append(routes, Route{
		Name:        "UserSelf",
		Method:      "GET",
		Pattern:     "/api/companies",
		HandlerFunc: handlers.EnterPriseUsers,
	})

	// orders routes.
	routes = append(routes, Route{
		Name:        "OrdersGet",
		Method:      "GET",
		Pattern:     "/api/orders",
		HandlerFunc: handlers.GetOrders,
	})

	routes = append(routes, Route{
		Name:        "OrderGetById",
		Method:      "GET",
		Pattern:     "/api/orders/{orderId}",
		HandlerFunc: handlers.GetOrderByID,
	})

	routes = append(routes, Route{
		Name:        "OrderDeleteById",
		Method:      "DELETE",
		Pattern:     "/api/orders/{orderId}",
		HandlerFunc: handlers.DeleteOrderByID,
	})

	routes = append(routes, Route{
		Name:        "BusinessGetByOrder",
		Method:      "GET",
		Pattern:     "/api/orders/{orderId}/business/",
		HandlerFunc: handlers.GetBusinessByOrderID,
	})

	routes = append(routes, Route{
		Name:        "BusinessPostByOrder",
		Method:      "POST",
		Pattern:     "/api/orders/{orderId}/business/",
		HandlerFunc: handlers.CreateBusinessByOrderID,
	})

	routes = append(routes, Route{
		Name:        "BusinessDelByOrder",
		Method:      "DELETE",
		Pattern:     "/api/orders/{orderId}/business/{businessId}",
		HandlerFunc: handlers.DeleteBusinessByOrderID,
	})

	routes = append(routes, Route{
		Name:        "BusinessGetByBusID",
		Method:      "GET",
		Pattern:     "/api/orders/{orderId}/business/{businessId}",
		HandlerFunc: handlers.GetBusinessByID,
	})

	routes = append(routes, Route{
		Name:        "OrderCreate",
		Method:      "POST",
		Pattern:     "/api/orders/",
		HandlerFunc: handlers.PostOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderProfitUpdate",
		Method:      "PUT",
		Pattern:     "/api/business/{businessId}/profits/",
		HandlerFunc: handlers.PutProfitInfoByOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderTaxPut",
		Method:      "PUT",
		Pattern:     "/api/business/{businessId}/tax/",
		HandlerFunc: handlers.PutTaxInfoByOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderCaptialUpdate",
		Method:      "PUT",
		Pattern:     "/api/business/{businessId}/capitals/",
		HandlerFunc: handlers.PutCapitalInfoByOrder,
	})

	routes = append(routes, Route{
		Name:        "BusinessDataFind",
		Method:      "GET",
		Pattern:     "/api/business/{businessId}/view/",
		HandlerFunc: handlers.GetBusinessInfoByOrder,
	})

	routes = append(routes, Route{
		Name:        "BusinessDataFind",
		Method:      "GET",
		Pattern:     "/api/business/{businessId}/feedback/",
		HandlerFunc: handlers.PutFeedbackByBusiness,
	})
	// 咨询接口 查询
	routes = append(routes, Route{
		Name:        "ConsultCreate",
		Method:      "POST",
		Pattern:     "/api/consults",
		HandlerFunc: handlers.PostConsults,
	})

	// 咨询接口， 新建
	routes = append(routes, Route{
		Name:        "ConsultGet",
		Method:      "GET",
		Pattern:     "/api/consults",
		HandlerFunc: handlers.GetConsults,
	})

	// weixin route
	routes = append(routes, Route{
		Name:        "WeixinLogin",
		Method:      "POST",
		Pattern:     "/api/weixin/login",
		HandlerFunc: handlers.LoginSession,
	})

}

// NewRouter ...
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	return router
}
