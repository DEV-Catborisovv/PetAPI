// NewUser API Method for PetAPI-Project
// Created by Catborisovv (c) 2020-2024

package newuser

import (
	errrorhandler "PetAPI/pkg/ErrrorHandler"
	"PetAPI/pkg/database"
	validdata "PetAPI/pkg/validData"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Обработчик события создания нового пользователя
func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Printf("Возникла ошибка при получении тела запроса:\n%v\n", err)
		}

		var nickname string = strings.TrimSpace(r.FormValue("name"))
		var pass string = strings.TrimSpace(r.FormValue("password"))
		var mail string = strings.TrimSpace(r.FormValue("mail"))

		err, u := database.SelectUserData(fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", nickname))
		if err != nil {
			log.Printf("Возникла ошибка при прочтении файла:\n%v\n", err)
		}
		// Проверка на существование пользователя
		if u.ID == 0 {
			errJson := errrorhandler.GetErrorJson(400, "Data is not valid")
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(errJson))
		} else {

			// Проверка переданных данных на валидность
			if validdata.DataValidation(nickname) == true && validdata.DataValidation(pass) == true && validdata.MailValidation(mail) && len(nickname) != 0 && len(pass) != 0 && len(mail) != 0 {
				err, _ := database.SelectUserData(fmt.Sprintf("SELECT * FROM users WHERE name = '%s' OR mail = '%s';", nickname, mail))
				if err != nil {
					log.Println(err)
				}
				// Регистрируем нового пользователя
				err = database.InsertData(fmt.Sprintf("INSERT INTO users (name, password, mail, admin) VALUES ('%s', encode(digest('%s', 'sha256'), 'hex'), '%s', '0');", nickname, pass, mail))
				if err != nil {
					log.Fatalf("Возникла ошибка при вставке значения в таблицу базы данных:\n%v\n", err)
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(500)
				}
				w.WriteHeader(200)
			} else {
				errJson := errrorhandler.GetErrorJson(400, "Data in URL is not valid")
				w.WriteHeader(400)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(errJson))
			}
		}
	} else {
		errJson := errrorhandler.GetErrorJson(405, fmt.Sprintf("The %s method is not supported for this API method", r.Method))
		w.WriteHeader(405)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(errJson))
	}
}
