package main

import (
	"GoMicroservices/connection"
	"GoMicroservices/repository"
	"context"
	"fmt"
)

func main() {
	// Инициализируем подключение к MongoDB
	if err := connection.InitMongoDB(); err != nil {
		fmt.Println("Ошибка при инициализации подключения к MongoDB:", err)
		return
	}

	// Создаем экземпляр репозитория
	paymentRepo := repository.NewPaymentRepository()

	// Получаем контекст
	ctx := context.Background()

	// Получаем все платежи из базы данных
	payments, err := paymentRepo.FindAll(ctx)
	if err != nil {
		fmt.Println("Ошибка при получении всех платежей:", err)
		return
	}

	// Выводим полученные платежи
	for _, payment := range payments {
		fmt.Println("ID:", payment.ID, "Сумма:", payment.Amount, "Статус:", payment.Status)
	}
}
