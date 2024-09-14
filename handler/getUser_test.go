package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/gabrielfmcoelho/abare-api.git/schemas"
)

func TestGetUserHandler(t *testing.T) {
	// Configurar o ambiente de teste do Gin
	r := gin.Default()
	r.GET("/users/:id", GetUserHandler)

	// Inicializar o banco de dados de teste
	db := setupTestDB()

	// Criar um usuário de teste
	testUser := schemas.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}
	db.Create(&testUser)

	// Criar uma requisição HTTP de teste
	req, _ := http.NewRequest("GET", "/users/1", nil)

	// Gravar a resposta
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verificar o código de status e a resposta
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John")
}
