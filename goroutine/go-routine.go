package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ввести до скольки поднимаемся")
	n := getnumber(reader)
	fmt.Println("Ввести количество горутин`Работников`")
	routine := getnumber(reader)
	results := make(chan int, routine)
	chunkSize := n / routine
	leftover := n % routine

	start := 1
	for i := 0; i < routine; i++ {
		end := start + chunkSize - 1
		if i == routine-1 {
			end += leftover
		}

		go thework(start, end, results)
		start = end + 1
	}
	total := 0
	for i := 0; i < routine; i++ {
		total += <-results
	}
	fmt.Printf("Сумма от 1 до %d = %d\n", n, total)
}
func getnumber(reader *bufio.Reader) int {
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Ошибка ввода", err)
		}
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil || number < 0 {
			log.Fatal("Введи число и чтобы оно было положительное", err)
		}
		return number
	}

}
func thework(start, end int, ch chan int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	ch <- sum
}
