package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/umutaraz/tradesman-app/internal/handlers"
)

func Setup(r *gin.Engine, h *handlers.Handler) {
	// Ana sayfa - Dashboard
	r.GET("/", h.Dashboard)
	r.GET("/dashboard", h.Dashboard)

	// Müşteriler
	r.GET("/customers", h.Customers)

	// Ürünler
	r.GET("/products", h.Products)

	// Siparişler
	r.GET("/orders", h.Orders)

	// Muhasebe
	r.GET("/accounting", h.Accounting)

	// Randevular
	r.GET("/appointments", h.Appointments)

	// Profil Sayfası
	r.GET("/profile", h.Profile)

	// Ayarlar Sayfası
	r.GET("/settings", h.Settings)

	// API Routes
	api := r.Group("/api/v1")
	{
		// Müşteri API'leri
		api.GET("/customers", h.GetCustomersAPI)
		api.POST("/customers", h.CreateCustomer)

		// Ürün API'leri
		// api.GET("/products", h.GetProductsAPI)
		// api.POST("/products", h.CreateProduct)

		// Sipariş API'leri
		// api.GET("/orders", h.GetOrdersAPI)
		// api.POST("/orders", h.CreateOrder)

		// Muhasebe API'leri
		// api.GET("/transactions", h.GetTransactionsAPI)
		// api.POST("/transactions", h.CreateTransaction)
	}
}
