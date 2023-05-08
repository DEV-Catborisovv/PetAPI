package errrorhandler

import (
	"encoding/json"
	"log"
)

type ErrResp struct {
	Code  uint
	Error string
}

func GetErrorJson(code uint, errorMsg string) []byte {
	errResp := ErrResp{
		code,
		errorMsg,
	}
	b, err := json.Marshal(errResp)
	if err != nil {
		log.Fatalf("Возникла ошибка создания JSON:\n%v\n", err)
	}
	return b
}
