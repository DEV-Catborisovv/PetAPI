// NewUser API Method for PetAPI-Project
// Created by Catborisovv (c) 2020-2024

package newuser

import (
	errrorhandler "PetAPI/pkg/ErrrorHandler"
	"PetAPI/pkg/database"
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

// Функция валидации почты пользователя
func MailValidation(mail string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, mail); !m {
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
		var mail string = strings.TrimSpace(r.FormValue("mail"))

		// Проверка переданных данных на валидность
		if dataValidation(nickname) == true && dataValidation(pass) == true && MailValidation(mail) && len(nickname) != 0 && len(pass) != 0 && len(mail) != 0 {
			err, u := database.SelectUserData(fmt.Sprintf("SELECT * FROM users WHERE name = '%s' OR mail = '%s';", nickname, mail))
			if err != nil {
				log.Println(err)
			}

			// Проверка на существование пользователя по никнейму / почте
			if u != nil {
				errJson := errrorhandler.GetErrorJson(409, "User already registered")
				w.WriteHeader(409)
				w.Write([]byte(errJson))
			} else {
				// Регистрируем нового пользователя
				err = database.InsertData(fmt.Sprintf("INSERT INTO users (name, password, mail) VALUES ('%s', %s, 'sha256'), '%s');", nickname, pass, mail))
				if err != nil {
					log.Fatalf("Возникла ошибка при вставке значения в таблицу базы данных:\n%v\n", err)
					w.WriteHeader(500)
				}
				w.WriteHeader(200)
			}
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
