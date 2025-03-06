package main

import (
	"api/internal/application"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	err := godotenv.Load("/Users/dfcarvalho/Documents/Impacta-book/impacta-book-api/impacta-book-api/.env")
	if err != nil {
		panic(err)
	}

	cfg := &application.ConfigApplicationDefault{
		Db: &mysql.Config{
			User:   os.Getenv("DB.USERNAME"),
			Passwd: os.Getenv("DB.PASSWORD"),
			Net:    "tcp",
			Addr:   "localhost" + os.Getenv("DB.ADDRESS"),
			DBName: os.Getenv("DB.NAME"),
		},
		Addr: "127.0.0.1" + os.Getenv("SERVER.PORT"),
	}
	app := application.NewApplicationDefault(cfg)

	err = app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

}
