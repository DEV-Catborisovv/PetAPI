// Пул подключений
// Created by Catborisovv (c) 2020-2024

package database

import (
	toml "PetAPI/pkg/TomlDecode"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"
)

type ConnectionPool struct {
	maxConnections     int
	currentConnections int
	connections        chan *sql.DB
	mutex              sync.Mutex
}

func NewConnectionPool(maxConnections int) *ConnectionPool {
	pool := &ConnectionPool{
		maxConnections: maxConnections,
		connections:    make(chan *sql.DB, maxConnections),
	}
	return pool
}

func (pool *ConnectionPool) GetConnection() (*sql.DB, error) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	if pool.connections == nil {
		return nil, errors.New("pool is closed")
	}

	if len(pool.connections) == 0 && pool.currentConnections < pool.maxConnections {
		err, config := toml.DecodeConfigTOML()
		if err != nil {
			log.Fatalf("Возникла ошибка при чтении данных из конфига:\n%v\n", err)
		}

		var dbinfo string = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.DataBase.DB_USER, config.DataBase.DB_PASS, config.DataBase.DB_NAME)
		db, err := sql.Open("postgres", dbinfo)
		if err != nil {
			return nil, err
		}
		pool.currentConnections++
		return db, nil
	}

	db := <-pool.connections
	return db, nil
}

func (pool *ConnectionPool) ReleaseConnection(db *sql.DB) error {
	if db == nil {
		return errors.New("invalid connection")
	}

	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	if pool.connections == nil {
		return db.Close()
	}

	pool.connections <- db
	return nil
}

func (pool *ConnectionPool) Close() {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	if pool.connections == nil {
		return
	}

	close(pool.connections)

	for db := range pool.connections {
		db.Close()
	}

	pool.connections = nil
}
