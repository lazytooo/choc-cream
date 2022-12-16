package config

import (
	"os"
	"strconv"
)

var (
	DBMaxOpenConns int
	DBMaxIdleConns int
	DBHost         string
	DBPort         string
	DBUser         string
	DBName         string
	DBPass         string
	HTTPAddr       string
	RedisAddr      string
	RedisPass      string
)

func init() {
	DBHost = os.Getenv("MYSQL_HOST")
	DBPort = os.Getenv("MYSQL_PORT")
	DBUser = os.Getenv("MYSQL_USER")
	DBName = os.Getenv("MYSQL_NAME")
	DBPass = os.Getenv("MYSQL_PASS")
	DBMaxOpenConns, _ = strconv.Atoi(os.Getenv("MYSQL_MAX_OPEN_CONNS"))
	DBMaxIdleConns, _ = strconv.Atoi(os.Getenv("MYSQL_MAX_IDLE_CONNS"))
	HTTPAddr = os.Getenv("HTTP_LISTEN_ADDR")
	RedisAddr = os.Getenv("REDIS_ADDR")
	RedisPass = os.Getenv("REDIS_PASS")
}
