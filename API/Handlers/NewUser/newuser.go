// NewUser API Method for PetAPI-Project
// Created by Catborisovv (c) 2020-2024

package newuser

import (
	errrorhandler "PetAPI/pkg/ErrrorHandler"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	// Запуск горутины (запуск функции в отдельном потоке, чтобы не останавливать работу {API})
	if r.Method == http.MethodPost {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		if r.URL.Query().Get("name") != "" || r.URL.Query().Get("password") != "" {

		} else {
			errJson := errrorhandler.GetErrorJson(400, "Data in URL is not valid")
			w.WriteHeader(400)
			w.Write([]byte(errJson))
		}
	} else {
		errJson := errrorhandler.GetErrorJson(405, "The GET method is not supported for this API method")
		w.WriteHeader(405)
		w.Write([]byte(errJson))
	}
}
