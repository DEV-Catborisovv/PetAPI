// REST API. PET-PROJECT
// Created by Catborisovv (c) 2020-2024

// Исходник точки входа в программу

package main

import (
	"PetAPI/internal/app"
	"PetAPI/pkg/database"
	"log"
)

func main() {
	err := app.Init()
	if err != nil {
		log.Fatalf("Возникла ошибка при развертке веб-сервера:\n%v\n", err)
	}

	// Сигнал-проекта завершения работы программы
	quit := make(chan struct{})
	go func() {
		quit <- struct{}{}
	}()

	<-quit
	database.DBClose()
}
