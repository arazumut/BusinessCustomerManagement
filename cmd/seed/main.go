package main

import (
	"log"
	"time"

	"github.com/umutaraz/tradesman-app/internal/database"
)

func main() {
	// Veritabanını başlat
	db, err := database.Initialize("./tradesman.db")
	if err != nil {
		log.Fatal("Veritabanı bağlantısı başarısız:", err)
	}
	defer db.Close()

	// Örnek kullanıcı
	_, err = db.Exec(`
		INSERT OR IGNORE INTO users (id, name, email, password_hash, business_name, phone, address) 
		VALUES (1, 'Ahmet Yılmaz', 'ahmet@example.com', 'hash123', 'Yılmaz Elektrik', '0532 123 45 67', 'İstanbul/Kadıköy')
	`)
	if err != nil {
		log.Printf("Kullanıcı ekleme hatası: %v", err)
	}

	// Örnek müşteriler
	customers := []struct {
		name    string
		email   string
		phone   string
		address string
		notes   string
	}{
		{"Mehmet Kaya", "mehmet@example.com", "0533 111 22 33", "İstanbul/Beşiktaş", "Düzenli müşteri"},
		{"Ayşe Demir", "ayse@example.com", "0534 222 33 44", "İstanbul/Şişli", "Toptan alım yapar"},
		{"Fatma Özkan", "fatma@example.com", "0535 333 44 55", "İstanbul/Bakırköy", "Hızlı ödeme"},
		{"Ali Çelik", "ali@example.com", "0536 444 55 66", "İstanbul/Ümraniye", "Kurumsal müşteri"},
		{"Zeynep Arslan", "zeynep@example.com", "0537 555 66 77", "İstanbul/Pendik", "İyi referans"},
	}

	for _, customer := range customers {
		_, err = db.Exec(`
			INSERT OR IGNORE INTO customers (user_id, name, email, phone, address, notes, created_at, updated_at) 
			VALUES (1, ?, ?, ?, ?, ?, ?, ?)
		`, customer.name, customer.email, customer.phone, customer.address, customer.notes, time.Now(), time.Now())
		if err != nil {
			log.Printf("Müşteri ekleme hatası: %v", err)
		}
	}

	// Örnek ürünler
	products := []struct {
		name        string
		description string
		price       float64
		category    string
		stock       int
		unit        string
	}{
		{"LED Ampul 12W", "Beyaz ışık LED ampul", 25.50, "Aydınlatma", 100, "adet"},
		{"Elektrik Kablosu 2.5mm", "NYA kablo 2.5mm²", 5.75, "Kablo", 500, "metre"},
		{"Priz Takımı", "Beyaz priz ve anahtar takımı", 35.00, "Elektrik Malzemesi", 50, "takım"},
		{"Elektrik Panosu", "6'lı sigorta panosu", 120.00, "Panel", 20, "adet"},
		{"Tesisat Hizmeti", "Ev elektrik tesisatı kurulumu", 500.00, "Hizmet", 0, "iş"},
		{"Spot LED", "3W spot LED", 15.00, "Aydınlatma", 80, "adet"},
		{"Kablo Kanalı", "16x16 beyaz kablo kanalı", 8.50, "Aksesuar", 200, "metre"},
		{"Dimmer Anahtar", "LED uyumlu dimmer", 85.00, "Elektrik Malzemesi", 25, "adet"},
	}

	for _, product := range products {
		_, err = db.Exec(`
			INSERT OR IGNORE INTO products (user_id, name, description, price, category, stock_quantity, unit, created_at, updated_at) 
			VALUES (1, ?, ?, ?, ?, ?, ?, ?, ?)
		`, product.name, product.description, product.price, product.category, product.stock, product.unit, time.Now(), time.Now())
		if err != nil {
			log.Printf("Ürün ekleme hatası: %v", err)
		}
	}

	// Örnek siparişler
	orders := []struct {
		customerID   int
		orderNumber  string
		status       string
		totalAmount  float64
		notes        string
		orderDate    time.Time
		deliveryDate *time.Time
	}{
		{1, "SIP-2024-001", "completed", 250.00, "Acil teslimat", time.Now().AddDate(0, 0, -5), nil},
		{2, "SIP-2024-002", "pending", 180.50, "Stok bekliyor", time.Now().AddDate(0, 0, -2), nil},
		{3, "SIP-2024-003", "processing", 95.00, "Hazırlanıyor", time.Now().AddDate(0, 0, -1), nil},
		{4, "SIP-2024-004", "shipped", 420.00, "Kargo verildi", time.Now(), nil},
		{1, "SIP-2024-005", "pending", 75.25, "Yeni sipariş", time.Now(), nil},
	}

	for _, order := range orders {
		_, err = db.Exec(`
			INSERT OR IGNORE INTO orders (user_id, customer_id, order_number, status, total_amount, notes, order_date, delivery_date, created_at, updated_at) 
			VALUES (1, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`, order.customerID, order.orderNumber, order.status, order.totalAmount, order.notes, order.orderDate, order.deliveryDate, time.Now(), time.Now())
		if err != nil {
			log.Printf("Sipariş ekleme hatası: %v", err)
		}
	}

	// Örnek gelir-gider kayıtları
	transactions := []struct {
		transactionType string
		category        string
		amount          float64
		description     string
		date            time.Time
	}{
		{"income", "Satış", 2500.00, "LED ampul satışları", time.Now().AddDate(0, 0, -10)},
		{"expense", "Alım", 1200.00, "Yeni stok alımı", time.Now().AddDate(0, 0, -8)},
		{"income", "Hizmet", 1500.00, "Tesisat kurulum işi", time.Now().AddDate(0, 0, -5)},
		{"expense", "Kira", 3000.00, "Mağaza kirası", time.Now().AddDate(0, 0, -3)},
		{"income", "Satış", 850.00, "Elektrik malzeme satışı", time.Now().AddDate(0, 0, -2)},
		{"expense", "Yakıt", 200.00, "Araç yakıtı", time.Now().AddDate(0, 0, -1)},
		{"income", "Satış", 450.00, "Bugünkü satışlar", time.Now()},
	}

	for _, tx := range transactions {
		_, err = db.Exec(`
			INSERT OR IGNORE INTO transactions (user_id, type, category, amount, description, transaction_date, created_at) 
			VALUES (1, ?, ?, ?, ?, ?, ?)
		`, tx.transactionType, tx.category, tx.amount, tx.description, tx.date, time.Now())
		if err != nil {
			log.Printf("İşlem ekleme hatası: %v", err)
		}
	}

	log.Println("Örnek veriler başarıyla eklendi!")
	log.Println("5 müşteri, 8 ürün, 5 sipariş ve 7 finansal işlem eklendi.")
}
