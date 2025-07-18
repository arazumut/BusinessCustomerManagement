package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/umutaraz/tradesman-app/internal/database"
	"github.com/umutaraz/tradesman-app/internal/models"
)

type Handler struct {
	db *database.DB
}

func New(db *database.DB) *Handler {
	return &Handler{db: db}
}

// Dashboard
func (h *Handler) Dashboard(c *gin.Context) {
	stats, err := h.getDashboardStats()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"stats":  stats,
		"title":  "Dashboard - Esnaf Yönetim Sistemi",
		"active": "dashboard",
	})
}

// Müşteriler
func (h *Handler) Customers(c *gin.Context) {
	customers, err := h.getCustomers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "customers.html", gin.H{
		"customers": customers,
		"title":     "Müşteriler - Esnaf Yönetim Sistemi",
		"active":    "customers",
	})
}

// Ürünler
func (h *Handler) Products(c *gin.Context) {
	products, err := h.getProducts()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "products.html", gin.H{
		"products": products,
		"title":    "Ürünler - Esnaf Yönetim Sistemi",
		"active":   "products",
	})
}

// Siparişler
func (h *Handler) Orders(c *gin.Context) {
	orders, err := h.getOrders()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "orders.html", gin.H{
		"orders": orders,
		"title":  "Siparişler - Esnaf Yönetim Sistemi",
		"active": "orders",
	})
}

// Muhasebe
func (h *Handler) Accounting(c *gin.Context) {
	transactions, err := h.getTransactions()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "accounting.html", gin.H{
		"transactions": transactions,
		"title":        "Muhasebe - Esnaf Yönetim Sistemi",
		"active":       "accounting",
	})
}

// Randevular
func (h *Handler) Appointments(c *gin.Context) {
	// İleride randevu bilgileri için getAppointments() fonksiyonu oluşturulabilir
	c.HTML(http.StatusOK, "appointments.html", gin.H{
		"title":  "Randevular - Esnaf Yönetim Sistemi",
		"active": "appointments",
	})
}

// Profil
func (h *Handler) Profile(c *gin.Context) {
	// Profil için kullanıcı bilgilerini al
	// İleride gerçek kullanıcı bilgileri için getUserProfile() fonksiyonu oluşturulabilir
	user := models.User{
		ID:           1,
		Name:         "Ahmet Yılmaz",
		Email:        "ahmet@esnaf.com",
		Role:         "Yönetici",
		BusinessName: "Yılmaz Ticaret",
		Phone:        "+90 (555) 123 4567",
		Address:      "Atatürk Cad. No:123 İstanbul",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"user":   user,
		"title":  "Profil - Esnaf Yönetim Sistemi",
		"active": "profile",
	})
}

// Profil Güncelle
func (h *Handler) UpdateProfile(c *gin.Context) {
	// Form verilerini al
	name := c.PostForm("name")
	businessName := c.PostForm("business_name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	address := c.PostForm("address")

	// Profil güncelleme işlemi burada yapılacak
	// Şimdilik sadece başarı mesajı döndürelim

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Profil başarıyla güncellendi",
		"data": gin.H{
			"name":          name,
			"business_name": businessName,
			"email":         email,
			"phone":         phone,
			"address":       address,
		},
	})
}

// Ayarlar
func (h *Handler) Settings(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.html", gin.H{
		"title":  "Ayarlar - Esnaf Yönetim Sistemi",
		"active": "settings",
	})
}

// Faturalar
func (h *Handler) Invoices(c *gin.Context) {
	// İleride fatura bilgileri için getInvoices() fonksiyonu oluşturulabilir
	c.HTML(http.StatusOK, "invoices.html", gin.H{
		"title":  "Faturalar - Esnaf Yönetim Sistemi",
		"active": "invoices",
	})
}

// Raporlar
func (h *Handler) Reports(c *gin.Context) {
	// İleride rapor bilgileri için getReports() fonksiyonu oluşturulabilir
	c.HTML(http.StatusOK, "reports.html", gin.H{
		"title":  "Raporlar - Esnaf Yönetim Sistemi",
		"active": "reports",
	})
}

