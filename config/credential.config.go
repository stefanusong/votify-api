package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetDSN() string {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	dbname := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port)
	return dsn
}

func GetJWTSecretKey() []byte {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	secretKey := viper.GetString("JWT_SECRET")
	return []byte(secretKey)
}
