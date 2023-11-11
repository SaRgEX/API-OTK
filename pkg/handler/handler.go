package handler

import (
	"github.com/SaRgEX/Diplom/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	products := router.Group("/products")
	{
		products.GET("/", h.findAll)
		products.POST("/", h.createProduct)

		products.GET(":id/", h.findById)
		products.PUT(":id/", h.updateProduct)
		products.DELETE(":id/", h.deleteProduct)
	}

	api := router.Group("/api", h.userIdentity)
	{
		address := api.Group("/address")
		{
			address.GET("/")
			address.POST("/create", h.createAddress)
			address.PUT("/")
			address.DELETE("/")
		}

		order := api.Group("/order")
		{
			order.GET("/", h.viewOrders)
			order.POST("/", h.createOrder)
			order.PUT("/")
			order.DELETE("/")

			order.GET(":purchase/")
			order.PUT(":purchase/")
			order.DELETE(":purchase/")
		}
	}
	return router
}
