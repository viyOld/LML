package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readValueDb() {
	// State := make(map[string]byte)
	// Target := make(map[string]byte)
	// OS := make(map[string]byte)
	// Media := make(map[string]byte)
	// Architecture := make(map[string]byte)
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
		st[1] = strings.TrimSpace(st[1])
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
	//fmt.Println(Target)
	//fmt.Println(OS)
	// fmt.Println(Media)
	// fmt.Println(Architecture)
}

// --------------------------------------------------------------------------------------------------
func readStartDb() {
	var (
		n uint64
		i = -1
	)

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
		if strings.HasPrefix(s, "Number:") {
			s = strings.TrimPrefix(s, "Number:")
			s := strings.TrimSpace(s)
			n, err = strconv.ParseUint(s, 10, 32)
			if err != nil {
				panic(err)
			}
			i++
			lmlDB = append(lmlDB, listMediaLive{})
			lmlDB[i].Number = int(n)
			continue
		}
		if strings.HasPrefix(s, "Name:") {
			s = strings.TrimPrefix(s, "Name:")
			s := strings.TrimSpace(s)
			lmlDB[i].Name = s
			continue
		}
		if strings.HasPrefix(s, "Homepage:") {
			s = strings.TrimPrefix(s, "Homepage:")
			s := strings.TrimSpace(s)
			lmlDB[i].homepage = s
			continue
		}
		if strings.HasPrefix(s, "Download:") {
			s = strings.TrimPrefix(s, "Download:")
			s := strings.TrimSpace(s)
			lmlDB[i].download = s
			continue
		}
		if strings.HasPrefix(s, "Wikipedia:") {
			s = strings.TrimPrefix(s, "Wikipedia:")
			s := strings.TrimSpace(s)
			lmlDB[i].wikipedia = s
			continue
		}
		if strings.HasPrefix(s, "Distrowatch:") {
			s = strings.TrimPrefix(s, "Distrowatch:")
			s := strings.TrimSpace(s)
			lmlDB[i].distrowatch = s
			continue
		}
		if strings.HasPrefix(s, "Size (mebibytes):") {
			if strings.Contains(s, "-") != true {
				continue
			}
			s = strings.TrimPrefix(s, "Size (mebibytes):")
			s = strings.TrimSpace(s)
			st := strings.Split(s, "-")
			n, err = strconv.ParseUint(st[0], 10, 32)
			lmlDB[i].sizeMin = int(n)
			if err != nil {
				panic(err)
			}
			n, err = strconv.ParseUint(st[1], 10, 32)
			lmlDB[i].sizeMax = int(n)
			if err != nil {
				panic(err)
			}
			continue
		}
		if strings.HasPrefix(s, "Last Stable Version:") {
			s = strings.TrimPrefix(s, "Last Stable Version:")
			s := strings.TrimSpace(s)
			lmlDB[i].stableVer = s
			//fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Last Release:") {
			s = strings.TrimPrefix(s, "Last Release:")
			s := strings.TrimSpace(s)
			lmlDB[i].lastRelease = s
			//fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Purpose:") {
			s = strings.TrimPrefix(s, "Purpose:")
			s := strings.TrimSpace(s)
			lmlDB[i].target = []int{}
			k := 0
			for j, v := range Target {
				if strings.Contains(s, j) {
					lmlDB[i].target = append(lmlDB[i].target, int(v))
					k++
				}

			}
			continue
		}
		if strings.HasPrefix(s, "Operating System:") {
			s = strings.TrimPrefix(s, "Operating System:")
			s = strings.TrimSpace(s)
			lmlDB[i].os = OS[s]
			continue
		}
		if strings.HasPrefix(s, "Primary Language(s): ") {
			s = strings.TrimPrefix(s, "Primary Language(s): ")
			//fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "State:") {
			s = strings.TrimPrefix(s, "State:")
			s = strings.TrimSpace(s)
			lmlDB[i].State = State[s]
			//fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Media:") {
			s = strings.TrimPrefix(s, "Media:")
			s := strings.TrimSpace(s)
			lmlDB[i].media = []int{}
			k := 0
			for j, v := range Media {
				if strings.Contains(s, j) {
					lmlDB[i].media = append(lmlDB[i].media, int(v))
					k++
				}

			}
			//fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Architecture:") {
			s = strings.TrimPrefix(s, "Architecture:") //architecture
			s := strings.TrimSpace(s)
			lmlDB[i].architecture = []int{}
			k := 0
			for j, v := range Architecture {
				if strings.Contains(s, j) {
					lmlDB[i].architecture = append(lmlDB[i].architecture, int(v))
					k++
				}

			}
			//fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Note:") {
			s = strings.TrimPrefix(s, "Note:")
			s := strings.TrimSpace(s)
			lmlDB[i].note = s
			continue
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, value := range lmlDB {
		fmt.Println("")
		fmt.Println(value.Number)
		fmt.Println(value.Name)
		fmt.Println(value.homepage)
		fmt.Println(value.download)
		fmt.Println(value.wikipedia)
		fmt.Println(value.distrowatch)
		fmt.Println(value.sizeMin)
		fmt.Println(value.sizeMax)
		fmt.Println(value.stableVer)
		fmt.Println(value.lastRelease)
		fmt.Println(value.os)
		fmt.Println(value.target)
		fmt.Println(value.State)
		fmt.Println(value.media)
		fmt.Println(value.architecture)
		fmt.Println(value.note)
		fmt.Println("___________________________________________________")
	}
	// fmt.Println(OS)
}
