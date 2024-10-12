package user

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserHandler(t *testing.T) {
	// Configurar o ambiente de teste do Gin
	r := gin.Default()
	r.POST("/users", CreateUserHandler)

	// Inicializar o banco de dados de teste
	setupTestDB()

	// Definir o corpo da requisição JSON
	userPayload := `{"first_name":"John", "last_name":"Doe", "email":"john@example.com", "password":"password123"}`

	// Criar uma requisição HTTP de teste
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(userPayload)))
	req.Header.Set("Content-Type", "application/json")

	// Gravar a resposta
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verificar o código de status e a resposta
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "User created successfully")
}
