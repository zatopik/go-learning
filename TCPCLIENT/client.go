package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		log.Fatal("вы ввели много значений")
	}
	num, _ := strconv.Atoi(os.Args[1])
	if num > 5 {
		log.Fatal("Нельзя вводить больше 5")
	}
	for range num {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			log.Fatal("Ошибка подключиться к серверу", err)
		}
		reader := bufio.NewReader(os.Stdin)
		defer conn.Close()
		messasger := messasge(reader)
		conn.Write([]byte(messasger))
		fmt.Println("Сообщение отправлено:", messasger)
	}
}
func messasge(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Ошибка ввода", err)
	}
	input = strings.TrimSpace(input)
	return input
}
