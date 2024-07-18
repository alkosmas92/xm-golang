package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/alkosmas92/xm-golang/internal/handlers"
	"github.com/alkosmas92/xm-golang/internal/middleware"
	"github.com/alkosmas92/xm-golang/internal/repository"
	"github.com/alkosmas92/xm-golang/internal/services"
	"github.com/sirupsen/logrus"
)

// Run starts the HTTP server.
func Run(logger *logrus.Logger, db *sql.DB) error {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService, logger)

	favoriteRepo := repository.NewCompanyRepository(db)
	favoriteService := services.NewCompanyService(favoriteRepo)
	favoriteHandler := handlers.NewCompanyHandler(favoriteService, logger)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)

	http.Handle("/company", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			favoriteHandler.CreateCompany(w, r)
		case http.MethodGet:
			favoriteHandler.GetCompanyByID(w, r)
		case http.MethodPut:
			favoriteHandler.UpdateCompany(w, r)
		case http.MethodDelete:
			favoriteHandler.DeleteCompany(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	logger.Info("Starting server on :8080")
	log.Print("Starting server on :8080")
	return http.ListenAndServe(":8080", nil)
}
