package tomldecode

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type Config struct {
	HttpSrv struct {
		Host string
		Port int
	}
	DataBase struct {
		DB_USER string
		DB_PASS string
		DB_NAME string
	}
}

func DecodeConfigTOML() (error, *Config) {
	// Создаем экземпляр структуры config
	var config Config

	// Read toml file
	TomlData, err := ioutil.ReadFile("configs/config.toml")
	if err != nil {
		return err, nil
	}

	// Decode TOML
	_, err = toml.Decode(string(TomlData), &config)
	if err != nil {
		return err, nil
	}
	return nil, &config
}
