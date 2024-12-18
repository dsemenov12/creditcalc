package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"strconv"
)

// Тестирование обработчика GetRequest
func TestGetRequest(t *testing.T) {
	// Создаем новый тестовый запрос (GET)
	req := httptest.NewRequest("GET", "/loan", nil)
	// Создаем новый записывающий ответ
	w := httptest.NewRecorder()

	// Вызываем обработчик
	GetRequest(w, req)

	// Проверяем статус ответа
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Проверяем, что в ответе присутствует заголовок "Кредитный калькулятор"
	if !strings.Contains(w.Body.String(), "Кредитный калькулятор") {
		t.Errorf("Expected body to contain 'Кредитный калькулятор', but got %q", w.Body.String())
	}
}

// Тестирование обработчика PostRequest
func TestPostRequest(t *testing.T) {
	tests := []struct {
		amount        string
		rate          string
		term          string
		expectedTotal float64
		expectedPayment float64
	}{
		{"1000", "5", "1", 1050, 86}, // Сумма 1000, ставка 5%, срок 1 год
		{"2000", "7", "2", 2280, 90}, // Сумма 2000, ставка 7%, срок 2 года
		//{"5000", "0", "3", 5000, 139}, // Ставка 0%
		{"1500", "10", "5", 2250, 32}, // Сумма 1500, ставка 10%, срок 5 лет
	}

	for _, tt := range tests {
		// Моделируем отправку POST-запроса с данными формы
		formData := "amount=" + tt.amount + "&rate=" + tt.rate + "&term=" + tt.term
		req := httptest.NewRequest("POST", "/loan", strings.NewReader(formData))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		// Вызываем обработчик
		PostRequest(w, req)

		// Проверяем статус ответа
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code 200, got %d", w.Code)
		}

		// Проверяем, что в ответе присутствуют корректные значения
		// Сумма кредита
		if !strings.Contains(w.Body.String(), tt.amount) {
			t.Errorf("Expected body to contain loan amount %s, but got %q", tt.amount, w.Body.String())
		}

		// Процентная ставка
		if !strings.Contains(w.Body.String(), tt.rate) {
			t.Errorf("Expected body to contain rate %s, but got %q", tt.rate, w.Body.String())
		}

		// Срок кредита
		if !strings.Contains(w.Body.String(), tt.term) {
			t.Errorf("Expected body to contain term %s, but got %q", tt.term, w.Body.String())
		}

		// Общая сумма к оплате
		expectedTotal := strconv.FormatFloat(tt.expectedTotal, 'f', 0, 64)
		if !strings.Contains(w.Body.String(), expectedTotal) {
			t.Errorf("Expected body to contain total %s, but got %q", expectedTotal, w.Body.String())
		}

		// Ежемесячный платеж
		expectedPayment := strconv.FormatFloat(tt.expectedPayment, 'f', 0, 64)
		if !strings.Contains(w.Body.String(), expectedPayment) {
			t.Errorf("Expected body to contain monthly payment %s, but got %q", expectedPayment, w.Body.String())
		}
	}
}