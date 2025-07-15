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
	r.GET("/products/detail/:id", h.ProductDetail)

	// Siparişler
	r.GET("/orders", h.Orders)
	r.GET("/orders/detail/:id", h.OrderDetail)

	// Stok Yönetimi
	r.GET("/stock", h.StockManagement)
	r.GET("/stock/movements", h.StockMovements)
	r.GET("/stock/barcode", h.BarcodeScanner)

	// Muhasebe
	r.GET("/accounting", h.Accounting)

	// Randevular
	r.GET("/appointments", h.Appointments)

	// Faturalar
	r.GET("/invoices", h.Invoices)

	// Raporlar
	r.GET("/reports", h.Reports)

	// Bildirimler
	r.GET("/notifications", h.Notifications)

	// Profil Sayfası
	r.GET("/profile", h.Profile)
	r.POST("/profile", h.UpdateProfile)

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

		// Stok Yönetimi API'leri
		api.POST("/stock/movements", h.AddStockMovementAPI)
		api.GET("/stock/products/barcode", h.GetProductByBarcodeAPI)
		api.GET("/stock/products/low", h.GetLowStockProductsAPI)
	}
}
