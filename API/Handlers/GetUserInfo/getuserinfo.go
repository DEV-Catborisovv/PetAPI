// Обработчик-метод получения данных пользователя.
// Created by Catborisovv (c) 2020-2024

package getuserinfo

import (
	errrorhandler "PetAPI/pkg/ErrrorHandler"
	"PetAPI/pkg/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Userinfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type JSONResponse struct {
	Code   int      `json:"code"`
	Status string   `json:"status"`
	UInfo  Userinfo `json:"uinfo"`
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Получение URL
		if r.URL.Path == "/users/" {
			errJson := errrorhandler.GetErrorJson(400, "Please provide a nickname in the URL")
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(errJson))
		} else {
			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) <= 0 {
				errJson := errrorhandler.GetErrorJson(400, "Please provide a nickname in the URL")
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(errJson))
			} else {
				// Проверка на пустую строку от пользователя
				if len(pathParts) < 2 {
					errJson := errrorhandler.GetErrorJson(400, "Please provide a nickname in the URL")
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")
					w.Write([]byte(errJson))
				} else {
					// Получаем последнюю часть пути как никнейм
					nickname := pathParts[len(pathParts)-1]

					err, User := database.SelectUserData(fmt.Sprintf("SELECT * FROM users WHERE name = '%s';", nickname))
					if err != nil {
						log.Fatalf("Возникла ошибка при выполнении SQL запроса:\n%v\n", err)
					}

					if User.ID == 0 {
						errJson := errrorhandler.GetErrorJson(404, "User is not found")
						w.WriteHeader(http.StatusBadRequest)
						w.Header().Set("Content-Type", "application/json")
						w.Write([]byte(errJson))
					} else {
						// Формируем JSON-Ответ
						var UiJson JSONResponse = JSONResponse{
							Code:   200,
							Status: "OK",
							UInfo: Userinfo{
								ID:   User.ID,
								Name: User.Name,
							},
						}

						JsonResp, err := json.Marshal(UiJson)
						if err != nil {
							log.Fatalf("Возникла ошибка при формировании JSON:\n%v\n", err)
						}

						w.WriteHeader(200)
						w.Header().Set("Content-Type", "application/json")
						w.Write([]byte(JsonResp))
					}
				}
			}
		}
	} else {
		errJson := errrorhandler.GetErrorJson(405, fmt.Sprintf("The %s method is not supported for this API method", r.Method))
		w.WriteHeader(405)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(errJson))
	}
}