// Bildirimler
func (h *Handler) Notifications(c *gin.Context) {
	// İleride bildirim bilgileri için getNotifications() fonksiyonu oluşturulabilir
	c.HTML(http.StatusOK, "notifications.html", gin.H{
		"title":  "Bildirimler - Esnaf Yönetim Sistemi",
		"active": "notifications",
	})
}

// API Endpoints
func (h *Handler) GetCustomersAPI(c *gin.Context) {
	customers, err := h.getCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *Handler) CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.UserID = 1 // Şimdilik sabit user ID
	id, err := h.insertCustomer(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customer.ID = id
	c.JSON(http.StatusCreated, customer)
}

// Database helper methods
func (h *Handler) getDashboardStats() (*models.DashboardStats, error) {
	stats := &models.DashboardStats{}

	// Toplam müşteri sayısı
	err := h.db.QueryRow("SELECT COUNT(*) FROM customers WHERE user_id = ?", 1).Scan(&stats.TotalCustomers)
	if err != nil {
		return nil, err
	}

	// Toplam ürün sayısı
	err = h.db.QueryRow("SELECT COUNT(*) FROM products WHERE user_id = ?", 1).Scan(&stats.TotalProducts)
	if err != nil {
		return nil, err
	}

	// Toplam sipariş sayısı
	err = h.db.QueryRow("SELECT COUNT(*) FROM orders WHERE user_id = ?", 1).Scan(&stats.TotalOrders)
	if err != nil {
		return nil, err
	}

	// Bekleyen sipariş sayısı
	err = h.db.QueryRow("SELECT COUNT(*) FROM orders WHERE user_id = ? AND status = 'pending'", 1).Scan(&stats.PendingOrders)
	if err != nil {
		return nil, err
	}

	// Bu ay gelir
	err = h.db.QueryRow(`
		SELECT COALESCE(SUM(amount), 0) 
		FROM transactions 
		WHERE user_id = ? AND type = 'income' 
		AND strftime('%Y-%m', transaction_date) = strftime('%Y-%m', 'now')
	`, 1).Scan(&stats.MonthlyRevenue)
	if err != nil {
		return nil, err
	}

	// Bu ay gider
	err = h.db.QueryRow(`
		SELECT COALESCE(SUM(amount), 0) 
		FROM transactions 
		WHERE user_id = ? AND type = 'expense' 
		AND strftime('%Y-%m', transaction_date) = strftime('%Y-%m', 'now')
	`, 1).Scan(&stats.MonthlyExpenses)
	if err != nil {
		return nil, err
	}

	stats.MonthlyProfit = stats.MonthlyRevenue - stats.MonthlyExpenses

	return stats, nil
}

