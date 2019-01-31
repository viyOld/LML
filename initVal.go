package main

import (
	"bufio"
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
		if strings.HasPrefix(s, "License:") {
			flag = 6
			continue
		}
		if strings.HasPrefix(s, "End:") {
			return
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
		case 6:
			License[st[1]] = byte(n)
		default:
			//fmt.Println(st)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// --------------------------------------------------------------------------------------------------
func readStartDb() {
	var (
		n uint64
		i = -1
	)

	file, err := os.Open("./db/StartDB.txt")
	if err != nil {
		panic(err)
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
			lmlDB[i].Homepage = s
			continue
		}
		if strings.HasPrefix(s, "Download:") {
			s = strings.TrimPrefix(s, "Download:")
			s := strings.TrimSpace(s)
			lmlDB[i].Download = s
			continue
		}
		if strings.HasPrefix(s, "Wikipedia:") {
			s = strings.TrimPrefix(s, "Wikipedia:")
			s := strings.TrimSpace(s)
			lmlDB[i].Wikipedia = s
			continue
		}
		if strings.HasPrefix(s, "Distrowatch:") {
			s = strings.TrimPrefix(s, "Distrowatch:")
			s := strings.TrimSpace(s)
			lmlDB[i].Distrowatch = s
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
			lmlDB[i].SizeMin = int(n)
			if err != nil {
				panic(err)
			}
			n, err = strconv.ParseUint(st[1], 10, 32)
			lmlDB[i].SizeMax = int(n)
			if err != nil {
				panic(err)
			}
			continue
		}
		if strings.HasPrefix(s, "Last Stable Version:") {
			s = strings.TrimPrefix(s, "Last Stable Version:")
			s := strings.TrimSpace(s)
			lmlDB[i].StableVer = s
			//fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Last Release:") {
			s = strings.TrimPrefix(s, "Last Release:")
			s := strings.TrimSpace(s)
			lmlDB[i].LastRelease = s
			//fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Purpose:") {
			s = strings.TrimPrefix(s, "Purpose:")
			s := strings.TrimSpace(s)
			lmlDB[i].Target = []int{}
			k := 0
			for j, v := range Target {
				if strings.Contains(s, j) {
					lmlDB[i].Target = append(lmlDB[i].Target, int(v))
					k++
				}

			}
			continue
		}
		if strings.HasPrefix(s, "Operating System:") {
			s = strings.TrimPrefix(s, "Operating System:")
			s = strings.TrimSpace(s)
			lmlDB[i].OS = OS[s]
			continue
		}
		if strings.HasPrefix(s, "Primary Language(s): ") {
			s = strings.TrimPrefix(s, "Primary Language(s): ")
			s = strings.TrimSpace(s)
			lmlDB[i].Language = s
			continue
		}
		if strings.HasPrefix(s, "State:") {
			s = strings.TrimPrefix(s, "State:")
			s = strings.TrimSpace(s)
			lmlDB[i].State = State[s]
			continue
		}
		if strings.HasPrefix(s, "Media:") {
			s = strings.TrimPrefix(s, "Media:")
			s := strings.TrimSpace(s)
			lmlDB[i].Media = []int{}
			k := 0
			for j, v := range Media {
				if strings.Contains(s, j) {
					lmlDB[i].Media = append(lmlDB[i].Media, int(v))
					k++
				}

			}
			//fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Architecture:") {
			s = strings.TrimPrefix(s, "Architecture:") //architecture
			s := strings.TrimSpace(s)
			lmlDB[i].Architecture = []int{}
			k := 0
			for j, v := range Architecture {
				if strings.Contains(s, j) {
					lmlDB[i].Architecture = append(lmlDB[i].Architecture, int(v))
					k++
				}

			}
			//fmt.Println(s)
			continue
		}
		if strings.HasPrefix(s, "Note:") {
			s = strings.TrimPrefix(s, "Note:")
			s := strings.TrimSpace(s)
			lmlDB[i].Note = s
			continue
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// for _, value := range lmlDB {
	// 	fmt.Println("")
	// 	fmt.Println(value.Number)
	// 	fmt.Println(value.Name)
	// 	fmt.Println(value.Homepage)
	// 	fmt.Println(value.Download)
	// 	fmt.Println(value.Wikipedia)
	// 	fmt.Println(value.Distrowatch)
	// 	fmt.Println(value.SizeMin)
	// 	fmt.Println(value.SizeMax)
	// 	fmt.Println(value.StableVer)
	// 	fmt.Println(value.LastRelease)
	// 	fmt.Println(value.OS)
	// 	fmt.Println(value.Target)
	// 	fmt.Println(value.State)
	// 	fmt.Println(value.Media)
	// 	fmt.Println(value.Architecture)
	// 	fmt.Println(value.Note)
	// 	fmt.Println("___________________________________________________")
	// }
	// fmt.Println(OS)
}

func writeLMLdb() {

	filename := "./db/DB.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
	} else {
		if err := os.Rename(filename, "./db/DB.bak"); err != nil {
			panic(err)
		}
		if _, err = os.Create(filename); err != nil {
			panic(err)
		}
	}

	fileDB, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer fileDB.Close()

	fileDB.WriteString("#comments time" + "\n" + "\n")
	for i, v := range lmlDB {
		if _, err = fileDB.WriteString("Number: " + strconv.Itoa(v.Number) + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Name: " + v.Name + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Homepage: " + v.Homepage + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Download: " + v.Download + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Wikipedia: " + v.Wikipedia + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Distrowatch: " + v.Distrowatch + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("SizeMin: " + strconv.Itoa(v.SizeMin) + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("SizeMax: " + strconv.Itoa(v.SizeMax) + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("StableVer: " + v.StableVer + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("LastRelease: " + v.LastRelease + "\n"); err != nil {
			panic(err)
		}
		str := ""
		for _, val := range lmlDB[i].Target {
			str = str + " " + strconv.Itoa(val)
		}
		if _, err = fileDB.WriteString("Target: " + str + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("OS: " + strconv.Itoa(int(v.OS)) + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("BasedOS: " + v.BasedOS + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("License: " + strconv.Itoa(int(v.License)) + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Language: " + v.Language + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("State: " + strconv.Itoa(int(v.State)) + "\n"); err != nil {
			panic(err)
		}
		str = ""
		for _, val := range lmlDB[i].Media {
			str = str + " " + strconv.Itoa(val)
		}
		if _, err = fileDB.WriteString("Media: " + str + "\n"); err != nil {
			panic(err)
		}
		str = ""
		for _, val := range lmlDB[i].Architecture {
			str = str + " " + strconv.Itoa(val)
		}
		if _, err = fileDB.WriteString("Architecture: " + str + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Note: " + v.Note + "\n"); err != nil {
			panic(err)
		}
		if _, err = fileDB.WriteString("Rating: " + strconv.Itoa(v.Rating) + "\n"); err != nil {
			panic(err)
		}
		fileDB.WriteString("\n")
	}
}

func readDB() {
	var (
		n   uint64
		num int
	)

	file, err := os.Open("./db/DB.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//читаємо строку з файла
		s := strings.TrimSpace(scanner.Text())
		if s == "" {
			continue
		}
		if strings.HasPrefix(s, "//") {
			continue
		}
		if strings.HasPrefix(s, "#") {
			continue
		}
		if strings.HasPrefix(s, "Number:") {
			s = strings.TrimPrefix(s, "Number:")
			s := strings.TrimSpace(s)
			n, err = strconv.ParseUint(s, 10, 32)
			if err != nil {
				panic(err)
			}
			lmlDB = append(lmlDB, listMediaLive{})
			num = len(lmlDB) - 1
			lmlDB[num].Number = int(n)
			continue
		}
		if strings.HasPrefix(s, "Name:") {
			s = strings.TrimPrefix(s, "Name:")
			s := strings.TrimSpace(s)
			lmlDB[num].Name = s
			continue
		}
		if strings.HasPrefix(s, "Homepage:") {
			s = strings.TrimPrefix(s, "Homepage:")
			s := strings.TrimSpace(s)
			lmlDB[num].Homepage = s
			continue
		}
		if strings.HasPrefix(s, "Download:") {
			s = strings.TrimPrefix(s, "Download:")
			s := strings.TrimSpace(s)
			lmlDB[num].Download = s
			continue
		}
		if strings.HasPrefix(s, "Wikipedia:") {
			s = strings.TrimPrefix(s, "Wikipedia:")
			s := strings.TrimSpace(s)
			lmlDB[num].Wikipedia = s
			continue
		}
		if strings.HasPrefix(s, "Distrowatch:") {
			s = strings.TrimPrefix(s, "Distrowatch:")
			s := strings.TrimSpace(s)
			lmlDB[num].Distrowatch = s
			continue
		}
		if strings.HasPrefix(s, "SizeMin:") {
			s = strings.TrimPrefix(s, "SizeMin:")
			s := strings.TrimSpace(s)
			n, err = strconv.ParseUint(s, 10, 32)
			if err != nil {
				panic(err)
			}
			lmlDB[num].SizeMin = int(n)
			continue
		}
		if strings.HasPrefix(s, "SizeMax:") {
			s = strings.TrimPrefix(s, "SizeMax:")
			s := strings.TrimSpace(s)
			n, err = strconv.ParseUint(s, 10, 32)
			if err != nil {
				panic(err)
			}
			lmlDB[num].SizeMax = int(n)
			continue
		}
		if strings.HasPrefix(s, "StableVer:") {
			s = strings.TrimPrefix(s, "StableVer:")
			s := strings.TrimSpace(s)
			lmlDB[num].StableVer = s
			continue
		}
		if strings.HasPrefix(s, "LastRelease:") {
			s = strings.TrimPrefix(s, "LastRelease:")
			s := strings.TrimSpace(s)
			lmlDB[num].LastRelease = s
			continue
		}

		if strings.HasPrefix(s, "Target:") {
			s = strings.TrimPrefix(s, "Target:")
			s := strings.TrimSpace(s)
			lmlDB[num].Target = []int{}
			//split
			st := strings.Split(s, " ")
			for _, val := range st {
				x, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				lmlDB[num].Target = append(lmlDB[num].Target, x)
			}
			continue
		}

		if strings.HasPrefix(s, "OS:") {
			s = strings.TrimPrefix(s, "OS:")
			s = strings.TrimSpace(s)
			x, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			lmlDB[num].OS = byte(x)
			continue
		}

		if strings.HasPrefix(s, "BasedOS:") {
			s = strings.TrimPrefix(s, "BasedOS:")
			s := strings.TrimSpace(s)
			lmlDB[num].BasedOS = s
			continue
		}

		if strings.HasPrefix(s, "License:") {
			s = strings.TrimPrefix(s, "License:")
			s = strings.TrimSpace(s)
			x, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			lmlDB[num].License = byte(x)
			continue
		}

		if strings.HasPrefix(s, "Language:") {
			s = strings.TrimPrefix(s, "Language:")
			s := strings.TrimSpace(s)
			lmlDB[num].Language = s
			continue
		}

		if strings.HasPrefix(s, "State:") {
			s = strings.TrimPrefix(s, "State:")
			s = strings.TrimSpace(s)
			x, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			lmlDB[num].State = byte(x)
			continue
		}

		if strings.HasPrefix(s, "Media:") {
			s = strings.TrimPrefix(s, "Media:")
			s := strings.TrimSpace(s)
			lmlDB[num].Media = []int{}
			st := strings.Split(s, " ")
			for _, val := range st {
				x, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				lmlDB[num].Media = append(lmlDB[num].Media, x)
			}
			continue
		}

		if strings.HasPrefix(s, "Architecture: ") {
			s = strings.TrimPrefix(s, "Architecture: ")
			s := strings.TrimSpace(s)
			lmlDB[num].Architecture = []int{}
			st := strings.Split(s, " ")
			for _, val := range st {
				x, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				lmlDB[num].Architecture = append(lmlDB[num].Architecture, x)
			}
			continue
		}

		if strings.HasPrefix(s, "Note:") {
			s = strings.TrimPrefix(s, "Note:")
			s := strings.TrimSpace(s)
			lmlDB[num].Note = s
			continue
		}

		if strings.HasPrefix(s, "Rating:") {
			s = strings.TrimPrefix(s, "Rating:")
			s := strings.TrimSpace(s)
			n, err = strconv.ParseUint(s, 10, 32)
			if err != nil {
				panic(err)
			}
			lmlDB[num].Rating = int(n)
			continue
		}

	}

}
