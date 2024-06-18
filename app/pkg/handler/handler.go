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
		products.GET(":id/", h.findById)
	}

	warehouse := router.Group("/warehouse")
	{
		warehouse.GET("/", h.findAllWarehouse)
	}

	my := router.Group("/my", h.userIdentity)
	{
		address := my.Group("/address")
		{
			address.GET("/", h.getAllAddress)
			address.POST("/", h.createAddress)
			address.PUT("/")
			address.DELETE("/")
		}

		profile := my.Group("/profile")
		{
			profile.GET("/", h.profile)
			profile.PUT("/", h.updateUser)
			profile.POST("/")
		}

		order := my.Group("/order")
		{
			order.GET("/", h.viewOrders)
			order.POST("/", h.createOrder)
			order.PUT("/")
			order.DELETE("/")

			order.GET(":id/", h.viewOrder)
			order.PUT(":purchase/")
			order.DELETE(":purchase/")
		}

		cart := my.Group("/cart")
		{
			cart.GET("/", h.viewCart)
			cart.POST("/", h.addToCart)
			cart.PUT("/", h.scaleProduct)
			cart.DELETE("/", h.clearCart)
			cart.DELETE(":id/", h.removeProduct)
		}

		favorite := my.Group("/favorite")
		{
			favorite.GET("/", h.findAllFavorite)
			favorite.POST("/", h.addToFavorite)
			favorite.PUT("/")
			favorite.DELETE("/", h.removeFavorite)
		}

	}

	admin := router.Group("/admin", h.adminIdentity)
	{
		products := admin.Group("/products")
		{
			products.POST("/", h.createProduct)
			products.PUT(":id/", h.updateProduct)
			products.DELETE(":id/", h.deleteProduct)
		}

		productStack := admin.Group("/product-stack")
		{
			productStack.POST("/")
			productStack.PUT(":id/")
			productStack.DELETE(":id/")
		}

		orders := admin.Group("/orders")
		{
			orders.GET("/", h.adminOrders)
			orders.PUT(":id/", h.updateOrder)

			status := orders.Group("/status")
			{
				status.GET("/", h.orderStatus)
			}
		}

		image := admin.Group("/image")
		{
			image.POST("/", h.uploadImage)
		}

		auth := admin.Group("/sign-up")
		{
			auth.POST("/", h.createUserWithRole)
		}

		category := admin.Group("/category")
		{
			category.GET("/", h.findAllCategory)
			category.POST("/", h.createCategory)
		}

		manufacturer := admin.Group("/manufacturer")
		{
			manufacturer.GET("/", h.findAllManufacturer)
			manufacturer.POST("/", h.createManufacturer)
		}
	}

	return router
}
