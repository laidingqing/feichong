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
		Pattern:     "/users",
		HandlerFunc: handlers.GetUsers,
	})

	routes = append(routes, Route{
		Name:        "UserGetById",
		Method:      "GET",
		Pattern:     "/users/{userId}",
		HandlerFunc: handlers.GetUserByID,
	})

	routes = append(routes, Route{
		Name:        "UserCreate",
		Method:      "POST",
		Pattern:     "/users",
		HandlerFunc: handlers.PostUser,
	})

	// orders routes.
	routes = append(routes, Route{
		Name:        "OrdersGet",
		Method:      "GET",
		Pattern:     "/orders",
		HandlerFunc: handlers.GetOrders,
	})

	routes = append(routes, Route{
		Name:        "OrderGetById",
		Method:      "GET",
		Pattern:     "/orders/{orderId}",
		HandlerFunc: handlers.GetOrderByID,
	})

	routes = append(routes, Route{
		Name:        "OrderCreate",
		Method:      "POST",
		Pattern:     "/orders",
		HandlerFunc: handlers.PostOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderCreate",
		Method:      "PUT",
		Pattern:     "/orders/{orderId}/views/",
		HandlerFunc: handlers.PutOrderByViews,
	})

	routes = append(routes, Route{
		Name:        "OrderProfitCreate",
		Method:      "POST",
		Pattern:     "/orders/{orderId}/profits/",
		HandlerFunc: handlers.AddProfitInfoByOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderProfitGet",
		Method:      "GET",
		Pattern:     "/orders/{orderId}/profits/{month}/",
		HandlerFunc: handlers.GetProfitInfosByOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderTaxCreate",
		Method:      "POST",
		Pattern:     "/orders/{orderId}/tax/",
		HandlerFunc: handlers.AddTaxInfoByOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderTaxGet",
		Method:      "GET",
		Pattern:     "/orders/{orderId}/tax/{month}/",
		HandlerFunc: handlers.GetTaxInfosByOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderCaptialCreate",
		Method:      "POST",
		Pattern:     "/orders/{orderId}/captials/",
		HandlerFunc: handlers.AddCapitalInfoByOrder,
	})

	routes = append(routes, Route{
		Name:        "OrderCaptialGet",
		Method:      "GET",
		Pattern:     "/orders/{orderId}/captials/{month}/",
		HandlerFunc: handlers.GetCapitalInfosByOrder,
	})

	// weixin route
	routes = append(routes, Route{
		Name:        "WeixinLogin",
		Method:      "POST",
		Pattern:     "/weixin/login/",
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
	return router
}
