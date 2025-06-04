package main

import (
	"fmt"
)

// calculateSquare возвращает квадрат числа
func calculateSquare(num float64) float64 {
	return num * num
}

func main() {
	fmt.Println("🟦 Программа для возведения в квадрат 🟦")

	var input float64
	fmt.Print("Введите число: ")

	if _, err := fmt.Scan(&input); err != nil {
		fmt.Println("🚫 Ошибка:", err)
		return
	}

	result := calculateSquare(input)
	fmt.Printf("✅ Результат: %.2f² = %.2f\n", input, result)
}
