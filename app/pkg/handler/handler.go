package handler

import (
	_ "github.com/SaRgEX/Diplom/docs"
	"github.com/SaRgEX/Diplom/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(CORSMiddleware())

	router.Static("/static", "../ui")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/logout", h.logout)
	}

	products := router.Group("/products")
	{
		products.GET("/", h.findAll)
		products.POST("/", h.createProduct)

		products.GET(":id/", h.findById)
		products.PUT(":id/", h.updateProduct)
		products.DELETE(":id/", h.deleteProduct)
	}

	category := router.Group("/category")
	{
		category.GET("/", h.findAllCategory)
		category.POST("/", h.createCategory)
	}

	manufacturer := router.Group("/manufacturer")
	{
		manufacturer.GET("/", h.findAllManufacturer)
		manufacturer.POST("/", h.createManufacturer)
	}

	warehouse := router.Group("/warehouse")
	{
		warehouse.GET("/", h.findAllWarehouse)
	}

	api := router.Group("/api", h.userIdentity)
	{
		address := api.Group("/address")
		{
			address.GET("/", h.getAllAddress)
			address.POST("/", h.createAddress)
			address.PUT("/")
			address.DELETE("/")
		}

		profile := api.Group("/profile")
		{
			profile.GET("/", h.profile)
			profile.PUT("/")
		}

		order := api.Group("/order")
		{
			order.GET("/", h.viewOrders)
			order.POST("/", h.createOrder)
			order.PUT("/")
			order.DELETE("/")

			order.GET(":id/", h.viewOrder)
			order.PUT(":purchase/")
			order.DELETE(":purchase/")
		}

		cart := api.Group("/cart")
		{
			cart.GET("/", h.viewCart)
			cart.POST("/", h.addToCart)
			cart.PUT("/", h.scaleProduct)
			cart.DELETE("/", h.clearCart)
			cart.DELETE(":id/", h.removeProduct)
		}

		favorite := api.Group("/favorite")
		{
			favorite.GET("/", h.findAllFavorite)
			favorite.POST("/", h.addToFavorite)
			favorite.PUT("/")
			favorite.DELETE("/", h.removeFavorite)
		}
	}
	return router
}
