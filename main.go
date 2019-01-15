// LML - LiveMediaList
// Service for choice LiveMedia
// Roadmap
// 1 - Pattern
// 2 - Read start data from file and write to DB
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {
	log.Println("Start Init func: ")
	readStartDb()
}

func main() {
	log.Println("Start Main func: ")
}

func readStartDb() {
	file, err := os.Open("./DB/StartDB.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//читаємо строку з файла
		s := strings.TrimSpace(scanner.Text())
		if s == "" {
			continue
		}
		if strings.HasPrefix(s, "Number: ") {
			//s = strings.TrimLeft(s, "Number:")
			s = strings.TrimPrefix(s, "Number: ")
			//i, err := strconv.Parse (s, 10, 64)
			if n, err := strconv.Atoi(s); err == nil {
				fmt.Printf("%T, %v\n", n, n)
			} else {
				fmt.Println(err)
			}
			continue
		}
		if strings.HasPrefix(s, "Name: ") {
			s = strings.TrimPrefix(s, "Name: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Homepage: ") {
			s = strings.TrimPrefix(s, "Homepage: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Download: ") {
			s = strings.TrimPrefix(s, "Download: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Wikipedia: ") {
			s = strings.TrimPrefix(s, "Wikipedia: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Distrowatch: ") {
			s = strings.TrimPrefix(s, "Distrowatch: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Size (mebibytes):") {
			s = strings.TrimPrefix(s, "Size (mebibytes):")
			s = strings.TrimSpace(s)
			// s = split  strings.Split("a,b,c", ",")
			mapStr := strings.Split(s, "-")
			fmt.Println(mapStr)
			continue
		}
		//fmt.Println("----------------------------")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
