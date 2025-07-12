package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/umutaraz/tradesman-app/internal/config"
	"github.com/umutaraz/tradesman-app/internal/database"
	"github.com/umutaraz/tradesman-app/internal/handlers"
	"github.com/umutaraz/tradesman-app/internal/middleware"
	"github.com/umutaraz/tradesman-app/internal/routes"
)

func main() {
	// Konfigürasyon yükle
	cfg := config.Load()

	// Veritabanını başlat
	db, err := database.Initialize(cfg.DatabasePath)
	if err != nil {
		log.Fatal("Veritabanı bağlantısı başarısız:", err)
	}
	defer db.Close()

	// Gin router'ı başlat
	r := gin.Default()

	// Static dosyalar
	r.Static("/assets", "./assets")

	// Şablon fonksiyonlarını tanımla
	r.SetFuncMap(template.FuncMap{
		"now": time.Now,
		"mul": func(a, b float64) float64 {
			return a * b
		},
		"float64": func(i int) float64 {
			return float64(i)
		},
	})

	r.LoadHTMLGlob("templates/*")

	// Middleware'ler
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())

	// Handler'ları başlat
	h := handlers.New(db)

	// Route'ları kaydet
	routes.Setup(r, h)

	// Sunucuyu başlat
	log.Printf("Esnaf Yönetim Uygulaması başlatılıyor... Port: %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal("Sunucu başlatma hatası:", err)
	}
}
