package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type he struct {
	number        int
	numbertwotime int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	number := getnumber(reader)
	result := stepen(generatenumbers(number))
	for i := range result {
		fmt.Printf("Число: %d, Квадрат: %d\n",
			i.number, i.numbertwotime)
	}
}

func generatenumbers(number int) <-chan int {
	out := make(chan (int))
	go func() {
		for i := 1; i <= number; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}
func stepen(numbers <-chan int) <-chan he {
	out := make(chan (he))
	go func() {
		for i := range numbers {
			out <- he{
				number:        i,
				numbertwotime: i * i,
			}

		}
		close(out)
	}()
	return out
}
func getnumber(reader *bufio.Reader) int {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		log.Fatal()
	}
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Ошибка переобразования в число", err)
	}
	return number
}
