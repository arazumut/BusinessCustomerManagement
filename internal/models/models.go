package models

import "time"

type User struct {
	ID           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Role         string    `json:"role" db:"role"`
	BusinessName string    `json:"business_name" db:"business_name"`
	Phone        string    `json:"phone" db:"phone"`
	Address      string    `json:"address" db:"address"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type Customer struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	Address   string    `json:"address" db:"address"`
	Notes     string    `json:"notes" db:"notes"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Product struct {
	ID            int       `json:"id" db:"id"`
	UserID        int       `json:"user_id" db:"user_id"`
	Name          string    `json:"name" db:"name"`
	Description   string    `json:"description" db:"description"`
	Price         float64   `json:"price" db:"price"`
	Category      string    `json:"category" db:"category"`
	StockQuantity int       `json:"stock_quantity" db:"stock_quantity"`
	Unit          string    `json:"unit" db:"unit"`
	Barcode       string    `json:"barcode" db:"barcode"`
	MinStockLevel int       `json:"min_stock_level" db:"min_stock_level"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type Order struct {
	ID           int         `json:"id" db:"id"`
	UserID       int         `json:"user_id" db:"user_id"`
	CustomerID   int         `json:"customer_id" db:"customer_id"`
	OrderNumber  string      `json:"order_number" db:"order_number"`
	Status       string      `json:"status" db:"status"`
	TotalAmount  float64     `json:"total_amount" db:"total_amount"`
	Notes        string      `json:"notes" db:"notes"`
	OrderDate    time.Time   `json:"order_date" db:"order_date"`
	DeliveryDate *time.Time  `json:"delivery_date" db:"delivery_date"`
	CreatedAt    time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at" db:"updated_at"`
	Customer     *Customer   `json:"customer,omitempty"`
	Items        []OrderItem `json:"items,omitempty"`
}

type OrderItem struct {
	ID         int      `json:"id" db:"id"`
	OrderID    int      `json:"order_id" db:"order_id"`
	ProductID  int      `json:"product_id" db:"product_id"`
	Quantity   int      `json:"quantity" db:"quantity"`
	UnitPrice  float64  `json:"unit_price" db:"unit_price"`
	TotalPrice float64  `json:"total_price" db:"total_price"`
	Product    *Product `json:"product,omitempty"`
}

type Transaction struct {
	ID              int       `json:"id" db:"id"`
	UserID          int       `json:"user_id" db:"user_id"`
	Type            string    `json:"type" db:"type"` // income, expense
	Category        string    `json:"category" db:"category"`
	Amount          float64   `json:"amount" db:"amount"`
	Description     string    `json:"description" db:"description"`
	TransactionDate time.Time `json:"transaction_date" db:"transaction_date"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

// Stok Hareketleri için model
type StockMovement struct {
	ID           int       `json:"id" db:"id"`
	UserID       int       `json:"user_id" db:"user_id"`
	ProductID    int       `json:"product_id" db:"product_id"`
	Type         string    `json:"type" db:"type"` // in, out, adjustment
	Quantity     int       `json:"quantity" db:"quantity"`
	Reference    string    `json:"reference" db:"reference"` // Sipariş no, fatura no, vb.
	Description  string    `json:"description" db:"description"`
	MovementDate time.Time `json:"movement_date" db:"movement_date"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	Product      *Product  `json:"product,omitempty"`
}

// Dashboard için özet veriler
type DashboardStats struct {
	TotalCustomers   int       `json:"total_customers"`
	TotalProducts    int       `json:"total_products"`
	TotalOrders      int       `json:"total_orders"`
	PendingOrders    int       `json:"pending_orders"`
	MonthlyRevenue   float64   `json:"monthly_revenue"`
	MonthlyExpenses  float64   `json:"monthly_expenses"`
	MonthlyProfit    float64   `json:"monthly_profit"`
	RecentOrders     []Order   `json:"recent_orders"`
	TopProducts      []Product `json:"top_products"`
	LowStockProducts []Product `json:"low_stock_products"`
}
