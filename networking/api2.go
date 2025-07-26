package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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
type FileContent struct {
	Content string `json:"content"`
}
type Name struct {
	V1 string `json:"name"`
}
type City struct {
	V2 string `json:"city"`
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
		fmt.Fprint(w, "Main page")
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
		city := r.URL.Query().Get("city")
		if city == "" {
			city = "неизвестного города"
		}
		truename := Name{V1: name}
		truecity := City{V2: city}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(truename)
		json.NewEncoder(w).Encode(truecity)

	})
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		reader := bufio.NewReader(os.Stdin)
		work, err := os.ReadFile("tasks.txt")
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		if string(work) != "" {
			otvet := &FileContent{Content: string(work)}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(otvet)
		} else {
			words := getwords(reader)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(words)
		}
	})
	http.HandleFunc("/test-form", func(w http.ResponseWriter, r *http.Request) {
		html := `
    <html>
    <body>
        <h2>Тест создания задачи</h2>
        <button onclick="createTask()">Создать задачу</button>
        <script>
            function createTask() {
                fetch('/tasks/add', {
                    method: 'POST',
                    headers: {'Content-Type': 'application/json'},
                    body: JSON.stringify({
                        id: Date.now(),
                        title: "Тестовая задача " + new Date().toLocaleTimeString(),
                        done: false
                    })
                })
                .then(response => response.json())
                .then(data => alert('Задача создана: ' + JSON.stringify(data)))
                .catch(error => alert('Ошибка: ' + error));
            }
        </script>
    </body>
    </html>
    `
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
	})
	http.HandleFunc("/tasks/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Используйте POST-запрос")
			return
		}
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			fmt.Fprint(w, "Ожидается application/json")
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Ошибка чтения тела: %v", err)
			return
		}
		defer r.Body.Close()
		var newTask Task
		if err := json.Unmarshal(body, &newTask); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Неверный JSON: %v", err)
			return
		}
		if newTask.ID <= 0 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprint(w, "ID должен быть положительным числом")
			return
		}
		fmt.Printf("Получена новая задача: %+v\n", newTask)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "created",
			"task":   newTask,
		})
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
