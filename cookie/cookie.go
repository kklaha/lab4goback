package cookie

import (
	"encoding/json"
	"errors"
	"lab4back/models"
	"net/http"
	"net/url"
)

const maxCookieSize = 3500

func SetCookie(w http.ResponseWriter, name string, value string) error {
	if len(value) > maxCookieSize {
		return errors.New("слишком большое значение cookie")
	}
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		Path:     "/",
		MaxAge:   31536000,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}
func GetCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	val, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return "", err
	}
	return val, nil
}
func DeleteCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}
func SaveErrorsToCookie(w http.ResponseWriter, errorMap map[string]string) error {
	cookieErrors, err := json.Marshal(errorMap)
	if err != nil {
		return err
	}
	return SetCookie(w, "form_errors", string(cookieErrors))
}

func GetFormDataFromCookies(r *http.Request) (models.Form, error) {
	cookies, err := GetCookie(r, "form_data")
	if err != nil {
		return models.Form{}, err
	}
	var data models.Form
	if err := json.Unmarshal([]byte(cookies), &data); err != nil {
		return models.Form{}, err

	}
	return data, nil
}
func SaveDataToCookies(w http.ResponseWriter, data models.Form) error {
	cookieData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return SetCookie(w, "form_data", string(cookieData))
}
func GetErrorsFromCookie(r *http.Request) (map[string]string, error) {
	cookie, err := GetCookie(r, "form_errors")
	if err != nil {
		return nil, err
	}
	var cookieErrors map[string]string
	if err := json.Unmarshal([]byte(cookie), &cookieErrors); err != nil {
		return nil, err
	}
	return cookieErrors, nil
}
