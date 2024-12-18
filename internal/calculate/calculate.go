package calculate

import "math"

// Функция для расчета общей суммы кредита
func CalculateTotal(amount, rate float64, term int) float64 {
	// Простая формула для расчета суммы с учетом процентов
	interest := rate / 100
	total := amount * (1 + (interest * float64(term)))
	return math.Round(total)
}

// Функция для расчета ежемесячного платежа
func CalculateMonthlyPayment(amount, rate float64, term int) float64 {
	// Преобразуем годовую ставку в месячную
	monthlyRate := rate / 100 / 12
	// Срок в месяцах
	months := float64(term * 12)

	// Формула для расчета ежемесячного платежа
	if monthlyRate == 0 {
		// Если процентная ставка 0, то платеж = сумма кредита / срок в месяцах
		return amount / months
	}

	// Аннуитетная формула
	payment := amount * (monthlyRate * math.Pow(1+monthlyRate, months)) / (math.Pow(1+monthlyRate, months) - 1)
	return math.Round(payment)
}