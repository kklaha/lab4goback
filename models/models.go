package models

import "time"

type Form struct {
	FIO          string
	PhoneNumber  string
	Email        string
	Dateofb      time.Time
	Gender       string
	Favlanguages []string
	Biography    string
	Accepted     bool
}
type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBname   string
}
type PageData struct {
	Data    Form
	Errors  map[string]string
	Success string
}
