package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID    int    `json:"ID"`
	Title string `json:"Title"`
	Done  bool   `json:"Done"`
}

func getwords(reader *bufio.Reader) Task {
	result := Task{}
	fmt.Print("Введите ID: ")
	idInput, _ := reader.ReadString('\n')
	result.ID, _ = strconv.Atoi(strings.TrimSpace(idInput))

	fmt.Print("Введите название: ")
	titleInput, _ := reader.ReadString('\n')
	result.Title = strings.TrimSpace(titleInput)

	fmt.Print("Завершено? (true/false): ")
	doneInput, _ := reader.ReadString('\n')
	result.Done, _ = strconv.ParseBool(strings.TrimSpace(doneInput))
	return result
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Главная страница")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "limbus company")
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "we@gmail.com")
	})
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Гость"
		}
		City := r.URL.Query().Get("City")
		if City == "" {
			City = "неизвестного города"
		}
		fmt.Fprintln(w, name, City)

	})
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		reader := bufio.NewReader(os.Stdin)
		work, err := os.ReadFile("tasks.txt")
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		if string(work) != "" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(work)
		} else {
			words := getwords(reader)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(words)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
