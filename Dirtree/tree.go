package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	countofdeep := 0
	if len(os.Args) > 2 {
		log.Fatal("Много символов в командной строке ")
	}
	reader := bufio.NewReader(os.Stdin)
	thedir := getthepath(reader)
	dirtree(countofdeep, 2, thedir)

}
func getthepath(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Ошибка ввода")
	}
	input = strings.TrimSpace(input)
	return input
}
func dirtree(countofdeep int, max int, fin string) {
	if countofdeep > max {
		log.Fatal("Больше лимита")
	}
	thedir, _ := os.ReadDir(fin)
	for i, dirs := range thedir {
		isLastEntry := i == len(thedir)-1
		if dirs.IsDir() {
			if countofdeep == 1 {
				fmt.Print("│")
				fmt.Print("\t"+"└──", dirs, "\n")
			} else {
				fmt.Print("└──", dirs, "\n")
				full := path.Join(fin, dirs.Name())
				dirtree(countofdeep+1, max, full)
			}
		}
		if countofdeep == 1 && dirs.IsDir() != true {
			if isLastEntry {
				fmt.Print("│")
				fmt.Print("\t"+"└──", dirs, "\n")

			} else {
				fmt.Print("│")
				fmt.Print("\t"+"├──", dirs, "\n")
			}
		}
		if countofdeep == 0 && dirs.IsDir() != true {
			if isLastEntry {
				fmt.Print("└──", dirs, "\n")
			} else {
				fmt.Print("├──", dirs, "\n")
			}
		}
	}
}
