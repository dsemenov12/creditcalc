package calculate

import (
	"testing"
	"math"
)

func TestCalculateTotal(t *testing.T) {
	tests := []struct {
		amount  float64
		rate    float64
		term    int
		want    float64
	}{
		{1000, 5, 1, 1050},         // Тест с суммой 1000, ставкой 5% и сроком 1 год
		{1000, 10, 2, 1200},        // Тест с суммой 1000, ставкой 10% и сроком 2 года
		{2000, 0, 3, 2000},         // Тест с процентной ставкой 0%
		{1500, 7, 5, 2025},         // Тест с суммой 1500, ставкой 7% и сроком 5 лет
		{5000, 3, 10, 6500},        // Тест с суммой 5000, ставкой 3% и сроком 10 лет
	}

	for _, tt := range tests {
		t.Run("CalculateTotal", func(t *testing.T) {
			got := CalculateTotal(tt.amount, tt.rate, tt.term)
			if math.Abs(got-tt.want) > 0.01 { // Используем погрешность 0.01 для сравнения с плавающими точками
				t.Errorf("CalculateTotal(%v, %v, %v) = %v; want %v", tt.amount, tt.rate, tt.term, got, tt.want)
			}
		})
	}
}

func TestCalculateMonthlyPayment(t *testing.T) {
	tests := []struct {
		amount  float64
		rate    float64
		term    int
		want    float64
	}{
		{1000, 5, 1, 86},         // Тест с суммой 1000, ставкой 5% и сроком 1 год
		{1000, 10, 2, 46},        // Тест с суммой 1000, ставкой 10% и сроком 2 года
		//{2000, 0, 3, 56},        // Тест с процентной ставкой 0%
		{1500, 7, 5, 30},         // Тест с суммой 1500, ставкой 7% и сроком 5 лет
		{5000, 3, 10, 48},        // Тест с суммой 5000, ставкой 3% и сроком 10 лет
	}

	for _, tt := range tests {
		t.Run("CalculateMonthlyPayment", func(t *testing.T) {
			got := CalculateMonthlyPayment(tt.amount, tt.rate, tt.term)
			if math.Abs(got-tt.want) > 0.01 { // Используем погрешность 0.01 для сравнения с плавающими точками
				t.Errorf("CalculateMonthlyPayment(%v, %v, %v) = %v; want %v", tt.amount, tt.rate, tt.term, got, tt.want)
			}
		})
	}
}
