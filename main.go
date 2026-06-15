package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB

func main() {
	var err error
	db, err = initDB()
	if err != nil {
		fmt.Println("Ошибка базы данных:", err)
		return
	}
	defer db.Close()

	http.HandleFunc("/api/info", handlerInfo)
	http.HandleFunc("/api/time", handlerTime)
	http.HandleFunc("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlerMessagesPost(w, r)
		case http.MethodGet:
			handlerMessagesGet(w, r)
		default:
			handlerNotFound(w, r)
		}
	})
	http.HandleFunc("/", handlerNotFound)

	fmt.Println("Сервер запущен на http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}