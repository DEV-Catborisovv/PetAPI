// Клиент базы данных PostgreSQL
// Created by Catborisovv (c) 2020-2024

package database

import (
	toml "PetAPI/pkg/TomlDecode"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Драйвер базы данных
)

// Стркутура пользователя
type User struct {
	ID       int
	Name     string
	Password string
	Mail     string
}

// Создание активной сессии БД
var database *sql.DB = DBConnect()

// Метод завершения активной сессии БД
func DBClose() {
	err := database.Close()
	if err != nil {
		log.Fatalf("Возникла ошибка при завершении активной сессии базы данных:\n%v\n", err)
	}
	log.Println("Сессия базы данных была завершена")
}

// Метод создания новой сессии базы данных
func DBConnect() *sql.DB {
	err, config := toml.DecodeConfigTOML()
	if err != nil {
		log.Fatalf("Возникла ошибка при подключении к БД:\n%v\n", err)
	}

	// Форматирование строки подключения
	var dbinfo string = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.DataBase.DB_USER, config.DataBase.DB_PASS, config.DataBase.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalf("Возникла ошибка при подключении к БД:\n%v\n", err)
	}
	log.Printf("Была создана новая сессия базы данных: %s\n", dbinfo)
	return db
}

// Метод вставки данных в таблицу
func InsertData(Request string) error {
	insert, err := database.Query(Request)
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}

// Метод получения данных пользователя
func SelectUserData(Request string) (error, *User) {
	u := User{}
	resp, err := database.Query(Request)
	if err != nil {
		return err, nil
	}
	if resp.Next() != true {
		return nil, nil
	}
	for resp.Next() {
		err = resp.Scan(&u.ID, &u.Name, &u.Password, &u.Mail)
		if err != nil {
			return err, nil
		}
	}
	defer resp.Close()
	return nil, &u
}
