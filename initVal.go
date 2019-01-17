package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//	State: Target: OS: Media: Architecture:
var (
	State        map[string]byte
	Target       map[string]byte
	OS           map[string]byte
	Media        map[string]byte
	Architecture map[string]byte
)

func readValueDb() {
	State := make(map[string]byte)
	Target := make(map[string]byte)
	OS := make(map[string]byte)
	Media := make(map[string]byte)
	Architecture := make(map[string]byte)
	// відкриваємо файл з значеннями
	file, err := os.Open("./db/Values.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	flag := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// читаємо строку з файла
		s := strings.TrimSpace(scanner.Text())
		// пусті строкі та коменти відхиляємо
		if s == "" {
			continue
		}
		if strings.HasPrefix(s, "//") {
			continue
		}
		if strings.HasPrefix(s, "#") {
			continue
		}
		// перевірка на новий список
		if strings.HasPrefix(s, "State:") {
			flag = 1
			continue
		}
		if strings.HasPrefix(s, "Target:") {
			flag = 2
			continue
		}
		if strings.HasPrefix(s, "OS:") {
			flag = 3
			continue
		}
		if strings.HasPrefix(s, "Media:") {
			flag = 4
			continue
		}
		if strings.HasPrefix(s, "Architecture:") {
			flag = 5
			continue
		}
		// ще одна контрольна перевірка
		if strings.Contains(s, ":") == false {
			continue
		}
		// розділяємо строку на дві
		st := strings.Split(s, ":")
		n, _ := strconv.ParseUint(st[0], 10, 8)
		//fmt.Println(st)
		switch flag {
		case 1:
			State[st[1]] = byte(n)
		case 2:
			Target[st[1]] = byte(n)
		case 3:
			OS[st[1]] = byte(n)
		case 4:
			Media[st[1]] = byte(n)
		case 5:
			Architecture[st[1]] = byte(n)
		default:
			//fmt.Println(st)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(State)
	// fmt.Println(Target)
	// fmt.Println(OS)
	// fmt.Println(Media)
	// fmt.Println(Architecture)
}

// --------------------------------------------------------------------------------------------------
func readStartDb() {
	lmlDB := []listMediaLive{listMediaLive{}}
	lmlDB[0].name = "test1"
	//lmlDB[1] = listMediaLive{listMediaLive{}}
	//numbers = append(numbers, 6)
	lmlDB = append(lmlDB, listMediaLive{})
	lmlDB[1].name = "test2"
	//lmlDB[1].name = "test3"
	fmt.Println(lmlDB[0].name)
	fmt.Println(lmlDB[1].name)

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
				//lmlDB[0].name = "xx"
			} else {
				fmt.Println(err)
			}

			continue
		}
		if strings.HasPrefix(s, "Name: ") {
			s = strings.TrimPrefix(s, "Name: ")
			//lmlDB[0].name = "nlknc"
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
	fmt.Println(lmlDB[0].name)
	fmt.Println(lmlDB[1].name)
}
