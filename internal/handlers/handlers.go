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

// Ayarlar
func (h *Handler) Settings(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.html", gin.H{
		"title":  "Ayarlar - Esnaf Yönetim Sistemi",
		"active": "settings",
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
