// Клиент базы данных PostgreSQL
// Created by Catborisovv (c) 2020-2024

// Необходимо ввести пул соединений

package database

import (
	"log"

	_ "github.com/lib/pq" // Драйвер базы данных
)

// Стркутура пользователя
type User struct {
	ID       int
	Name     string
	Password string
	Mail     string
	Admin    int
}

var pool *ConnectionPool = NewConnectionPool(10)

// Метод вставки данных в таблицу
func InsertData(Request string) error {
	db, err := pool.GetConnection()
	if err != nil {
		log.Printf("Возникла ошибка при получении пула:\n%v\n", err)
	}

	insert, err := db.Query(Request)
	if err != nil {
		return err
	}
	defer insert.Close()
	defer pool.ReleaseConnection(db)
	return nil
}

// Метод получения данных пользователя
func SelectUserData(Request string) (error, *User) {
	db, err := pool.GetConnection()
	if err != nil {
		log.Printf("Возникла ошибка при получении пула:\n%v\n", err)
	}

	UserS := User{}
	resp, err := db.Query(Request)
	if err != nil {
		return err, nil
	}

	for resp.Next() {
		err = resp.Scan(&UserS.ID, &UserS.Name, &UserS.Password, &UserS.Mail, &UserS.Admin)
		if err != nil {
			return err, nil
		}
	}
	defer resp.Close()
	defer pool.ReleaseConnection(db)
	return nil, &UserS
}
