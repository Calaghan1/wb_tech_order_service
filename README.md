Для запуска программы укажите в env файле значения для этих переменных окружения
POSTRGRES_CONNECTION=
NATS_STRIMNG_CLASTER=
NATS_STRIMNG_CLIENT_ID_PROD=
NATS_STRIMNG_CLIENT_ID_SUB=
POSTRGRES__TEST_CONNECTION=
Затем поднимите nats-streaming в Docker
docker run -d -p 4222:4222 -p 8222:8222 nats-streaming
И выполните go run main.go