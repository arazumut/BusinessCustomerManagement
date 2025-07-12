package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func Initialize(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("veritabanı açma hatası: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("veritabanı bağlantı testi hatası: %w", err)
	}

	database := &DB{DB: db}

	// Tabloları oluştur
	if err := database.createTables(); err != nil {
		return nil, fmt.Errorf("tablo oluşturma hatası: %w", err)
	}

	return database, nil
}

func (db *DB) createTables() error {
	// Kullanıcılar tablosu
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		role TEXT DEFAULT 'user',
		business_name TEXT,
		phone TEXT,
		address TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	// Müşteriler tablosu
	customersTable := `
	CREATE TABLE IF NOT EXISTS customers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		email TEXT,
		phone TEXT,
		address TEXT,
		notes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	// Ürünler/Hizmetler tablosu
	productsTable := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		description TEXT,
		price DECIMAL(10,2) NOT NULL,
		category TEXT,
		stock_quantity INTEGER DEFAULT 0,
		unit TEXT DEFAULT 'adet',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	// Siparişler tablosu
	ordersTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		customer_id INTEGER NOT NULL,
		order_number TEXT UNIQUE NOT NULL,
		status TEXT DEFAULT 'pending',
		total_amount DECIMAL(10,2) NOT NULL,
		notes TEXT,
		order_date DATETIME DEFAULT CURRENT_TIMESTAMP,
		delivery_date DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (customer_id) REFERENCES customers(id)
	);`

	// Sipariş detayları tablosu
	orderItemsTable := `
	CREATE TABLE IF NOT EXISTS order_items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		order_id INTEGER NOT NULL,
		product_id INTEGER NOT NULL,
		quantity INTEGER NOT NULL,
		unit_price DECIMAL(10,2) NOT NULL,
		total_price DECIMAL(10,2) NOT NULL,
		FOREIGN KEY (order_id) REFERENCES orders(id),
		FOREIGN KEY (product_id) REFERENCES products(id)
	);`

	// Gelir/Gider takibi tablosu
	transactionsTable := `
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		type TEXT NOT NULL CHECK (type IN ('income', 'expense')),
		category TEXT NOT NULL,
		amount DECIMAL(10,2) NOT NULL,
		description TEXT,
		transaction_date DATETIME DEFAULT CURRENT_TIMESTAMP,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	tables := []string{
		usersTable,
		customersTable,
		productsTable,
		ordersTable,
		orderItemsTable,
		transactionsTable,
	}

	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			return fmt.Errorf("tablo oluşturma hatası: %w", err)
		}
	}

	return nil
}
