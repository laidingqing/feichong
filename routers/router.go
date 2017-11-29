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
		Name:        "BusinessGetByOrder",
		Method:      "GET",
		Pattern:     "/api/orders/{orderId}/business/",
		HandlerFunc: handlers.GetBusinessByOrderID,
	})

	routes = append(routes, Route{
		Name:        "OrderCreate",
		Method:      "POST",
		Pattern:     "/api/orders/",
		HandlerFunc: handlers.PostOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderUpdateView",
		Method:      "PUT",
		Pattern:     "/api/orders/{orderId}/views/",
		HandlerFunc: handlers.PutOrderByViews,
	})

	routes = append(routes, Route{
		Name:        "OrderProfitCreate",
		Method:      "POST",
		Pattern:     "/api/orders/{orderId}/profits/",
		HandlerFunc: handlers.AddProfitInfoByOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderProfitGet",
		Method:      "GET",
		Pattern:     "/api/orders/{orderId}/profits/{month}/",
		HandlerFunc: handlers.GetProfitInfosByOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderTaxCreate",
		Method:      "POST",
		Pattern:     "/api/orders/{orderId}/tax",
		HandlerFunc: handlers.AddTaxInfoByOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderTaxGet",
		Method:      "GET",
		Pattern:     "/api/orders/{orderId}/tax/{month}",
		HandlerFunc: handlers.GetTaxInfosByOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderCaptialCreate",
		Method:      "POST",
		Pattern:     "/api/orders/{orderId}/captials",
		HandlerFunc: handlers.AddCapitalInfoByOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderCaptialGet",
		Method:      "GET",
		Pattern:     "/api/orders/{orderId}/captials/{month}",
		HandlerFunc: handlers.GetCapitalInfosByOrder,
	})

	// 咨询接口
	routes = append(routes, Route{
		Name:        "ConsultCreate",
		Method:      "POST",
		Pattern:     "/api/consults",
		HandlerFunc: handlers.AddCapitalInfoByOrder,
	})

	routes = append(routes, Route{
		Name:        "ConsultGet",
		Method:      "GET",
		Pattern:     "/api/consults",
		HandlerFunc: handlers.GetCapitalInfosByOrder,
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
