package handlers

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/alkosmas92/xm-golang/internal/services"
	"github.com/alkosmas92/xm-golang/internal/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	Service services.UserService
	Logger  *logrus.Logger
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(service services.UserService, logger *logrus.Logger) *UserHandler {
	return &UserHandler{Service: service, Logger: logger}
}

// Register handles user registration.
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	newUser := models.NewUser(user.Username, string(hashedPassword), user.FirstName, user.LastName)
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.Service.RegisterUser(ctx, newUser); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("user registered", zap.String("username", user.Username))
	w.WriteHeader(http.StatusCreated)
}

// Login handles user login.
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	user, err := h.Service.AuthenticateUser(ctx, credentials.Username, credentials.Password)
	if err != nil {
		h.Logger.Warn("failed to authenticate user", zap.String("username", credentials.Username), zap.Error(err))
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		h.Logger.Warn("invalid password", zap.String("username", credentials.Username))
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.UserID, user.Username)
	if err != nil {
		h.Logger.Error("failed to generate JWT", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("user logged in", zap.String("username", user.Username))
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"token": token}); err != nil {
		h.Logger.Error("failed to encode response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
