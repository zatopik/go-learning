package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type he struct {
	number          int
	numbertwotime   int
	numberthreetime int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	number := getnumber(reader)
	cx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result := stepen(generatenumbers(number, cx), cx)
	for i := range result {
		fmt.Printf("Число: %d, Квадрат: %d, Куб: %d\n",
			i.number, i.numbertwotime, i.numberthreetime)
	}
}

func generatenumbers(number int, cx context.Context) <-chan int {
	out := make(chan (int))
	go func() {
		for i := 1; i <= number; i++ {
			out <- i
		}
		select {
		case <-cx.Done():
			log.Fatal("программа сликшом долго работает")
		default:
		}
		defer close(out)
	}()
	return out
}
func stepen(numbers <-chan int, cx context.Context) <-chan he {
	out := make(chan (he))
	go func() {
		for i := range numbers {
			if cx.Err() != nil {
				log.Println("Обработка прервана")
				return
			}
			out <- he{
				number:          i,
				numbertwotime:   i * i,
				numberthreetime: i * i * i,
			}
			select {
			case <-cx.Done():
				log.Fatal("программа сликшом долго работает")
			default:
			}

		}
		defer close(out)
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
