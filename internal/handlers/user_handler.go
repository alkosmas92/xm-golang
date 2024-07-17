package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/alkosmas92/xm-golang/internal/services"
	"github.com/alkosmas92/xm-golang/internal/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	Service services.UserService
	Logger  *zap.Logger
}

func NewUserHandler(service services.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{Service: service, Logger: logger}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.Logger.Error("failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		h.Logger.Error("failed to hash password", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	newUser := models.NewUser(user.Username, string(hashedPassword), user.FirstName, user.LastName)

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.Service.RegisterUser(ctx, newUser); err != nil {
		h.Logger.Error("failed to register user", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("user registered", zap.String("username", user.Username))
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		h.Logger.Error("failed to decode request body", zap.Error(err))
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
	if err := json.NewEncoder(w).Encode(map[string]string{"token": token, "user_id": user.UserID}); err != nil {
		h.Logger.Error("failed to encode response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
