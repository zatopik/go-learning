package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type guy struct {
	Name    string `json:"Иван"`
	Subname string `json:"Иванов"`
}

func main() {
	chanel := make(chan (string))
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"name": "Иван"}
		json.NewEncoder(w).Encode(response)
	})
	http.HandleFunc("/greet", fetchWords(chanel))
	go func() {
		file, err := os.OpenFile("names.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal("Ошибка открытия файла:", err)
		}
		defer file.Close()

		for name := range chanel {
			_, err := file.WriteString(name + "\n")
			if err != nil {
				log.Println("Ошибка записи в файл:", err)
			}
		}
	}()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
	log.Println("Сервер запущен")
}
func fetchWords(ch chan<- string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, `{"error":"Параметр name обязателен"}`, http.StatusBadRequest)
			return
		}
		ch <- name
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Привет, " + name + "!",
		})
	}
}
