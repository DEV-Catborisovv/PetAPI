// NewUser API Method for PetAPI-Project
// Created by Catborisovv (c) 2020-2024

package newuser

import (
	errrorhandler "PetAPI/pkg/ErrrorHandler"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// Функция валидации данных пользователя
func dataValidation(name string) bool {
	pattern := `^[a-zA-Z0-9]+$`
	re := regexp.MustCompile(pattern)

	if re.MatchString(name) {
		return true
	} else {
		return false
	}
}

// Обработчик события создания нового пользователя
func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Printf("Возникла ошибка при получении тела запроса:\n%v\n", err)
		}

		var nickname string = strings.TrimSpace(r.FormValue("name"))
		var pass string = strings.TrimSpace(r.FormValue("password"))

		// Проверка переданных данных на валидность
		if dataValidation(nickname) == true && dataValidation(pass) == true && len(nickname) != 0 && len(pass) != 0 {
			w.WriteHeader(200)
			fmt.Println(r)
		} else {
			errJson := errrorhandler.GetErrorJson(400, "Data in URL is not valid")
			w.WriteHeader(400)
			w.Write([]byte(errJson))
		}
	} else {
		errJson := errrorhandler.GetErrorJson(405, fmt.Sprintf("The %s method is not supported for this API method", r.Method))
		w.WriteHeader(405)
		w.Write([]byte(errJson))
	}
}
