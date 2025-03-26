package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Запитуємо ім'я
	fmt.Print("Введіть ваше ім'я: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	if name == "" {
		fmt.Println("Помилка: Ім'я не може бути порожнім")
		return
	}

	// Запитуємо стать
	fmt.Print("Введіть вашу стать (ч/ж): ")
	scanner.Scan()
	gender := strings.ToLower(strings.TrimSpace(scanner.Text()))
	if gender != "ч" && gender != "ж" {
		fmt.Println("Помилка: Введіть 'ч' для чоловіка або 'ж' для жінки")
		return
	}

	// Запитуємо вік
	fmt.Print("Введіть ваш вік: ")
	var age int
	_, err := fmt.Scan(&age)
	if err != nil || age <= 0 || age > 150 {
		fmt.Println("Помилка: Введіть коректний вік (число від 1 до 150)")
		return
	}

	// Запитуємо вагу
	fmt.Print("Введіть вашу вагу (кг): ")
	var weight float64
	_, err = fmt.Scan(&weight)
	if err != nil || weight <= 0 || weight > 500 {
		fmt.Println("Помилка: Введіть коректну вагу (число від 0 до 500 кг)")
		return
	}

	// Запитуємо зріст
	fmt.Print("Введіть ваш зріст (см): ")
	var height float64
	_, err = fmt.Scan(&height)
	if err != nil || height <= 0 || height > 300 {
		fmt.Println("Помилка: Введіть коректний зріст (число від 0 до 300 см)")
		return
	}

	// Очищаємо буфер перед наступним введенням
	scanner = bufio.NewScanner(os.Stdin)

	// Переводимо зріст з см в метри
	heightInMeters := height / 100

	// Розрахунок ІМТ
	bmi := weight / (heightInMeters * heightInMeters)

	// Визначаємо повну назву статі
	genderFull := "чоловік"
	if gender == "ж" {
		genderFull = "жінка"
	}

	// Виводимо результат в консоль
	fmt.Printf("\n%s, ваш ІМТ: %.1f\n", name, bmi)

	// Інтерпретація результату
	switch {
	case bmi < 18.5:
		fmt.Println("Категорія: Нестача ваги")
	case bmi < 25:
		fmt.Println("Категорія: Нормальна вага")
	case bmi < 30:
		fmt.Println("Категорія: Надлишкова вага")
	default:
		fmt.Println("Категорія: Ожиріння")
	}

	// Запитуємо шлях для збереження файлу
	fmt.Print("\nВведіть шлях для збереження файлу (наприклад, C:/Users/ім'я/bmi.txt): ")
	scanner.Scan()
	filePath := strings.TrimSpace(scanner.Text())
	if filePath == "" {
		fmt.Println("Помилка: Шлях до файлу не може бути порожнім")
		return
	}

	// Створюємо вміст для файлу
	content := fmt.Sprintf("Ім'я: %s\nСтать: %s\nВік: %d\nВага: %.1f кг\nЗріст: %.1f см\nІМТ: %.1f\n",
		name, genderFull, age, weight, height, bmi)

	// Записуємо у файл
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Помилка при створенні файлу:", err)
		return
	}

	fmt.Println("Інформацію успішно збережено у файл:", filePath)

	// Очищаємо буфер і чекаємо вводу
	fmt.Println("\nНатисніть Enter для завершення програми...")
	// Використовуємо низькорівневе читання вводу
	buf := make([]byte, 1)
	os.Stdin.Read(buf)
}
