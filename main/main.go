package main

import (
	"bufio"
	"fmt"
	"lab4back/datab"
	"lab4back/form"
	"net/http"
	"net/http/cgi"
	"os"
	"strings"
)

func loadEnv(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			os.Setenv(key, value)
		}
	}
}

func main() {
	loadEnv("/home/u82188/www/lab3/cgi-bin/.env")

	DB, err := datab.ConnectToDB()
	if err != nil {
		return
	}
	defer DB.Close()

	http.HandleFunc("/", form.FormHandler(DB))
	if err := cgi.Serve(nil); err != nil {
		fmt.Println("Status: 500 Internal Server Error")
		fmt.Println("Content-Type: text/plain")
		fmt.Println()
		fmt.Println("CGI error:", err)
	}

}
