// Обработчик события авторизации пользователя через API

package autotizate

import (
	errrorhandler "PetAPI/pkg/ErrrorHandler"
	"PetAPI/pkg/database"
	generatetoken "PetAPI/pkg/generateToken"
	validdata "PetAPI/pkg/validData"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Session struct {
	Code   int
	Status string
	Token  string
	Name   string
}

func AutotizateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Парсинг тела запроса
		err := r.ParseForm()
		if err != nil {
			log.Printf("Возникла ошибка при получении тела запроса:\n%v\n", err)
		}

		// Получение данных из тела запроса
		var nickname string = strings.TrimSpace(r.FormValue("name"))
		var pass string = strings.TrimSpace(r.FormValue("password"))

		// Валидация входных данных
		if validdata.DataValidation(nickname) == true && validdata.DataValidation(pass) == true && len(nickname) > 0 && len(pass) > 0 {
			err, u := database.SelectUserData(fmt.Sprintf("SELECT * FROM users WHERE name = '%s' AND password = encode(digest('%s', 'sha256'), 'hex');", nickname, pass))
			if err != nil {
				log.Println(err)
			}

			// Проверка на существование пользователя по никнейм
			if u == nil {
				errJson := errrorhandler.GetErrorJson(404, "This user does not exist")
				w.WriteHeader(404)
				w.Write([]byte(errJson))
			} else {
				uToken, err := generatetoken.GenerateAPIKey(92)
				if err != nil {
					log.Fatalf("Возникла ошибка при генерации API-ключа:\n%v\n", err)
				}
				err = database.InsertData(fmt.Sprintf("INSERT INTO sessions (TOKEN, name) VALUES ('%s', '%s');", uToken, nickname))
				if err != nil {
					log.Fatalf("Возникла ошибка обработки создания сессии пользователя:\n%v\n", err)
				}

				// Генерация JSON-Ответа
				s := Session{200, "ok", uToken, nickname}
				sJson, err := json.Marshal(s)

				if err != nil {
					log.Fatalf("Возникла ошибка при создании JSON-ответа:\n%v\n", err)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(200)
				w.Write([]byte(sJson))
			}
		}
	} else {
		errJson := errrorhandler.GetErrorJson(405, fmt.Sprintf("The %s method is not supported for this API method", r.Method))
		w.WriteHeader(405)
		w.Write([]byte(errJson))
	}
}
