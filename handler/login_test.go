package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gabrielfmcoelho/abare-api.git/config"
	"github.com/gabrielfmcoelho/abare-api.git/schemas"
)

func setupTestDB() *gorm.DB {
	// Criar um banco de dados SQLite em memória para os testes
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&schemas.User{})
	config.SetSQLite(db) // Configurar o SQLite para os testes
	return db
}

func TestLoginHandler(t *testing.T) {
	// Configurar o ambiente de teste do Gin
	r := gin.Default()
	r.POST("/login", LoginHandler)

	// Inicializar o banco de dados de teste
	db := setupTestDB()

	// Criar um usuário de teste com senha criptografada
	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	testUser := schemas.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  string(password),
	}
	db.Create(&testUser)

	// Definir o corpo da requisição JSON para login
	loginPayload := `{"email":"john@example.com", "password":"password123"}`

	// Criar uma requisição HTTP de teste
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(loginPayload)))
	req.Header.Set("Content-Type", "application/json")

	// Gravar a resposta
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verificar o código de status e a resposta
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}
