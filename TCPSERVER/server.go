package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	user, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("не удалось запустить сервер ", err)
	}
	answer(user)
}
func answer(user net.Listener) {
	for {
		conne, err := user.Accept()
		if err != nil {
			log.Println("Ошибка добавлении пользователя ", err)
			continue
		}
		go workwithclient(conne)
	}
}
func workwithclient(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Println("не удалось прочитать пользователя ", err)
	}
	fmt.Printf("Получено: %s\n", strings.TrimSpace(string(buf)))
	defer conn.Close()
}
