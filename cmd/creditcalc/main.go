package main

import (
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

// Структура для хранения данных формы
type LoanData struct {
	Amount          float64
	Rate            float64
	Term            int
	Total           float64
	MonthlyPayment  float64
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

// Обработчик для отображения формы и расчета кредита
func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Парсим данные из формы
		amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
		rate, _ := strconv.ParseFloat(r.FormValue("rate"), 64)
		term, _ := strconv.Atoi(r.FormValue("term"))

		// Расчет общей суммы кредита и ежемесячного платежа
		total := calculateTotal(amount, rate, term)
		monthlyPayment := calculateMonthlyPayment(amount, rate, term)

		// Подготовка данных для вывода
		loanData := LoanData{
			Amount:         amount,
			Rate:           rate,
			Term:           term,
			Total:          total,
			MonthlyPayment: monthlyPayment,
		}

		// Загрузка шаблона и передача данных в шаблон
		tmpl, err := template.ParseFiles("../../templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, loanData)
	} else {
		// Загружаем шаблон с пустыми данными при первом запуске
		tmpl, err := template.ParseFiles("../../templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
}

// Функция для расчета общей суммы кредита
func calculateTotal(amount, rate float64, term int) float64 {
	// Простая формула для расчета суммы с учетом процентов
	interest := rate / 100
	total := amount * (1 + (interest * float64(term)))
	return total
}

// Функция для расчета ежемесячного платежа
func calculateMonthlyPayment(amount, rate float64, term int) float64 {
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
	return payment
}
