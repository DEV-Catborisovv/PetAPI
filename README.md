# PET-API
Доброго времени суток. Это один из моих Pet-проектов на Go: REST-API.  

Я постарался соблюсти главные концепции разработки таких проектов. В этом проекте создана оптимальная и расширяемая структура.  

## Основные методы

Этот API не имеет большого функционала и является лишь демонстрационным проектом


### [POST]: Метод создания нового пользователя в базе данных  
```
/newuser
```  
В теле запроса требуется передать три значения:  
```
name=value
password=value
mail=value@mail.domain
```
curl
```
curl -d "name=VALUE&password=VALUE&mail=VALUE@mail.domain" -X POST http://localhost:7777/newuser
```
После выполнения метода в базе данных будет создан новый пользователь. Пароль хешируется при помощи хшш-функции SHA256

### [POST]: Метод авторизации пользователя
```
/login
```
В теле запроса передается два значения:
```
name=value
password=value
```
curl
```
curl -d "name=value&password=value" -X POST http://localhost:7777/login
```
Возвращаемые значения (JSON):  
Ошибка или информация о пользователе и токен авторизации.
### [GET]: Метод получения данных пользователя
```
/users/%username%
```

Чтобы получить обедоступную информацию о пользователе Вы можете использовать этот метод API

curl
```
curl -X GET http://localhost:7777/users/%username%                                                                
```
Возвращаемые значения (JSON):  
Ошибка или информация о пользователе.

# Настойка конфига
В этом проекте я решил использовать TOML для конфигурации всего проекта, найти его можно по пути

```
%PROJECT%/confgis/config.toml
```
Стандартный порт для запуска веб-сервера: 7777
```
# This is a TOML Config for my REST-API app
# Created by Catborisovv (c) 2020-2024

# Конфигурация HTTP-Сервера
[HttpSrv]
Host = "localhost"
Port = 7777

# Конфигурация Базы-Данных PostgreSQL
[DataBase]
DB_USER = ""
DB_PASS = ""
DB_NAME = ""
```

# Создание таблиц БД
Создание таблицы пользователей
```
CREATE TABLE users (
  id SERIAL,
  name VARCHAR(50) NOT NULL,
  password VARCHAR(256) NOT NULL,
  mail VARCHAR(50) NOT NULL,
  admin INT NOT NULL
);
```
Создание таблицы активных сессий
```
CREATE TABLE sessions (
  id SERIAL,
  TOKEN VARCHAR(256) NOT NULL,
  name VARCHAR(256) NOT NULL
);
```

# Авторство
Created by Catborisovv (c) 2020-2024
