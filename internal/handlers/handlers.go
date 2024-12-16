package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/dsemenov12/creditcalc/internal/calculate"
)

// Структура для хранения данных формы
type LoanData struct {
	Amount          float64
	Rate            float64
	Term            int
	Total           float64
	MonthlyPayment  float64
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	// Загружаем шаблон с пустыми данными при первом запуске
	tmpl, err := template.ParseFiles("../../templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	// Парсим данные из формы
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
	rate, _ := strconv.ParseFloat(r.FormValue("rate"), 64)
	term, _ := strconv.Atoi(r.FormValue("term"))

	// Расчет общей суммы кредита и ежемесячного платежа
	total := calculate.CalculateTotal(amount, rate, term)
	monthlyPayment := calculate.CalculateMonthlyPayment(amount, rate, term)

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
}