package validation

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var fioVal = regexp.MustCompile(`^[a-zA-Zа-яА-Я\s\-\']+$`)
var emailVal = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
var numVal = regexp.MustCompile(`^\+?[0-9]{10,15}$`)
var validLanguages = map[string]bool{
	"Pascal": true, "C": true, "C++": true, "JavaScript": true,
	"PHP": true, "Python": true, "Java": true, "Haskell": true,
	"Clojure": true, "Prolog": true, "Scala": true, "Go": true,
}

func ValidateGender(gender string) error {
	validGenders := map[string]bool{"male": true, "female": true}
	if !validGenders[strings.ToLower(gender)] {
		return fmt.Errorf("Некорректный пол")
	}
	return nil
}
func ValidateFio(fio string) error {
	if len(fio) == 0 {
		return fmt.Errorf("Имя не может быть пустым")
	}
	if len(fio) >= 150 {
		return fmt.Errorf("Некорректная длина")

	}
	if !fioVal.MatchString(fio) {
		return fmt.Errorf("Некорректные символы, доступны только буквы и -")
	}
	return nil
}

func ValidateLanguages(langs []string) error {
	if len(langs) == 0 {
		return fmt.Errorf("Выберите хотя бы один язык")
	}
	for _, lang := range langs {
		if !validLanguages[lang] {
			return fmt.Errorf("Недопустимый язык программирования: %s , выберите из списка", lang)
		}
	}
	return nil
}
func ValiDateOfBirthday(dateob time.Time) error {
	if dateob.After(time.Now()) {
		return fmt.Errorf("Дата рождения не может быть в будущем")
	}
	if dateob.IsZero() {
		return fmt.Errorf("Некорректная дата")
	}
	return nil
}
func ValidatePhoneNumber(num string) error {
	if len(num) == 0 {
		return fmt.Errorf("Некорректная длина номера")
	}
	if !numVal.MatchString(num) {
		return fmt.Errorf("Некорректный номер телефона, для ввода доступны только цифры")
	}
	return nil
}
func ValidateEmail(email string) error {
	if len(email) == 0 {
		return fmt.Errorf("Email обязателен")
	}
	if !emailVal.MatchString(email) {
		return fmt.Errorf("Некорректный email, для ввода достпуны только: латиница верхнего и нижнего регистра,цифры, . и @")
	}
	return nil
}
func ValidateAccept(accepted bool) error {
	if !accepted {
		return fmt.Errorf("Ознакомьтесь с контрактом")
	}
	return nil
}
func ValidateBio(biography string) error {
	if len(biography) > 1000 {
		return fmt.Errorf("Сделайте биографию короче")
	}
	return nil
}
