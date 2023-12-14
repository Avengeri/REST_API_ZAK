# Программа для тренировки замены БД.

## Реализуется по щелчку пальцев заменой одной переменной

### Порядок запуска программы:


Для работы с БД типа PostgresQL:

1. Установите Docker на свой компьютер. Инструкцию по установке можно найти [здесь](https://www.docker.com/)
2. Проверьте установлен ли `Docker Compose` с помощью команды  `docker compose version`. Если он не установлен, то не мои проблемы, решите сами этот вопрос! :thinking:
3. В файле main.go закомментируйте Redis и раскомментируйте Postgres
4. 
```Go
	db, err := postgres.NewPostgresDB(postgres.Config{
Host:     viper.GetString("db.host"),
Port:     viper.GetString("db.port"),
Username: viper.GetString("db.username"),
DBName:   viper.GetString("db.dbname"),
SSLMode:  viper.GetString("db.sslmode"),
Password: os.Getenv("POSTGRES_PASSWORD"),
})

if err != nil {
log.Fatalf("Ошибка создания Postgres: %s", err.Error())
}
repo := repository.NewStorageUsersPostgres(db)
```

Для работы с БД типа Redis:

1. Установите Docker на свой компьютер. Инструкцию по установке можно найти [здесь](https://www.docker.com/)
2. Проверьте установлен ли `Docker Compose` с помощью команды  `docker compose version`. Если он не установлен, то не мои проблемы, решите сами этот вопрос! :thinking:
3.  В файле main.go закомментируйте Postgres и раскомментируйте Redis

```Go
	rdb, err := redis_storage.NewRedisClient(redis_storage.Config{
Addr: viper.GetString("rdb.address"),
})
repo := repository.NewStorageUsersRedis(rdb)
```

## Инструкция по подключению
Замените название файла `.env.example` на `.env` и укажите значения переменных

## Команды  для запуска
```shell
docker pull redis:latest  # Скачать образ Redis
```
```shell
docker pull postgres:latest # Скачать образ PostgresQL
```
```shell
make start # Создает миграции и запускает docker compose
```
```shell
make stop # Останавливает docker и удаляет миграции
```
```shell
make restart # Перезапускает docker и миграции
```
```shell
make build # Собирает приложение в exe файл
```
```shell
make run # Собирает приложение в exe файл и запускает его
```
```shell
make migrate_up # Создает миграции
```
```shell
make migrate_down # Удаляет миграции
```

## Инструкция как запустить в продакшне:
- Я хз
## Описание Endpoints
### Добавить пользователя
POST http://localhost:8080/user/
```json
{
"id": 1,
"name": "ZAK-pipisya"
}
```
### Добавить пользователя
GET http://localhost:8080/user/:id - заменить :id на действующий id

### Проверить пользователя
GET http://localhost:8080/user/check/:id - заменить :id на действующий id

### Удалить пользователя
DELETE http://localhost:8080/user/:id - заменить :id на действующий id

### Получить все id пользователей
GET http://localhost:8080/user/get_all

> Если вы хотите пожаловаться, то обязательно пишите [сюда](https://t.me/zak47) 