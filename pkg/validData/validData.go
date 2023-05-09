// Пакет валидации данных
// Created by Catborisovv (c) 2020-2024

package validdata

import "regexp"

// Функция валидации данных пользователя
func DataValidation(data string) bool {
	pattern := `^[a-zA-Z0-9]+$`
	re := regexp.MustCompile(pattern)

	if re.MatchString(data) {
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
