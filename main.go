// LML - LiveMediaList
// Service for choice LiveMedia
// Roadmap
// 1 - Pattern
// 2 - Read start data from file and write to DB
// 3 - Read config file
// 4 - Read file variable
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	//"gopkg.in/yaml.v2"
)

func init() {
	log.Println("Start Init func: ")
	readStartDb()
	readValueDb()

}

func main() {
	log.Println(" ")
	log.Println("Start Main func: ")
	log.Println(" ")
}

func readStartDb() {
	file, err := os.Open("./db/StartDB.txt")
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
			fmt.Println("----------------------------")
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
		if strings.HasPrefix(s, "Last Stable Version: ") {
			s = strings.TrimPrefix(s, "Last Stable Version: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Last Release: ") {
			s = strings.TrimPrefix(s, "Last Release: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Purpose: ") {
			s = strings.TrimPrefix(s, "Purpose: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Operating System: ") {
			s = strings.TrimPrefix(s, "Operating System: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Primary Language(s): ") {
			s = strings.TrimPrefix(s, "Primary Language(s): ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "State: ") {
			s = strings.TrimPrefix(s, "State: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Media: ") {
			s = strings.TrimPrefix(s, "Media: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Architecture: ") {
			s = strings.TrimPrefix(s, "Architecture: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Note: ") {
			s = strings.TrimPrefix(s, "Note: ")
			fmt.Println(s)
			continue
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