func (h *Handler) getCustomers() ([]models.Customer, error) {
	rows, err := h.db.Query("SELECT * FROM customers WHERE user_id = ? ORDER BY created_at DESC", 1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.ID, &customer.UserID, &customer.Name, &customer.Email,
			&customer.Phone, &customer.Address, &customer.Notes, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (h *Handler) getProducts() ([]models.Product, error) {
	rows, err := h.db.Query("SELECT * FROM products WHERE user_id = ? ORDER BY created_at DESC", 1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.UserID, &product.Name, &product.Description,
			&product.Price, &product.Category, &product.StockQuantity, &product.Unit,
			&product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (h *Handler) getOrders() ([]models.Order, error) {
	rows, err := h.db.Query(`
		SELECT o.*, c.name as customer_name 
		FROM orders o 
		JOIN customers c ON o.customer_id = c.id 
		WHERE o.user_id = ? 
		ORDER BY o.created_at DESC
	`, 1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		var customerName string
		err := rows.Scan(&order.ID, &order.UserID, &order.CustomerID, &order.OrderNumber,
			&order.Status, &order.TotalAmount, &order.Notes, &order.OrderDate,
			&order.DeliveryDate, &order.CreatedAt, &order.UpdatedAt, &customerName)
		if err != nil {
			return nil, err
		}
		order.Customer = &models.Customer{Name: customerName}
		orders = append(orders, order)
	}

	return orders, nil
}

func (h *Handler) getTransactions() ([]models.Transaction, error) {
	rows, err := h.db.Query("SELECT * FROM transactions WHERE user_id = ? ORDER BY transaction_date DESC", 1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Type,
			&transaction.Category, &transaction.Amount, &transaction.Description,
			&transaction.TransactionDate, &transaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (h *Handler) insertCustomer(customer *models.Customer) (int, error) {
	result, err := h.db.Exec(`
		INSERT INTO customers (user_id, name, email, phone, address, notes)
		VALUES (?, ?, ?, ?, ?, ?)
	`, customer.UserID, customer.Name, customer.Email, customer.Phone, customer.Address, customer.Notes)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Ürün Detayı
func (h *Handler) ProductDetail(c *gin.Context) {
	id := c.Param("id")

	// Ürün detayını veritabanından al
	var product models.Product
	err := h.db.QueryRow(`
		SELECT id, user_id, name, description, price, category, stock_quantity, unit, created_at, updated_at
		FROM products WHERE id = ? AND user_id = ?
	`, id, 1).Scan(&product.ID, &product.UserID, &product.Name, &product.Description,
		&product.Price, &product.Category, &product.StockQuantity, &product.Unit,
		&product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"error": "Ürün bulunamadı"})
		return
	}

	c.HTML(http.StatusOK, "product_detail.html", gin.H{
		"product": product,
		"title":   "Ürün Detayı - " + product.Name,
		"active":  "products",
	})
}

// Sipariş Detayı
func (h *Handler) OrderDetail(c *gin.Context) {
	id := c.Param("id")

	// Sipariş detayını veritabanından al
	var order models.Order
	err := h.db.QueryRow(`
		SELECT o.id, o.user_id, o.customer_id, o.order_number, o.status, o.total_amount, 
		       o.notes, o.order_date, o.delivery_date, o.created_at, o.updated_at,
		       c.name as customer_name, c.email as customer_email, c.phone as customer_phone
		FROM orders o 
		JOIN customers c ON o.customer_id = c.id 
		WHERE o.id = ? AND o.user_id = ?
	`, id, 1).Scan(&order.ID, &order.UserID, &order.CustomerID, &order.OrderNumber,
		&order.Status, &order.TotalAmount, &order.Notes, &order.OrderDate,
		&order.DeliveryDate, &order.CreatedAt, &order.UpdatedAt,
		&order.Customer.Name, &order.Customer.Email, &order.Customer.Phone)

	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"error": "Sipariş bulunamadı"})
		return
	}

	// Sipariş kalemlerini al
	items, err := h.getOrderItems(order.ID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}
	order.Items = items

	c.HTML(http.StatusOK, "order_detail.html", gin.H{
		"order":  order,
		"title":  "Sipariş Detayı - " + order.OrderNumber,
		"active": "orders",
	})
}

// Sipariş kalemlerini getir
func (h *Handler) getOrderItems(orderID int) ([]models.OrderItem, error) {
	rows, err := h.db.Query(`
		SELECT oi.id, oi.order_id, oi.product_id, oi.quantity, oi.unit_price, oi.total_price,
		       p.name as product_name, p.unit as product_unit
		FROM order_items oi
		JOIN products p ON oi.product_id = p.id
		WHERE oi.order_id = ?
	`, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.OrderItem
	for rows.Next() {
		var item models.OrderItem
		var productName, productUnit string
		err := rows.Scan(&item.ID, &item.OrderID, &item.ProductID, &item.Quantity,
			&item.UnitPrice, &item.TotalPrice, &productName, &productUnit)
		if err != nil {
			return nil, err
		}
		item.Product = &models.Product{Name: productName, Unit: productUnit}
		items = append(items, item)
	}

	return items, nil
}
