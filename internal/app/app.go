// Pet-API project. Main-APP source file
// Created by Catborisovv (c) 2020-2024

package app

import (
	newuser "PetAPI/API/Handlers/NewUser"
	toml "PetAPI/pkg/TomlDecode"
	"fmt"
	"log"
	"net/http"
)

func Init() error {
	// Создание экземпляра структуры конфига
	err, config := toml.DecodeConfigTOML()

	if err != nil {
		return err
	}

	// Настройка функций обработчиков:
	http.HandleFunc("/newuser", newuser.NewUserHandler)

	// Развертка HTTP-Cервера
	var configAddr string = fmt.Sprintf("%s:%d", config.HttpSrv.Host, config.HttpSrv.Port)
	log.Println(configAddr)

	err = http.ListenAndServe(configAddr, nil)
	if err != nil {
		return err
	}
	return nil
}
