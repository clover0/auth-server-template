package db

import (
	"fmt"
	"log"

	"auth465/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB() *sqlx.DB {
	user := config.Conf.DB.User
	host := config.Conf.DB.Host
	port := config.Conf.DB.Port
	password := config.Conf.DB.Password
	dbName := config.Conf.DB.DbName
	maxIdleConn := config.Conf.DB.MaxIdleConn
	maxOpenConn := config.Conf.DB.MaxOpenConn

	ds := fmt.Sprintf(
		"user=%s host=%s port=%d password=%s dbname=%s sslmode=disable",
		user,
		host,
		port,
		password,
		dbName,
	)
	db, err := sqlx.Connect("postgres", ds)
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)
	log.Printf("database connected: user=%s host=%s port=%d password=*** dbname=%s sslmode=disable maxidle=%s maxopen=%s",
		user,
		host,
		port,
		dbName,
		maxIdleConn,
		maxOpenConn,
	)

	return db
}
