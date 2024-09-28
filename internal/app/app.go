package app

import (
	"github.com/ankodd/demoexam/core/internal/handlers"
	"github.com/ankodd/demoexam/core/internal/handlers/cors"
	"github.com/ankodd/demoexam/core/internal/storage"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"net/http"
	"os"
)

// Run initializes handlers, storage, etc. Starts the app
func Run(logger *slog.Logger) error {
	// Load dotenv
	err := godotenv.Load("../.env")
	if err != nil {
		return err
	}

	// Get path to storage.db from .env
	storagePath := os.Getenv("STORAGE_PATH")
	if storagePath == "" {
		log.Fatal("STORAGE_PATH env variable not set")
	}

	// Initial storages
	userStorage, err := storage.NewUserStorage(logger, storagePath)
	if err != nil {
		return err
	}

	defer userStorage.Close()

	orderStorage, err := storage.NewOrderStorage(logger, storagePath)
	if err != nil {
		return err
	}
	defer orderStorage.Close()

	// Get port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initial handlers
	uHandler := handlers.NewUserHandler(userStorage, logger)
	http.HandleFunc("/api/users", cors.Middleware(uHandler.FetchUsers))
	http.HandleFunc("/api/users/", cors.Middleware(uHandler.FetchUser))
	http.HandleFunc("/api/users/update/", cors.Middleware(uHandler.UpdateUser))
	http.HandleFunc("/api/users/delete/", cors.Middleware(uHandler.DeleteUser))
	http.HandleFunc("/api/users/registration", cors.Middleware(uHandler.Register))
	http.HandleFunc("/api/users/login", cors.Middleware(uHandler.Login))

	oHandler := handlers.NewOrderHandler(orderStorage, logger)
	http.HandleFunc("/api/orders", cors.Middleware(oHandler.FetchOrders))
	http.HandleFunc("/api/orders/create", cors.Middleware(oHandler.AddOrder))
	http.HandleFunc("/api/orders/", cors.Middleware(oHandler.FetchOrder))
	http.HandleFunc("/api/orders/update/", cors.Middleware(oHandler.UpdateOrder))
	http.HandleFunc("/api/orders/delete/", cors.Middleware(oHandler.DeleteOrder))
	http.HandleFunc("/api/statistics", cors.Middleware(oHandler.Statistics))

	// Start server
	logger.Info("server listening", slog.String("port", port))
	err = http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		return err
	}

	return nil
}
