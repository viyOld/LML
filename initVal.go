package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	file, err := os.Open("./db/Values.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	flag := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//читаємо строку з файла
		s := strings.TrimSpace(scanner.Text())
		//пусті строкі та коменти відхиляємо
		if s == "" {
			continue
		}
		if strings.HasPrefix(s, "//") {
			continue
		}
		if strings.HasPrefix(s, "#") {
			continue
		}

		if strings.HasPrefix(s, "State:") {
			flag = 1
			// s = strings.TrimPrefix(s, "State: ")
			fmt.Println(s)
			continue
		}

		if strings.HasPrefix(s, "Target:") {
			flag = 2
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Name: ") {
			s = strings.TrimPrefix(s, "Name: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Name: ") {
			s = strings.TrimPrefix(s, "Name: ")
			fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Name: ") {
			s = strings.TrimPrefix(s, "Name: ")
			fmt.Println(s)
			continue
		}
		if flag == 1 {
			st := strings.Split(s, ":")
			//ip, port := s[0], s[1]
			//State[st[1]] = 0
			fmt.Println(st)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(State)
}
