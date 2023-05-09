// Пакет генерации API-Ключей

package generatetoken

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateAPIKey(length int) (string, error) {
	// Вычисляем количество байт, необходимых для ключа нужной длины
	numBytes := length / 4 * 3
	if length%4 > 0 {
		numBytes += 3
	}

	// Генерируем случайную последовательность байт нужной длины
	bytes := make([]byte, numBytes)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Кодируем байты в base64 и обрезаем до нужной длины
	key := base64.StdEncoding.EncodeToString(bytes)[:length]

	return key, nil
}
