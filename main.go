package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight, userKg float64

	// Запитуємо зріст
	fmt.Print("Введіть ваш зріст (у метрах, наприклад, 1.75): ")
	fmt.Scan(&userHeight)

	// Запитуємо вагу
	fmt.Print("Введіть вашу вагу (у кілограмах, наприклад, 70): ")
	fmt.Scan(&userKg)

	// Обчислюємо ІМТ
	var imt float64 = userKg / math.Pow(userHeight, 2)

	// Виводимо результат
	fmt.Printf("Ваш ІМТ: %.2f\n", imt)

	// Інтерпретація
	if imt < 18.5 {
		fmt.Println("Недостатня вага")
	} else if imt >= 18.5 && imt < 25 {
		fmt.Println("Нормальна вага")
	} else if imt >= 25 && imt < 30 {
		fmt.Println("Надлишкова вага")
	} else {
		fmt.Println("Ожиріння")
	}
}
