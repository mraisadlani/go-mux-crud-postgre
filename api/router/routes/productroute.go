package routes

import (
	"github.com/vanilla/go-mux-postgre/api/controller"
	"net/http"
)

var productRoute = []Route{
	Route{
		URI: "/getproducts",
		Method: http.MethodGet,
		Handler: controller.GetAllProduct,
		AuthRequired: false,
	},
	Route{
		URI: "/get_product/{id}",
		Method: http.MethodGet,
		Handler: controller.GetProduct,
		AuthRequired: false,
	},
	Route{
		URI: "/create_product",
		Method: http.MethodPost,
		Handler: controller.CreateProduct,
		AuthRequired: false,
	},
	Route{
		URI: "/update_product/{id}",
		Method: http.MethodPut,
		Handler: controller.UpdateProduct,
		AuthRequired: false,
	},
	Route{
		URI: "/delete_product/{id}",
		Method: http.MethodDelete,
		Handler: controller.DeleteProduct,
		AuthRequired: false,
	},
}