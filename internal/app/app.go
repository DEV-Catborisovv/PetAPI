// Pet-API project. Main-APP source file
// Created by Catborisovv (c) 2020-2024

package app

import (
	autotizate "PetAPI/API/Handlers/Autotizate"
	getuserinfo "PetAPI/API/Handlers/GetUserInfo"
	newuser "PetAPI/API/Handlers/NewUser"
	toml "PetAPI/pkg/TomlDecode"
	"fmt"
	"log"
	"net/http"
)

// Метод запуска Вэб-Сервера (Старт-Функция)
func Init() error {
	// Создание экземпляра структуры конфига
	err, config := toml.DecodeConfigTOML()

	if err != nil {
		return err
	}

	// Настройка функций обработчиков
	http.HandleFunc("/newuser", newuser.NewUserHandler)
	http.HandleFunc("/login", autotizate.AutotizateHandler)

	// Обработчик получения данных пользователя
	http.HandleFunc("/", getuserinfo.GetUserInfo)

	// Развертка HTTP-Cервера
	var configAddr string = fmt.Sprintf("%s:%d", config.HttpSrv.Host, config.HttpSrv.Port)
	log.Println(configAddr)

	err = http.ListenAndServe(configAddr, nil)
	if err != nil {
		return err
	}
	return nil
}
