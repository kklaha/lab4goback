package datab

import (
	"database/sql"
	"fmt"
	"lab4back/models"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*sql.DB, error) {
	cfg := models.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBname:   os.Getenv("DB_NAME"),
	}
	if cfg.User == "" {
		return nil, fmt.Errorf("Пустое имя пользователя")
	}
	if cfg.Password == "" {
		return nil, fmt.Errorf("Пароль не может быть пустым")
	}
	if cfg.Host == "" {
		cfg.Host = "127.0.0.1"
	}
	if cfg.Port == "" {
		cfg.Port = "3306"
	}
	if cfg.DBname == "" {
		cfg.DBname = cfg.User
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&tls=false", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBname)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := DB.Ping(); err != nil {
		DB.Close()
		return nil, err
	}
	return DB, nil

}
func SaveToDB(db *sql.DB, data models.Form) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		INSERT INTO form (user_FIO,phonenumb,email,dateofb,gender,biography,accepted) VALUES(?,?,?,?,?,?,?)
		`, data.FIO, data.PhoneNumber, data.Email, data.Dateofb, data.Gender, data.Biography, data.Accepted)
	if err != nil {
		return err
	}
	user_ID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	for _, lang := range data.Favlanguages {
		_, err := tx.Exec(`
		INSERT INTO languages(user_id,language) VALUES(?,?)
		`, user_ID, lang)
		if err != nil {
			return err
		}
	}

	return tx.Commit()

}
