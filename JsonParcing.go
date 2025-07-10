package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsonreader() {
	File, err := os.ReadFile("")
	if err != nil {
		log.Fatal("Ошибка открытия", err)
	}
	var users []User
	err = json.Unmarshal(File, &users)
	if err != nil {
		log.Fatal("Ошибка парсинга", err)
	}
	adduser := User{Name: "Stepan", Age: 18}
	users = append(users, adduser)
	newppl, err := json.MarshalIndent(users, "", "")
	err = os.WriteFile("users", newppl, 0644)
	if err != nil {
		log.Fatal("Ошибка записывания", err)
	}
	fmt.Println("User был успешно добавлен и записан в файл")
}
