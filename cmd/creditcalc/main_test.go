package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/dsemenov12/creditcalc/internal/handlers"
	"github.com/stretchr/testify/assert"
)

// Функция для создания тестового роутера
func setupRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", handlers.GetRequest)
	router.Post("/", handlers.PostRequest)
	return router
}

// Тест для маршрута GET "/"
func TestGetRoute(t *testing.T) {
	router := setupRouter()

	// Создаем запрос GET на корневой маршрут
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	// Выполняем запрос
	router.ServeHTTP(rec, req)

	// Проверяем код ответа
	assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200")

	// Проверяем содержимое ответа
	body := rec.Body.String()
	assert.Contains(t, body, "Кредитный калькулятор", "Response should contain 'Кредитный калькулятор'")
}

// Тест для маршрута POST "/"
func TestPostRoute(t *testing.T) {
	router := setupRouter()

	// Подготовка данных формы для POST-запроса
	formData := "amount=1000&rate=5&term=1"
	req := httptest.NewRequest("POST", "/", strings.NewReader(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()

	// Выполняем запрос
	router.ServeHTTP(rec, req)

	// Проверяем код ответа
	assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200")

	// Проверяем содержимое ответа
	body := rec.Body.String()

	assert.Contains(t, body, "1000 руб.", "Response should contain '1000 руб.'")
	assert.Contains(t, body, "5%", "Response should contain '5%'")
	assert.Contains(t, body, "1 лет", "Response should contain '1 лет'")
	assert.Contains(t, body, "Общая сумма к оплате", "Response should contain 'Общая сумма к оплате'")
	assert.Contains(t, body, "Ежемесячный платеж", "Response should contain 'Ежемесячный платеж'")
}
