package form

import (
	"database/sql"
	"html/template"
	"lab4back/cookie"
	"lab4back/datab"
	"lab4back/models"
	"lab4back/validation"
	"log"
	"net/http"
	"os"
	"time"
)

func ParseForm(r *http.Request) (models.Form, map[string]string) {
	r.ParseForm()
	data := models.Form{
		FIO:          r.FormValue("fio"),
		PhoneNumber:  r.FormValue("phone"),
		Email:        r.FormValue("email"),
		Gender:       r.FormValue("gender"),
		Favlanguages: r.Form["languages"],
		Biography:    r.FormValue("biography"),
		Accepted:     r.Form.Has("accepted"),
	}
	birthDate := r.FormValue("birthdate")
	if birthDate != "" {
		date, err := time.Parse("2006-01-02", birthDate)
		if err == nil {
			data.Dateofb = date
		}
	}
	errors := make(map[string]string)
	if err := validation.ValidateFio(data.FIO); err != nil {
		errors["fio"] = err.Error()
	}
	if err := validation.ValidateEmail(data.Email); err != nil {
		errors["email"] = err.Error()
	}
	if err := validation.ValidateGender(data.Gender); err != nil {
		errors["gender"] = err.Error()
	}
	if err := validation.ValiDateOfBirthday(data.Dateofb); err != nil {
		errors["birthdate"] = err.Error()
	}
	if err := validation.ValidateAccept(data.Accepted); err != nil {
		errors["accepted"] = err.Error()
	}
	if err := validation.ValidateBio(data.Biography); err != nil {
		errors["bio"] = err.Error()
	}
	if err := validation.ValidatePhoneNumber(data.PhoneNumber); err != nil {
		errors["phone"] = err.Error()
	}
	if err := validation.ValidateLanguages(data.Favlanguages); err != nil {
		errors["languages"] = err.Error()
	}
	return data, errors
}
func FormHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		redirectPath := os.Getenv("SCRIPT_NAME")
		if redirectPath == "" {
			redirectPath = "/lab4/cgi-bin/cookie.cgi"

		}
		tmpl := template.New("index.html").Funcs(template.FuncMap{
			"contains": func(slice []string, item string) bool {
				for _, s := range slice {
					if s == item {
						return true
					}
				}
				return false
			},
		})

		tmpl, err := tmpl.ParseFiles("/home/u82188/www/lab4/html/index.html")
		if err != nil {
			http.Error(w, "Template parse error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if r.Method == "GET" {
			formData, _ := cookie.GetFormDataFromCookies(r)
			errors, _ := cookie.GetErrorsFromCookie(r)
			cookie.DeleteCookie(w, "form_errors")
			successMsg := ""
			if r.URL.Query().Get("saved") == "1" {
				successMsg = "<div style='color:green; margin: 20px 0;'><h1>Данные сохранены!</h1><a href='/'>← На главную</a></div>"
			}
			err := tmpl.Execute(w, models.PageData{
				Data:    formData,
				Errors:  errors,
				Success: successMsg,
			})
			if err != nil {
				http.Error(w, "Template execute error: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if r.Method == "POST" {
			data, errors := ParseForm(r)
			if len(errors) > 0 {
				if err := cookie.SaveErrorsToCookie(w, errors); err != nil {
					log.Printf("save errors cookie: %v", err)
				}
				if err := cookie.SaveDataToCookies(w, data); err != nil {
					log.Printf("save data cookie: %v", err)
				}
				http.Redirect(w, r, redirectPath, http.StatusSeeOther)
				return
			}
			if err := datab.SaveToDB(db, data); err != nil {
				log.Printf("DB save error: %v", err)
				http.Error(w, "Ошибка сохранения", http.StatusInternalServerError)
				return
			}
			cookie.DeleteCookie(w, "form_data")
			cookie.DeleteCookie(w, "form_errors")
			http.Redirect(w, r, redirectPath+"?saved=1", http.StatusSeeOther)
			return
		}
	}
}
