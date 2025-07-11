package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type guy struct {
	Name    string `json:"Иван"`
	Subname string `json:"Иванов"`
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"name": "Иван"}
		json.NewEncoder(w).Encode(response)
	})
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		answer := r.URL.Query().Get("name")
		if answer == "" {
			http.Error(w, `{"error":"Параметр name не найден"}`, http.StatusBadRequest)
		}
		_, err := fmt.Fprint(w, answer)
		if err != nil {
			http.Error(w, `{"error":"неполучилось вывести ответ"}`, http.StatusBadRequest)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
	log.Println("Сервер запущен")
}
